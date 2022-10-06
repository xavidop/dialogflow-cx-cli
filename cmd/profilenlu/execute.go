package env

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/xavidop/dialogflow-cx-test-runner/cmd/cmdutils"
	"github.com/xavidop/dialogflow-cx-test-runner/internal/global"
	"github.com/xavidop/dialogflow-cx-test-runner/pkg/profilenlu"
)

// executeCmd represents the execute command
var executeCmd = &cobra.Command{
	Use:     "execute",
	Aliases: []string{"execute", "e", "exe", "exec"},
	Short:   "Execute a suite",
	Run: func(cmd *cobra.Command, args []string) {
		suite, _ := cmd.Flags().GetString("suite")

		if err := profilenlu.ExecuteSuite(suite); err != nil {
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
	profilenluCmd.AddCommand(executeCmd)

	executeCmd.Flags().StringP("suite", "s", "suite.yaml", "Suite yaml file")
}
