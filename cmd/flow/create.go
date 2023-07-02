package flow

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/xavidop/dialogflow-cx-cli/cmd/cmdutils"
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	"github.com/xavidop/dialogflow-cx-cli/pkg/flow"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:     "create [flow-name]",
	Aliases: []string{"creates", "c"},
	Short:   "Creates a flow in an agent",
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

		if err := flow.Create(flowName, description, locationID, projectID, agentName, localeId, nluClassificationThreshold, nluModelType, nluModelTrainingMode); err != nil {
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
	flowCmd.AddCommand(createCmd)

	createCmd.Flags().StringP("project-id", "p", "", "Dialogflow CX Project ID (required)")
	if err := createCmd.MarkFlagRequired("project-id"); err != nil {
		global.Log.Errorf(err.Error())
		os.Exit(1)
	}
	createCmd.Flags().StringP("location-id", "l", "", "Dialogflow CX Location ID of the Project (required)")
	if err := createCmd.MarkFlagRequired("location-id"); err != nil {
		global.Log.Errorf(err.Error())
		os.Exit(1)
	}
	createCmd.Flags().StringP("agent-name", "a", "", "Dialogflow CX Agent Name (required)")
	if err := createCmd.MarkFlagRequired("agent-name"); err != nil {
		global.Log.Errorf(err.Error())
		os.Exit(1)
	}
	createCmd.Flags().StringP("locale", "e", "", "Locale of the flow. Default: agent locale (optional)")
	createCmd.Flags().StringP("description", "d", "", "Description for this flow (optional)")
	createCmd.Flags().StringP("nlu-classification-threshold", "s", "0.3", "NLU Classification Threshold. From 0.0 (completely uncertain) to 1.0 (completely certain). Default: 0.3 (optional)")
	createCmd.Flags().StringP("nlu-model-type", "m", "standard", "NLU Model Type. Possible values: advanced or standard. Default: standard (optional)")
	createCmd.Flags().StringP("nlu-model-training-mode", "t", "manual", "NLU Model training mode. Possible values: automatic or manual. Default: manual (optional)")

}
