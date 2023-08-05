package main

import (
	"bufio"
	_ "embed"
	"fmt"
	day7 "rkalmar/aoc2018/day7/internal/pkg"
	"strings"

	"golang.org/x/exp/maps"
)

//go:embed input.txt
var input string

func main() {
	fileScanner := bufio.NewScanner(strings.NewReader(input))
	fileScanner.Split(bufio.ScanLines)
	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, strings.TrimSpace(fileScanner.Text()))
	}
	steps := day7.ParseInput(lines)
	fmt.Println("Step order: ", day7.DetermineStepOrder(maps.Clone(steps)))
	fmt.Println("Total time: ", day7.CalculateTotalTime(maps.Clone(steps), 5, 60))
}
