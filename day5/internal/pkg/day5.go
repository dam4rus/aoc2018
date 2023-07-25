package day5

import (
	"strings"
	"unicode"
)

func ReducePolymer(polymer string) string {
	reducedPolymer := []rune{rune(polymer[0])}
	for _, char := range polymer[1:] {
		if len(reducedPolymer) > 0 {
			lastUnit := reducedPolymer[len(reducedPolymer)-1]
			if unicode.ToLower(lastUnit) == unicode.ToLower(char) && unicode.IsLower(lastUnit) != unicode.IsLower(char) {
				reducedPolymer = reducedPolymer[:len(reducedPolymer)-1]
				continue
			}
		}

		reducedPolymer = append(reducedPolymer, char)
	}
	return string(reducedPolymer)
}

func FindShortestReducedPolymer(polymer string) string {
	shortestPolymer := polymer
	for _, char := range "qwertyuiopasdfghjklzxcvbnm" {
		replacer := strings.NewReplacer(string(char), "", string(unicode.ToUpper(char)), "")
		reducedPolymer := ReducePolymer(replacer.Replace(polymer))
		if len(reducedPolymer) < len(shortestPolymer) {
			shortestPolymer = reducedPolymer
		}
	}
	return shortestPolymer
}
