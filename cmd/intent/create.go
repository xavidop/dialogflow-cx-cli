package intent

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/xavidop/dialogflow-cx-cli/cmd/cmdutils"
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	"github.com/xavidop/dialogflow-cx-cli/pkg/intent"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:     "create [intent-name]",
	Aliases: []string{"creates", "c"},
	Short:   "Creates an intent in an agent",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Get the information
		locationID, _ := cmd.Flags().GetString("location-id")
		projectID, _ := cmd.Flags().GetString("project-id")
		agentName, _ := cmd.Flags().GetString("agent-name")
		localeId, _ := cmd.Flags().GetString("locale")
		trainingPhrases, _ := cmd.Flags().GetStringSlice("training-phrases")
		description, _ := cmd.Flags().GetString("description")
		intentName := args[0]

		if err := intent.Create(intentName, description, locationID, projectID, agentName, localeId, trainingPhrases); err != nil {
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
	intentCmd.AddCommand(createCmd)

	createCmd.Flags().StringP("project-id", "p", "", "Dialogflow CX Project ID (required)")
	if err := createCmd.MarkFlagRequired("project-id"); err != nil {
		global.Log.Errorf("%s", err.Error())
		os.Exit(1)
	}
	createCmd.Flags().StringP("location-id", "l", "", "Dialogflow CX Location ID of the Project (required)")
	if err := createCmd.MarkFlagRequired("location-id"); err != nil {
		global.Log.Errorf("%s", err.Error())
		os.Exit(1)
	}
	createCmd.Flags().StringP("agent-name", "a", "", "Dialogflow CX Agent Name (required)")
	if err := createCmd.MarkFlagRequired("agent-name"); err != nil {
		global.Log.Errorf("%s", err.Error())
		os.Exit(1)
	}
	createCmd.Flags().StringSliceP("training-phrases", "t", []string{}, "List of the training phrases for this intent, comma separated. Entities, add @entity-type to the word: word@entity-type in the training phrase. Example: hello,hi how are you today@sys.date,morning!. (required)")
	if err := createCmd.MarkFlagRequired("training-phrases"); err != nil {
		global.Log.Errorf("%s", err.Error())
		os.Exit(1)
	}
	createCmd.Flags().StringP("locale", "e", "", "Locale of the intent. Default: agent locale (optional)")
	createCmd.Flags().StringP("description", "d", "", "Description for this intent (optional)")

}
