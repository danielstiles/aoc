package operators

import (
	"strconv"
)

func Process2B(lines []string) (total int) {
	for _, line := range lines {
		testNum, operands := parseLine(line)
		if canMakeB(operands, testNum) > 0 {
			total += testNum
		}
	}
	return
}

func canMakeB(operands []int, num int) (ways int) {
	if len(operands) == 0 {
		return
	}
	if len(operands) == 1 {
		if operands[0] == num {
			ways = 1
		} else {
			ways = 0
		}
		return
	}
	last := operands[len(operands)-1]
	next := operands[:len(operands)-1]
	ways += canMakeB(next, num-last)
	if num%last == 0 {
		ways += canMakeB(next, num/last)
	}
	if remainder, works := endsWith(num, last); works {
		ways += canMakeB(next, remainder)
	}
	return
}

func endsWith(a, last int) (remains int, works bool) {
	if last > a {
		return
	}
	aStr := strconv.Itoa(a)
	lastStr := strconv.Itoa(last)
	if aStr[len(aStr)-len(lastStr):len(aStr)] != lastStr {
		return
	}
	remains, _ = strconv.Atoi(aStr[:len(aStr)-len(lastStr)])
	works = true
	return
}
