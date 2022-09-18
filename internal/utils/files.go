package utils

import (
	"encoding/csv"
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

func ReadCsvFile(filePath string) (*csv.Reader, *os.File, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, nil, err
	}

	csvReader := csv.NewReader(f)

	return csvReader, f, nil
}

func WriteFile(b []byte, file string) error {
	f, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(file, b, 0644)
	if err != nil {
		return err
	}

	return nil
}
