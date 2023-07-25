package main

import (
	"bufio"
	_ "embed"
	"fmt"
	day4 "rkalmar/aoc2020/day4/internal/pkg"
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
	entries, err := day4.ParseInput(lines)
	if err != nil {
		panic(fmt.Sprintln("Failed to parse input", err))
	}
	guardSleepPeriods, err := day4.ProcessEntries(entries)
	if err != nil {
		panic(fmt.Sprintln("Failed to process entries", err))
	}
	sleepiestGuardId := day4.FindSleepiestGuardId(guardSleepPeriods)
	mostAsleep := day4.FindMostAsleepMinute(guardSleepPeriods[sleepiestGuardId])
	fmt.Println("ID multiplied by minute at asleep: ", sleepiestGuardId*mostAsleep.Minute)
	guardId, guardMostAsleepAtCertainMinute := day4.FindGuardMostAsleepAtCertainMinute(guardSleepPeriods)
	fmt.Println("ID multiplied by minute for guard most asleep at certain minute: ", guardId*guardMostAsleepAtCertainMinute.Minute)
}
