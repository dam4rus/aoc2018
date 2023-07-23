package main

import (
	day1 "rkalmar/aoc2020/day1/internal/pkg"
	"testing"

	"github.com/stretchr/testify/assert"
)

var INPUT = []string{"+1", "-2", "+3", "+1"}

func TestPart1(t *testing.T) {
	frequencies, err := day1.ParseInput(INPUT)
	if err != nil {
		t.Fatalf("Failed to parse input: %d", err)
	}
	frequency := day1.CalculateFrequency(frequencies)
	assert.Equal(t, 3, frequency, "Wrong frequency")
}

func TestPart2(t *testing.T) {
	frequencies, err := day1.ParseInput(INPUT)
	if err != nil {
		t.Fatalf("Failed to parse input: %d", err)
	}
	duplicateFrequency := day1.FindDuplicateFrequency(frequencies)
	assert.Equal(t, 2, duplicateFrequency, "Wrong duplicate frequency")
}
