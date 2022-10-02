package cmd

import (
	"github.com/spf13/cobra"
	"github.com/xavidop/dialogflow-cx-test-runner/cmd/cmdutils"
)

// VersionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Get cxtester version",
	Run: func(cmd *cobra.Command, args []string) {
		// Not check in development
		cmdutils.CheckUpdate(true)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
