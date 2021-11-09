package clusters_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/weaveworks/weave-gitops/cmd/gitops/root"
)

func TestEndpointNotSet(t *testing.T) {
	client := resty.New()
	cmd := root.RootCmd(client)
	cmd.SetArgs([]string{
		"delete", "cluster",
		"dev-cluster",
	})

	err := cmd.Execute()
	assert.EqualError(t, err, "the Weave GitOps Enterprise HTTP API endpoint flag (--endpoint) has not been set")
}

func TestPayload(t *testing.T) {
	client := resty.New()

	httpmock.ActivateNonDefault(client.GetClient())
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder(
		http.MethodDelete,
		"http://localhost:8000/v1/clusters",
		func(r *http.Request) (*http.Response, error) {
			var f interface{}
			err := json.NewDecoder(r.Body).Decode(&f)
			assert.NoError(t, err)

			m := f.(map[string]interface{})
			assert.Contains(t, m, "repositoryUrl")
			assert.Contains(t, m, "headBranch")
			assert.Contains(t, m, "baseBranch")
			assert.Contains(t, m, "title")
			assert.Contains(t, m, "description")
			assert.Contains(t, m, "clusterNames")
			assert.Contains(t, m, "commitMessage")
			assert.Contains(t, m, "credentials")

			return httpmock.NewJsonResponse(http.StatusOK, httpmock.File("../../../../pkg/adapters/testdata/pull_request_created.json"))
		},
	)

	cmd := root.RootCmd(client)
	cmd.SetArgs([]string{
		"delete", "cluster",
		"dev-cluster",
		"--endpoint", "http://localhost:8000",
	})

	err := cmd.Execute()
	assert.NoError(t, err)
}
