package flowversion

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/xavidop/dialogflow-cx-cli/cmd/cmdutils"
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	"github.com/xavidop/dialogflow-cx-cli/pkg/flowversion"
)

// deleteCmd represents the delete version set command
var deleteCmd = &cobra.Command{
	Use:     "delete [name]",
	Aliases: []string{"d", "del", "remove", "deletes", "removes"},
	Short:   "delete a version",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Get the information
		locationID, _ := cmd.Flags().GetString("location-id")
		projectID, _ := cmd.Flags().GetString("project-id")
		agentName, _ := cmd.Flags().GetString("agent-name")
		startFlow, _ := cmd.Flags().GetString("start-flow")
		name := args[0]

		if err := flowversion.Delete(name, startFlow, locationID, projectID, agentName); err != nil {
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
	flowversionCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().StringP("start-flow", "s", "", "Start Flow name to create the version (required)")
	if err := deleteCmd.MarkFlagRequired("start-flow"); err != nil {
		global.Log.Errorf("%s", err.Error())
		os.Exit(1)
	}
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

}
