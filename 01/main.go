package main

import (
	"fmt"
	"strconv"
	"strings"

	. "github.com/fredrikln/advent-of-code-2023/parsing"
)

func main() {
	input := ReadRowsToStringSlice("input.txt")

	Part1(input)
	Part2(input)
}

func Part1(lines []string) {
	var total int

	for _, line := range lines {
		digits := getDigits(line)
		total += getNumber(digits)
	}

	fmt.Println("Part 1:", total)
}

func getNumber(digits []int) int {
	return 10*digits[0] + digits[len(digits)-1]
}

func getDigits(line string) []int {
	var digits []int

	for _, letter := range line {
		if number, err := strconv.Atoi(string(letter)); err == nil {
			digits = append(digits, number)
		}
	}

	return digits
}

func Part2(lines []string) {
	var total int

	for _, line := range lines {
		digits := getPart2Digits(line)
		total += getNumber(digits)
	}

	fmt.Println("Part 1:", total)
}

func getPart2Digits(line string) []int {
	out := []int{}

	numbers := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	for i := 0; i < len(line); i++ {
		rest := line[i:]
		letter := line[i]

		for key, number := range numbers {
			if strings.HasPrefix(rest, key) {
				out = append(out, number)
			}
		}

		if number, err := strconv.Atoi(string(letter)); err == nil {
			out = append(out, number)
		}
	}

	return out
}
