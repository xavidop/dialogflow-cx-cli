package intent

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/xavidop/dialogflow-cx-cli/cmd/cmdutils"
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	"github.com/xavidop/dialogflow-cx-cli/pkg/intent"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:     "delete [intent-name]",
	Aliases: []string{"deletes", "d", "del", "remove"},
	Short:   "Deletes an intent in an agent",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Get the information
		locationID, _ := cmd.Flags().GetString("location-id")
		projectID, _ := cmd.Flags().GetString("project-id")
		agentName, _ := cmd.Flags().GetString("agent-name")
		intentName := args[0]

		if err := intent.Delete(intentName, locationID, projectID, agentName); err != nil {
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
	intentCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().StringP("project-id", "p", "", "Dialogflow CX Project ID")
	deleteCmd.Flags().StringP("location-id", "l", "", "Dialogflow CX Location ID of the Project")
	deleteCmd.Flags().StringP("agent-name", "a", "", "Dialogflow CX Agent Name")
}
