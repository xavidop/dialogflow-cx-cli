package types

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Test struct {
	Name        string  `yaml:"name"`
	Description string  `yaml:"description"`
	LocaleID    string  `yaml:"localeId"`
	Checks      []Check `yaml:"checks"`
}

func NewTest(file string) (*Test, error) {
	test := &Test{}

	yfile, err := os.ReadFile(file)
	if err != nil {
		return test, err
	}

	err = yaml.Unmarshal([]byte(yfile), &test)
	if err != nil {
		return test, err
	}
	return test, nil
}
