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
	m := &grid.Maze{
		Grid:    g,
		Blocker: 1,
	}
	m.CalcMoveMap(true)
	var paths queue.PriorityQueue[path]
	paths.Push(0, path{
		pos:  start,
		dir:  grid.Right,
		path: "",
		cost: 0,
	})
	p := walk(m, paths, end)
	total = p[0].cost
	return
}

func Process2(lines []string) (total int) {
	g, locs := grid.Load(lines, key)
	start := locs['S'][0]
	end := locs['E'][0]
	m := &grid.Maze{
		Grid:    g,
		Blocker: 1,
	}
	m.CalcMoveMap(true)
	var paths queue.PriorityQueue[path]
	paths.Push(0, path{
		pos:  start,
		dir:  grid.Right,
		path: "",
		cost: 0,
	})
	best := walk(m, paths, end)
	visited := grid.Record(make([]int, g.Size.Row*g.Size.Col))
	for _, b := range best {
		curr := start
		currDir := grid.Right
		visited.Visit(g, grid.Step{
			Dest:    curr,
			DestDir: currDir,
		})
		for _, r := range b.path {
			switch r {
			case 'F':
				curr = curr.Move(currDir, 1)
			case 'C':
				currDir = currDir.TurnCW()
			case 'W':
				currDir = currDir.TurnCCW()
			}
			visited.Visit(g, grid.Step{
				Dest:    curr,
				DestDir: currDir,
			})
		}
	}
	for i := 0; i < g.Size.Row; i += 1 {
		for j := 0; j < g.Size.Col; j += 1 {
			pos := grid.Vec2{Row: i, Col: j}
			if visited.Get(pos.Loc(g.Size)) > 0 {
				total += 1
			}
		}
	}
	return
}

type path struct {
	pos  grid.Vec2
	dir  grid.Dir
	path string
	cost int
}

const turnCost = 1000

func walk(m *grid.Maze, paths queue.PriorityQueue[path], end grid.Vec2) (best []path) {
	found := make(map[int][]grid.Dir)
	for {
		curr := paths.Pop()
		loc := curr.pos.Loc(m.Grid.Size)
		found[loc] = append(found[loc], curr.dir)
		if curr.pos.Row == end.Row && curr.pos.Col == end.Col {
			bestCost := curr.cost
			for curr.cost == bestCost {
				if curr.pos.Row == end.Row && curr.pos.Col == end.Col {
					best = append(best, curr)
				}
				if paths.Len() == 0 {
					break
				}
				curr = paths.Pop()
			}
			return
		}
		next, ok := m.MoveMap[loc]
		if ok && (len(curr.path) == 0 || curr.path[len(curr.path)-1] == 'F') {
			_, right := next[int(curr.dir.TurnCW())]
			_, left := next[int(curr.dir.TurnCCW())]
			costAfterTurn := curr.cost + turnCost
			if right && !slices.Contains(found[loc], curr.dir.TurnCW()) {
				paths.Push(costAfterTurn, path{
					pos:  curr.pos,
					dir:  curr.dir.TurnCW(),
					path: curr.path + "C",
					cost: costAfterTurn,
				})
			}
			if left && !slices.Contains(found[loc], curr.dir.TurnCCW()) {
				paths.Push(costAfterTurn, path{
					pos:  curr.pos,
					dir:  curr.dir.TurnCCW(),
					path: curr.path + "W",
					cost: costAfterTurn,
				})
			}
		}
		nextStep, forward := next[int(curr.dir)]
		if !forward {
			continue
		}
		if endStep, passes := nextStep.Passes(end); passes {
			nextStep = endStep
		}
		nextCost := curr.cost + calcCost(nextStep)
		if !slices.Contains(found[nextStep.Dest.Loc(m.Grid.Size)], nextStep.DestDir) {
			paths.Push(nextCost, path{
				pos:  nextStep.Dest,
				dir:  nextStep.DestDir,
				path: curr.path + nextStep.Path,
				cost: nextCost,
			})
		}
	}
}

func calcCost(s grid.Step) (total int) {
	for _, r := range s.Path {
		switch r {
		case 'F':
			total += 1
		case 'C':
			total += 1000
		case 'W':
			total += 1000
		}
	}
	return
}
