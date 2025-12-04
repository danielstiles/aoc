package parts

import (
	"fmt"
	"strconv"
)

const (
	R = 'R'
	L = 'L'
)

func Process1(lines []string) (total int) {
	size := 100
	pos := 50
	for _, line := range lines {
		dir := 1
		if line[0] == L {
			dir = -1
		}
		dist, _ := strconv.Atoi(line[1:])
		pos += dist * dir
		pos = pos % size
		if pos < 0 {
			pos += size
		}
		if pos == 0 {
			total += 1
		}
	}
	return
}

func Process2(lines []string) (total int) {
	size := 100
	pos := 50
	fmt.Printf("pos: %d\n", pos)
	for _, line := range lines {
		dir := 1
		if line[0] == L {
			dir = -1
		}
		dist, _ := strconv.Atoi(line[1:])
		if dist == 0 {
			continue
		}
		prevZero := pos == 0
		pos += dist * dir
		if pos < 0 {
			// Went from positive to negative
			if pos%size == 0 {
				total += 1
			}
			// Don't double count a previous exact match
			if prevZero {
				pos += size
			}
		}
		if pos == 0 {
			total += 1
		}
		for pos >= size {
			pos -= size
			total += 1
		}
		for pos < 0 {
			pos += size
			total += 1
		}
	}
	return
}
