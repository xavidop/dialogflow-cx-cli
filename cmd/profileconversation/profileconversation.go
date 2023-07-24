package profileconversation

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/xavidop/dialogflow-cx-cli/cmd/cmdutils"
)

// profileconversationCmd represents the profile-conversation root command
var profileconversationCmd = &cobra.Command{
	Use:     "profile-conversation",
	Aliases: []string{"test-conversation", "tc", "conversation-tests"},
	Short:   "Actions on conversation testing",
	Run: func(cmd *cobra.Command, args []string) {
		if err := cmd.Help(); err != nil {
			os.Exit(1)
		}
		os.Exit(0)
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		cmdutils.PreRun(cmd.Name())
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
	},
}

func Register(rootCmd *cobra.Command) {
	rootCmd.AddCommand(profileconversationCmd)
}
