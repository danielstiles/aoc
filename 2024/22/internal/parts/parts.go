package parts

import (
	"github.com/danielstiles/aoc/pkg/math"
	"github.com/danielstiles/aoc/pkg/parse"
)

func Process1(lines []string) (total int) {
	for _, line := range lines {
		seed := parse.FindAllInt(line)[0]
		for i := 0; i < 2000; i++ {
			seed = getNext(seed)
		}
		total += seed
	}
	return
}

func Process2(lines []string) (total int) {
	options := make(map[int]int)
	mod := math.Pow(37, 4)
	for _, line := range lines {
		seqs := make(map[int]int)
		seed := parse.FindAllInt(line)[0]
		changes := 0
		for i := 0; i < 2000; i++ {
			next := getNext(seed)
			diff := (next%10 - seed%10) + 18
			changes = (changes*37 + diff) % mod
			if _, ok := seqs[changes]; i >= 3 && !ok {
				seqs[changes] = next % 10
			}
			seed = next
		}
		for opt, val := range seqs {
			options[opt] += val
		}
	}
	for _, val := range options {
		total = max(total, val)
	}
	return
}

func getNext(seed int) int {
	seed ^= seed << 6
	seed %= 16777216
	seed ^= seed >> 5
	seed %= 16777216
	seed ^= seed << 11
	seed %= 16777216
	return seed
}
