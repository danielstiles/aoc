package grid

import (
	"strings"
)

type Maze struct {
	Grid    *Grid
	Blocker int

	MoveMap map[int]map[int]Step
}

func (m *Maze) Size() Vec2 {
	return m.Grid.Size
}

// GetNeighbors returns the next steps from pos. Uses the move map if it has been initialized,
// otherwise moves one square in the given direciton.
func (m *Maze) GetNeighbors(pos Vec2, end Vec2) (next []Step) {
	for dir := range DirVecs {
		if m.MoveMap != nil {
			s, ok := m.MoveMap[pos.Loc(m.Size())][int(dir)]
			if !ok {
				continue
			}
			endStep, ends := s.Passes(end)
			if ends {
				next = append(next, endStep)
				continue
			}
			next = append(next, s)
			continue
		}
		nextPos := pos.Move(dir, 1)
		if !m.Grid.CheckBounds(nextPos) || m.Grid.Get(nextPos) == m.Blocker {
			continue
		}
		next = append(next, Step{
			Start:    pos,
			StartDir: dir,
			Dest:     nextPos,
			DestDir:  dir,
			Len:      1,
			Path:     string(Forward),
		})
	}
	return
}

// CalcMoveMap determines how far from each space the maze allows movement in
// each direction before hitting a blocker (or intersection if desired),
// then stores the results for later use.
func (m *Maze) CalcMoveMap(intersections bool) {
	m.MoveMap = make(map[int]map[int]Step)
	for i := 0; i < m.Grid.Size.Row; i += 1 {
		m.CalcMoveMapRow(i, intersections)
	}
	for j := 0; j < m.Grid.Size.Col; j += 1 {
		m.CalcMoveMapCol(j, intersections)
	}
	m.CondenseMap()
}

// CalcMoveMapRow updates the move map for a given row for the Left and Right directions.
func (m *Maze) CalcMoveMapRow(row int, intersections bool) {
	prev := -1
	prevType := 1 // 0 for intersection, 1 for blocker
	for j := 0; j < m.Grid.Size.Col; j += 1 {
		pos := Vec2{Row: row, Col: j}
		isIntersection := intersections && m.checkIntersection(pos)
		if isIntersection || m.Grid.Get(pos) == m.Blocker {
			currType := 1
			if isIntersection {
				currType = 0
				m.updateMap(row, j, j-prev-prevType, Left)
			}
			if prevType == 0 {
				m.updateMap(row, prev, j-prev-currType, Right)
			}
			for j2 := prev + 1; j2 < j; j2 += 1 {
				m.updateMap(row, j2, j2-prev-prevType, Left)
				m.updateMap(row, j2, j-j2-currType, Right)
			}
			prev = j
			prevType = currType
		}
	}
	for j2 := prev + 1; j2 < m.Grid.Size.Col; j2 += 1 {
		m.updateMap(row, j2, j2-prev-prevType, Left)
		m.updateMap(row, j2, m.Grid.Size.Col-j2-1, Right)
	}
}

// CalcMoveMapCol updates the move map for a given column for the Up and Down directions.
func (m *Maze) CalcMoveMapCol(col int, intersections bool) {
	prev := -1
	prevType := 1 // 0 for intersection, 1 for blocker
	for i := 0; i < m.Grid.Size.Row; i += 1 {
		pos := Vec2{Row: i, Col: col}
		isIntersection := intersections && m.checkIntersection(pos)
		if isIntersection || m.Grid.Get(pos) == m.Blocker {
			currType := 1
			if isIntersection {
				currType = 0
				m.updateMap(i, col, i-prev-prevType, Down)
			}
			if prevType == 0 {
				m.updateMap(prev, col, i-prev-currType, Up)
			}
			for i2 := prev + 1; i2 < i; i2 += 1 {
				m.updateMap(i2, col, i2-prev-prevType, Down)
				m.updateMap(i2, col, i-i2-currType, Up)
			}
			prev = i
			prevType = currType
		}
	}
	for i2 := prev + 1; i2 < m.Grid.Size.Row; i2 += 1 {
		m.updateMap(i2, col, i2-prev-prevType, Down)
		m.updateMap(i2, col, m.Grid.Size.Row-i2-1, Up)
	}
}

// updateMap sets the MoveMap for the row, col to forwards len cells
func (m *Maze) updateMap(row, col, len int, dir Dir) {
	if len == 0 {
		return
	}
	pos := Vec2{Row: row, Col: col}
	loc := pos.Loc(m.Grid.Size)
	if _, ok := m.MoveMap[loc]; !ok {
		m.MoveMap[loc] = make(map[int]Step)
	}
	m.MoveMap[loc][int(dir)] = Step{
		Start:    pos,
		StartDir: dir,
		Dest:     pos.Move(dir, len),
		DestDir:  dir,
		Len:      len,
		Path:     strings.Repeat(string(Forward), len),
	}
}

// checkIntersection returns true if the maze has a choice of paths at this point.
func (m *Maze) checkIntersection(pos Vec2) bool {
	if m.Grid.Get(pos) == m.Blocker {
		return false
	}
	corridors := 0
	neighbors := []Vec2{
		pos.Move(Up, 1),
		pos.Move(Right, 1),
		pos.Move(Down, 1),
		pos.Move(Left, 1),
	}
	for _, n := range neighbors {
		if m.Grid.CheckBounds(n) && m.Grid.Get(n) != m.Blocker {
			corridors += 1
		}
	}
	return corridors > 2
}

// CondenseMap updates the map to follow any passages where there are no choices
// to be made.
func (m *Maze) CondenseMap() {
	newMap := make(map[int]map[int]Step)
	for loc, dirMap := range m.MoveMap {
		newMap[loc] = make(map[int]Step)
		for dir := range dirMap {
			newMap[loc][dir] = m.follow(loc, dir, loc, dir)
		}
	}
	m.MoveMap = newMap
}

// follow creates a new step from the start position/direction going until it hits
// a dead end, an intersection, or itself
func (m *Maze) follow(startLoc, startDir, loc, dir int) Step {
	s := m.MoveMap[loc][dir]
	destLoc := s.Dest.Loc(m.Grid.Size)
	count := 0
	var nextDir Dir
	for d := range DirVecs {
		if d == s.DestDir.TurnCW().TurnCW() {
			continue
		}
		if _, ok := m.MoveMap[destLoc][int(d)]; ok {
			count += 1
			nextDir = d
		}
	}
	if count != 1 || (startLoc == destLoc && startDir == int(nextDir)) {
		return s
	}
	nextStep := m.follow(startLoc, startDir, s.Dest.Loc(m.Grid.Size), int(nextDir))
	nextStep, _ = s.Move(nextStep)
	return nextStep
}
