package main

import (
	day3 "rkalmar/aoc2018/day3/internal/pkg"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	assert := assert.New(t)
	input := []string{
		"#1 @ 1,3: 4x4",
		"#2 @ 3,1: 4x4",
		"#3 @ 5,5: 2x2",
	}
	claims, err := day3.ParseInput(input)
	if err != nil {
		t.Fatal("Failed to parse input", err)
	}
	overlappingPoints := day3.OverlappingPoints(claims)
	assert.Equal(4, len(overlappingPoints))
}

func TestPart2(t *testing.T) {
	input := []string{
		"#1 @ 1,3: 4x4",
		"#2 @ 3,1: 4x4",
		"#3 @ 5,5: 2x2",
	}
	claims, err := day3.ParseInput(input)
	if err != nil {
		t.Fatal("Failed to parse input", err)
	}
	id, err := day3.NonOverlapping(claims)
	if err != nil {
		t.Fatal("Finding non overlapping claim failed", err)
	}
	assert.Equal(t, 3, id)
}

func TestParseClaim(t *testing.T) {
	assert := assert.New(t)
	claim, err := day3.ParseClaim("#1 @ 1,3: 4x4")
	if err != nil {
		t.Fatal("Failed to parse claim", err)
	}
	assert.Equal(day3.Claim{
		Id: 1,
		Rect: day3.Rect{
			Position: day3.Point{
				X: 1,
				Y: 3,
			},
			Width:  4,
			Height: 4,
		},
	}, *claim)
}
