package operators

import (
	"regexp"
	"strconv"
	"strings"
)

var lineRegex = regexp.MustCompile("(\\d+):([0-9 ]+)")

func Process1(lines []string) (total int) {
	for _, line := range lines {
		testNum, operands := parseLine(line)
		if canMake(operands, testNum, false) > 0 {
			total += testNum
		}
	}
	return
}

func Process2(lines []string) (total int) {
	for _, line := range lines {
		testNum, operands := parseLine(line)
		if canMake(operands, testNum, true) > 0 {
			total += testNum
		}
	}
	return
}

func parseLine(line string) (testNum int, operands []int) {
	parts := lineRegex.FindStringSubmatch(line)
	testNum, _ = strconv.Atoi(parts[1])
	ops := strings.Split(parts[2], " ")
	for _, op := range ops {
		num, _ := strconv.Atoi(op)
		operands = append(operands, num)
	}
	return
}

func canMake(operands []int, num int, concatAllowed bool) (ways int) {
	if len(operands) == 0 {
		return
	}
	max := 1 << (len(operands) - 1)
	if concatAllowed {
		max = pow(3, len(operands)-1)
	}
	for i := 0; i < max; i++ {
		total := operands[0]
		for j := 0; j < len(operands)-1; j++ {
			op := 0
			if !concatAllowed {
				op = (i >> j) & 1
			} else {
				op = (i / pow(3, j)) % 3
			}
			switch op {
			case 0:
				total += operands[j+1]
			case 1:
				total *= operands[j+1]
			case 2:
				total = concat(total, operands[j+1])
			}
		}
		if total == num {
			ways += 1
		}
	}
	return
}

func pow(base, exp int) (res int) {
	res = 1
	for i := 0; i < exp; i++ {
		res *= base
	}
	return
}

func concat(a, b int) (res int) {
	aStr := strconv.Itoa(a)
	bStr := strconv.Itoa(b)
	res, _ = strconv.Atoi(aStr + bStr)
	return
}
