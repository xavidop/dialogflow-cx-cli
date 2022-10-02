package cmd

import (
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	cmdcicd "github.com/xavidop/dialogflow-cx-test-runner/cmd/cicd"
	"github.com/xavidop/dialogflow-cx-test-runner/cmd/cmdutils"
	cmdtest "github.com/xavidop/dialogflow-cx-test-runner/cmd/test"
	"github.com/xavidop/dialogflow-cx-test-runner/internal/global"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cxtester",
	Short: "Google Action Type Importer",
	Long: `Welcome to cxtester!
	
This utility provides you with an easy way to run tests 
for your Dialogflow CX agents. 

You can find the documentation at https://github.com/xavidop/dialogflow-cx-test-runner.

Please file all bug reports on Github at https://github.com/xavidop/dialogflow-cx-test-runner/issues.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
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

	// Add the subcommands
	cmdtest.Register(rootCmd)
	cmdcicd.Register(rootCmd)

	// Add the subcommands
	rootCmd.PersistentFlags().BoolVarP(&global.Verbose, "verbose", "v", false, "verbose error output (with stack trace)")
	rootCmd.PersistentFlags().StringVarP(&global.Credentials, "credentials", "c", "", "verbose error output (with stack trace)")
	rootCmd.PersistentFlags().BoolVarP(&global.SkipUpdate, "skip-update-check", "u", false, "Skip the check for updates check run before every command")
	rootCmd.PersistentFlags().StringVarP(&global.Output, "output", "o", "text", "Output Format")

}
