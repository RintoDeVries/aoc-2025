package utils

import (
	"bufio"
	"os"
)

func FileToStringList(filepath string) ([]string, error) {
	result := []string{}

	file, err := os.Open(filepath)
	if err != nil {
		return result, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return result, err
	}

	return result, nil
}
