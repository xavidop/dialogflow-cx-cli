package generator

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/xavidop/dialogflow-cx-cli/cmd/cmdutils"
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	"github.com/xavidop/dialogflow-cx-cli/pkg/generator"
)

// updateCmd represents the update generator set command
var updateCmd = &cobra.Command{
	Use:     "update [name]",
	Aliases: []string{"updates", "u"},
	Short:   "update a generator",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Get the information
		locationID, _ := cmd.Flags().GetString("location-id")
		projectID, _ := cmd.Flags().GetString("project-id")
		agentName, _ := cmd.Flags().GetString("agent-name")
		prompt, _ := cmd.Flags().GetString("prompt")
		localeId, _ := cmd.Flags().GetString("locale")
		name := args[0]

		if err := generator.Update(name, locationID, projectID, agentName, prompt, localeId); err != nil {
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
	generatorCmd.AddCommand(updateCmd)

	updateCmd.Flags().StringP("prompt", "r", "", "Prompt (required)")
	if err := updateCmd.MarkFlagRequired("prompt"); err != nil {
		global.Log.Errorf("%s", err.Error())
		os.Exit(1)
	}
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

	updateCmd.Flags().StringP("locale", "e", "", "Locale of the generator. Default: agent locale (optional)")

}
