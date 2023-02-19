package entitytype

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/xavidop/dialogflow-cx-cli/cmd/cmdutils"
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	"github.com/xavidop/dialogflow-cx-cli/pkg/entitytype"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:     "update [entity-type-name]",
	Aliases: []string{"updates", "u"},
	Short:   "Updates an entity type in an agent",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Get the information
		locationID, _ := cmd.Flags().GetString("location-id")
		projectID, _ := cmd.Flags().GetString("project-id")
		agentName, _ := cmd.Flags().GetString("agent-name")
		localeId, _ := cmd.Flags().GetString("locale")
		entities, _ := cmd.Flags().GetStringSlice("entities")
		redacted, _ := cmd.Flags().GetBool("redacted")
		entityTypeName := args[0]

		if err := entitytype.Update(entityTypeName, locationID, projectID, agentName, localeId, entities, redacted); err != nil {
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
	entitytypeCmd.AddCommand(updateCmd)

	updateCmd.Flags().StringP("project-id", "p", "", "Dialogflow CX Project ID")
	updateCmd.Flags().StringP("location-id", "l", "", "Dialogflow CX Location ID of the Project")
	updateCmd.Flags().StringP("agent-name", "a", "", "Dialogflow CX Agent Name")
	updateCmd.Flags().StringP("locale", "e", "", "Optional. Locale of the intent. Default: agent locale")
	updateCmd.Flags().BoolP("redacted", "r", false, "Optional. Set the entity type as redacted")
	updateCmd.Flags().StringSliceP("entities", "n", []string{}, "List of the entities for this entity type, comma separated. Format: entity1@synonym1|synonym2,entity2@synonym1|synonym2. Example: pikachu@25|pika,charmander@3")
}
