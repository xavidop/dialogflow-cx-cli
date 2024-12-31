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
	flowCmd.AddCommand(updateCmd)

	updateCmd.Flags().StringP("project-id", "p", "", "Dialogflow CX Project ID (required)")
	if err := updateCmd.MarkFlagRequired("project-id"); err != nil {
		global.Log.Errorf("%s", err.Error())
		os.Exit(1)
	}
	updateCmd.Flags().StringP("location-id", "l", "", "Dialogflow CX Location ID of the Project (required)")
	if err := updateCmd.MarkFlagRequired("location-id"); err != nil {
		global.Log.Errorf("%s", err.Error())
		os.Exit(1)
	}
	updateCmd.Flags().StringP("agent-name", "a", "", "Dialogflow CX Agent Name (required)")
	if err := updateCmd.MarkFlagRequired("agent-name"); err != nil {
		global.Log.Errorf("%s", err.Error())
		os.Exit(1)
	}
	updateCmd.Flags().StringP("locale", "e", "", "Locale of the flow. Default: agent locale (optional)")
	updateCmd.Flags().StringP("description", "d", "", "Description for this flow (optional)")
	updateCmd.Flags().StringP("nlu-classification-threshold", "s", "", "NLU Classification Threshold. From 0.0 (completely uncertain) to 1.0 (completely certain). (optional)")
	updateCmd.Flags().StringP("nlu-model-type", "m", "", "NLU Model Type. Possible values: advanced or standard (optional)")
	updateCmd.Flags().StringP("nlu-model-training-mode", "t", "", "NLU Model training mode. Possible values: automatic or manual (optional)")

}
