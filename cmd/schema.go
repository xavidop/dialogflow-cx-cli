package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/invopop/jsonschema"
	"github.com/spf13/cobra"
	"github.com/xavidop/dialogflow-cx-cli/internal/types"
)

type schemaCmd struct {
	cmd    *cobra.Command
	output string
}

func newSchemaCmd() *schemaCmd {
	root := &schemaCmd{}
	cmd := &cobra.Command{
		Use:           "jsonschema",
		Aliases:       []string{"schema"},
		Short:         "outputs cxcli's JSON schema",
		SilenceUsage:  true,
		SilenceErrors: true,
		Args:          cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			suiteSchema := jsonschema.Reflect(&types.Suite{})
			suiteSchema.Definitions["Tests"] = jsonschema.Reflect(&types.Tests{})
			suiteSchema.Description = "cxcli suite definition file"

			testSchema := jsonschema.Reflect(&types.Test{})
			testSchema.Description = "cxcli test definition file"

			testBts, err := json.MarshalIndent(testSchema, "	", "	")
			if err != nil {
				return fmt.Errorf("failed to create test jsonschema: %w", err)
			}

			suiteBts, err := json.MarshalIndent(suiteSchema, "	", "	")
			if err != nil {
				return fmt.Errorf("failed to create suite jsonschema: %w", err)
			}
			if root.output == "-" {
				fmt.Println(string(suiteBts))
				fmt.Println(string(testBts))
				return nil
			}
			if err := os.MkdirAll(filepath.Dir(root.output), 0o755); err != nil {
				return fmt.Errorf("failed to write jsonschema file: %w", err)
			}

			if err := os.WriteFile(root.output+"/suite.json", suiteBts, 0o666); err != nil {
				return fmt.Errorf("failed to write jsonschema file: %w", err)
			}
			if err := os.WriteFile(root.output+"/test.json", testBts, 0o666); err != nil {
				return fmt.Errorf("failed to write jsonschema file: %w", err)
			}
			return nil
		},
	}

	cmd.Flags().StringVarP(&root.output, "output-folder", "f", "-", "Where to save the JSONSchema file")
	_ = cmd.Flags().SetAnnotation("output-file", cobra.BashCompFilenameExt, []string{"json"})

	root.cmd = cmd
	return root
}

func init() {
	rootCmd.AddCommand(newSchemaCmd().cmd)
}
