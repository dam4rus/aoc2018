package main

import (
	day5 "rkalmar/aoc2018/day5/internal/pkg"
	"testing"

	"github.com/stretchr/testify/assert"
)

const INPUT = "dabAcCaCBAcCcaDA"

func TestPart1(t *testing.T) {
	reducedPolymer := day5.ReducePolymer(INPUT)
	assert.Equal(t, "dabCBAcaDA", reducedPolymer)
}

func TestPart2(t *testing.T) {
	reducedPolymer := day5.FindShortestReducedPolymer(INPUT)
	assert.Equal(t, "daDA", reducedPolymer)
}
