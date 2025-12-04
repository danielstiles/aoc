package parts

import "strconv"

func Process1(lines []string) (total int) {
	for _, line := range lines {
		var digits []int
		for c := range line {
			num, _ := strconv.Atoi(line[c : c+1])
			digits = append(digits, num)
		}
		first := 0
		second := 0
		for i, d := range digits {
			if i < len(digits)-1 && d > first {
				first = d
				second = 0
				continue
			}
			if d > second {
				second = d
			}
		}
		total += first*10 + second
	}
	return
}

func Process2(lines []string) (total int) {
	for _, line := range lines {
		var digits []int
		for c := range line {
			num, _ := strconv.Atoi(line[c : c+1])
			digits = append(digits, num)
		}
		var joltage [12]int
		for i, d := range digits {
			for j := range 12 {
				if i+12-j <= len(digits) && d > joltage[j] {
					joltage[j] = d
					for k := j + 1; k < 12; k++ {
						joltage[k] = 0
					}
					break
				}
			}
		}
		found := 0
		for i := range 12 {
			found *= 10
			found += joltage[i]
		}
		total += found
	}
	return
}
