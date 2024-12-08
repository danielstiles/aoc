package memory

import (
	"regexp"
	"strconv"
)

var mulRegex = regexp.MustCompile("mul\\((\\d{1,3}),(\\d{1,3})\\)")
var partsRegex = regexp.MustCompile("mul\\((\\d{1,3}),(\\d{1,3})\\)|do\\(\\)|don't\\(\\)")

func Process1(lines []string) (total int) {
	for _, line := range lines {
		matches := mulRegex.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			a, _ := strconv.Atoi(match[1])
			b, _ := strconv.Atoi(match[2])
			total += a * b
		}
	}
	return
}

func Process2(lines []string) (total int) {
	enabled := true
	for _, line := range lines {
		matches := partsRegex.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			switch match[0] {
			case "do()":
				enabled = true
			case "don't()":
				enabled = false
			default:
				if !enabled {
					continue
				}
				a, _ := strconv.Atoi(match[1])
				b, _ := strconv.Atoi(match[2])
				total += a * b
			}
		}
	}
	return
}
