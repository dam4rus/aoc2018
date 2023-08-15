package day8

import (
	"bufio"
	"errors"
	"strconv"
	"strings"
)

type Node struct {
	ChildNodes      []Node
	MetadataEntries []int
}

func (n *Node) SumMetadataEntries() int {
	var sum int
	for _, entry := range n.MetadataEntries {
		sum += entry
	}
	return sum
}

func (n *Node) SumMetadataEntriesRecursively() int {
	sum := n.SumMetadataEntries()
	for _, childNode := range n.ChildNodes {
		sum += childNode.SumMetadataEntriesRecursively()
	}
	return sum
}

func (n *Node) CalculateValue() int {
	if len(n.ChildNodes) == 0 {
		return n.SumMetadataEntries()
	}
	var value int
	for _, childIndex := range n.MetadataEntries {
		if childIndex == 0 || childIndex > len(n.ChildNodes) {
			continue
		}
		value += n.ChildNodes[childIndex-1].CalculateValue()
	}
	return value
}

func ParseInput(input string) (*Node, error) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)
	return parseNode(scanner)
}

func parseNode(scanner *bufio.Scanner) (*Node, error) {
	if !scanner.Scan() {
		return nil, errors.New("unexpected end of node input")
	}
	childCount, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return nil, err
	}
	if !scanner.Scan() {
		return nil, errors.New("unexpected end of node input")
	}
	metadataEntryCount, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return nil, err
	}
	node := new(Node)

	childNodes, err := parseChildNodes(scanner, childCount)
	if err != nil {
		return nil, err
	}
	node.ChildNodes = childNodes

	metadataEntries, err := parseMetadataEntries(scanner, metadataEntryCount)
	if err != nil {
		return nil, err
	}
	node.MetadataEntries = metadataEntries
	return node, nil
}

func parseChildNodes(scanner *bufio.Scanner, numberOfChilds int) ([]Node, error) {
	var nodes []Node
	for i := 0; i < numberOfChilds; i++ {
		childNode, err := parseNode(scanner)
		if err != nil {
			return nil, err
		}
		nodes = append(nodes, *childNode)
	}
	return nodes, nil
}

func parseMetadataEntries(scanner *bufio.Scanner, numberOfMetadata int) ([]int, error) {
	var metadatas []int
	for i := 0; i < numberOfMetadata; i++ {
		if !scanner.Scan() {
			return nil, errors.New("unexpected end of node input")
		}
		metadata, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		metadatas = append(metadatas, metadata)
	}
	return metadatas, nil
}
