package main

import (
	_ "embed"
	"fmt"
	day5 "rkalmar/aoc2020/day5/internal/pkg"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println("Reduced polymer length: ", len(day5.ReducePolymer(input)))
	fmt.Println("Shortest reduced polymer length: ", len(day5.FindShortestReducedPolymer(input)))
}
