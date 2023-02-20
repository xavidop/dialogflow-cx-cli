package utils

import "strconv"

func ParseFloat(s string) (float32, error) {
	value, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return 0.0, err
	}
	return float32(value), nil
}
