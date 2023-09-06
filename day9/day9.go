package main

import (
	"fmt"
	day9 "rkalmar/aoc2018/day9/internal/pkg"
)

func main() {
	fmt.Println("High score: ", day9.PlayGame(419, 72164))
	fmt.Println("High score: ", day9.PlayGame(419, 72164*100))
}
