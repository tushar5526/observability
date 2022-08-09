package cmd

import (
	"os"
	"os/user"
	"path/filepath"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "installer",
	Short: "Installs an observability stack focused on monitoring Gitpod.",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

var rootOpts struct {
	VersionMF         string
	StrictConfigParse bool
}

func init() {
	rootCmd.PersistentFlags().StringVar(&rootOpts.VersionMF, "debug-version-file", "", "path to a version manifest - not intended for production use")
	rootCmd.PersistentFlags().BoolVar(&rootOpts.StrictConfigParse, "strict-parse", true, "toggle strict configuration parsing")
}

type kubeConfig struct {
	Config string
}

// checkKubeConfig performs validation on the Kubernetes struct and fills in default values if necessary
func checkKubeConfig(kube *kubeConfig) error {
	if kube.Config == "" {
		kube.Config = os.Getenv("KUBECONFIG")
	}
	if kube.Config == "" {
		u, err := user.Current()
		if err != nil {
			return err
		}
		kube.Config = filepath.Join(u.HomeDir, ".kube", "config")
	}

	return nil
}