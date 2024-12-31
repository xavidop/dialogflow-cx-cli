package agent

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/xavidop/dialogflow-cx-cli/cmd/cmdutils"
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	"github.com/xavidop/dialogflow-cx-cli/internal/utils"
	"github.com/xavidop/dialogflow-cx-cli/pkg/agent"
)

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:     "export [agent-name]",
	Aliases: []string{"export", "e", "ex"},
	Short:   "Executes an export for a specific agent",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Get the information
		locationID, _ := cmd.Flags().GetString("location-id")
		projectID, _ := cmd.Flags().GetString("project-id")
		output, _ := cmd.Flags().GetString("output-file")
		exportFormat, _ := cmd.Flags().GetString("export-format")
		if err := utils.ValidateExportFormat(exportFormat); err != nil {
			global.Log.Errorf("%s", err.Error())
			os.Exit(1)
		}
		if err := utils.ValidateExportOutputFileAndFormatCorrelation(output, exportFormat); err != nil {
			global.Log.Errorf("%s", err.Error())
			os.Exit(1)
		}
		agentName := args[0]

		if err := agent.Export(locationID, projectID, agentName, output, exportFormat); err != nil {
			global.Log.Errorf("%s", err.Error())
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
	agentCmd.AddCommand(exportCmd)

	exportCmd.Flags().StringP("project-id", "p", "", "Dialogflow CX Project ID (required)")
	if err := exportCmd.MarkFlagRequired("project-id"); err != nil {
		global.Log.Errorf("%s", err.Error())
		os.Exit(1)
	}
	exportCmd.Flags().StringP("location-id", "l", "", "Dialogflow CX Location ID of the Project (required)")
	if err := exportCmd.MarkFlagRequired("location-id"); err != nil {
		global.Log.Errorf("%s", err.Error())
		os.Exit(1)
	}
	exportCmd.Flags().StringP("output-file", "f", "agent.blob", "Output file name. Default: agent.blob (optional)")
	exportCmd.Flags().StringP("export-format", "t", "blob", "Export format type: json or blob. blob by default. Default: blob (optional)")

}
