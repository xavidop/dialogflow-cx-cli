package utils

import (
	"errors"
	"os"
	"path"
	"path/filepath"
)

func CheckIfFileExists(file string) error {

	if _, err := os.Stat(file); errors.Is(err, os.ErrNotExist) {
		return err
	}

	return nil
}

func WriteFile(b []byte, file string) error {
	f, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	err = os.WriteFile(file, b, 0644)
	if err != nil {
		return err
	}

	return nil
}

func GetRelativeFilePathFromParentFile(parentFile string, file string) string {
	base := filepath.Dir(parentFile)

	if !filepath.IsAbs(file) {
		return path.Join(base, file)
	} else {
		return file
	}
}
