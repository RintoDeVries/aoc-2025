package main

import (
	"aoc-2025/utils"
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Range struct {
	Min int
	Max int
}

func parseRangeLine(line string) Range {
	minStr, maxStr, ok := strings.Cut(line, "-")
	if !ok {
		panic("invalid range format")
	}

	min, err := strconv.Atoi(minStr)
	if err != nil {
		panic("invalid min value")
	}

	max, err := strconv.Atoi(maxStr)
	if err != nil {
		panic("invalid max value")
	}

	return Range{
		Min: min,
		Max: max,
	}
}

func parseIdLine(line string) int {
	num, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}
	return num
}

func parseInput(input []string) ([]Range, []int) {
	var id_ranges []Range
	var ids []int
	id_rangesDone := false

	for _, line := range input {
		if line == "" {
			id_rangesDone = true
			continue
		}

		if !id_rangesDone {
			id_ranges = append(id_ranges, parseRangeLine(line))
		} else {
			ids = append(ids, parseIdLine(line))
		}
	}

	return id_ranges, ids
}

func part1(id_ranges []Range, ids []int) int {
	result := 0
	for _, id := range ids {
		for _, id_rng := range id_ranges {
			if id >= id_rng.Min && id <= id_rng.Max {
				result++
				break
			}
		}
	}
	return result
}

func sortIdRange(a, b Range) int {
	return cmp.Compare(a.Min, b.Min)
}

func part2(id_ranges []Range) int {
	slices.SortFunc(id_ranges, sortIdRange)
	currMax := id_ranges[0].Max
	currMin := id_ranges[0].Min

	sum := 0
	for _, r := range id_ranges[1:] {
		// no overlap
		if r.Min > currMax {
			sum += currMax - currMin + 1
			currMin = r.Min
			currMax = r.Max
		} else { //overlap
			if r.Max > currMax {
				currMax = r.Max
			}
		}
	}
	sum += currMax - currMin + 1
	return sum
}

func main() {
	input, err := utils.FileToStringList("input.txt")
	if err != nil {
		panic(err)
	}

	id_ranges, ids := parseInput(input)

	part1 := part1(id_ranges, ids)
	part2 := part2(id_ranges)

	fmt.Println("The result for part 1:", part1)
	fmt.Println("The result for part 2:", part2)

}
