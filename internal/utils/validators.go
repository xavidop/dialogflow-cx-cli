package utils

import (
	"errors"
	"path/filepath"
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
