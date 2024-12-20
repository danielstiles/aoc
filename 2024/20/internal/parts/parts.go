package parts

import (
	"slices"

	"github.com/danielstiles/aoc/pkg/grid"
	"github.com/danielstiles/aoc/pkg/queue"
)

var key = map[rune]int{
	'#': 1,
}

func Process1(lines []string) (total int) {
	g, locs := grid.Load(lines, key)
	start := locs['S'][0]
	end := locs['E'][0]
	paths := queue.PriorityQueue[[]grid.Vec2]{}
	paths.Push(1, []grid.Vec2{start})
	path := walk(g, paths, end)
	for i := range path {
		cheats := tryCheat(g, path, i)
		for _, cheat := range cheats {
			if cheat >= 100 {
				total += 1
			}
		}
	}
	return
}

func Process2(lines []string) (total int) {
	g, locs := grid.Load(lines, key)
	start := locs['S'][0]
	end := locs['E'][0]
	paths := queue.PriorityQueue[[]grid.Vec2]{}
	paths.Push(1, []grid.Vec2{start})
	path := walk(g, paths, end)
	cheatLen := 20
	for i := range path {
		for j := range path {
			colsNeeded := path[i].Col - path[j].Col
			if colsNeeded < 0 {
				colsNeeded = -colsNeeded
			}
			rowsNeeded := path[i].Row - path[j].Row
			if rowsNeeded < 0 {
				rowsNeeded = -rowsNeeded
			}
			totalNeeded := colsNeeded + rowsNeeded
			if totalNeeded > cheatLen {
				continue
			}
			if j-i-totalNeeded >= 100 {
				total += 1
			}
		}
	}
	return
}

func walk(g *grid.Grid, paths queue.PriorityQueue[[]grid.Vec2], end grid.Vec2) []grid.Vec2 {
	visited := make(map[int]bool)
	for {
		if paths.Len() == 0 {
			return nil
		}
		curr := paths.Pop()
		pos := curr[len(curr)-1]
		if pos.Row == end.Row && pos.Col == end.Col {
			return curr
		}
		for dir := range grid.DirVecs {
			next := pos.Move(dir, 1)
			loc := next.Loc(g.Size)
			if _, ok := visited[loc]; !g.CheckBounds(next) || g.Get(next) == 1 || ok {
				continue
			}
			visited[loc] = true
			nextPath := slices.Clone(curr)
			nextPath = append(nextPath, next)
			paths.Push(len(nextPath), nextPath)
			continue
		}
	}
}

func tryCheat(g *grid.Grid, path []grid.Vec2, i int) map[int]int {
	ret := make(map[int]int)
	for dir := range grid.DirVecs {
		next := path[i].Move(dir, 2)
		if !g.CheckBounds(next) || g.Get(next) == 1 {
			ret[int(dir)] = -1
			continue
		}
		j := find(next, path)
		if j == -1 {
			ret[int(dir)] = -1
		} else {
			ret[int(dir)] = j - i - 2
		}
	}
	return ret
}

func find(pos grid.Vec2, path []grid.Vec2) int {
	for i, c := range path {
		if pos.Row == c.Row && pos.Col == c.Col {
			return i
		}
	}
	return -1
}
