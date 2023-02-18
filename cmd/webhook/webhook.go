package webhook

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/xavidop/dialogflow-cx-cli/cmd/cmdutils"
)

// webhookCmd represents the webhook root command
var webhookCmd = &cobra.Command{
	Use:     "webhook",
	Aliases: []string{"w"},
	Short:   "Actions on webhook commands",
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
	rootCmd.AddCommand(webhookCmd)
}
