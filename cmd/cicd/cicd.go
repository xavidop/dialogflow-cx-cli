package cicd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/xavidop/dialogflow-cx-cli/cmd/cmdutils"
)

// cicdCmd represents the cicd root command
var cicdCmd = &cobra.Command{
	Use:     "cicd",
	Aliases: []string{"cicd", "ci", "cd"},
	Short:   "Actions on CICD testings",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		cmdutils.PreRun(cmd.Name())
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
	},
}

func Register(rootCmd *cobra.Command) {
	rootCmd.AddCommand(cicdCmd)
}
