package webhook

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/xavidop/dialogflow-cx-cli/cmd/cmdutils"
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	"github.com/xavidop/dialogflow-cx-cli/pkg/webhook"
)

// updateCmd represents the update webhook set command
var updateCmd = &cobra.Command{
	Use:     "update [name]",
	Aliases: []string{"updates", "u"},
	Short:   "update a webhook",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Get the information
		locationID, _ := cmd.Flags().GetString("location-id")
		projectID, _ := cmd.Flags().GetString("project-id")
		agentName, _ := cmd.Flags().GetString("agent-name")
		environment, _ := cmd.Flags().GetString("environment")
		url, _ := cmd.Flags().GetString("url")
		name := args[0]

		if err := webhook.Update(name, url, locationID, projectID, agentName, environment); err != nil {
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
	webhookCmd.AddCommand(updateCmd)

	updateCmd.Flags().StringP("url", "r", "", "Webhook URL")
	updateCmd.Flags().StringP("agent-name", "a", "", "Dialogflow CX Agent Name")
	updateCmd.Flags().StringP("project-id", "p", "", "Dialogflow CX Project ID")
	updateCmd.Flags().StringP("location-id", "l", "", "Dialogflow CX Location ID of the Project")
	updateCmd.Flags().StringP("environment", "e", "global", "Optional. Environment where you want to set the webhook url. Default: global")
}
