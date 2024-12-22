package grid

type Record []int

func (r Record) Visit(g *Grid, s Step) bool {
	destLoc := s.Dest.Loc(g.Size)
	if r[destLoc]&int(s.DestDir) > 0 {
		return false
	}
	r[destLoc] |= int(s.DestDir)
	return true
}

func (r Record) Get(loc int) int {
	return r[loc]
}
