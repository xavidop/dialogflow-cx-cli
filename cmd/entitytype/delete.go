package entitytype

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/xavidop/dialogflow-cx-cli/cmd/cmdutils"
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	"github.com/xavidop/dialogflow-cx-cli/pkg/entitytype"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:     "delete [entity-type-name]",
	Aliases: []string{"deletes", "d", "remove", "del"},
	Short:   "Deletes an entity type in an agent",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Get the information
		locationID, _ := cmd.Flags().GetString("location-id")
		projectID, _ := cmd.Flags().GetString("project-id")
		agentName, _ := cmd.Flags().GetString("agent-name")
		force, _ := cmd.Flags().GetBool("force")
		entityTypeName := args[0]

		if err := entitytype.Delete(entityTypeName, locationID, projectID, agentName, force); err != nil {
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
	entitytypeCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().StringP("project-id", "p", "", "Dialogflow CX Project ID")
	deleteCmd.Flags().StringP("location-id", "l", "", "Dialogflow CX Location ID of the Project")
	deleteCmd.Flags().StringP("agent-name", "a", "", "Dialogflow CX Agent Name")
	deleteCmd.Flags().BoolP("force", "f", false, "Optional. Forces to delete the Entity type. NOTE: it will delete all any references to the entity type")
}
