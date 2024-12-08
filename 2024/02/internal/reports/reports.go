package reports

import (
	"strconv"
	"strings"
)

func Process1(lines []string) (total int) {
	reports := getReports(lines)
	for _, report := range reports {
		total += isSafe(report, false)
	}
	return
}

func Process2(lines []string) (total int) {
	reports := getReports(lines)
	for _, report := range reports {
		check := isSafe(report, true)
		total += check
	}
	return
}

func getReports(lines []string) (reports [][]int) {
	for _, line := range lines {
		split := strings.Split(line, " ")
		var report []int
		for _, s := range split {
			num, _ := strconv.Atoi(s)
			report = append(report, num)
		}
		reports = append(reports, report)
	}
	return
}

func isSafe(report []int, dampener bool) int {
	if isSafeDir(report, true, dampener) == 1 {
		return 1
	} else if isSafeDir(report, false, dampener) == 1 {
		return 1
	}
	return 0
}

func isSafeDir(report []int, increasing, dampener bool) int {
	skipped := !dampener
	for i := 0; i < len(report)-1; i += 1 {
		good := checkDiff(report[i], report[i+1], increasing)
		if !good {
			if skipped {
				return 0
			}
			skipped = true
			if i == len(report)-2 || checkDiff(report[i], report[i+2], increasing) {
				i += 1
				continue
			}
			if i == 0 || checkDiff(report[i-1], report[i+1], increasing) {
				continue
			}
			return 0
		}
	}
	return 1
}

func checkDiff(a, b int, increasing bool) bool {
	diff := a - b
	if increasing {
		diff = -diff
	}
	return (diff <= 3 && diff >= 1)
}
