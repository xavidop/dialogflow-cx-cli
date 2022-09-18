package cmd

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/xavidop/dialogflow-cx-test-runner/pkg"
)

// import represents the init command
var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Imports a type",
	Run: func(cmd *cobra.Command, args []string) {
		typeName, _ := cmd.Flags().GetString("type-name")
		file, _ := cmd.Flags().GetString("file")
		header, _ := cmd.Flags().GetBool("header")
		if err := pkg.ImportType(typeName, file, header); err != nil {
			log.Error(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(importCmd)
	importCmd.Flags().StringP("type-name", "t", "type", "Type to create")
	importCmd.Flags().StringP("file", "f", "file.csv", "CSV to read")
	importCmd.Flags().BoolP("header", "e", false, "Specifies if the CSV contains headers or not")
}
