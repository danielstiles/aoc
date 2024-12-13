package parts

import (
	"regexp"
	"strconv"
)

var numRegex = regexp.MustCompile("\\d+")

func Process1(lines []string) (total int) {
	for i := 0; i < len(lines); i += 4 {
		a := findAllInt(lines[i])
		b := findAllInt(lines[i+1])
		goal := findAllInt(lines[i+2])
		combs := getCombinations(a, b, goal, 100)
		if len(combs) > 0 {
			cost1 := combs[0][0]*3 + combs[0][1]
			cost2 := combs[len(combs)-1][0]*3 + combs[len(combs)-1][1]
			if cost1 < cost2 {
				total += cost1
			} else {
				total += cost2
			}
		}
	}
	return
}

func Process2(lines []string) (total int) {
	for i := 0; i < len(lines); i += 4 {
		a := findAllInt(lines[i])
		b := findAllInt(lines[i+1])
		goal := findAllInt(lines[i+2])
		goal[0] += 10000000000000
		goal[1] += 10000000000000
		combs := getCombinations(a, b, goal, -1)
		if len(combs) > 0 {
			cost1 := combs[0][0]*3 + combs[0][1]
			cost2 := combs[len(combs)-1][0]*3 + combs[len(combs)-1][1]
			if cost1 < cost2 {
				total += cost1
			} else {
				total += cost2
			}
		}
	}
	return
}

func findAllInt(line string) (nums []int) {
	matches := numRegex.FindAllString(line, -1)
	for _, m := range matches {
		num, _ := strconv.Atoi(m)
		nums = append(nums, num)
	}
	return
}

func getCombinations(a, b, goal []int, cap int) (combs [][]int) {
	numA := goal[0] / a[0]
	if cap != -1 && numA > cap {
		numA = cap
	}
	found := make(map[int][]int)
	var next []int
	for ; numA >= 0; numA -= 1 {
		numB := (goal[0] - a[0]*numA) / b[0]
		if cap != -1 && numB > cap {
			break
		}
		diffx := goal[0] - (a[0]*numA + b[0]*numB)
		if _, ok := found[diffx]; ok {
			next = []int{numA, numB}
			break
		}
		found[diffx] = []int{numA, numB}
	}
	if _, ok := found[0]; !ok {
		return
	}
	if len(next) == 0 {
		c := found[0]
		diffy := goal[1] - (a[1]*c[0] + b[1]*c[1])
		if diffy == 0 {
			combs = append(combs, c)
		}
		return
	}
	diffx := goal[0] - (a[0]*next[0] + b[0]*next[1])
	diffy0 := goal[1] - (a[1]*found[0][0] + b[1]*found[0][1])
	diffy1 := goal[1] - (a[1]*found[diffx][0] + b[1]*found[diffx][1])
	diffy2 := goal[1] - (a[1]*next[0] + b[1]*next[1])
	minusv := diffy1 - diffy2
	diffa := found[diffx][0] - next[0]
	diffb := found[diffx][1] - next[1]
	if minusv == 0 {
		if diffy0 != 0 {
			return
		}
		combs = append(combs, found[0])
		iters := found[0][0] / diffa
		combs = append(combs, []int{
			found[0][0] - iters*diffa,
			found[0][1] - iters*diffb,
		})
		return
	}
	if diffy0/minusv < 0 || diffy0%minusv != 0 {
		return
	}
	iters := diffy0 / minusv
	combs = append(combs, []int{
		found[0][0] - iters*diffa,
		found[0][1] - iters*diffb,
	})
	return
}
