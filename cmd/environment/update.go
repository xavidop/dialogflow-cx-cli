package environment

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/xavidop/dialogflow-cx-cli/cmd/cmdutils"
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	"github.com/xavidop/dialogflow-cx-cli/pkg/environment"
)

// updateCmd represents the update environment set command
var updateCmd = &cobra.Command{
	Use:     "update [name]",
	Aliases: []string{"updates", "u"},
	Short:   "update an environment",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Get the information
		// Get the information
		locationID, _ := cmd.Flags().GetString("location-id")
		projectID, _ := cmd.Flags().GetString("project-id")
		agentName, _ := cmd.Flags().GetString("agent-name")
		flowVersions, _ := cmd.Flags().GetStringSlice("flow-versions")
		description, _ := cmd.Flags().GetString("description")
		name := args[0]

		if err := environment.Update(name, description, locationID, projectID, agentName, flowVersions); err != nil {
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
	environmentCmd.AddCommand(updateCmd)

	updateCmd.Flags().StringSliceP("flow-versions", "s", []string{}, "List of Flow and its version to be added to this environment, comma separated. Format: flowName1@version1,flowName2|version2. Example: Default Start Flow@v1.0.0|Buy Flow@v1.0.1 (required)")
	updateCmd.MarkFlagRequired("flow-versions")
	updateCmd.Flags().StringP("agent-name", "a", "", "Dialogflow CX Agent Name (required)")
	updateCmd.MarkFlagRequired("agent-name")
	updateCmd.Flags().StringP("project-id", "p", "", "Dialogflow CX Project ID (required)")
	updateCmd.MarkFlagRequired("project-id")
	updateCmd.Flags().StringP("location-id", "l", "", "Dialogflow CX Location ID of the Project (required)")
	updateCmd.MarkFlagRequired("location-id")
	updateCmd.Flags().StringP("description", "d", "", "Optional. Description for this environment")
}
