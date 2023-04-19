package environment

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/xavidop/dialogflow-cx-cli/cmd/cmdutils"
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	"github.com/xavidop/dialogflow-cx-cli/pkg/environment"
)

// createCmd represents the create environment set command
var createCmd = &cobra.Command{
	Use:     "create [name]",
	Aliases: []string{"creates", "s"},
	Short:   "create an environment",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Get the information
		locationID, _ := cmd.Flags().GetString("location-id")
		projectID, _ := cmd.Flags().GetString("project-id")
		agentName, _ := cmd.Flags().GetString("agent-name")
		flowVersions, _ := cmd.Flags().GetStringSlice("flow-versions")
		description, _ := cmd.Flags().GetString("description")
		name := args[0]

		if err := environment.Create(name, description, locationID, projectID, agentName, flowVersions); err != nil {
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
	environmentCmd.AddCommand(createCmd)

	createCmd.Flags().StringSliceP("flow-versions", "s", []string{}, "List of Flow and its version to be added to this environment, comma separated. Format: flowName1@version1,flowName2|version2. Example: Default Start Flow@v1.0.0|Buy Flow@v1.0.1 (required)")
	createCmd.MarkFlagRequired("flow-versions")
	createCmd.Flags().StringP("agent-name", "a", "", "Dialogflow CX Agent Name (required)")
	createCmd.MarkFlagRequired("agent-name")
	createCmd.Flags().StringP("project-id", "p", "", "Dialogflow CX Project ID (required)")
	createCmd.MarkFlagRequired("project-id")
	createCmd.Flags().StringP("location-id", "l", "", "Dialogflow CX Location ID of the Project (required)")
	createCmd.MarkFlagRequired("location-id")
	createCmd.Flags().StringP("description", "d", "", "Optional. Description for this environment")
}
