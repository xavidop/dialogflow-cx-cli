package webhook

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/xavidop/dialogflow-cx-cli/cmd/cmdutils"
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	"github.com/xavidop/dialogflow-cx-cli/pkg/webhook"
)

// createCmd represents the create webhook set command
var createCmd = &cobra.Command{
	Use:     "create [name]",
	Aliases: []string{"creates", "s"},
	Short:   "create a webhook",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Get the information
		locationID, _ := cmd.Flags().GetString("location-id")
		projectID, _ := cmd.Flags().GetString("project-id")
		agentName, _ := cmd.Flags().GetString("agent-name")
		environment, _ := cmd.Flags().GetString("environment")
		url, _ := cmd.Flags().GetString("url")
		name := args[0]

		if err := webhook.Create(name, url, locationID, projectID, agentName, environment); err != nil {
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
	webhookCmd.AddCommand(createCmd)

	createCmd.Flags().StringP("url", "r", "", "Webhook URL (required)")
	if err := createCmd.MarkFlagRequired("url"); err != nil {
		global.Log.Errorf(err.Error())
		os.Exit(1)
	}
	createCmd.Flags().StringP("project-id", "p", "", "Dialogflow CX Project ID (required)")
	if err := createCmd.MarkFlagRequired("project-id"); err != nil {
		global.Log.Errorf(err.Error())
		os.Exit(1)
	}
	createCmd.Flags().StringP("location-id", "l", "", "Dialogflow CX Location ID of the Project (required)")
	if err := createCmd.MarkFlagRequired("location-id"); err != nil {
		global.Log.Errorf(err.Error())
		os.Exit(1)
	}
	createCmd.Flags().StringP("agent-name", "a", "", "Dialogflow CX Agent Name (required)")
	if err := createCmd.MarkFlagRequired("agent-name"); err != nil {
		global.Log.Errorf(err.Error())
		os.Exit(1)
	}
	createCmd.Flags().StringP("environment", "e", "global", "Environment where you want to set the webhook url. Default: global (optional)")
}
