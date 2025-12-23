package utils

import (
	"bufio"
	"os"
)

func FileToStringList(filepath string) ([]string, error) {
	result := []string{}

	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result, scanner.Err()
}
