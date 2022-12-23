package cicd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/xavidop/dialogflow-cx-cli/cmd/cmdutils"
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	"github.com/xavidop/dialogflow-cx-cli/pkg/cicd"
)

// executeCmd represents the execute CICD command
var executeCmd = &cobra.Command{
	Use:     "execute [environment]",
	Aliases: []string{"execute", "e", "exe", "exec"},
	Short:   "Executes a CICD pipeline for a specific environment",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Get the information
		locationID, _ := cmd.Flags().GetString("location-id")
		projectID, _ := cmd.Flags().GetString("project-id")
		agentName, _ := cmd.Flags().GetString("agent-name")
		envName := args[0]

		if err := cicd.ExecutePipeline(envName, locationID, projectID, agentName); err != nil {
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
	cicdCmd.AddCommand(executeCmd)

	executeCmd.Flags().StringP("agent-name", "a", "", "Dialogflow CX Agent Name")
	executeCmd.Flags().StringP("project-id", "p", "", "Dialogflow CX Project ID")
	executeCmd.Flags().StringP("location-id", "l", "", "Dialogflow CX Location ID of the Project")
}
