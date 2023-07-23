package day2

import (
	"strings"
)

func CalculateChecksum(lines []string) int {
	var containsTwoCounter int
	var containsThreeCounter int
	for _, line := range lines {
		containsTwo, containsThree := containsExactlyTwoOrThreeOfTheSameLetter(line)
		if containsTwo {
			containsTwoCounter += 1
		}
		if containsThree {
			containsThreeCounter += 1
		}
	}
	return containsTwoCounter * containsThreeCounter
}

func FindCommonLettersOfBoxesWithPrototypeFabric(lines []string) *string {
	for i, id := range lines {
		for _, otherId := range lines[i+1:] {
			if len(id) != len(otherId) {
				panic("IDs have differing length")
			}
			var differingIndexes []int
			for j := 0; j < len(strings.TrimSpace(id)); j += 1 {
				if id[j] != otherId[j] {
					differingIndexes = append(differingIndexes, j)
				}
			}
			if len(differingIndexes) == 1 {
				differingIndex := differingIndexes[0]
				commonLetters := id[:differingIndex] + id[differingIndex+1:]
				return &commonLetters
			}
		}
	}
	return nil
}

func containsExactlyTwoOrThreeOfTheSameLetter(id string) (bool, bool) {
	charCount := make(map[rune]int)
	for _, c := range id {
		charCount[c] += 1
	}
	var containsTwo bool
	var containsThree bool
	for _, count := range charCount {
		if count == 2 {
			containsTwo = true
		}
		if count == 3 {
			containsThree = true
		}
		if containsTwo && containsThree {
			return true, true
		}
	}
	return containsTwo, containsThree
}
