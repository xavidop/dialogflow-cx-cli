package tts

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/xavidop/dialogflow-cx-test-runner/cmd/cmdutils"
	"github.com/xavidop/dialogflow-cx-test-runner/internal/global"
	"github.com/xavidop/dialogflow-cx-test-runner/pkg/tts"
)

// synthesizeCmd represents the execute synthesis command
var synthesizeCmd = &cobra.Command{
	Use:     "synthesize [input]",
	Aliases: []string{"synth", "s"},
	Short:   "Executes a CICD pipeline for a specific environment",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Get the information
		locale, _ := cmd.Flags().GetString("locale")
		input := strings.Join(args, " ")
		output, _ := cmd.Flags().GetString("output")

		if err := tts.Synthesize(input, locale, output); err != nil {
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
	ttsCmd.AddCommand(synthesizeCmd)

	synthesizeCmd.Flags().StringP("locale", "l", "", "Input locale")
	synthesizeCmd.Flags().StringP("input", "i", "", "Input text to synthesize")
	synthesizeCmd.Flags().StringP("output", "o", "output.mp3", "Output file name")
}
