package search

import (
	"github.com/danielstiles/aoc/pkg/grid"
	"github.com/danielstiles/aoc/pkg/queue"
)

type Arena interface {
	Size() grid.Vec2
	GetNext(pos grid.Vec2, dir grid.Dir, end grid.Vec2) (next grid.Step, ok bool)
}

type Path[PT any] interface {
	Pos() grid.Vec2
	Cost() int

	Move(nextStep grid.Step) (next PT, cost int)
	Finish() (cost int, finished PT)
	Finished() bool
}

type Record []int

func (r Record) Visit(loc int, dir grid.Dir) bool {
	if r[loc]&int(dir) > 0 {
		return false
	}
	r[loc] |= int(dir)
	return true
}

func (r Record) Get(loc int) int {
	return r[loc]
}

func BFS[A Arena, P Path[P]](a A, start P, end grid.Vec2, all bool) (cost int, best []P) {
	var paths queue.PriorityQueue[P]
	visited := Record(make([]int, a.Size().Row*a.Size().Col))
	finished := false
	bestCost := 0
	paths.Push(0, start)
	for {
		if paths.Len() == 0 {
			return -1, nil
		}
		curr := paths.Pop()
		if finished && (!all || curr.Cost() > bestCost) {
			return
		}
		if end.Distance(curr.Pos()) == 0 {
			if curr.Finished() {
				finished = true
				bestCost = curr.Cost()
				best = append(best, curr)
				continue
			}
			nextCost, nextPath := curr.Finish()
			paths.Push(nextCost, nextPath)
			continue
		}
		for dir := range grid.DirVecs {
			nextStep, ok := a.GetNext(curr.Pos(), dir, end)
			if !ok || !visited.Visit(nextStep.Dest.Loc(a.Size()), nextStep.DestDir) {
				continue
			}
			nextPath, nextCost := curr.Move(nextStep)
			if nextCost != -1 {
				paths.Push(nextCost, nextPath)
			}
		}
	}
}
