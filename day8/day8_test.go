package main

import (
	day8 "rkalmar/aoc2018/day8/internal/pkg"
	"testing"

	"github.com/stretchr/testify/assert"
)

const INPUT = "2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2"

func TestNodeParse(t *testing.T) {
	leafNode, err := day8.ParseInput("0 1 99")
	if err != nil {
		t.Fatal("Failed to parse leaf node: ", err)
	}
	assert.Equal(t, &day8.Node{MetadataEntries: []int{99}}, leafNode)
}

func TestChildNodeParse(t *testing.T) {
	nodeTree, err := day8.ParseInput("1 1 0 1 99 2")
	if err != nil {
		t.Fatal("Failed to parse node tree: ", err)
	}
	assert.Equal(t, &day8.Node{
		ChildNodes: []day8.Node{
			{MetadataEntries: []int{99}},
		},
		MetadataEntries: []int{2},
	}, nodeTree)
}

func TestWholeInputParse(t *testing.T) {
	nodeTree, err := day8.ParseInput(INPUT)
	if err != nil {
		t.Fatal("Failed to parse node tree: ", err)
	}
	assert.Equal(t, &day8.Node{
		ChildNodes: []day8.Node{
			{
				MetadataEntries: []int{10, 11, 12},
			},
			{
				ChildNodes: []day8.Node{
					{
						MetadataEntries: []int{99},
					},
				},
				MetadataEntries: []int{2},
			},
		},
		MetadataEntries: []int{1, 1, 2},
	}, nodeTree)
}

func TestPart1(t *testing.T) {
	nodeTree, err := day8.ParseInput(INPUT)
	if err != nil {
		t.Fatal("Failed to parse node tree: ", err)
	}
	assert.Equal(t, 138, nodeTree.SumMetadataEntriesRecursively())
}

func TestPart2(t *testing.T) {
	nodeTree, err := day8.ParseInput(INPUT)
	if err != nil {
		t.Fatal("Failed to parse node tree: ", err)
	}
	assert.Equal(t, 66, nodeTree.CalculateValue())
}
