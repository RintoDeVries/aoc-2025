package main

import (
	"aoc-2025/utils"
	"fmt"
)

func parseInput(input []string) [][]int {
	result := [][]int{}
	for _, line := range input {
		row := make([]int, len(line))
		for idx, char := range line {
			if char == '@' {
				row[idx] = 1
			}
		}
		result = append(result, row)
	}
	return result
}

func countAdjacent(grid [][]int, r, c int) int {
	rows := len(grid)
	cols := len(grid[0])
	count := 0

	dirs := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	for _, d := range dirs {
		nr := r + d[0]
		nc := c + d[1]

		if nr >= 0 && nr < rows && nc >= 0 && nc < cols {
			if grid[nr][nc] == 1 {
				count++
			}
		}
	}

	return count
}

func part1(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	result := 0

	for r := range rows {
		for c := range cols {
			if grid[r][c] == 1 && countAdjacent(grid, r, c) < 4 {
				result++
			}
		}
	}
	return result
}

func part2(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	result := 0
	prevResult := 0

	for {
		for r := range rows {
			for c := range cols {
				if grid[r][c] == 1 && countAdjacent(grid, r, c) < 4 {
					result++
					grid[r][c] = 0
				}
			}
		}
		if result == prevResult {
			break
		}
		prevResult = result
	}
	return result
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
