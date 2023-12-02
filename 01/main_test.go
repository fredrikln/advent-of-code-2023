package main

import (
	"fmt"
	"testing"
)

func TestGetDigits(t *testing.T) {
	testCases := []struct {
		line     string
		expected []int
	}{
		{"1abc2", []int{1, 2}},
		{"pqr3stu8vwx", []int{3, 8}},
		{"a1b2c3d4e5f", []int{1, 2, 3, 4, 5}},
		{"treb7uchet", []int{7}},
	}
	for i, tC := range testCases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			digits := getDigits(tC.line)

			if !Equal(tC.expected, digits) {
				t.Errorf("expected %d, got %d", tC.expected, digits)
			}
		})
	}
}

func TestGetNumber(t *testing.T) {
	testCases := []struct {
		input    []int
		expected int
	}{
		{[]int{1, 2}, 12},
		{[]int{3, 8}, 38},
		{[]int{1, 2, 3, 4, 5}, 15},
		{[]int{7}, 77},
	}
	for i, tC := range testCases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			got := getNumber(tC.input)

			if tC.expected != got {
				t.Errorf("expected %d, got %d", tC.expected, got)
			}
		})
	}
}

func TestGetPart2Digits(t *testing.T) {
	testCases := []struct {
		input    string
		expected []int
	}{
		{"two1nine", []int{2, 1, 9}},
		{"eightwothree", []int{8, 2, 3}},
		{"abcone2threexyz", []int{1, 2, 3}},
		{"xtwone3four", []int{2, 1, 3, 4}},
		{"4nineeightseven2", []int{4, 9, 8, 7, 2}},
		{"zoneight234", []int{1, 8, 2, 3, 4}},
		{"7pqrstsixteen", []int{7, 6}},
	}
	for i, tC := range testCases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			digits := getPart2Digits(tC.input)

			if !Equal(tC.expected, digits) {
				t.Errorf("expected %d, got %d", tC.expected, digits)
			}
		})
	}
}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, val := range a {
		if val != b[i] {
			return false
		}
	}

	return true
}
