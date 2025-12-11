package parts

import (
	"strings"

	"github.com/danielstiles/aoc/pkg/math"
	"github.com/danielstiles/aoc/pkg/parse"
)

func Process1(lines []string) (total int) {
	for _, line := range lines {
		lights, buttons, _ := parseLine(line)
		minPresses := len(buttons) + 1
		for mask := range 1 << len(buttons) {
			res := lights
			presses := 0
			for i, button := range buttons {
				if mask&(1<<i) != 0 {
					presses += 1
					res ^= button
				}
			}
			if res == 0 && presses < minPresses {
				minPresses = presses
			}
		}
		total += minPresses
	}
	return
}

func Process2(lines []string) (total int) {
	for _, line := range lines {
		_, buttons, joltages := parseLine(line)
		m := math.NewMatrix(len(joltages), len(buttons))
		maxes := make([]int, len(buttons))
		for i, button := range buttons {
			buttonVec := make([]int, len(joltages))
			miniMax := -1
			for j := range joltages {
				if button&(1<<j) != 0 {
					buttonVec[j] = 1
					if miniMax == -1 || joltages[j] < miniMax {
						miniMax = joltages[j]
					}
				}
			}
			maxes[i] = miniMax
			m.SetCol(i, buttonVec)
		}
		res := m.Solve(joltages)
		best := findPresses(m, res, maxes, make([]int, len(buttons)), m.Rows-1, 0)
		total += sum(best)
	}
	return
}

func parseLine(line string) (lights int, buttons []int, joltages []int) {
	parts := strings.Split(line, " ")
	lightStr := parts[0]
	lightStr = lightStr[1 : len(lightStr)-1]
	for i := len(lightStr) - 1; i >= 0; i-- {
		lights = lights << 1
		if lightStr[i] == '#' {
			lights |= 1
		}
	}
	for _, buttonStr := range parts[1 : len(parts)-1] {
		bits := parse.FindAllInt(buttonStr)
		button := 0
		for _, bit := range bits {
			button |= 1 << bit
		}
		buttons = append(buttons, button)
	}
	joltages = parse.FindAllInt(parts[len(parts)-1])
	return
}

func findPresses(m *math.Matrix, rhs, maxes, current []int, row, locked int) []int {
	minimum := -1
	best := make([]int, len(current))
	toDefine := m.LeadZeros(row)
	goal := rhs[row]
	for col := m.Cols - 1; col > toDefine; col-- {
		if locked&(1<<col) != 0 {
			// Variable is locked in, make the substitution.
			goal -= m.Get(row, col) * current[col]
			continue
		}
		if m.Get(row, col) == 0 {
			// Variable isn't in play for this row.
			continue
		}
		// Found a free variable, try all possible values up to its limit.
		locked |= 1 << col
		for val := range maxes[col] + 1 {
			current[col] = val
			try := findPresses(m, rhs, maxes, current, row, locked)
			if try == nil {
				continue
			}
			presses := sum(try)
			if minimum == -1 || presses < minimum {
				minimum = presses
				copy(best, try)
			}
		}
		if minimum == -1 {
			return nil
		}
		return best
	}
	if toDefine == m.Cols {
		// A row of all 0s, the goal better be 0 or the input is impossible.
		if goal == 0 {
			return findPresses(m, rhs, maxes, current, row-1, locked)
		}
		return nil
	}
	if goal%m.Get(row, toDefine) != 0 {
		// If we can't reach the goal with an integer number of presses, the input is impossible.
		return nil
	}
	val := goal / m.Get(row, toDefine)
	if val < 0 {
		// Can't press a negative number of times.
		return nil
	}
	if locked&(1<<toDefine) != 0 {
		// If we've already locked in this variable, it better be the same value we just computed.
		if val == current[toDefine] {
			return findPresses(m, rhs, maxes, current, row-1, locked)
		}
		return nil
	}
	current[toDefine] = val
	if row == 0 {
		// If we've satisfied every equation, we're done.
		return current
	}
	locked |= 1 << toDefine
	return findPresses(m, rhs, maxes, current, row-1, locked)
}

func sum(slice []int) (total int) {
	for _, val := range slice {
		total += val
	}
	return
}
