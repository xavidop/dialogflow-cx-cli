package environment

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/xavidop/dialogflow-cx-cli/cmd/cmdutils"
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	"github.com/xavidop/dialogflow-cx-cli/pkg/environment"
)

// executeCicdCmd represents the execute CICD command
var executeCicdCmd = &cobra.Command{
	Use:     "execute-cicd [environment]",
	Aliases: []string{"execute", "e", "exe", "exec"},
	Short:   "Executes a CI/CD pipeline for a specific environment",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Get the information
		locationID, _ := cmd.Flags().GetString("location-id")
		projectID, _ := cmd.Flags().GetString("project-id")
		agentName, _ := cmd.Flags().GetString("agent-name")
		envName := args[0]

		if err := environment.ExecutePipeline(envName, locationID, projectID, agentName); err != nil {
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
	environmentCmd.AddCommand(executeCicdCmd)

	executeCicdCmd.Flags().StringP("agent-name", "a", "", "Dialogflow CX Agent Name (required)")
	executeCicdCmd.MarkFlagRequired("agent-name")
	executeCicdCmd.Flags().StringP("project-id", "p", "", "Dialogflow CX Project ID (required)")
	executeCicdCmd.MarkFlagRequired("project-id")
	executeCicdCmd.Flags().StringP("location-id", "l", "", "Dialogflow CX Location ID of the Project (required)")
	executeCicdCmd.MarkFlagRequired("location-id")

}
