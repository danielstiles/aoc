package parts

import (
	"github.com/danielstiles/aoc/pkg/grid"
)

var key = map[rune]int{
	'0': 0,
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
}

func Process1(lines []string) (total int) {
	g, locs := grid.Load(lines, key)
	starts := locs['0']
	for _, start := range starts {
		ends := hikeUp(g, start)
		total += len(ends)
	}
	return
}

func Process2(lines []string) (total int) {
	g, locs := grid.Load(lines, key)
	starts := locs['0']
	for _, start := range starts {
		ends := hikeUp(g, start)
		for _, count := range ends {
			total += count
		}
	}
	return
}

func hikeUp(g *grid.Grid, start grid.Vec2) (ends map[int]int) {
	currHeight := g.Get(start)
	if currHeight == 9 {
		ends = map[int]int{
			start.Loc(g.Size): 1,
		}
		return
	}
	next := []grid.Vec2{
		start.Move(grid.Up, 1),
		start.Move(grid.Right, 1),
		start.Move(grid.Down, 1),
		start.Move(grid.Left, 1),
	}
	for _, n := range next {
		if !g.CheckBounds(n) || g.Get(n) != currHeight+1 {
			continue
		}
		if ends == nil {
			ends = make(map[int]int)
		}
		for loc, count := range hikeUp(g, n) {
			currCount, ok := ends[loc]
			if !ok {
				currCount = 0
			}
			ends[loc] = count + currCount
		}
	}
	return
}
