package datastore

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/xavidop/dialogflow-cx-cli/cmd/cmdutils"
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	"github.com/xavidop/dialogflow-cx-cli/pkg/datastore"
)

// deleteCmd represents the search datastore set command
var deleteCmd = &cobra.Command{
	Use:     "delete [name]",
	Aliases: []string{"deletes", "d", "del"},
	Short:   "delete a datastore",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Get the information
		locationID, _ := cmd.Flags().GetString("location-id")
		projectID, _ := cmd.Flags().GetString("project-id")
		name := args[0]

		if err := datastore.Delete(name, locationID, projectID); err != nil {
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
	datastoreCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().StringP("project-id", "p", "", "Data Store Project ID (required)")
	if err := deleteCmd.MarkFlagRequired("project-id"); err != nil {
		global.Log.Errorf(err.Error())
		os.Exit(1)
	}
	deleteCmd.Flags().StringP("location-id", "l", "", "Data Store Location ID of the Project (required)")
	if err := deleteCmd.MarkFlagRequired("location-id"); err != nil {
		global.Log.Errorf(err.Error())
		os.Exit(1)
	}

}
