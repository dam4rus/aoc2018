package main

import (
	"bufio"
	_ "embed"
	"fmt"
	day1 "rkalmar/aoc2020/day1/internal/pkg"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fileScanner := bufio.NewScanner(strings.NewReader(input))
	fileScanner.Split(bufio.ScanLines)
	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	frequencies, err := day1.ParseInput(lines)
	if err != nil {
		panic("Failed to parse input")
	}
	fmt.Println("Frequency: ", day1.CalculateFrequency(frequencies))
	fmt.Println("First duplicate frequency: ", day1.FindDuplicateFrequency(frequencies))
}
