package lotto

import (
	"regexp"

	"slices"
)

var setExpr = regexp.MustCompile("\\d[0-9 ]*")
var numExpr = regexp.MustCompile("\\d+")

func GetValue(line string) int {
	var total int
	sets := setExpr.FindAllString(line, -1)
	winning := numExpr.FindAllString(sets[1], -1)
	draw := numExpr.FindAllString(sets[2], -1)
	for _, d := range draw {
		if slices.Contains(winning, d) {
			if total == 0 {
				total = 1
			} else {
				total *= 2
			}
		}
	}
	return total
}

func GetWinners(line string) int {
	var total int
	sets := setExpr.FindAllString(line, -1)
	winning := numExpr.FindAllString(sets[1], -1)
	draw := numExpr.FindAllString(sets[2], -1)
	for _, d := range draw {
		if slices.Contains(winning, d) {
			total += 1
		}
	}
	return total
}
