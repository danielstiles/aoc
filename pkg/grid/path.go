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

func (p *Path) Copy() (new *Path) {
	new = &Path{
		Visited: make([]int, len(p.Visited)),
		Size:    p.Size,
	}
	for i, val := range p.Visited {
		new.Visited[i] = val
	}
	return
}

// Visit marks the position and direction as visited in this path.
// Returns whether or not the position and direction have been previously visited.
func (p *Path) Visit(pos Vec2, dir Dir) bool {
	loc := pos.Loc(p.Size)
	if p.Visited[loc]&int(dir) > 0 {
		return true
	}
	p.Visited[loc] |= int(dir)
	return false
}

func (p *Path) Get(pos Vec2) int {
	return p.Visited[pos.Loc(p.Size)]
}
