package parts

import (
	"regexp"
	"strconv"

	"github.com/danielstiles/aoc/pkg/parse"
)

const (
	mult = '*'
	add  = '+'
)

var opsRegex = regexp.MustCompile(`\*|\+`)

func Process1(lines []string) (total int) {
	var nums [][]int
	var ops []rune
	for _, line := range lines {
		parts := parse.FindAllInt(line)
		if len(parts) == 0 {
			matches := opsRegex.FindAllString(line, -1)
			for _, m := range matches {
				ops = append(ops, []rune(m)[0])
			}
			break
		}
		nums = append(nums, parts)
	}
	for i, op := range ops {
		res := 0
		if op == mult {
			res = 1
		}
		for _, list := range nums {
			switch op {
			case add:
				res += list[i]
			case mult:
				res *= list[i]
			}
		}
		total += res
	}
	return
}

func Process2(lines []string) (total int) {
	var ops []rune
	end := 0
	for i, line := range lines {
		parts := parse.FindAllInt(line)
		if len(parts) == 0 {
			matches := opsRegex.FindAllString(line, -1)
			for _, m := range matches {
				ops = append(ops, []rune(m)[0])
			}
			end = i
			break
		}
	}
	pos := 0
	res := 0
	if ops[pos] == mult {
		res = 1
	}
	for j := range lines[0] {
		num := 0
		allBlank := true
		for i := range end {
			digit, err := strconv.Atoi(lines[i][j : j+1])
			if err != nil {
				continue
			}
			allBlank = false
			num = num*10 + digit
		}
		if allBlank {
			total += res
			pos += 1
			res = 0
			if ops[pos] == mult {
				res = 1
			}
			continue
		}
		switch ops[pos] {
		case add:
			res += num
		case mult:
			res *= num
		}
	}
	if pos != len(ops) {
		total += res
	}
	return
}
