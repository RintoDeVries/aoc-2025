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

func notAllEqualParts(id int, parts int) bool {
	s := strconv.Itoa(id)
	length := len(s)
	chunk := length / parts

	slices := make([]string, 0, parts)
	for i := 0; i < parts; i++ {
		start := i * chunk
		end := start + chunk

		if i == parts-1 {
			end = length
		}

		slices = append(slices, s[start:end])
	}

	for i := 1; i < len(slices); i++ {
		if slices[i] != slices[0] {
			return true
		}
	}
	return false
}

func isValidIdPt2(id int) bool {
	return notAllEqualParts(id, 2) && notAllEqualParts(id, 3) && notAllEqualParts(id, 5) && notAllEqualParts(id, 7)
}

func part2(input []id_range) int {
	result := 0
	for _, id_range := range input {
		for id := id_range.min; id <= id_range.max; id++ {
			if !isValidIdPt2(id) {
				result += id
			}
		}
	}
	return result
}

func main() {
	parsedInput, err := parseInput("input.txt")
	if err != nil {
		panic(err)
	}

	part1 := part1(parsedInput)
	part2 := part2(parsedInput)
	fmt.Println("The result for part 1:", part1)
	fmt.Println("The result for part 2:", part2)
}
