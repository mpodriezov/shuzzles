package utils

import "strconv"

func ParseBitField(input []string) (byte, error) {
	var result byte
	for _, num := range input {
		roleNum, err := strconv.ParseUint(num, 10, 8)
		if err != nil {
			return 0, err
		}
		result |= byte(roleNum)
	}
	return result, nil
}
