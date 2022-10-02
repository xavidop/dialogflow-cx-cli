package utils

import (
	"errors"
	"io/ioutil"
	"os"
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
	err = ioutil.WriteFile(file, b, 0644)
	if err != nil {
		return err
	}

	return nil
}
