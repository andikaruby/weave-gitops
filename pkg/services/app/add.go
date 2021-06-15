package app

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/pkg/errors"
	"github.com/weaveworks/weave-gitops/pkg/git"
	"github.com/weaveworks/weave-gitops/pkg/kube"
)

type DeploymentType string
type SourceType string
type ConfigType string

const (
	DeployTypeKustomize DeploymentType = "kustomize"
	DeployTypeHelm      DeploymentType = "helm"

	SourceTypeGit  SourceType = "git"
	SourceTypeHelm SourceType = "helm"

	ConfigTypeUserRepo ConfigType = ""
	ConfigTypeNone     ConfigType = "NONE"
)

type AddParams struct {
	Dir                  string
	Name                 string
	Url                  string
	Path                 string
	Branch               string
	PrivateKey           string
	PrivateKeyPass       string
	DeploymentType       string
	Chart                string
	SourceType           string
	CommitManifests      bool
	AutomationRepo       string
	AutomationRepoPath   string
	AutomationRepoBranch string
	Namespace            string
	DryRun               bool
}

func (a *App) Add(params AddParams) error {
	fmt.Print("Updating parameters from environment... ")
	params, err := a.updateParametersIfNecessary(params)
	if err != nil {
		return errors.Wrap(err, "could not update parameters")
	}

	fmt.Print("done\n\n")
	fmt.Print("Checking cluster status... ")

	clusterStatus := a.kube.GetClusterStatus()
	fmt.Printf("%s\n\n", clusterStatus)

	switch clusterStatus {
	case kube.Unmodified:
		return errors.New("WeGO not installed... exiting")
	case kube.Unknown:
		return errors.New("WeGO can not determine cluster status... exiting")
	}

	clusterName, err := a.kube.GetClusterName()
	if err != nil {
		return err
	}

	fmt.Println("Generating deploy key...")
	secretRef, err := a.createAndUploadDeployKey([]string{params.Url, params.AutomationRepo}, clusterName, params.Namespace, params.DryRun)
	if err != nil {
		return errors.Wrap(err, "could not generate deploy key")
	}

	if params.AutomationRepo != "" {
		return a.addAppWithConfigInExternalRepo(params, clusterName, secretRef)
	}

	if params.CommitManifests {
		return a.addAppWithConfigInAppRepo(params, clusterName, secretRef)
	}

	return a.addAppWithNoConfigRepo(params, clusterName, secretRef)
}

func (a *App) updateParametersIfNecessary(params AddParams) (AddParams, error) {
	params.SourceType = string(SourceTypeGit)

	if params.Chart != "" {
		params.SourceType = string(SourceTypeHelm)
		params.DeploymentType = string(DeployTypeHelm)
		params.Name = params.Chart

		return params, nil
	}

	// Identifying repo url if not set by the user
	if params.Url == "" {
		url, err := a.getGitRemoteUrl(params)
		if err != nil {
			return params, err
		}

		params.Url = url
	} else {
		// making sure url is in the correct format
		params.Url = sanitizeRepoUrl(params.Url)

		// resetting Dir param since Url has priority over it
		params.Dir = ""
	}

	fmt.Printf("using URL: '%s' of origin from git config...\n\n", params.Url)

	if params.Name == "" {
		params.Name = generateResourceName(params.Url)
	}

	return params, nil
}

func (a *App) getGitRemoteUrl(params AddParams) (string, error) {
	repo, err := a.git.Open(params.Dir)
	if err != nil {
		return "", errors.Wrapf(err, "failed to open repository: %s", params.Dir)
	}

	remote, err := repo.Remote("origin")
	if err != nil {
		return "", errors.Wrapf(err, "failed to find the origin remote in the repository")
	}

	urls := remote.Config().URLs
	if len(urls) == 0 {
		return "", errors.Errorf("remote config in %s does not have an url", params.Dir)
	}

	return sanitizeRepoUrl(urls[0]), nil
}

func (a *App) addAppWithNoConfigRepo(params AddParams, clusterName string, secretRef string) error {
	// Returns the source, app spec and kustomization
	source, appGoat, appSpec, err := a.generateAppManifests(params, secretRef, clusterName)
	if err != nil {
		return errors.Wrap(err, "could not generate application GitOps Automation manifests")
	}

	fmt.Println("Applying manifests to the cluster...")
	return a.applyToCluster(params, source, appGoat, appSpec)
}

func (a *App) addAppWithConfigInAppRepo(params AddParams, clusterName string, secretRef string) error {
	// Returns the source, app spec and kustomization
	source, appGoat, appSpec, err := a.generateAppManifests(params, secretRef, clusterName)
	if err != nil {
		return errors.Wrap(err, "could not generate application GitOps Automation manifests")
	}

	fmt.Println("Applying manifests to the cluster...")
	if err := a.applyToCluster(params, source, appGoat, appSpec); err != nil {
		return errors.Wrap(err, "could not apply manifests to the cluster")
	}

	// a local directory has not been passed, so we clone the repo passed in the --url
	if params.Dir == "" {
		fmt.Printf("Cloning %s...\n", params.Url)
		if err := a.cloneRepo(params.Url, params.Branch, params.DryRun); err != nil {
			return errors.Wrap(err, "failed to clone application repo")
		}
	}

	fmt.Println("Writing manifests to disk...")
	if !params.DryRun {
		if err := a.writeAppYaml(".wego", params.Name, appSpec); err != nil {
			return errors.Wrap(err, "failed writing app.yaml to disk")
		}

		if err := a.writeAppGoats(".wego", params.Name, clusterName, source, appGoat); err != nil {
			return errors.Wrap(err, "failed writing app.yaml to disk")
		}
	}

	return a.commitAndPush(params, func(fname string) bool {
		return strings.Contains(fname, ".wego")
	})
}

func (a *App) addAppWithConfigInExternalRepo(params AddParams, clusterName string, secretRef string) error {
	// making sure the url is in good format
	params.AutomationRepo = sanitizeRepoUrl(params.AutomationRepo)

	// Returns the source, app spec and kustomization
	appSource, appGoat, appSpec, err := a.generateAppManifests(params, secretRef, clusterName)
	if err != nil {
		return errors.Wrap(err, "could not generate application GitOps Automation manifests")
	}

	targetSource, targetGoat, err := a.generateTargetManifests(params, secretRef, clusterName)
	if err != nil {
		return errors.Wrap(err, "could not generate target GitOps Automation manifests")
	}

	fmt.Println("Applying manifests to the cluster...")
	if err := a.applyToCluster(params, appSource, appGoat, appSpec, targetSource, targetGoat); err != nil {
		return errors.Wrapf(err, "could not apply manifests to the cluster")
	}

	if err := a.cloneRepo(params.AutomationRepo, params.AutomationRepoBranch, params.DryRun); err != nil {
		return errors.Wrap(err, "failed to clone application repo")
	}

	fmt.Println("Writing manifests to disk...")
	if !params.DryRun {
		if err := a.writeAppYaml(params.AutomationRepoPath, params.Name, appSpec); err != nil {
			return errors.Wrap(err, "failed writing app.yaml to disk")
		}

		if err := a.writeAppGoats(params.AutomationRepoPath, params.Name, clusterName, appSource, appGoat); err != nil {
			return errors.Wrap(err, "failed writing app.yaml to disk")
		}
	}

	return a.commitAndPush(params, func(fname string) bool {
		return strings.Contains(fname, strings.TrimPrefix(params.AutomationRepoPath, "./"))
	})
}

func (a *App) generateAppManifests(params AddParams, secretRef string, clusterName string) ([]byte, []byte, []byte, error) {
	var sourceManifest, appManifest, appGoatManifest []byte
	var err error
	fmt.Println("Generating Source manifest...")
	sourceManifest, err = a.generateSource(params, secretRef)
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "could not set up GitOps for user repository")
	}

	fmt.Println("Generating GitOps automation manifests...")
	appGoatManifest, err = a.generateApplicationGoat(params)
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, fmt.Sprintf("could not create GitOps automation for '%s'", params.Name))
	}

	fmt.Println("Generating Application spec manifest...")
	appManifest, err = generateAppYaml(params)
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, fmt.Sprintf("could not create app.yaml for '%s'", params.Name))
	}

	return sourceManifest, appGoatManifest, appManifest, nil
}

func (a *App) generateTargetManifests(params AddParams, secretRef string, clusterName string) ([]byte, []byte, error) {
	targetSource, err := a.generateTargetSource(params, secretRef)
	if err != nil {
		return nil, nil, errors.Wrap(err, "could not generate target source manifests")
	}

	targetGoat, err := a.generateTargetGoat(params, clusterName)
	if err != nil {
		return nil, nil, errors.Wrap(err, "could not generate target goat manifests")
	}

	return targetSource, targetGoat, nil
}

func (a *App) generateTargetSource(params AddParams, secretRef string) ([]byte, error) {
	repoName := generateResourceName(params.AutomationRepo)

	return a.flux.CreateSourceGit(repoName, params.AutomationRepo, params.AutomationRepoBranch, secretRef, params.Namespace)
}

func (a *App) generateTargetGoat(params AddParams, clusterName string) ([]byte, error) {
	repoName := urlToRepoName(params.AutomationRepo)

	targetPath := filepath.Join(params.AutomationRepoPath, "targets", clusterName)

	return a.flux.CreateKustomization(fmt.Sprintf("weave-gitops-%s", clusterName), repoName, targetPath, params.Namespace)
}

func (a *App) commitAndPush(params AddParams, filters ...func(string) bool) error {
	fmt.Println("Commiting and pushing wego resources for application...")
	if params.DryRun {
		return nil
	}

	_, err := a.git.Commit(git.Commit{
		Author:  git.Author{Name: "Weave Gitops", Email: "weave-gitops@weave.works"},
		Message: "Add App manifests",
	}, filters...)
	if err != nil && err != git.ErrNoStagedFiles {
		return fmt.Errorf("failed to commit sync manifests: %w", err)
	}

	if err == nil {
		fmt.Println("Pushing app manifests to repository...")
		if err = a.git.Push(context.Background()); err != nil {
			return fmt.Errorf("failed to push manifests: %w", err)
		}
	} else {
		fmt.Println("App manifests are up to date")
	}

	return nil
}

func (a *App) createAndUploadDeployKey(reposUrls []string, clusterName string, namespace string, dryRun bool) (string, error) {
	secretRef := fmt.Sprintf("weave-gitops-%s", clusterName)
	if dryRun {
		return secretRef, nil
	}

	for _, repoUrl := range reposUrls {
		if repoUrl == "" {
			continue
		}

		repoUrl = sanitizeRepoUrl(repoUrl)

		deployKey, err := a.flux.CreateSecretGit(secretRef, repoUrl, namespace)
		if err != nil {
			return "", errors.Wrap(err, "could not create git secret")
		}

		owner, err := getOwnerFromUrl(repoUrl)
		if err != nil {
			return "", err
		}

		repoName := urlToRepoName(repoUrl)
		if err := a.gitProviders.UploadDeployKey(owner, repoName, deployKey); err != nil {
			return "", errors.Wrap(err, "error uploading deploy key")
		}
	}

	return secretRef, nil
}

func (a *App) generateSource(params AddParams, secretRef string) ([]byte, error) {
	switch SourceType(params.SourceType) {
	case SourceTypeGit:
		sourceManifest, err := a.flux.CreateSourceGit(params.Name, params.Url, params.Branch, secretRef, params.Namespace)
		if err != nil {
			return nil, errors.Wrap(err, "could not create git source")
		}

		return sourceManifest, nil
	case SourceTypeHelm:
		return a.flux.CreateSourceHelm(params.Name, params.Url, params.Namespace)
	default:
		return nil, fmt.Errorf("unknown source type: %v", params.SourceType)
	}
}

func (a *App) generateApplicationGoat(params AddParams) ([]byte, error) {
	switch params.DeploymentType {
	case string(DeployTypeKustomize):
		return a.flux.CreateKustomization(params.Name, params.Name, params.Path, params.Namespace)
	case string(DeployTypeHelm):
		switch params.SourceType {
		case string(SourceTypeHelm):
			return a.flux.CreateHelmReleaseHelmRepository(params.Name, params.Chart, params.Namespace)
		case string(SourceTypeGit):
			return a.flux.CreateHelmReleaseGitRepository(params.Name, params.Name, params.Path, params.Namespace)
		default:
			return nil, fmt.Errorf("invalid source type: %v", params.SourceType)
		}
	default:
		return nil, fmt.Errorf("invalid deployment type: %v", params.DeploymentType)
	}
}

func (a *App) applyToCluster(params AddParams, manifests ...[]byte) error {
	if params.DryRun {
		for _, manifest := range manifests {
			fmt.Printf("%s\n", manifest)
		}
		return nil
	}

	for _, manifest := range manifests {
		if out, err := a.kube.Apply(manifest, params.Namespace); err != nil {
			return errors.Wrap(err, fmt.Sprintf("could not apply manifest: %s", string(out)))
		}
	}

	return nil
}

func (a *App) cloneRepo(url string, branch string, dryRun bool) error {
	if dryRun {
		return nil
	}

	url = sanitizeRepoUrl(url)

	repoDir, err := ioutil.TempDir("", "user-repo-")
	if err != nil {
		return errors.Wrap(err, "failed creating temp. directory to clone repo")
	}

	_, err = a.git.Clone(context.Background(), repoDir, url, branch)
	if err != nil {
		return errors.Wrapf(err, "failed cloning user repo: %s", url)
	}

	return nil
}

func (a *App) writeAppYaml(basePath string, name string, manifest []byte) error {
	manifestPath := filepath.Join(basePath, "apps", name, "app.yaml")

	return a.git.Write(manifestPath, manifest)
}

func (a *App) writeAppGoats(basePath string, name string, clusterName string, manifests ...[]byte) error {
	goatPath := filepath.Join(basePath, "targets", clusterName, name, fmt.Sprintf("%s-gitops-runtime.yaml", name))

	goat := bytes.Join(manifests, []byte(""))
	return a.git.Write(goatPath, goat)
}

func generateAppYaml(params AddParams) ([]byte, error) {
	const appYamlTemplate = `---
apiVersion: wego.weave.works/v1alpha1
kind: Application
metadata:
  name: {{ .AppName }}
spec:
  path: {{ .AppPath }}
  url: {{ .AppURL }}
`
	// Create app.yaml
	t, err := template.New("appYaml").Parse(appYamlTemplate)
	if err != nil {
		return nil, errors.Wrap(err, "could not parse app yaml template")
	}

	var populated bytes.Buffer
	err = t.Execute(&populated, struct {
		AppName string
		AppPath string
		AppURL  string
	}{params.Name, params.Path, params.Url})
	if err != nil {
		return nil, errors.Wrap(err, "could not execute populated template")
	}
	return populated.Bytes(), nil
}

func generateResourceName(url string) string {
	return strings.ReplaceAll(urlToRepoName(url), "_", "-")
}

func getOwnerFromUrl(url string) (string, error) {
	parts := strings.Split(url, "/")
	if len(parts) < 2 {
		return "", fmt.Errorf("could not get owner from url %s", url)
	}
	return parts[len(parts)-2], nil
}

func urlToRepoName(url string) string {
	return strings.TrimSuffix(filepath.Base(url), ".git")
}

func sanitizeRepoUrl(url string) string {
	trimmed := ""

	if !strings.HasSuffix(url, ".git") {
		url = url + ".git"
	}

	sshPrefix := "git@github.com:"
	if strings.HasPrefix(url, sshPrefix) {
		trimmed = strings.TrimPrefix(url, sshPrefix)
	}

	httpsPrefix := "https://github.com/"
	if strings.HasPrefix(url, httpsPrefix) {
		trimmed = strings.TrimPrefix(url, httpsPrefix)
	}

	if trimmed != "" {
		return "ssh://git@github.com/" + trimmed
	}

	return url
}

// NOTE: ready to save the targets automation in phase 2
// func (a *App) writeTargetGoats(basePath string, name string, manifests ...[]byte) error {
// 	goatPath := filepath.Join(basePath, "targets", fmt.Sprintf("%s-gitops-runtime.yaml", name))

// 	goat := bytes.Join(manifests, []byte(""))
// 	return a.git.Write(goatPath, goat)
// }
