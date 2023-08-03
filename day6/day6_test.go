package main

import (
	day6 "rkalmar/aoc2018/day6/internal/pkg"
	"testing"

	"github.com/stretchr/testify/assert"
)

var INPUT = []string{
	"1, 1",
	"1, 6",
	"8, 3",
	"3, 4",
	"5, 5",
	"8, 9",
}

func TestInputParse(t *testing.T) {
	points, err := day6.ParsePoints(INPUT)
	if err != nil {
		t.Fatal(err)
	}
	expected := []day6.Point{
		{X: 1, Y: 1},
		{X: 1, Y: 6},
		{X: 8, Y: 3},
		{X: 3, Y: 4},
		{X: 5, Y: 5},
		{X: 8, Y: 9},
	}
	assert.Equal(t, expected, points)
}

func TestPart1(t *testing.T) {
	assert := assert.New(t)
	points, err := day6.ParsePoints(INPUT)
	if err != nil {
		t.Fatal(err)
	}
	grid := day6.NewGrid(points)
	assert.Equal(17, grid.FindLargestArea())
}

func TestPart2(t *testing.T) {
	points, err := day6.ParsePoints(INPUT)
	if err != nil {
		t.Fatal(err)
	}
	grid := day6.NewGrid(points)
	assert.Equal(t, 16, grid.CalculateSafeAreaSize(32))
}
