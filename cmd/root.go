package cmd

import (
	"os"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/xavidop/dialogflow-cx-test-runner/internal/global"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dialogflow-cx-test-runner",
	Short: "Google Action Type Importer",
	Long: `Welcome to dialogflow-cx-test-runner!
	
This utility provides you with an easy way to create custom types 
for your Google Actions projects importing those values from files. 

You can find the documentation at https://github.com/xavidop/dialogflow-cx-test-runner.

Please file all bug reports on Github at https://github.com/xavidop/dialogflow-cx-test-runner/issues.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if global.VersionString == "" {
			global.VersionString = "development"
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Error(errors.Errorf("%s", err))
		os.Exit(1)
	}
}

func init() {
	// Add the subcommands
	rootCmd.PersistentFlags().BoolVarP(&global.Verbose, "verbose", "v", false, "verbose error output (with stack trace)")

}
