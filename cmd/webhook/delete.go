package webhook

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/xavidop/dialogflow-cx-cli/cmd/cmdutils"
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	"github.com/xavidop/dialogflow-cx-cli/pkg/webhook"
)

// deleteCmd represents the delete webhook set command
var deleteCmd = &cobra.Command{
	Use:     "delete [name]",
	Aliases: []string{"d", "del", "remove", "deletes", "removes"},
	Short:   "delete a webhook",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Get the information
		locationID, _ := cmd.Flags().GetString("location-id")
		projectID, _ := cmd.Flags().GetString("project-id")
		agentName, _ := cmd.Flags().GetString("agent-name")
		force, _ := cmd.Flags().GetString("force")
		name := args[0]

		if err := webhook.Delete(name, locationID, projectID, agentName, force); err != nil {
			global.Log.Errorf("%s", err.Error())
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
	webhookCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().StringP("project-id", "p", "", "Dialogflow CX Project ID (required)")
	if err := deleteCmd.MarkFlagRequired("project-id"); err != nil {
		global.Log.Errorf("%s", err.Error())
		os.Exit(1)
	}
	deleteCmd.Flags().StringP("location-id", "l", "", "Dialogflow CX Location ID of the Project (required)")
	if err := deleteCmd.MarkFlagRequired("location-id"); err != nil {
		global.Log.Errorf("%s", err.Error())
		os.Exit(1)
	}
	deleteCmd.Flags().StringP("agent-name", "a", "", "Dialogflow CX Agent Name (required)")
	if err := deleteCmd.MarkFlagRequired("agent-name"); err != nil {
		global.Log.Errorf("%s", err.Error())
		os.Exit(1)
	}
	deleteCmd.Flags().StringP("force", "f", "", "Forces to delete the webhook and its references in environments and flows. Possible values: true or false (optional)")

}
