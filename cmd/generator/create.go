package generator

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/xavidop/dialogflow-cx-cli/cmd/cmdutils"
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	"github.com/xavidop/dialogflow-cx-cli/pkg/generator"
)

// createCmd represents the create generator set command
var createCmd = &cobra.Command{
	Use:     "create [name]",
	Aliases: []string{"creates", "s"},
	Short:   "create a generator",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Get the information
		locationID, _ := cmd.Flags().GetString("location-id")
		projectID, _ := cmd.Flags().GetString("project-id")
		agentName, _ := cmd.Flags().GetString("agent-name")
		prompt, _ := cmd.Flags().GetString("prompt")
		localeId, _ := cmd.Flags().GetString("locale")
		name := args[0]

		if err := generator.Create(name, locationID, projectID, agentName, prompt, localeId); err != nil {
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
	generatorCmd.AddCommand(createCmd)

	createCmd.Flags().StringP("prompt", "r", "", "Prompt (required)")
	if err := createCmd.MarkFlagRequired("prompt"); err != nil {
		global.Log.Errorf("%s", err.Error())
		os.Exit(1)
	}

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

	createCmd.Flags().StringP("locale", "e", "", "Locale of the generator. Default: agent locale (optional)")
}
