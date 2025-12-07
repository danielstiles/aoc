package parts

import (
	"github.com/danielstiles/aoc/pkg/grid"
)

var tiles = map[rune]int{
	'.': 0,
	'S': 2,
	'^': 1,
	'|': 2,
}

func Process1(lines []string) (total int) {
	g, locs := grid.Load(lines, tiles)
	counts := grid.New(g.Size)
	start := locs['S'][0]
	total += trace(g, counts, start, false)
	return
}

func Process2(lines []string) (total int) {
	g, locs := grid.Load(lines, tiles)
	counts := grid.New(g.Size)
	start := locs['S'][0]
	total = 1
	total += trace(g, counts, start, true)
	return
}

func trace(g *grid.Grid, counts *grid.Grid, start grid.Vec2, quantum bool) (total int) {
	for {
		start = start.Move(grid.Down, 1)
		if !g.CheckBounds(start) {
			return 0
		}
		if g.Get(start) != tiles['^'] {
			continue
		}
		count := counts.Get(start)
		if count > 0 {
			if quantum {
				return count
			}
			return 0
		}
		left := start.Move(grid.Left, 1)
		if g.CheckBounds(left) {
			total += trace(g, counts, left, quantum)
		}
		right := start.Move(grid.Right, 1)
		if g.CheckBounds(right) {
			total += trace(g, counts, right, quantum)
		}
		counts.Set(start, total+1)
		return total + 1
	}
}
