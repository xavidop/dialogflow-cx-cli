package utils

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	"golang.org/x/exp/slices"
)

func ValidateFileType(file string, fileType string) error {
	if err := CheckIfFileExists(file); err != nil {
		return err
	}

	if err := ValidFileTypes(file, fileType); err != nil {
		return err
	}

	return nil
}

func ValidFileTypes(file string, fileType string) error {

	fileExtension := filepath.Ext(file)

	if fileExtension != fileType {
		return errors.New("file extension ins't equal to " + fileType)
	}

	return nil
}

func ValidateExportFormat(exportFormat string) error {
	if !slices.Contains(global.AgentExportFormats, exportFormat) {
		return fmt.Errorf("invalid export format. These are the valid ones: %v", global.AgentExportFormats)
	}
	return nil
}

func ValidateAgentFileType(exportOutputFile string) error {
	if !slices.Contains(global.AgentExportOutputFile, filepath.Ext(exportOutputFile)) {
		return fmt.Errorf("invalid file type. These are the valid ones: %v", global.AgentExportOutputFile)
	}
	return nil
}

func ValidateExportOutputFileAndFormatCorrelation(exportOutputFile, exportFormat string) error {
	fileExtension := filepath.Ext(exportOutputFile)
	if err := ValidateAgentFileType(exportOutputFile); err != nil {
		return err
	}
	switch exportFormat {
	case "json":
		if fileExtension != ".zip" {
			return errors.New("the file has to be a .zip file")
		}
	case "blob":
		if fileExtension != ".blob" {
			return errors.New("the file has to be a .blob file")
		}
	}
	return nil
}
