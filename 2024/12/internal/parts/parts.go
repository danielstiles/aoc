package parts

import (
	"github.com/danielstiles/aoc/pkg/grid"
)

type area struct {
	val       int
	locs      []grid.Vec2
	area      int
	perimeter int
}

func Process1(lines []string) (total int) {
	key := make(map[rune]int)
	for i := 1; i < 27; i += 1 {
		key[rune(i-1+'A')] = i
	}
	g, _ := grid.Load(lines, key)
	areas := getAreas(g)
	for _, area := range areas {
		total += area.area * area.perimeter
	}
	return
}

func Process2(lines []string) (total int) {
	key := make(map[rune]int)
	for i := 1; i < 27; i += 1 {
		key[rune(i-1+'A')] = i
	}
	g, _ := grid.Load(lines, key)
	areas := getAreas(g)
	for _, area := range areas {
		total += area.area * numSides(g, area)
	}
	return
}

func getAreas(g *grid.Grid) (res []area) {
	matched := g.Copy()
	for i := 0; i < g.Size.Row; i += 1 {
		for j := 0; j < g.Size.Col; j += 1 {
			pos := grid.Vec2{Row: i, Col: j}
			currVal := matched.Get(pos)
			if currVal == 0 {
				continue
			}
			newArea := fill(matched, pos, currVal)
			res = append(res, newArea)
			for _, interior := range newArea.locs {
				matched.Set(interior, 0)
			}
		}
	}
	return
}

func fill(g *grid.Grid, start grid.Vec2, val int) (res area) {
	res.val = val
	checked := make(map[int]struct{})
	toCheck := []grid.Vec2{start}
	for ; len(toCheck) > 0; toCheck = toCheck[1:] {
		curr := toCheck[0]
		if _, ok := checked[curr.Loc(g.Size)]; ok {
			continue
		}
		if !g.CheckBounds(curr) || g.Get(curr) != val {
			res.perimeter += 1
			continue
		}
		res.locs = append(res.locs, curr)
		res.area += 1
		checked[curr.Loc(g.Size)] = struct{}{}
		toCheck = append(toCheck, curr.Move(grid.Up, 1))
		toCheck = append(toCheck, curr.Move(grid.Right, 1))
		toCheck = append(toCheck, curr.Move(grid.Down, 1))
		toCheck = append(toCheck, curr.Move(grid.Left, 1))
	}
	return
}

func numSides(g *grid.Grid, a area) (sides int) {
	vertices := make(map[int]int)
	interior := make(map[int]struct{})
	for _, pos := range a.locs {
		interior[pos.Loc(g.Size)] = struct{}{}
	}
	for _, pos := range a.locs {
		curr := pos
		vertices[curr.Loc(g.Size)] = checkVertex(g, interior, curr)
		curr = curr.Move(grid.Up, 1)
		vertices[curr.Loc(g.Size)] = checkVertex(g, interior, curr)
		curr = curr.Move(grid.Right, 1)
		vertices[curr.Loc(g.Size)] = checkVertex(g, interior, curr)
		curr = curr.Move(grid.Down, 1)
		vertices[curr.Loc(g.Size)] = checkVertex(g, interior, curr)
	}
	for _, c := range vertices {
		sides += c
	}
	return
}

func checkVertex(g *grid.Grid, interior map[int]struct{}, pos grid.Vec2) int {
	alternating := true
	var inside bool
	var totalInside int
	curr := pos
	if _, ok := interior[curr.Loc(g.Size)]; ok {
		inside = true
		totalInside += 1
	} else {
		inside = false
	}
	curr = curr.Move(grid.Down, 1)
	if _, ok := interior[curr.Loc(g.Size)]; ok {
		alternating = alternating && !inside
		inside = true
		totalInside += 1
	} else {
		alternating = alternating && inside
		inside = false
	}
	curr = curr.Move(grid.Left, 1)
	if _, ok := interior[curr.Loc(g.Size)]; ok {
		alternating = alternating && !inside
		inside = true
		totalInside += 1
	} else {
		alternating = alternating && inside
		inside = false
	}
	curr = curr.Move(grid.Up, 1)
	if _, ok := interior[curr.Loc(g.Size)]; ok {
		alternating = alternating && !inside
		inside = true
		totalInside += 1
	} else {
		alternating = alternating && inside
		inside = false
	}
	if alternating {
		return 2
	} else if totalInside%2 == 0 {
		return 0
	}
	return 1
}
