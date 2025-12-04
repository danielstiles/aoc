package parse

import (
	"regexp"
	"strconv"
)

var NumRegex = regexp.MustCompile(`-?\d+`)

func FindAllInt(line string) (nums []int) {
	matches := NumRegex.FindAllString(line, -1)
	for _, m := range matches {
		num, _ := strconv.Atoi(m)
		nums = append(nums, num)
	}
	return
}
