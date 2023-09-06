package main

import (
	day9 "rkalmar/aoc2018/day9/internal/pkg"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(32, day9.PlayGame(9, 25))
	assert.Equal(8317, day9.PlayGame(10, 1618))
	assert.Equal(146373, day9.PlayGame(13, 7999))
	assert.Equal(2764, day9.PlayGame(17, 1104))
	assert.Equal(54718, day9.PlayGame(21, 6111))
	assert.Equal(37305, day9.PlayGame(30, 5807))
}
