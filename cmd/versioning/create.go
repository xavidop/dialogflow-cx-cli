package versioning

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/xavidop/dialogflow-cx-cli/cmd/cmdutils"
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	"github.com/xavidop/dialogflow-cx-cli/pkg/versioning"
)

// createCmd represents the create version set command
var createCmd = &cobra.Command{
	Use:     "create [name]",
	Aliases: []string{"creates", "s"},
	Short:   "create a version",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Get the information
		locationID, _ := cmd.Flags().GetString("location-id")
		projectID, _ := cmd.Flags().GetString("project-id")
		agentName, _ := cmd.Flags().GetString("agent-name")
		startFlow, _ := cmd.Flags().GetString("start-flow")
		description, _ := cmd.Flags().GetString("description")
		name := args[0]

		if err := versioning.Create(name, description, startFlow, locationID, projectID, agentName); err != nil {
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
	versioningCmd.AddCommand(createCmd)

	createCmd.Flags().StringP("start-flow", "s", "", "Start Flow name to create the version")
	createCmd.Flags().StringP("agent-name", "a", "", "Dialogflow CX Agent Name")
	createCmd.Flags().StringP("project-id", "p", "", "Dialogflow CX Project ID")
	createCmd.Flags().StringP("location-id", "l", "", "Dialogflow CX Location ID of the Project")
	createCmd.Flags().StringP("description", "d", "", "Optional. Description for this version")
}
