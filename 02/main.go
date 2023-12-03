package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fredrikln/advent-of-code-2023/parsing"
)

func main() {
	input := parsing.ReadRowsToStringSlice("input.txt")

	fmt.Println("Part 1:", Part1(input))

	fmt.Println("Part 2:", Part2(input))
}

func Part1(lines []string) int {
	maxRed := 12
	maxGreen := 13
	maxBlue := 14

	var total int

	for _, line := range lines {
		parts := strings.Split(line, ": ")

		title := parts[0]

		roundNo, err := strconv.Atoi(strings.Split(title, " ")[1])
		if err != nil {
			panic("Invalid roundNo")
		}

		subsets := strings.Split(parts[1], "; ")

		possible := true

		for _, subset := range subsets {
			if !possible {
				break
			}

			cubeInfos := strings.Split(subset, ", ")

			for _, cubeInfo := range cubeInfos {
				cubeInfoParts := strings.Split(cubeInfo, " ")

				color := cubeInfoParts[1]

				count, err := strconv.Atoi(cubeInfoParts[0])
				if err != nil {
					panic("invalid cube count")
				}

				if color == "red" && count > maxRed {
					possible = false
				} else if color == "green" && count > maxGreen {
					possible = false
				} else if color == "blue" && count > maxBlue {
					possible = false
				}

				if !possible {
					break
				}
			}

		}

		if possible {
			total += roundNo
		}
	}

	return total
}

func Part2(lines []string) int {
	var totalPower int

	for _, line := range lines {
		var red, green, blue int

		parts := strings.Split(line, ": ")

		// title := parts[0]

		// roundNo, err := strconv.Atoi(strings.Split(title, " ")[1])
		// if err != nil {
		// 	panic("Invalid roundNo")
		// }

		subsets := strings.Split(parts[1], "; ")

		for _, subset := range subsets {

			cubeInfos := strings.Split(subset, ", ")

			for _, cubeInfo := range cubeInfos {
				cubeInfoParts := strings.Split(cubeInfo, " ")

				color := cubeInfoParts[1]

				count, err := strconv.Atoi(cubeInfoParts[0])
				if err != nil {
					panic("invalid cube count")
				}

				if color == "red" && count > red {
					red = count
				} else if color == "green" && count > green {
					green = count
				} else if color == "blue" && count > blue {
					blue = count
				}
			}
		}

		// fmt.Println(roundNo, "red", red, "green", green, "blue", blue)
		power := red * green * blue

		totalPower += power
	}

	return totalPower
}
