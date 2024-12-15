package parts

import (
	"regexp"

	"github.com/danielstiles/aoc/pkg/grid"
)

var key = map[rune]int{
	'#': 1,
	'O': 2,
	'@': 3,
}

var printKey = map[int]rune{
	0: '.',
	1: '#',
	2: 'O',
	3: '@',
	4: '[',
	5: ']',
}

var dirs = map[rune]grid.Dir{
	'^': grid.Up,
	'>': grid.Right,
	'v': grid.Down,
	'<': grid.Left,
}

var dirRegex = regexp.MustCompile("[>v<^]+")

func Process1(lines []string) (total int) {
	g, pos, end := parse(lines)
	for _, line := range lines[end+1:] {
		pos = followInstructions(g, pos, line)
	}
	for i := 0; i < g.Size.Row; i += 1 {
		for j := 0; j < g.Size.Col; j += 1 {
			if g.Get(grid.Vec2{Row: i, Col: j}) == 2 {
				total += j + 100*(g.Size.Row-1-i)
			}
		}
	}
	return
}

func Process2(lines []string) (total int) {
	g, pos, end := parse(lines)
	g = makeWide(g)
	pos.Col *= 2
	for _, line := range lines[end+1:] {
		pos = followInstructions(g, pos, line)
	}
	for i := 0; i < g.Size.Row; i += 1 {
		for j := 0; j < g.Size.Col; j += 1 {
			if g.Get(grid.Vec2{Row: i, Col: j}) == 4 {
				total += j + 100*(g.Size.Row-1-i)
			}
		}
	}
	return
}

func parse(lines []string) (g *grid.Grid, pos grid.Vec2, end int) {
	end = len(lines) - 1
	for ; end >= 0; end -= 1 {
		if dirRegex.FindString(lines[end]) == "" {
			break
		}
	}
	var locs map[rune][]grid.Vec2
	g, locs = grid.Load(lines[:end], key)
	pos = locs['@'][0]
	return
}

func makeWide(g *grid.Grid) (newGrid *grid.Grid) {
	newGrid = &grid.Grid{
		Grid: make([]int, len(g.Grid)*2),
		Size: grid.Vec2{
			Row: g.Size.Row,
			Col: g.Size.Col * 2,
		},
	}
	for i := range g.Grid {
		switch g.Grid[i] {
		case 0:
			newGrid.Grid[2*i] = 0
			newGrid.Grid[2*i+1] = 0
		case 1:
			newGrid.Grid[2*i] = 1
			newGrid.Grid[2*i+1] = 1
		case 2:
			newGrid.Grid[2*i] = 4
			newGrid.Grid[2*i+1] = 5
		case 3:
			newGrid.Grid[2*i] = 3
			newGrid.Grid[2*i+1] = 0
		}
	}
	return
}

func followInstructions(g *grid.Grid, pos grid.Vec2, line string) grid.Vec2 {
	var success bool
	for _, c := range line {
		moveDir := dirs[c]
		success = tryMove(g, pos, moveDir)
		if success {
			pos = move(g, pos, moveDir)
		}
	}
	return pos
}

func tryMove(g *grid.Grid, pos grid.Vec2, dir grid.Dir) bool {
	nextPos := pos.Move(dir, 1)
	if !g.CheckBounds(nextPos) {
		return false
	}
	success := true
	switch g.Get(nextPos) {
	case 1:
		success = false
	case 2:
		success = tryMove(g, nextPos, dir)
	case 4:
		success = tryMove(g, nextPos, dir)
		if dir == grid.Up || dir == grid.Down {
			success = success && tryMove(g, nextPos.Move(grid.Right, 1), dir)
		}
	case 5:
		success = tryMove(g, nextPos, dir)
		if dir == grid.Up || dir == grid.Down {
			success = success && tryMove(g, nextPos.Move(grid.Left, 1), dir)
		}
	}
	return success
}

func move(g *grid.Grid, pos grid.Vec2, dir grid.Dir) grid.Vec2 {
	nextPos := pos.Move(dir, 1)
	switch g.Get(nextPos) {
	case 2:
		move(g, nextPos, dir)
	case 4:
		move(g, nextPos, dir)
		if dir == grid.Up || dir == grid.Down {
			move(g, nextPos.Move(grid.Right, 1), dir)
		}
	case 5:
		move(g, nextPos, dir)
		if dir == grid.Up || dir == grid.Down {
			move(g, nextPos.Move(grid.Left, 1), dir)
		}
	}
	g.Set(nextPos, g.Get(pos))
	g.Set(pos, 0)
	return nextPos
}
