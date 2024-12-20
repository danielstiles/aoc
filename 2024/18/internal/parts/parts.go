package parts

import (
	"strconv"

	"github.com/danielstiles/aoc/pkg/grid"
	"github.com/danielstiles/aoc/pkg/parse"
	"github.com/danielstiles/aoc/pkg/queue"
)

type path struct {
	length int
	pos    grid.Vec2
}

var size = grid.Vec2{
	Row: 71,
	Col: 71,
}

func Process1(lines []string) (total int) {
	g := grid.New(size)
	start := grid.Vec2{Row: 70, Col: 0}
	end := grid.Vec2{Row: 0, Col: 70}
	for i, line := range lines {
		if i >= 1024 {
			break
		}
		coords := parse.FindAllInt(line)
		byteCoords := grid.Vec2{Row: size.Row - 1 - coords[1], Col: coords[0]}
		g.Set(byteCoords, 1)
	}
	var paths queue.PriorityQueue[path]
	paths.Push(0, path{
		pos: start,
	})
	p := walk(g, paths, end)
	total = p.length
	return
}

func Process2(lines []string) (coord string) {
	g := grid.New(size)
	start := grid.Vec2{Row: 70, Col: 0}
	end := grid.Vec2{Row: 0, Col: 70}
	for _, line := range lines {
		coords := parse.FindAllInt(line)
		byteCoords := grid.Vec2{Row: size.Row - 1 - coords[1], Col: coords[0]}
		g.Set(byteCoords, 1)
		var paths queue.PriorityQueue[path]
		paths.Push(0, path{
			pos: start,
		})
		p := walk(g, paths, end)
		if p == nil {
			coord += strconv.Itoa(size.Row - 1 - byteCoords.Row)
			coord += ","
			coord += strconv.Itoa(byteCoords.Col)
			return
		}
	}
	return
}

func walk(g *grid.Grid, paths queue.PriorityQueue[path], end grid.Vec2) *path {
	found := make(map[int]struct{})
	for {
		if paths.Len() == 0 {
			return nil
		}
		curr := paths.Pop()
		if curr.pos.Row == end.Row && curr.pos.Col == end.Col {
			return &curr
		}
		for dir := range grid.DirVecs {
			next := curr.pos.Move(dir, 1)
			if _, ok := found[next.Loc(g.Size)]; ok {
				continue
			}
			if !g.CheckBounds(next) || g.Get(next) == 1 {
				continue
			}
			found[next.Loc(g.Size)] = struct{}{}
			paths.Push(curr.length+1, path{
				pos:    next,
				length: curr.length + 1,
			})
		}
	}
}
