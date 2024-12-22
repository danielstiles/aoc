package parts

import (
	"github.com/danielstiles/aoc/pkg/parse"
)

func Process1(lines []string) (total int) {
	for _, line := range lines {
		seed := int(parse.FindAllInt(line)[0])
		for i := 0; i < 2000; i++ {
			seed = getNext(seed)
		}
		total += int(seed)
	}
	return
}

const (
	changeMask = (1 << 20) - 1
)

func Process2(lines []string) (total int) {
	options := make([]int, changeMask+1)
	currOptions := make([]bool, changeMask+1)
	for _, line := range lines {
		seed := parse.FindAllInt(line)[0]
		changes := 0
		for i := 0; i < 2000; i++ {
			next := getNext(seed)
			diff := (next%10 - seed%10) + 9
			changes = (changes<<5 + diff) & changeMask
			if i >= 3 && currOptions[changes] == false {
				currOptions[changes] = true
				options[changes] += next % 10
			}
			seed = next
		}
		clear(currOptions)
	}
	for _, val := range options {
		total = max(total, val)
	}
	return
}

const pruneMask = (1 << 24) - 1

func getNext(seed int) int {
	seed ^= seed << 6
	seed &= pruneMask
	seed ^= seed >> 5
	seed ^= seed << 11
	seed &= pruneMask
	return seed
}
