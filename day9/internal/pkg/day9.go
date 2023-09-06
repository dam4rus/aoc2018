package day9

import (
	"golang.org/x/exp/slices"
)

func PlayGame(playerCount, lastMarble int) int {
	marbles := []int{0}
	scores := make([]int, playerCount)
	currentMarble := 0
	for i := 1; i <= lastMarble; i++ {
		if i%23 == 0 {
			currentMarble -= 7
			if currentMarble < 0 {
				currentMarble += len(marbles)
			}
			currentPlayer := (i - 1) % playerCount
			scores[currentPlayer] += marbles[currentMarble] + i
			marbles = slices.Delete(marbles, currentMarble, currentMarble+1)
			continue
		}
		currentMarble += 2
		if currentMarble == len(marbles) {
			marbles = append(marbles, i)
		} else {
			currentMarble = currentMarble % len(marbles)
			marbles = slices.Insert(marbles, currentMarble, i)
		}
	}
	return slices.Max(scores)
}
