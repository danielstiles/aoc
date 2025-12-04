package parts

import (
	"github.com/danielstiles/aoc/pkg/math"
	"github.com/danielstiles/aoc/pkg/parse"
)

func Process1(lines []string) (total int) {
	nums := parse.FindAllInt(lines[0])
	for i := 0; i < len(nums); i += 2 {
		for n := nums[i]; n <= -nums[i+1]; n++ {
			digits := math.Digits(n, 10)
			if digits%2 == 1 {
				continue
			}
			start, end := math.IntSplit(n, 10, digits/2)
			if start == end {
				total += n
			}
		}
	}
	return
}

func Process2(lines []string) (total int) {
	nums := parse.FindAllInt(lines[0])
	for i := 0; i < len(nums); i += 2 {
		for n := nums[i]; n <= -nums[i+1]; n++ {
			match := findMatch(n)
			if match {
				total += n
			}
		}
	}
	return
}

func findMatch(n int) (match bool) {
	digits := math.Digits(n, 10)
	for length := 1; length <= digits/2; length++ {
		if digits%length != 0 {
			continue
		}
		start, test := math.IntSplit(n, 10, length)
		end := 0
		same := true
		for start > 0 {
			start, end = math.IntSplit(start, 10, length)
			if end != test {
				same = false
				break
			}
		}
		if same {
			return true
		}
	}
	return false
}
