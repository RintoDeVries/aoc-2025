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

func isValidId(id int) bool {
	s_id := strconv.Itoa(id)
	if len(s_id)%2 != 0 {
		return true
	}

	for idx := 1; idx <= len(s_id)/2; idx++ {
		if s_id[idx-1] != s_id[len(s_id)/2+idx-1] {
			return true
		}
	}
	return false
}

func part1(input []id_range) int {
	result := 0
	for _, id_range := range input {
		for id := id_range.min; id <= id_range.max; id++ {
			if !isValidId(id) {
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
