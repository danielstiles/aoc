package display

import "github.com/danielstiles/aoc/pkg/grid"

// Prints the grid with the givevn map, returning a slice of the printed lines.
func PrintGrid(g *grid.Grid, key map[int]string) (lines []string) {
	for r := g.Size.Row - 1; r >= 0; r -= 1 {
		var str string
		for c := 0; c < g.Size.Col; c += 1 {
			pos := grid.Vec2{Row: r, Col: c}
			if val, ok := key[g.Get(pos)]; ok {
				str += val
			} else {
				str += "?"
			}
		}
		lines = append(lines, str)
	}
	return
}

var condensedChars = map[int]rune{
	0:  ' ', 1:  '🬀', 2:  '🬁', 3:  '🬂', 4:  '🬃', 5:  '🬄', 6:  '🬅', 7:  '🬆',
	8:  '🬇', 9:  '🬈', 10: '🬉', 11: '🬊', 12: '🬋', 13: '🬌', 14: '🬍', 15: '🬎',
	16: '🬏', 17: '🬐', 18: '🬑', 19: '🬒', 20: '🬓', 21: '▌', 22: '🬔', 23: '🬕',
	24: '🬖', 25: '🬗', 26: '🬘', 27: '🬙', 28: '🬚', 29: '🬛', 30: '🬜', 31: '🬝',
	32: '🬞', 33: '🬟', 34: '🬠', 35: '🬡', 36: '🬢', 37: '🬣', 38: '🬤', 39: '🬥',
	40: '🬦', 41: '🬧', 42: '▐', 43: '🬨', 44: '🬩', 45: '🬪', 46: '🬫', 47: '🬬',
	48: '🬭', 49: '🬮', 50: '🬯', 51: '🬰', 52: '🬱', 53: '🬲', 54: '🬳', 55: '🬴',
	56: '🬵', 57: '🬶', 58: '🬷', 59: '🬸', 60: '🬹', 61: '🬺', 62: '🬻', 63: '█',
}

// Prints a grid containing only 1s and 0s in a condensed form for display.
func PrintCondensedGrid(g *grid.Grid) (lines []string) {
	for r := g.Size.Row - 1; r >= 0; r -= 3 {
		var str string
		for c := 0; c < g.Size.Col; c += 2 {
			pos := grid.Vec2{Row: r, Col: c}
			nextColumn := c < g.Size.Col-1
			cell := g.Get(pos)
			if nextColumn {
				cell += 2 * g.Get(pos.Move(grid.Right, 1))
			}
			if r-1 < 0 {
				continue
			}
			pos = pos.Move(grid.Down, 1)
			cell += 4 * g.Get(pos)
			if nextColumn {
				cell += 8 * g.Get(pos.Move(grid.Right, 1))
			}
			pos = pos.Move(grid.Down, 1)
			cell += 16 * g.Get(pos)
			if nextColumn {
				cell += 32 * g.Get(pos.Move(grid.Right, 1))
			}
			str += string(condensedChars[cell])
		}
		lines = append(lines, str)
	}
	return
}
