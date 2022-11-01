package tts

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/xavidop/dialogflow-cx-test-runner/cmd/cmdutils"
)

// ttsCmd represents the tts root command
var ttsCmd = &cobra.Command{
	Use:     "tts",
	Aliases: []string{"tts", "text-to-speech", "cd"},
	Short:   "Actions on text-to-speech commands",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		cmdutils.PreRun(cmd.Name())
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
	},
}

func Register(rootCmd *cobra.Command) {
	rootCmd.AddCommand(ttsCmd)
}
