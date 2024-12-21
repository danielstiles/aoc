package grid

import "github.com/danielstiles/aoc/pkg/math"

type Vec2 struct {
	Row int
	Col int
}

// Add adds the other vector to this one and returns the result.
func (v Vec2) Add(o Vec2) (res Vec2) {
	res.Row = v.Row + o.Row
	res.Col = v.Col + o.Col
	return
}

// Sub subtracts the other vector from this one and returns the result.
func (v Vec2) Sub(o Vec2) (res Vec2) {
	res.Row = v.Row - o.Row
	res.Col = v.Col - o.Col
	return
}

// Mul multiplies this vector by the given length and returns the result.
func (v Vec2) Mul(len int) (res Vec2) {
	res.Row = v.Row * len
	res.Col = v.Col * len
	return
}

// Get the 1D location for this vector in a grid of the given size with rows concatenated.
func (v Vec2) Loc(size Vec2) int {
	return v.Row*size.Col + v.Col
}

// Get the position of a vector in a grid based off its 1D location
func FromLoc(loc int, size Vec2) (res Vec2) {
	res.Col = loc % size.Col
	res.Row = loc / size.Col
	return
}

// Get the Manhattan distance from this vector to another
func (v Vec2) Distance(o Vec2) int {
	rowDist := math.Abs(v.Row - o.Row)
	colDist := math.Abs(v.Col - o.Col)
	return rowDist + colDist
}

type Dir int

const (
	Unknown = Dir(0)
	Up      = Dir(1)
	Right   = Dir(2)
	Down    = Dir(4)
	Left    = Dir(8)
)

var DirVecs = map[Dir]Vec2{
	Up:    Vec2{Row: 1, Col: 0},
	Right: Vec2{Row: 0, Col: 1},
	Down:  Vec2{Row: -1, Col: 0},
	Left:  Vec2{Row: 0, Col: -1},
}

// TurnCW returns the direction clockwise from this one.
func (d Dir) TurnCW() Dir {
	switch d {
	case Up:
		return Right
	case Right:
		return Down
	case Down:
		return Left
	case Left:
		return Up
	}
	return Unknown
}

// TurnCCW returns the direction counter-clockwise from this one.
func (d Dir) TurnCCW() Dir {
	switch d {
	case Up:
		return Left
	case Right:
		return Up
	case Down:
		return Right
	case Left:
		return Down
	}
	return Unknown
}

// Move moves the current vector in the given direction len times and returns the result.
func (v Vec2) Move(d Dir, len int) (new Vec2) {
	return v.Add(DirVecs[d].Mul(len))
}
