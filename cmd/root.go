package cmd

import (
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/tmc/langchaingo/llms/vertexai"
	cobracompletefig "github.com/withfig/autocomplete-tools/integrations/cobra"
	cmdagent "github.com/xavidop/dialogflow-cx-cli/cmd/agent"
	"github.com/xavidop/dialogflow-cx-cli/cmd/cmdutils"
	cmdentitytype "github.com/xavidop/dialogflow-cx-cli/cmd/entitytype"
	environmentCmd "github.com/xavidop/dialogflow-cx-cli/cmd/environment"
	cmdflow "github.com/xavidop/dialogflow-cx-cli/cmd/flow"
	cmdflowversion "github.com/xavidop/dialogflow-cx-cli/cmd/flowversion"
	cmdintent "github.com/xavidop/dialogflow-cx-cli/cmd/intent"
	cmdprofileconversation "github.com/xavidop/dialogflow-cx-cli/cmd/profileconversation"
	cmdprofilenlu "github.com/xavidop/dialogflow-cx-cli/cmd/profilenlu"
	cmdstt "github.com/xavidop/dialogflow-cx-cli/cmd/stt"
	cmdtts "github.com/xavidop/dialogflow-cx-cli/cmd/tts"
	cmdwebhook "github.com/xavidop/dialogflow-cx-cli/cmd/webhook"
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cxcli",
	Short: "Dialogflow CX CLI",
	Long: `Welcome to cxcli!

This utility provides you with an easy way to interact
with your Dialogflow CX agents.

You can find the documentation at https://github.com/xavidop/dialogflow-cx-cli.

Please file all bug reports on GitHub at https://github.com/xavidop/dialogflow-cx-cli/issues.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			if err := cmd.Help(); err != nil {
				os.Exit(1)
			}
			os.Exit(0)
		}
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		cmdutils.PreRun(cmd.Name())
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		global.Log.Error(errors.Errorf("%s", err))
		os.Exit(1)
	}
}

func init() {

	vertexai.NewChat()

	// Add the subcommands
	cmdprofileconversation.Register(rootCmd)
	cmdprofilenlu.Register(rootCmd)
	environmentCmd.Register(rootCmd)
	cmdtts.Register(rootCmd)
	cmdagent.Register(rootCmd)
	cmdintent.Register(rootCmd)
	cmdentitytype.Register(rootCmd)
	cmdstt.Register(rootCmd)
	cmdwebhook.Register(rootCmd)
	cmdflow.Register(rootCmd)
	cmdflowversion.Register(rootCmd)
	rootCmd.AddCommand(cobracompletefig.CreateCompletionSpecCommand())

	// Add the subcommands
	rootCmd.PersistentFlags().BoolVarP(&global.Verbose, "verbose", "v", false, "verbose error output (with stack trace) (optional)")
	rootCmd.PersistentFlags().StringVarP(&global.Credentials, "credentials", "c", "", "Google Cloud credentials JSON file path (optional)")
	rootCmd.PersistentFlags().BoolVarP(&global.SkipUpdate, "skip-update-check", "u", false, "Skip the check for updates check run before every command (optional)")
	rootCmd.PersistentFlags().StringVarP(&global.Output, "output-format", "o", "text", "Output Format. Options: text, json. Default: text (optional)")

}
