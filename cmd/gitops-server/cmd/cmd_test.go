package ui

import (
	"net/http"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/weaveworks/weave-gitops/cmd/gitops/cmderrors"
)

func TestNoIssuerURL(t *testing.T) {
	os.Setenv("WEAVE_GITOPS_AUTH_ENABLED", "true")
	defer os.Unsetenv("WEAVE_GITOPS_AUTH_ENABLED")

	cmd := NewCommand()
	cmd.SetArgs([]string{
		"ui", "run",
	})

	err := cmd.Execute()
	assert.ErrorIs(t, err, cmderrors.ErrNoIssuerURL)
}

func TestNoClientID(t *testing.T) {
	os.Setenv("WEAVE_GITOPS_AUTH_ENABLED", "true")
	defer os.Unsetenv("WEAVE_GITOPS_AUTH_ENABLED")

	cmd := NewCommand()
	cmd.SetArgs([]string{
		"ui", "run",
		"--oidc-issuer-url=http://weave.works",
	})

	err := cmd.Execute()
	assert.ErrorIs(t, err, cmderrors.ErrNoClientID)
}

func TestNoClientSecret(t *testing.T) {
	os.Setenv("WEAVE_GITOPS_AUTH_ENABLED", "true")
	defer os.Unsetenv("WEAVE_GITOPS_AUTH_ENABLED")

	cmd := NewCommand()
	cmd.SetArgs([]string{
		"ui", "run",
		"--oidc-issuer-url=http://weave.works",
		"--oidc-client-id=client-id",
	})

	err := cmd.Execute()
	assert.ErrorIs(t, err, cmderrors.ErrNoClientSecret)
}

func TestNoRedirectURL(t *testing.T) {
	os.Setenv("WEAVE_GITOPS_AUTH_ENABLED", "true")
	defer os.Unsetenv("WEAVE_GITOPS_AUTH_ENABLED")

	cmd := NewCommand()
	cmd.SetArgs([]string{
		"ui", "run",
		"--oidc-issuer-url=http://weave.works",
		"--oidc-client-id=client-id",
		"--oidc-client-secret=client-secret",
	})

	err := cmd.Execute()
	assert.ErrorIs(t, err, cmderrors.ErrNoRedirectURL)
}

func TestMissingTLSKeyOrCert(t *testing.T) {
	log := logrus.New()
	err := run.ListenAndServe(&http.Server{}, false, "foo", "", log)
	assert.ErrorIs(t, err, cmderrors.ErrNoTLSCertOrKey)

	err = run.ListenAndServe(&http.Server{}, false, "", "bar", log)
	assert.ErrorIs(t, err, cmderrors.ErrNoTLSCertOrKey)
}
