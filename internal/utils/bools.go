package utils

import "strconv"

func ParseBool(s string) (bool, error) {
	boolValue, err := strconv.ParseBool(s)
	if err != nil {
		return false, err
	}
	return boolValue, nil
}
