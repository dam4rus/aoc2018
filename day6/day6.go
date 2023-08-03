package main

import (
	"bufio"
	_ "embed"
	"fmt"
	day6 "rkalmar/aoc2018/day6/internal/pkg"
	"strings"
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
	points, err := day6.ParsePoints(lines)
	if err != nil {
		panic(err)
	}
	grid := day6.NewGrid(points)

	fmt.Println("Largest area: ", grid.FindLargestArea())
	fmt.Println("Safe area size: ", grid.CalculateSafeAreaSize(10000))
}
