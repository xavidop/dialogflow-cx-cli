package flow

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/xavidop/dialogflow-cx-cli/cmd/cmdutils"
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	"github.com/xavidop/dialogflow-cx-cli/pkg/flow"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:     "update [flow-name]",
	Aliases: []string{"updates", "u"},
	Short:   "Updates a flow in an agent",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Get the information
		locationID, _ := cmd.Flags().GetString("location-id")
		projectID, _ := cmd.Flags().GetString("project-id")
		agentName, _ := cmd.Flags().GetString("agent-name")
		localeId, _ := cmd.Flags().GetString("locale")
		description, _ := cmd.Flags().GetString("description")
		nluClassificationThreshold, _ := cmd.Flags().GetString("nlu-classification-threshold")
		nluModelType, _ := cmd.Flags().GetString("nlu-model-type")
		nluModelTrainingMode, _ := cmd.Flags().GetString("nlu-model-training-mode")
		flowName := args[0]

		if err := flow.Update(flowName, description, locationID, projectID, agentName, localeId, nluClassificationThreshold, nluModelType, nluModelTrainingMode); err != nil {
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
	flowCmd.AddCommand(updateCmd)

	updateCmd.Flags().StringP("project-id", "p", "", "Dialogflow CX Project ID")
	updateCmd.Flags().StringP("location-id", "l", "", "Dialogflow CX Location ID of the Project")
	updateCmd.Flags().StringP("agent-name", "a", "", "Dialogflow CX Agent Name")
	updateCmd.Flags().StringP("locale", "e", "", "Optional. Locale of the flow. Default: agent locale")
	updateCmd.Flags().StringP("description", "d", "", "Optional. Description for this flow")
	updateCmd.Flags().StringP("nlu-classification-threshold", "s", "", "Optional. NLU Classification Threshold. From 0.0 (completely uncertain) to 1.0 (completely certain)")
	updateCmd.Flags().StringP("nlu-model-type", "m", "", "Optional. NLU Model Type. Possible values: advanced or standard")
	updateCmd.Flags().StringP("nlu-model-training-mode", "t", "", "Optional. NLU Model training mode. Possible values: automatic or manual")

}
