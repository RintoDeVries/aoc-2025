package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type id_range struct {
	min int
	max int
}

func parseInput(filename string) ([]id_range, error) {
	result := []id_range{}

	data, err := os.ReadFile(filename)
	if err != nil {
		return result, err
	}

	re := regexp.MustCompile(`(\d+)-(\d+)`)
	matches := re.FindAllStringSubmatch(string(data), -1)

	for _, nums := range matches {
		min, _ := strconv.Atoi(nums[1])
		max, _ := strconv.Atoi(nums[2])

		result = append(result, id_range{
			min: min,
			max: max,
		})
	}
	return result, nil
}

func isValidIdPt1(id int) bool {
	s_id := strconv.Itoa(id)
	half_len := len(s_id) / 2
	return s_id[:half_len] != s_id[half_len:]
}

func part1(input []id_range) int {
	result := 0
	for _, id_range := range input {
		for id := id_range.min; id <= id_range.max; id++ {
			if !isValidIdPt1(id) {
				result += id
			}
		}
	}
	return result
}

func main() {
	input, err := parseInput("input.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(part1(input))
}
