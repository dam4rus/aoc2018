package main

import (
	"bufio"
	_ "embed"
	"fmt"
	day3 "rkalmar/aoc2018/day3/internal/pkg"
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
	claims, err := day3.ParseInput(lines)
	if err != nil {
		panic(fmt.Sprintln("Failed to parse input", err))
	}
	overlappingPoints := day3.OverlappingPoints(claims)
	fmt.Println("Overlapping inches: ", len(overlappingPoints))
	id, err := day3.NonOverlapping(claims)
	if err != nil {
		panic(fmt.Sprintln("Finding non overlapping claim failed", err))
	}
	fmt.Println("Non overlapping id: ", id)
}
