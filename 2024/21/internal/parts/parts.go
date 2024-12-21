package parts

import (
	"fmt"
	"log/slog"
	"slices"

	"github.com/danielstiles/aoc/pkg/grid"
	"github.com/danielstiles/aoc/pkg/parse"
	"github.com/danielstiles/aoc/pkg/queue"
)

var numpad = &grid.Grid{
	Grid: []int{
		-1, 0, 10,
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
	},
	Size: grid.Vec2{
		Row: 4,
		Col: 3,
	},
}

var arrows = &grid.Grid{
	Grid: []int{
		12, 13, 14,
		-1, 11, 10,
	},
	Size: grid.Vec2{
		Row: 2,
		Col: 3,
	},
}

const A = 10

func getNumpadPos(r rune) grid.Vec2 {
	switch r {
	case '0':
		return grid.Vec2{Row: 0, Col: 1}
	case 'A':
		return grid.Vec2{Row: 0, Col: 2}
	case '1':
		return grid.Vec2{Row: 1, Col: 0}
	case '2':
		return grid.Vec2{Row: 1, Col: 1}
	case '3':
		return grid.Vec2{Row: 1, Col: 2}
	case '4':
		return grid.Vec2{Row: 2, Col: 0}
	case '5':
		return grid.Vec2{Row: 2, Col: 1}
	case '6':
		return grid.Vec2{Row: 2, Col: 2}
	case '7':
		return grid.Vec2{Row: 3, Col: 0}
	case '8':
		return grid.Vec2{Row: 3, Col: 1}
	case '9':
		return grid.Vec2{Row: 3, Col: 2}
	}
	return grid.Vec2{}
}

func getArrowPos(arrow int) grid.Vec2 {
	switch arrow {
	case A:
		return grid.Vec2{Row: 1, Col: 2}
	case int(grid.Up):
		return grid.Vec2{Row: 1, Col: 1}
	case int(grid.Right):
		return grid.Vec2{Row: 0, Col: 2}
	case int(grid.Down):
		return grid.Vec2{Row: 0, Col: 1}
	case int(grid.Left):
		return grid.Vec2{Row: 0, Col: 0}
	}
	return grid.Vec2{}
}

func getKey(start, end, size grid.Vec2) int {
	startLoc := start.Loc(arrows.Size)
	endLoc := end.Loc(arrows.Size)
	len := size.Row * size.Col
	return startLoc*len + endLoc
}

func Process1(lines []string) (total int) {
	layer1 := getLayer1()
	layer2 := getLayer(arrows, layer1)
	layer3 := getLayer(numpad, layer2)
	for _, line := range lines {
		cost := 0
		curr := getNumpadPos('A')
		num := parse.FindAllInt(line)[0]
		for _, r := range line {
			next := getNumpadPos(r)
			cost += layer3[getKey(curr, next, numpad.Size)]
			curr = next
		}
		total += cost * num
	}
	return
}

func Process2(lines []string) (total int) {
	layer := getLayer1()
	for i := 1; i < 25; i++ {
		layer = getLayer(arrows, layer)
	}
	layer = getLayer(numpad, layer)
	for _, line := range lines {
		cost := 0
		curr := getNumpadPos('A')
		num := parse.FindAllInt(line)[0]
		for _, r := range line {
			next := getNumpadPos(r)
			cost += layer[getKey(curr, next, numpad.Size)]
			curr = next
		}
		total += cost * num
	}
	return
}

func getLayer1() map[int]int {
	layer1 := make(map[int]int)
	keys := []int{
		A, int(grid.Up), int(grid.Right), int(grid.Down), int(grid.Left),
	}
	for _, a := range keys {
		start := getArrowPos(a)
		for _, b := range keys {
			end := getArrowPos(b)
			rowDiff := abs(start.Row - end.Row)
			colDiff := abs(start.Col - end.Col)
			layer1[getKey(start, end, arrows.Size)] = rowDiff + colDiff + 1
		}
	}
	return layer1
}

func getLayer(g *grid.Grid, arrowCost map[int]int) map[int]int {
	layer := make(map[int]int)
	for i1 := 0; i1 < g.Size.Row; i1 += 1 {
		for j1 := 0; j1 < g.Size.Col; j1 += 1 {
			start := grid.Vec2{Row: i1, Col: j1}
			if g.Get(start) == -1 {
				continue
			}
			for i2 := 0; i2 < g.Size.Row; i2 += 1 {
				for j2 := 0; j2 < g.Size.Col; j2 += 1 {
					end := grid.Vec2{Row: i2, Col: j2}
					if g.Get(end) == -1 {
						continue
					}
					var p path
					p = walk(g, start, end, arrowCost)
					layer[getKey(start, end, g.Size)] = p.len
				}
			}
		}
	}
	return layer
}

type path struct {
	pos        grid.Vec2
	prevButton []int
	len        int
}

func walk(g *grid.Grid, start, end grid.Vec2, cost map[int]int) path {
	var paths queue.PriorityQueue[path]
	paths.Push(1, path{
		pos:        start,
		prevButton: []int{A},
		len:        0,
	})
	for {
		if paths.Len() == 0 {
			return path{}
		}
		curr := paths.Pop()
		prevButton := curr.prevButton[len(curr.prevButton)-1]
		prevPos := getArrowPos(prevButton)
		if curr.pos.Row == end.Row && curr.pos.Col == end.Col {
			nextPos := getArrowPos(A)
			cost := cost[getKey(prevPos, nextPos, arrows.Size)]
			if prevButton == A {
				if curr.len == 0 {
					return path{
						pos:        curr.pos,
						prevButton: append(curr.prevButton, A),
						len:        curr.len + cost,
					}
				}
				return curr
			}
			paths.Push(curr.len+cost, path{
				pos:        curr.pos,
				prevButton: append(curr.prevButton, A),
				len:        curr.len + cost,
			})
			continue
		}
		for dir := range grid.DirVecs {
			next := curr.pos.Move(dir, 1)
			if !g.CheckBounds(next) || g.Get(next) == -1 {
				continue
			}
			nextPos := getArrowPos(int(dir))
			cost := cost[getKey(prevPos, nextPos, arrows.Size)]
			p := slices.Clone(curr.prevButton)
			paths.Push(curr.len+cost, path{
				pos:        next,
				prevButton: append(p, int(dir)),
				len:        curr.len + cost,
			})
		}
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func PrintLayer(layer map[int]int, isArrows bool) {
	var fields []any
	numpadMap := map[int]string{
		1:  "0",
		2:  "A",
		3:  "7",
		4:  "8",
		5:  "9",
		6:  "4",
		7:  "5",
		8:  "6",
		9:  "1",
		10: "2",
		11: "3",
	}
	arrowMap := map[int]string{
		0: "<",
		1: "v",
		2: ">",
		4: "^",
		5: "A",
	}
	for k, v := range layer {
		var from, to string
		if isArrows {
			size := arrows.Size.Row * arrows.Size.Col
			from = arrowMap[k/size]
			to = arrowMap[k%size]
		} else {
			size := numpad.Size.Row * numpad.Size.Col
			from = numpadMap[k/size]
			to = numpadMap[k%size]
		}
		fields = append(fields, slog.Int(fmt.Sprintf("%s -> %s", from, to), v))
	}
	slog.Info("layer", fields...)
}
