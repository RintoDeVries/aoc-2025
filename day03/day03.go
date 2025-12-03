package main

import (
	"aoc-2025/utils"
	"fmt"
	"slices"
	"strconv"
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

func maxJoltage(bank []int) int {
	max1, idx := returnMaxAndIndex(bank[:len(bank)-1]) // deliberatly cut off the last
	max2, _ := returnMaxAndIndex(bank[idx+1:])
	s := strconv.Itoa(max1) + strconv.Itoa(max2)
	maxJoltage, _ := strconv.Atoi(s)
	return maxJoltage
}

func part1(input [][]int) int {
	part1 := 0
	for _, bank := range input {
		part1 += maxJoltage(bank)
	}
	return part1
}

func main() {
	input, err := utils.FileToStringList("input.txt")
	if err != nil {
		panic(err)
	}

	parsedInput := parseInput(input)

	part1 := part1(parsedInput)
	fmt.Println("The result for part 1:", part1)
}
