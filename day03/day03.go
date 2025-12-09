package main

import (
	"aoc-2025/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func parseInput(input []string) [][]int {
	result := [][]int{}
	for _, string_bank := range input {
		bank := []int{}
		for _, char := range string_bank {
			bank = append(bank, int(char-'0'))
		}
		result = append(result, bank)
	}
	return result
}

func returnMaxAndIndex(slice []int) (int, int) {
	max := slices.Max(slice)
	for idx, val := range slice {
		if val == max {
			return max, idx
		}
	}
	panic("this shouldn't happen")
}

// Builds the largest number of length numDigits by greedily picking the max digit
// from the remaining slice of bank while preserving order.
func maxJoltage(bank []int, numDigits int) int {
	var b strings.Builder
	highestIndex := -1
	b.Grow(numDigits)
	for i := range numDigits {
		max, newIndex := returnMaxAndIndex(bank[highestIndex+1 : len(bank)-numDigits+1+i])
		highestIndex = newIndex + highestIndex + 1
		b.WriteRune(rune('0' + max))
	}
	maxJoltage, err := strconv.Atoi(b.String())
	if err != nil {
		panic(err)
	}
	return maxJoltage
}

func part1(input [][]int) int {
	part1 := 0
	numDigits := 2
	for _, bank := range input {
		part1 += maxJoltage(bank, numDigits)
	}
	return part1
}

func part2(input [][]int) int {
	part2 := 0
	numDigits := 12
	for _, bank := range input {
		part2 += maxJoltage(bank, numDigits)
	}
	return part2
}

func main() {
	input, err := utils.FileToStringList("input.txt")
	if err != nil {
		panic(err)
	}

	parsedInput := parseInput(input)

	part1 := part1(parsedInput)
	part2 := part2(parsedInput)

	fmt.Println("The result for part 1:", part1)
	fmt.Println("The result for part 2:", part2)

}
