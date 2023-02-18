package versioning

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/xavidop/dialogflow-cx-cli/cmd/cmdutils"
)

// versioningCmd represents the versioning root command
var versioningCmd = &cobra.Command{
	Use:     "versioning",
	Aliases: []string{"ver"},
	Short:   "Actions on flow versions commands",
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
	rootCmd.AddCommand(versioningCmd)
}
