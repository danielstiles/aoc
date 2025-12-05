package parts

import (
	"github.com/danielstiles/aoc/pkg/parse"
)

func Process1(lines []string) (total int) {
	var ranges [][]int
	var start int
	for i, line := range lines {
		if line == "" {
			start = i + 1
			break
		}
		vals := parse.FindAllInt(line)
		ranges = append(ranges, []int{vals[0], -vals[1]})
	}
	for _, line := range lines[start:] {
		vals := parse.FindAllInt(line)
		num := vals[0]
		var fresh bool
		for _, r := range ranges {
			if num >= r[0] && num <= r[1] {
				fresh = true
				break
			}
		}
		if fresh {
			total += 1
		}
	}
	return
}

func Process2(lines []string) (total int) {
	var ranges [][]int
	for _, line := range lines {
		if line == "" {
			break
		}
		vals := parse.FindAllInt(line)
		start := vals[0]
		end := -vals[1]
		var newRanges [][]int
		for _, r := range ranges {
			if start >= r[0] && start <= r[1] ||
				r[0] >= start && r[0] <= end {
				start = min(start, r[0])
				end = max(end, r[1])
				continue
			}
			newRanges = append(newRanges, r)
		}
		ranges = newRanges
		ranges = append(ranges, []int{start, end})
	}
	for _, r := range ranges {
		total += r[1] - r[0] + 1
	}
	return
}
