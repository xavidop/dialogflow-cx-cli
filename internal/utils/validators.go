package utils

import (
	"errors"
	"path/filepath"
)

func ValidateFileType(file string) error {
	if err := CheckIfFileExists(file); err != nil {
		return err
	}

	if err := ValidFileTypes(file); err != nil {
		return err
	}

	return nil
}

func ValidFileTypes(file string) error {

	fileExtension := filepath.Ext(file)

	if fileExtension != ".csv" {
		return errors.New("File extension ins't equal to .csv")
	}

	return nil
}
