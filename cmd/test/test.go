package env

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/xavidop/dialogflow-cx-test-runner/cmd/cmdutils"
	"github.com/xavidop/dialogflow-cx-test-runner/internal/global"
)

// testCmd represents the test root command
var testCmd = &cobra.Command{
	Use:     "test",
	Aliases: []string{"test", "t", "tests"},
	Short:   "Actions on testing",
	Run: func(cmd *cobra.Command, args []string) {
		global.Log.Infof("test")
		os.Exit(0)
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		cmdutils.PreRun(cmd.Name())
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
	},
}

func Register(rootCmd *cobra.Command) {
	rootCmd.AddCommand(testCmd)
}
