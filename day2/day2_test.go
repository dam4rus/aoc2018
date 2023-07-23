package main

import (
	day2 "rkalmar/aoc2020/day2/internal/pkg"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	input := []string{"abcdef", "bababc", "abbcde", "abcccd", "aabcdd", "abcdee", "ababab"}
	checksum := day2.CalculateChecksum(input)
	assert.Equal(t, 4*3, checksum)
}

func TestPart2(t *testing.T) {
	input := []string{
		"abcde",
		"fghij",
		"klmno",
		"pqrst",
		"fguij",
		"axcye",
		"wvxyz",
	}
	commonLetters := day2.FindCommonLettersOfBoxesWithPrototypeFabric(input)
	assert.Equal(t, "fgij", *commonLetters)
}
