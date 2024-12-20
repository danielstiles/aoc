package parts

import (
	"regexp"
	"unicode/utf8"
)

var towelRegex = regexp.MustCompile("[wburg]+")

func Process1(lines []string) (total int) {
	towels := towelRegex.FindAllString(lines[0], -1)
	sortedTowels := map[rune]map[string]int{
		'w': make(map[string]int),
		'b': make(map[string]int),
		'u': make(map[string]int),
		'r': make(map[string]int),
		'g': make(map[string]int),
	}
	for _, towel := range towels {
		r, _ := utf8.DecodeRuneInString(towel)
		sortedTowels[r][towel] = 1
	}
	available := map[rune]map[string]int{
		'w': make(map[string]int),
		'b': make(map[string]int),
		'u': make(map[string]int),
		'r': make(map[string]int),
		'g': make(map[string]int),
	}
	for _, line := range lines[2:] {
		if countPossible(line, sortedTowels, available) > 0 {
			total += 1
		}
	}
	return
}

func Process2(lines []string) (total int) {
	towels := towelRegex.FindAllString(lines[0], -1)
	sortedTowels := map[rune]map[string]int{
		'w': make(map[string]int),
		'b': make(map[string]int),
		'u': make(map[string]int),
		'r': make(map[string]int),
		'g': make(map[string]int),
	}
	for _, towel := range towels {
		r, _ := utf8.DecodeRuneInString(towel)
		sortedTowels[r][towel] = 1
	}
	available := map[rune]map[string]int{
		'w': make(map[string]int),
		'b': make(map[string]int),
		'u': make(map[string]int),
		'r': make(map[string]int),
		'g': make(map[string]int),
	}
	for _, line := range lines[2:] {
		total += countPossible(line, sortedTowels, available)
	}
	return
}

func countPossible(towel string, basic, available map[rune]map[string]int) int {
	r, _ := utf8.DecodeRuneInString(towel)
	num, ok := available[r][towel]
	if ok {
		return num
	} else {
		available[r][towel] = 0
	}
	for c := range basic[r] {
		if len(towel) >= len(c) && towel[:len(c)] == c {
			remainder := towel[len(c):]
			if remainder == "" {
				available[r][towel] += 1
				continue
			}
			available[r][towel] += countPossible(remainder, basic, available)
		}
	}
	return available[r][towel]
}
