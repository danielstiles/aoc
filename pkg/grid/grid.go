package grid

import (
	"slices"
)

type Grid struct {
	Grid []int
	Size Vec2
}

func New(size Vec2) *Grid {
	return &Grid{
		Grid: make([]int, size.Row*size.Col),
		Size: size,
	}
}

// Load takes a list of lines as loaded from a file and converts them to a Grid.
// The Grid treats the last line as row 0, to represent the top left quadrant of a graph.
// It will store any values listed in the key in the grid, and stores 0 for any not in the key.
// It also returns a map of the locations of all the different characters present in the grid.
func Load(lines []string, key map[rune]int) (g *Grid, locs map[rune][]Vec2) {
	if len(lines) == 0 {
		return
	}
	locs = make(map[rune][]Vec2)
	g = &Grid{
		Size: Vec2{
			Row: len(lines),
			Col: len(lines[0]),
		},
	}
	for row, line := range slices.Backward(lines) {
		for col, r := range line {
			if val, ok := key[r]; ok {
				g.Grid = append(g.Grid, val)
			} else {
				g.Grid = append(g.Grid, 0)
			}
			locs[r] = append(
				locs[r],
				Vec2{
					Row: len(lines) - 1 - row,
					Col: col,
				},
			)
		}
	}
	return
}

// Copies the Grid.
func (g *Grid) Copy() (new *Grid) {
	new = &Grid{
		Grid: make([]int, len(g.Grid)),
		Size: g.Size,
	}
	for i, val := range g.Grid {
		new.Grid[i] = val
	}
	return
}

// Get gets the value of the Grid at the specified location.
func (g *Grid) Get(pos Vec2) int {
	return g.Grid[pos.Loc(g.Size)]
}

// Set sets the value of the Grid at the specified location to the given int.
func (g *Grid) Set(pos Vec2, val int) {
	g.Grid[pos.Loc(g.Size)] = val
}

// CheckBounds checks if the given position lies within the grid.
func (g *Grid) CheckBounds(pos Vec2) bool {
	return pos.Row >= 0 && pos.Row < g.Size.Row && pos.Col >= 0 && pos.Col < g.Size.Col
}
