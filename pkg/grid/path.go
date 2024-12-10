package grid

type Path struct {
	Visited []int
	Size    Vec2
}

// NewPath creates a new path that will lie on the given grid.
func NewPath(g *Grid) *Path {
	return &Path{
		Visited: make([]int, len(g.Grid)),
		Size:    g.Size,
	}
}

// Visit marks the position and direction as visited in this path.
func (p *Path) Visit(pos Vec2, dir Dir) bool {
	loc := pos.Loc(p.Size)
	if p.Visited[loc]&int(dir) > 0 {
		return true
	}
	p.Visited[loc] |= int(dir)
	return false
}
