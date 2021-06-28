package version

import (
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/weaveworks/go-checkpoint"
	"github.com/weaveworks/weave-gitops/pkg/fluxops"

	"github.com/spf13/cobra"
)

// The current wego version
var Version = "v0.0.0"
var GitCommit = ""
var Branch = ""
var BuildTime = ""

var Cmd = &cobra.Command{
	Use:   "version",
	Short: "Display wego version",
	Run:   runCmd,
	PostRun: func(cmd *cobra.Command, args []string) {
		CheckVersion(CheckpointParams())
	},
}

func runCmd(cmd *cobra.Command, args []string) {
	fmt.Println("Current Version:", Version)
	fmt.Println("GitCommit:", GitCommit)
	fmt.Println("BuildTime:", BuildTime)
	fmt.Println("Branch:", Branch)
	version, err := CheckFluxVersion()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Flux Version:", version)
}

// CheckVersion looks to see if there is a newer version of the software available
func CheckVersion(p *checkpoint.CheckParams) {
	checkResponse, err := checkpoint.Check(p)
	if err == nil && checkResponse.Outdated {
		log.Infof("wego version %s is available; please update at %s",
			checkResponse.CurrentVersion, checkResponse.CurrentDownloadURL)
	}
}

// CheckpointParams creates the structure to pass to CheckVersion
func CheckpointParams() *checkpoint.CheckParams {
	return &checkpoint.CheckParams{
		Product: "weave-gitops",
		Version: Version,
	}
}

// CheckpointParamsWithFlags adds the object and command from the arguments list to the checkpoint parameters
func CheckpointParamsWithFlags(params *checkpoint.CheckParams, c *cobra.Command) *checkpoint.CheckParams {
	// wego uses noun verb command syntax and the parent command will have the noun and the command passed in will be the verb
	params.Flags = map[string]string{
		"object":  c.Parent().Name(),
		"command": c.Name(),
	}
	return params
}
func CheckFluxVersion() (string, error) {
	fluxops.SetFluxHandler(&fluxops.QuietFluxHandler{})
	output, err := fluxops.CallFlux("-v")
	if err != nil {
		return "", err
	}

	// change string format to match other version info
	version := strings.ReplaceAll(string(output), "flux version ", "v")

	return version, nil
}
