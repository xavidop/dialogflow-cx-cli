package types

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Suite struct {
	Name        string  `yaml:"name"`
	Description string  `yaml:"description"`
	ProjectID   string  `yaml:"projectId"`
	LocationID  string  `yaml:"locationId"`
	AgentName   string  `yaml:"agentName"`
	Tests       []Tests `yaml:"tests"`
}

func NewSuite(file string) (*Suite, error) {
	suite := &Suite{}

	yfile, err := ioutil.ReadFile(file)
	if err != nil {
		return suite, err
	}

	err = yaml.Unmarshal([]byte(yfile), &suite)
	if err != nil {
		return suite, err
	}
	return suite, nil
}
