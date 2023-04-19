package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/xavidop/dialogflow-cx-cli/cmd/cmdutils"
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	"github.com/xavidop/dialogflow-cx-cli/pkg/dialog"
)

// VersionCmd represents the version command
var dialogCmd = &cobra.Command{
	Use:     "dialog",
	Aliases: []string{"d"},
	Short:   "Test your CX Agent interactively directly from your terminal",
	Run: func(cmd *cobra.Command, args []string) {
		if err := dialog.Dialog(); err != nil {
			global.Log.Errorf(err.Error())
			os.Exit(1)
		}
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		cmdutils.PreRun(cmd.Name())
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(dialogCmd)

	dialogCmd.Flags().StringP("project-id", "p", "", "Dialogflow CX Project ID")
	dialogCmd.Flags().StringP("location-id", "l", "", "Dialogflow CX Location ID of the Project")
	dialogCmd.Flags().StringP("agent-name", "a", "", "Dialogflow CX Agent Name")
	dialogCmd.Flags().StringP("locale", "e", "", "Optional. Locale of the intent. Default: agent locale")
}
