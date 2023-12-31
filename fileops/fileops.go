package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatfromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return 1000, errors.New("Failed to find balance file.")
	}

	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return 1000, errors.New("Failed to parse stored value.")
	}

	return value, nil
}

func WritefloatToFile(value float64, fileName string) {
	valuetext := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valuetext), 0644)
}
