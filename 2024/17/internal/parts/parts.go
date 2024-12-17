package parts

import (
	"regexp"
	"strconv"
)

type machine struct {
	a      int
	b      int
	c      int
	reader int
	tape   []int
}

func Process1(lines []string) (total string) {
	m := parse(lines)
	total = m.run()
	return
}

func Process2(lines []string) (total int) {
	tape := findAllInt(lines[4])
	possible := []int{0}
	for i := len(tape) - 1; i >= 0; i -= 1 {
		var nextPossible []int
		for _, p := range possible {
			for num := 0; num < 8; num++ {
				m := &machine{
					a:    (p << 3) + num,
					tape: tape,
				}
				for j := 0; j < len(tape)/2-2; j += 1 {
					m.step()
				}
				out := m.step()
				if out == tape[i] {
					nextPossible = append(nextPossible, (p<<3)+num)
				}
			}
		}
		possible = nextPossible
	}
	return possible[0]
}

var numRegex = regexp.MustCompile("-?\\d+")

func parse(lines []string) *machine {
	return &machine{
		a:      findAllInt(lines[0])[0],
		b:      findAllInt(lines[1])[0],
		c:      findAllInt(lines[2])[0],
		reader: 0,
		tape:   findAllInt(lines[4]),
	}
}

func findAllInt(line string) (nums []int) {
	matches := numRegex.FindAllString(line, -1)
	for _, m := range matches {
		num, _ := strconv.Atoi(m)
		nums = append(nums, num)
	}
	return
}

func (m *machine) run() (out string) {
	for m.reader < len(m.tape)-1 {
		if next := m.step(); next != -1 {
			out += strconv.Itoa(next) + ","
		}
	}
	out = out[:len(out)-1]
	return
}

func (m *machine) getCombo(val int) (combo int) {
	combo = val
	switch combo {
	case 4:
		combo = m.a
	case 5:
		combo = m.b
	case 6:
		combo = m.c
	}
	return
}

func (m *machine) step() (out int) {
	out = -1
	op := m.tape[m.reader]
	val := m.tape[m.reader+1]
	m.reader += 2
	combo := m.getCombo(val)
	switch op {
	case 0:
		m.a /= 1 << combo
	case 1:
		m.b ^= val
	case 2:
		m.b = combo % 8
	case 3:
		if m.a != 0 {
			m.reader = val
		}
	case 4:
		m.b ^= m.c
	case 5:
		out = combo % 8
	case 6:
		m.b = m.a / (1 << combo)
	case 7:
		m.c = m.a / (1 << combo)
	}
	return
}
