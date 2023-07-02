package agent

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/xavidop/dialogflow-cx-cli/cmd/cmdutils"
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	"github.com/xavidop/dialogflow-cx-cli/internal/utils"
	"github.com/xavidop/dialogflow-cx-cli/pkg/agent"
)

// restoreCmd represents the restore command
var restoreCmd = &cobra.Command{
	Use:     "restore [agent-name]",
	Aliases: []string{"restore", "r", "re"},
	Short:   "Executes a restore action for a specific agent",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Get the information
		locationID, _ := cmd.Flags().GetString("location-id")
		projectID, _ := cmd.Flags().GetString("project-id")
		input, _ := cmd.Flags().GetString("input")
		if err := utils.ValidateAgentFileType(input); err != nil {
			global.Log.Errorf(err.Error())
			os.Exit(1)
		}
		agentName := args[0]

		if err := agent.Restore(locationID, projectID, agentName, input); err != nil {
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
	agentCmd.AddCommand(restoreCmd)

	restoreCmd.Flags().StringP("project-id", "p", "", "Dialogflow CX Project ID (required)")
	if err := restoreCmd.MarkFlagRequired("project-id"); err != nil {
		global.Log.Errorf(err.Error())
		os.Exit(1)
	}
	restoreCmd.Flags().StringP("location-id", "l", "", "Dialogflow CX Location ID of the Project (required)")
	if err := restoreCmd.MarkFlagRequired("location-id"); err != nil {
		global.Log.Errorf(err.Error())
		os.Exit(1)
	}
	restoreCmd.Flags().StringP("input", "i", "agent.blob", "Input file name. Default: agent.blob (optional)")

}
