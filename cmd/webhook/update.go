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
		flexible, _ := cmd.Flags().GetString("flexible")
		requestBody, _ := cmd.Flags().GetString("request-body")
		parametersMapping, _ := cmd.Flags().GetString("parameters-mapping")
		name := args[0]

		if err := webhook.Update(name, url, locationID, projectID, agentName, environment, flexible, requestBody, parametersMapping); err != nil {
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

	updateCmd.Flags().StringP("url", "r", "", "Webhook URL (required)")
	if err := updateCmd.MarkFlagRequired("url"); err != nil {
		global.Log.Errorf(err.Error())
		os.Exit(1)
	}
	updateCmd.Flags().StringP("project-id", "p", "", "Dialogflow CX Project ID (required)")
	if err := updateCmd.MarkFlagRequired("project-id"); err != nil {
		global.Log.Errorf(err.Error())
		os.Exit(1)
	}
	updateCmd.Flags().StringP("location-id", "l", "", "Dialogflow CX Location ID of the Project (required)")
	if err := updateCmd.MarkFlagRequired("location-id"); err != nil {
		global.Log.Errorf(err.Error())
		os.Exit(1)
	}
	updateCmd.Flags().StringP("agent-name", "a", "", "Dialogflow CX Agent Name (required)")
	if err := updateCmd.MarkFlagRequired("agent-name"); err != nil {
		global.Log.Errorf(err.Error())
		os.Exit(1)
	}
	updateCmd.Flags().StringP("environment", "e", "global", "Environment where you want to set the webhook url. Default: global (optional)")
	updateCmd.Flags().StringP("flexible", "f", "false", "Creates a flexible webhook. Possible values: true or false (optional)")
	updateCmd.Flags().StringP("request-body", "t", "", "Creates a request body for flexible webhook. It has to be in JSON Format. This only applies to flexible webhooks (optional)")
	updateCmd.Flags().StringP("parameters-mapping", "m", "", "Creates a parameter mapping for flexible webhook, comma separated. The format is parameter@json-path,paramter2@json-path2. Example: my-param@$.fully.qualified.path.to.field. This only applies to flexible webhooks (optional)")

}
