package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/invopop/jsonschema"
	"github.com/spf13/cobra"
	conersationypes "github.com/xavidop/dialogflow-cx-cli/internal/types/profileconversation"
	nlutypes "github.com/xavidop/dialogflow-cx-cli/internal/types/profilenlu"
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

			if err := generateProfilenluSchema(root.output); err != nil {
				return err
			}

			if err := generateProfileconversationSchema(root.output); err != nil {
				return err
			}

			return nil
		},
	}

	cmd.Flags().StringVarP(&root.output, "output-folder", "f", "-", "Where to save the JSONSchema file")
	_ = cmd.Flags().SetAnnotation("output-file", cobra.BashCompFilenameExt, []string{"json"})

	root.cmd = cmd
	return root
}

func generateProfilenluSchema(output string) error {
	suiteSchema := jsonschema.Reflect(&nlutypes.Suite{})
	suiteSchema.Definitions["Tests"] = jsonschema.Reflect(&nlutypes.Tests{})
	suiteSchema.Description = "cxcli NLU Profiler suite definition file"

	testSchema := jsonschema.Reflect(&nlutypes.Test{})
	testSchema.Description = "cxcli NLU Profiler test definition file"

	testBts, err := json.MarshalIndent(testSchema, "	", "	")
	if err != nil {
		return fmt.Errorf("failed to create test jsonschema: %w", err)
	}

	suiteBts, err := json.MarshalIndent(suiteSchema, "	", "	")
	if err != nil {
		return fmt.Errorf("failed to create suite jsonschema: %w", err)
	}
	if output == "-" {
		fmt.Println(string(suiteBts))
		fmt.Println(string(testBts))
		return nil
	}
	if err := os.MkdirAll(filepath.Dir(output), 0o755); err != nil {
		return fmt.Errorf("failed to write jsonschema file: %w", err)
	}

	if err := os.WriteFile(output+"/nlusuite.json", suiteBts, 0o666); err != nil {
		return fmt.Errorf("failed to write jsonschema file: %w", err)
	}
	if err := os.WriteFile(output+"/nlutest.json", testBts, 0o666); err != nil {
		return fmt.Errorf("failed to write jsonschema file: %w", err)
	}
	return nil
}

func generateProfileconversationSchema(output string) error {
	suiteSchema := jsonschema.Reflect(&conersationypes.Suite{})
	suiteSchema.Definitions["Tests"] = jsonschema.Reflect(&conersationypes.Tests{})
	suiteSchema.Description = "cxcli Conversation Profiler suite definition file"

	testSchema := jsonschema.Reflect(&conersationypes.Test{})
	testSchema.Description = "cxcli Conversation Profiler test definition file"

	testBts, err := json.MarshalIndent(testSchema, "	", "	")
	if err != nil {
		return fmt.Errorf("failed to create test jsonschema: %w", err)
	}

	suiteBts, err := json.MarshalIndent(suiteSchema, "	", "	")
	if err != nil {
		return fmt.Errorf("failed to create suite jsonschema: %w", err)
	}
	if output == "-" {
		fmt.Println(string(suiteBts))
		fmt.Println(string(testBts))
		return nil
	}
	if err := os.MkdirAll(filepath.Dir(output), 0o755); err != nil {
		return fmt.Errorf("failed to write jsonschema file: %w", err)
	}

	if err := os.WriteFile(output+"/conversationsuite.json", suiteBts, 0o666); err != nil {
		return fmt.Errorf("failed to write jsonschema file: %w", err)
	}
	if err := os.WriteFile(output+"/conversationtest.json", testBts, 0o666); err != nil {
		return fmt.Errorf("failed to write jsonschema file: %w", err)
	}
	return nil
}

func init() {
	rootCmd.AddCommand(newSchemaCmd().cmd)
}
