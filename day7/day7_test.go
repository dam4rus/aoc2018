package main

import (
	"fmt"
	day7 "rkalmar/aoc2018/day7/internal/pkg"
	"testing"

	"github.com/stretchr/testify/assert"
)

var INPUT = []string{
	"Step C must be finished before step A can begin.",
	"Step C must be finished before step F can begin.",
	"Step A must be finished before step B can begin.",
	"Step A must be finished before step D can begin.",
	"Step B must be finished before step E can begin.",
	"Step D must be finished before step E can begin.",
	"Step F must be finished before step E can begin.",
}

type Step struct {
	requiredBy []rune
	requires   []rune
}

func (s Step) String() string {
	return fmt.Sprintf("requiredBy: %v, requires: %v", s.requiredBy, s.requires)
}

func TestPart1(t *testing.T) {
	steps := day7.ParseInput(INPUT)
	assert.Equal(t, "CABDFE", string(day7.DetermineStepOrder(steps)))
}

func TestPart2(t *testing.T) {
	steps := day7.ParseInput(INPUT)

	assert.Equal(t, 15, day7.CalculateTotalTime(steps, 2, 0))
	// assert.Equal(t, "CABDFE", string(day7.DetermineStepOrder(steps)))
}
