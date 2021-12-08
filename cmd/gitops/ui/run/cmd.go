package run

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/go-logr/zapr"
	"github.com/mattn/go-isatty"
	"github.com/pkg/browser"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/weaveworks/weave-gitops/pkg/kube"
	"github.com/weaveworks/weave-gitops/pkg/server"
	"github.com/weaveworks/weave-gitops/pkg/server/auth"
	"go.uber.org/zap"
)

// Options contains all the options for the `ui run` command.
type Options struct {
	Port              string
	HelmRepoNamespace string
	HelmRepoName      string
	Path              string
	LoggingEnabled    bool
	OIDC              OIDCAuthenticationOptions
}

// OIDCAuthenticationOptions contains the OIDC authentication options for the
// `ui run` command.
type OIDCAuthenticationOptions struct {
	IssuerURL      string
	ClientID       string
	ClientSecret   string
	RedirectURL    string
	CookieDuration string
}

var options Options

// Command returns the `ui run` command
func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run [--log]",
		Short: "Runs gitops ui",
		RunE:  runCmd,
	}

	options = Options{}

	cmd.Flags().BoolVarP(&options.LoggingEnabled, "log", "l", false, "enable logging for the ui")
	cmd.Flags().StringVar(&options.Port, "port", "9001", "UI port")
	cmd.Flags().StringVar(&options.Path, "path", "", "Path url")
	cmd.Flags().StringVar(&options.HelmRepoNamespace, "helm-repo-namespace", "default", "the namespace of the Helm Repository resource to scan for profiles")
	cmd.Flags().StringVar(&options.HelmRepoName, "helm-repo-name", "weaveworks-charts", "the name of the Helm Repository resource to scan for profiles")

	if server.AuthEnabled() {
		cmd.Flags().StringVar(&options.OIDC.IssuerURL, "oidc-issuer-url", "", "The URL of the OpenID Connect issuer")
		cmd.Flags().StringVar(&options.OIDC.ClientID, "oidc-client-id", "", "The client ID for the OpenID Connect client")
		cmd.Flags().StringVar(&options.OIDC.ClientSecret, "oidc-client-secret", "", "The client secret to use with OpenID Connect issuer")
		cmd.Flags().StringVar(&options.OIDC.RedirectURL, "oidc-redirect-url", "", "The OAuth2 redirect URL")
		cmd.Flags().StringVar(&options.OIDC.CookieDuration, "oidc-cookie-duration", "1h", "The duration of the ID token cookie. It should be set in the format: number + time unit (s,m,h) e.g., 20m")
	}

	return cmd
}

func runCmd(cmd *cobra.Command, args []string) error {
	var log = logrus.New()

	mux := http.NewServeMux()

	mux.Handle("/health/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("ok"))

		if err != nil {
			log.Errorf("error writing health check: %s", err)
		}
	}))

	assetFS := getAssets()
	assetHandler := http.FileServer(http.FS(assetFS))
	redirector := createRedirector(assetFS, log)

	appConfig, err := server.DefaultApplicationsConfig()
	if err != nil {
		return fmt.Errorf("could not create http client: %w", err)
	}

	if !options.LoggingEnabled {
		appConfig.Logger = zapr.NewLogger(zap.NewNop())
	}

	rest, clusterName, err := kube.RestConfig()
	if err != nil {
		return fmt.Errorf("could not create client config: %w", err)
	}

	_, rawClient, err := kube.NewKubeHTTPClientWithConfig(rest, clusterName)
	if err != nil {
		return fmt.Errorf("could not create kube http client: %w", err)
	}

	profilesConfig := server.NewProfilesConfig(rawClient, options.HelmRepoNamespace, options.HelmRepoName)

	var authConfig *auth.AuthConfig

	if server.AuthEnabled() {
		_, err := url.Parse(options.OIDC.IssuerURL)
		if err != nil {
			return fmt.Errorf("invalid issuer URL: %w", err)
		}

		redirectURL, err := url.Parse(options.OIDC.RedirectURL)
		if err != nil {
			return fmt.Errorf("invalid redirect URL: %w", err)
		}

		var oidcIssueSecureCookies bool
		if redirectURL.Scheme == "https" {
			oidcIssueSecureCookies = true
		}

		oidcCookieDuration, err := time.ParseDuration(options.OIDC.CookieDuration)
		if err != nil {
			return fmt.Errorf("invalid cookie duration: %w", err)
		}

		cfg, err := auth.NewAuthConfig(cmd.Context(), options.OIDC.IssuerURL,
			options.OIDC.ClientID, options.OIDC.ClientSecret, options.OIDC.RedirectURL,
			oidcIssueSecureCookies, oidcCookieDuration, http.DefaultClient,
			appConfig.Logger)
		if err != nil {
			return fmt.Errorf("could not create auth config: %w", err)
		}

		cfg.Logger().Info("Registering callback route")
		// Register /callback handler with mux
		auth.RegisterAuthHandler(mux, "/oauth2", cfg)

		authConfig = cfg
	}

	appAndProfilesHandlers, err := server.NewHandlers(context.Background(), &server.Config{AppConfig: appConfig, ProfilesConfig: profilesConfig, AuthConfig: authConfig})
	if err != nil {
		return fmt.Errorf("could not create handler: %w", err)
	}

	mux.Handle("/v1/", appAndProfilesHandlers)

	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// Assume anything with a file extension in the name is a static asset.
		extension := filepath.Ext(req.URL.Path)
		// We use the golang http.FileServer for static file requests.
		// This will return a 404 on normal page requests, ie /some-page.
		// Redirect all non-file requests to index.html, where the JS routing will take over.
		if extension == "" {
			if server.AuthEnabled() {
				auth.WithWebAuth(redirector, authConfig).ServeHTTP(w, req)
			} else {
				redirector(w, req)
			}
			return
		}
		assetHandler.ServeHTTP(w, req)
	}))

	addr := "0.0.0.0:" + options.Port
	srv := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	go func() {
		log.Infof("Serving on port %s", options.Port)

		if err := srv.ListenAndServe(); err != nil {
			log.Error(err, "server exited")
			os.Exit(1)
		}
	}()

	if isatty.IsTerminal(os.Stdout.Fd()) {
		url := fmt.Sprintf("http://%s/%s", addr, options.Path)

		log.Printf("Opening browser at %s", url)

		if err := browser.OpenURL(url); err != nil {
			return fmt.Errorf("failed to open the browser: %w", err)
		}
	}

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer func() {
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		return fmt.Errorf("Server Shutdown Failed: %w", err)
	}

	return nil
}

//go:embed dist/*
var static embed.FS

func getAssets() fs.FS {
	f, err := fs.Sub(static, "dist")

	if err != nil {
		panic(err)
	}

	return f
}

// A redirector ensures that index.html always gets served.
// The JS router will take care of actual navigation once the index.html page lands.
func createRedirector(fsys fs.FS, log logrus.FieldLogger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		indexPage, err := fsys.Open("index.html")

		if err != nil {
			log.Error(err, "could not open index.html page")
			w.WriteHeader(http.StatusInternalServerError)

			return
		}

		stat, err := indexPage.Stat()
		if err != nil {
			log.Error(err, "could not get index.html stat")
			w.WriteHeader(http.StatusInternalServerError)

			return
		}

		bt := make([]byte, stat.Size())
		_, err = indexPage.Read(bt)

		if err != nil {
			log.Error(err, "could not read index.html")
			w.WriteHeader(http.StatusInternalServerError)

			return
		}

		_, err = w.Write(bt)

		if err != nil {
			log.Error(err, "error writing index.html")
			w.WriteHeader(http.StatusInternalServerError)

			return
		}
	}
}
