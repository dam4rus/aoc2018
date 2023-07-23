package day1

import (
	"fmt"
	"strconv"
)

func CalculateFrequency(input []int) int {
	sum := 0
	for _, frequency := range input {
		sum += frequency
	}
	return sum
}

func FindDuplicateFrequency(input []int) int {
	frequencySet := make(map[int]bool)
	currentFrequency := 0
	for {
		for _, frequency := range input {
			currentFrequency += frequency
			contains := frequencySet[currentFrequency]
			if contains {
				return currentFrequency
			}
			frequencySet[currentFrequency] = true
		}
	}
}

func ParseInput(lines []string) ([]int, error) {
	var frequencies []int
	for _, line := range lines {
		intValue, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		frequencies = append(frequencies, intValue)
	}
	return frequencies, nil
}

func Foo() {
	fmt.Println("foo")
}
