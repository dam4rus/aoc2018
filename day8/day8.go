package main

import (
	_ "embed"
	"fmt"
	day8 "rkalmar/aoc2018/day8/internal/pkg"
)

//go:embed input.txt
var input string

func main() {
	nodeTree, err := day8.ParseInput(input)
	if err != nil {
		panic(err)
	}
	fmt.Println("Sum of metadata entries: ", nodeTree.SumMetadataEntriesRecursively())
	fmt.Println("Value of root node: ", nodeTree.CalculateValue())
}
