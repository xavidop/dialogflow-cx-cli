package flowversion

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/xavidop/dialogflow-cx-cli/cmd/cmdutils"
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	"github.com/xavidop/dialogflow-cx-cli/pkg/flowversion"
)

// updateCmd represents the update version set command
var updateCmd = &cobra.Command{
	Use:     "update [name]",
	Aliases: []string{"updates", "u"},
	Short:   "update a version",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Get the information
		locationID, _ := cmd.Flags().GetString("location-id")
		projectID, _ := cmd.Flags().GetString("project-id")
		agentName, _ := cmd.Flags().GetString("agent-name")
		startFlow, _ := cmd.Flags().GetString("start-flow")
		description, _ := cmd.Flags().GetString("description")
		name := args[0]

		if err := flowversion.Update(name, description, startFlow, locationID, projectID, agentName); err != nil {
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
	flowversionCmd.AddCommand(updateCmd)

	updateCmd.Flags().StringP("start-flow", "s", "", "Start Flow name to create the version (required)")
	if err := updateCmd.MarkFlagRequired("start-flow"); err != nil {
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
	updateCmd.Flags().StringP("description", "d", "", "Description for this version (optional)")
}
