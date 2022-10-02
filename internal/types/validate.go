package types

import (
	"errors"
)

type Validate struct {
	Intent     string      `yaml:"intent"`
	Parameters []Parameter `yaml:"parameters"`
}

func FindParameterByName(validate Validate, name string) (Parameter, error) {
	for _, param := range validate.Parameters {
		if param.Parameter == name {
			return param, nil
		}
	}
	return Parameter{}, errors.New("parameter not found")
}

func RemoveParameterByName(parameters []Parameter, name string) []Parameter {
	for i, param := range parameters {
		if param.Parameter == name {
			return remove(parameters, i)
		}
	}
	return parameters
}

func remove(slice []Parameter, s int) []Parameter {
	return append(slice[:s], slice[s+1:]...)
}
