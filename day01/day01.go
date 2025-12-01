package main

import (
	"aoc-2025/utils"
	"fmt"
	"log"
	"strconv"
)

const startDialPosition = 50

type Direction int

const (
	Left  Direction = -1
	Right Direction = 1
)

type rotation struct {
	direction Direction
	amount    int
}

func parseInput(input []string) []rotation {
	result := []rotation{}

	for _, v := range input {
		num, err := strconv.Atoi(v[1:])
		if err != nil {
			log.Fatal(err)
		}

		var dir Direction
		switch v[0] {
		case 'R':
			dir = Right
		case 'L':
			dir = Left
		default:
			log.Fatal("unexpected direction: ", string(v[0]))
		}

		result = append(result, rotation{
			direction: dir,
			amount:    num,
		})
	}

	return result
}

func part1(input []rotation) int {
	dialPosition := startDialPosition
	numberOfTimesZero := 0
	for _, rotation := range input {
		dialPosition = (((dialPosition + rotation.amount*int(rotation.direction)) % 100) + 100) % 100
		if dialPosition == 0 {
			numberOfTimesZero += 1
		}
	}
	return numberOfTimesZero
}

func part2(input []rotation) int {
	dialPosition := startDialPosition
	numberOfTimesZero := 0
	for _, rotation := range input {
		for i := 0; i < rotation.amount; i++ {
			dialPosition = (((dialPosition + int(rotation.direction)) % 100) + 100) % 100
			if dialPosition == 0 {
				numberOfTimesZero += 1
			}
		}
	}
	return numberOfTimesZero
}

func main() {
	input, err := utils.FileToStringList("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	parsedInput := parseInput(input)

	part1 := part1(parsedInput)
	part2 := part2(parsedInput)
	fmt.Println("The result for part 1:", part1)
	fmt.Println("The result for part 2:", part2)

}
