package main

import (
	"bufio"
	_ "embed"
	"fmt"
	day2 "rkalmar/aoc2018/day2/internal/pkg"
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
	fmt.Println("Checksum: ", day2.CalculateChecksum(lines))
	fmt.Println("Common letters: ", *day2.FindCommonLettersOfBoxesWithPrototypeFabric(lines))
}
