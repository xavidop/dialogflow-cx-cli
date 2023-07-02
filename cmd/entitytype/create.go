package entitytype

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/xavidop/dialogflow-cx-cli/cmd/cmdutils"
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	"github.com/xavidop/dialogflow-cx-cli/pkg/entitytype"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:     "create [entity-type-name]",
	Aliases: []string{"creates", "c"},
	Short:   "Creates an entity type in an agent",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Get the information
		locationID, _ := cmd.Flags().GetString("location-id")
		projectID, _ := cmd.Flags().GetString("project-id")
		agentName, _ := cmd.Flags().GetString("agent-name")
		localeId, _ := cmd.Flags().GetString("locale")
		entities, _ := cmd.Flags().GetStringSlice("entities")
		redacted, _ := cmd.Flags().GetString("redacted")
		entityTypeName := args[0]

		if err := entitytype.Create(entityTypeName, locationID, projectID, agentName, localeId, entities, redacted); err != nil {
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
	entitytypeCmd.AddCommand(createCmd)

	createCmd.Flags().StringP("agent-name", "a", "", "Dialogflow CX Agent Name (required)")
	if err := createCmd.MarkFlagRequired("agent-name"); err != nil {
		global.Log.Errorf(err.Error())
		os.Exit(1)
	}
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
	createCmd.Flags().StringSliceP("entities", "n", []string{}, "List of the entities for this entity type, comma separated. Format: entity1@synonym1|synonym2,entity2@synonym1|synonym2. Example: pikachu@25|pika,charmander@3 (required)")
	if err := createCmd.MarkFlagRequired("entities"); err != nil {
		global.Log.Errorf(err.Error())
		os.Exit(1)
	}

	createCmd.Flags().StringP("locale", "e", "", "Locale of the intent. Default: agent locale (optional)")
	createCmd.Flags().StringP("redacted", "r", "", "Set the entity type as redacted. Possible values: true or false (optional)")
}
