package environment

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/xavidop/dialogflow-cx-cli/cmd/cmdutils"
)

// environmentCmd represents the environment root command
var environmentCmd = &cobra.Command{
	Use:     "environment",
	Aliases: []string{"env", "e"},
	Short:   "Actions on environment",
	Run: func(cmd *cobra.Command, args []string) {
		if err := cmd.Help(); err != nil {
			os.Exit(1)
		}
		os.Exit(0)
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		cmdutils.PreRun(cmd.Name())
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
	},
}

func Register(rootCmd *cobra.Command) {
	rootCmd.AddCommand(environmentCmd)
}
