package datastore

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/xavidop/dialogflow-cx-cli/cmd/cmdutils"
	cmdducment "github.com/xavidop/dialogflow-cx-cli/cmd/datastore/document"
)

// datastoreCmd represents the datastore root command
var datastoreCmd = &cobra.Command{
	Use:     "datastore",
	Aliases: []string{"ds"},
	Short:   "Actions on datastore commands",
	Run: func(cmd *cobra.Command, args []string) {
		if err := cmd.Help(); err != nil {
			os.Exit(1)
		}
		os.Exit(0)
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		cmdutils.PreRun(cmd.Name())
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
	},
}

func Register(rootCmd *cobra.Command) {
	rootCmd.AddCommand(datastoreCmd)
}

func init() {
	cmdducment.Register(datastoreCmd)
}