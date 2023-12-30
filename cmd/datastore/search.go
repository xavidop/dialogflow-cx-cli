package datastore

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/xavidop/dialogflow-cx-cli/cmd/cmdutils"
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	"github.com/xavidop/dialogflow-cx-cli/pkg/datastore"
)

// searchCmd represents the search datastore set command
var searchCmd = &cobra.Command{
	Use:     "search [name]",
	Aliases: []string{"searches", "s"},
	Short:   "search in a datastore",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Get the information
		locationID, _ := cmd.Flags().GetString("location-id")
		projectID, _ := cmd.Flags().GetString("project-id")
		query, _ := cmd.Flags().GetString("query")
		name := args[0]

		if err := datastore.Search(name, locationID, projectID, query); err != nil {
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
	datastoreCmd.AddCommand(searchCmd)

	searchCmd.Flags().StringP("query", "r", "", "Query to search (required)")
	if err := searchCmd.MarkFlagRequired("query"); err != nil {
		global.Log.Errorf(err.Error())
		os.Exit(1)
	}
	searchCmd.Flags().StringP("project-id", "p", "", "Data Store Project ID (required)")
	if err := searchCmd.MarkFlagRequired("project-id"); err != nil {
		global.Log.Errorf(err.Error())
		os.Exit(1)
	}
	searchCmd.Flags().StringP("location-id", "l", "", "Data Store Location ID of the Project (required)")
	if err := searchCmd.MarkFlagRequired("location-id"); err != nil {
		global.Log.Errorf(err.Error())
		os.Exit(1)
	}

}
