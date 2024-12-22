package search

import (
	"github.com/danielstiles/aoc/pkg/queue"
)

type Node[NT any] interface {
	// Distance should return 0 if equal and not zero otherwise.
	Distance(other NT) int
}

type Path[N Node[N], PT any] interface {
	// Returns the endpoint of the path.
	Pos() N
	// Returns the cost of the path.
	Cost() int

	// Combines this path with nextStep, and returns the resulting path and cost.
	Move(nextStep PT) (next PT, cost int)
	// Finish this path, returning the completed path and cost.
	Finish() (finished PT, cost int)
	// Check if this path is finished.
	Finished() bool
}

type Graph[N Node[N], P Path[N, P]] interface {
	// Returns a slice of all the next paths for the current position.
	GetNeighbors(pos N, end N) []P
}

type Record[N Node[N], P Path[N, P], G Graph[N, P]] interface {
	// Record that the endpoint of P has been visited.
	Visit(g G, curr P) bool
}

func BFS[N Node[N], P Path[N, P], G Graph[N, P], R Record[N, P, G]](g G, start P, end N, record R, all bool) (cost int, best []P) {
	var paths queue.PriorityQueue[P]
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
			nextPath, nextCost := curr.Finish()
			paths.Push(nextCost, nextPath)
			continue
		}
		neighbors := g.GetNeighbors(curr.Pos(), end)
		for _, n := range neighbors {
			if !record.Visit(g, n) {
				continue
			}
			nextPath, nextCost := curr.Move(n)
			if nextCost != -1 {
				paths.Push(nextCost, nextPath)
			}
		}
	}
}
