package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/xavidop/dialogflow-cx-test-runner/internal/global"
)

// VersionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Get dialogflow-cx-test-runner version",
	Run: func(cmd *cobra.Command, args []string) {
		log.Infof("Current version: %s", global.VersionString)
		// Not check in development
		if global.VersionString != "" {
			// Don't ignore errors
			// TODO check update
		}

	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
