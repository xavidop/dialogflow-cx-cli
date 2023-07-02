package agent

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/xavidop/dialogflow-cx-cli/cmd/agent/types"
	"github.com/xavidop/dialogflow-cx-cli/cmd/cmdutils"
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	"github.com/xavidop/dialogflow-cx-cli/pkg/agent"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:     "update [agent-name]",
	Aliases: []string{"update", "u"},
	Short:   "Update an agent",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Get the information
		locationID, _ := cmd.Flags().GetString("location-id")
		projectID, _ := cmd.Flags().GetString("project-id")
		description, _ := cmd.Flags().GetString("description")
		timezone, _ := cmd.Flags().GetString("timezone")
		supportedLocales, _ := cmd.Flags().GetStringSlice("supported-locales")
		avatarURI, _ := cmd.Flags().GetString("avatar-uri")
		enableStackdriverLogging, _ := cmd.Flags().GetString("enable-stackdriver-logging")
		enableInteractionLogging, _ := cmd.Flags().GetString("enable-interaction-logging")
		enableSpeechAdaptation, _ := cmd.Flags().GetString("enable-speech-adaptation")
		enableSpellCorrection, _ := cmd.Flags().GetString("enable-spell-correction")

		agentName := args[0]

		updateInput := &types.CreateUpdateAgent{
			Description:              description,
			TimeZone:                 timezone,
			SupportedLanguageCodes:   supportedLocales,
			AvatarURI:                avatarURI,
			EnableStackdriverLogging: enableStackdriverLogging,
			EnableInteractionLogging: enableInteractionLogging,
			EnableSpellCorrection:    enableSpellCorrection,
			EnableSpeechAdaptation:   enableSpeechAdaptation,
		}

		if err := agent.Update(agentName, locationID, projectID, updateInput); err != nil {
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
	agentCmd.AddCommand(updateCmd)

	updateCmd.Flags().StringP("project-id", "p", "", "Dialogflow CX Project ID (required)")
	if err := updateCmd.MarkFlagRequired("project-id"); err != nil {
		global.Log.Errorf(err.Error())
		os.Exit(1)
	}
	updateCmd.Flags().StringP("location-id", "l", "", "Dialogflow CX Location ID of the Project (required)")
	if err := updateCmd.MarkFlagRequired("location-id"); err != nil {
		global.Log.Errorf(err.Error())
		os.Exit(1)
	}
	updateCmd.Flags().StringP("timezone", "t", "", "Timezone of the agent from the time zone database https://www.iana.org/time-zones. Example: America/New_York, Europe/Madrid (optional)")
	updateCmd.Flags().StringP("description", "d", "", "Description of the agent (optional)")
	updateCmd.Flags().StringSliceP("supported-locales", "x", []string{}, "Supported locales of the agent, comma separated. Example: fr,es,de (optional)")
	updateCmd.Flags().StringP("avatar-uri", "r", "", "Avatar URI of the agent (optional)")
	updateCmd.Flags().StringP("enable-stackdriver-logging", "a", "", "Enable Google Stackdriver logging for this agent. Possible values: true or false (optional)")
	updateCmd.Flags().StringP("enable-interaction-logging", "b", "", "Enable interaction logging for this agent. Possible values: true or false (optional)")
	updateCmd.Flags().StringP("enable-speech-adaptation", "s", "", "Enable speech adaptation for this agent. Possible values: true or false (optional)")
	updateCmd.Flags().StringP("enable-spell-correction", "n", "", "Enable spell correction for this agent. Possible values: true or false (optional)")

}
