package parts

import (
	"github.com/danielstiles/aoc/pkg/grid"
	"github.com/danielstiles/aoc/pkg/parse"
	"github.com/danielstiles/aoc/pkg/queue"
)

type Rectangle struct {
	Left   int
	Right  int
	Top    int
	Bottom int
	area   int
}

func NewRect(corner1, corner2 grid.Vec2) *Rectangle {
	var ret Rectangle
	ret.Left = min(corner1.Col, corner2.Col)
	ret.Right = max(corner1.Col, corner2.Col)
	ret.Bottom = min(corner1.Row, corner2.Row)
	ret.Top = max(corner1.Row, corner2.Row)
	return &ret
}

func (r *Rectangle) Area() int {
	if r.area != 0 {
		return r.area
	}
	r.area = (r.Top - r.Bottom + 1)
	r.area *= (r.Right - r.Left + 1)
	return r.area
}

func (r *Rectangle) Interstects(o *Rectangle) bool {
	return !(o.Left >= r.Right || o.Right <= r.Left ||
		o.Bottom >= r.Top || o.Top <= r.Bottom)
}

func Process1(lines []string) (total int) {
	_, rects := getRects(lines)
	best := rects.Pop()
	total = best.Area()
	return
}

func Process2(lines []string) (total int) {
	tiles, rects := getRects(lines)
	shape := getShape(tiles)
	var best *Rectangle
	for best = rects.Pop(); !isValid(best, shape); best = rects.Pop() {
	}
	total = best.Area()
	return
}

func getRects(lines []string) ([]grid.Vec2, *queue.PriorityQueue[*Rectangle]) {
	var tiles []grid.Vec2
	rects := &queue.PriorityQueue[*Rectangle]{}
	for _, line := range lines {
		nums := parse.FindAllInt(line)
		newTile := grid.Vec2{
			Row: nums[0],
			Col: nums[1],
		}
		for _, tile := range tiles {
			newRect := NewRect(tile, newTile)
			rects.Push(-newRect.Area(), newRect)
		}
		tiles = append(tiles, newTile)
	}
	return tiles, rects
}

func getShape(tiles []grid.Vec2) []*Rectangle {
	var shape []*Rectangle
	start := tiles[len(tiles)-1]
	for _, end := range tiles {
		shape = append(shape, NewRect(start, end))
		start = end
	}
	return shape
}

func isValid(r *Rectangle, shape []*Rectangle) bool {
	var insideCheck *Rectangle
	if r.Top-r.Bottom > 1 && r.Right-r.Left > 1 {
		point := grid.Vec2{
			Row: r.Bottom + 1,
			Col: r.Left + 1,
		}
		left := grid.Vec2{
			Row: r.Bottom + 1,
			Col: 0,
		}
		insideCheck = NewRect(left, point)
	}
	var inside bool
	for _, edge := range shape {
		if r.Interstects(edge) {
			return false
		}
		if insideCheck != nil && insideCheck.Interstects(edge) {
			inside = !inside
		}
	}
	return inside
}
