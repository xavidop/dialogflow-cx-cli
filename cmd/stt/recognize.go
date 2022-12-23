package stt

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/xavidop/dialogflow-cx-cli/cmd/cmdutils"
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	"github.com/xavidop/dialogflow-cx-cli/pkg/stt"
)

// recognizeCmd represents the execute recognize command
var recognizeCmd = &cobra.Command{
	Use:     "recognize [input]",
	Aliases: []string{"rec", "recognise"},
	Short:   "Transforms audio into text",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Get the information
		locale, _ := cmd.Flags().GetString("locale")
		input := args[0]

		if err := stt.Recognize(input, locale); err != nil {
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
	sttCmd.AddCommand(recognizeCmd)

	recognizeCmd.Flags().StringP("locale", "l", "", "Input locale")
}
