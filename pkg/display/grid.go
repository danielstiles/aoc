package display

import "github.com/danielstiles/aoc/pkg/grid"

// Prints the grid with the givevn map, returning a slice of the printed lines.
func PrintGrid(g *grid.Grid, key map[int]rune) (lines []string) {
	for r := g.Size.Row - 1; r >= 0; r -= 1 {
		var str string
		for c := 0; c < g.Size.Col; c += 1 {
			pos := grid.Vec2{Row: r, Col: c}
			if val, ok := key[g.Get(pos)]; ok {
				str += string(val)
			} else {
				str += "?"
			}
		}
		lines = append(lines, str)
	}
	return
}

var condensedChars = map[int]rune{
	0:  ' ',
	1:  '▘',
	2:  '▝',
	3:  '▀',
	4:  '▖',
	5:  '▌',
	6:  '▞',
	7:  '▛',
	8:  '▗',
	9:  '▚',
	10: '▐',
	11: '▜',
	12: '▄',
	13: '▙',
	14: '▟',
	15: '█',
}

// Prints a grid containing only 1s and 0s in a condensed form for display.
func PrintCondensedGrid(g *grid.Grid) (lines []string) {
	for r := g.Size.Row - 1; r >= 0; r -= 2 {
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
			str += string(condensedChars[cell])
		}
		lines = append(lines, str)
	}
	return
}
