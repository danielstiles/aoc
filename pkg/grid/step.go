package grid

const (
	Forward          = 'F'
	Clockwise        = 'C'
	CounterClockwise = 'W'
)

type Step struct {
	Start    Vec2
	StartDir Dir
	Dest     Vec2
	DestDir  Dir
	Len      int
	Path     string
}

// Checks if this step intersects the given position, and returns the steps to get there
func (s Step) Passes(pos Vec2) (Step, bool) {
	curr := s.Start
	currDir := s.StartDir
	traced := ""
	len := 0
	for _, r := range s.Path {
		switch r {
		case Forward:
			curr = curr.Move(currDir, 1)
			len += 1
		case Clockwise:
			currDir = currDir.TurnCW()
		case CounterClockwise:
			currDir = currDir.TurnCCW()
		}
		traced += string(r)
		if curr.Row == pos.Row && curr.Col == pos.Col {
			return Step{
				Start:    s.Start,
				StartDir: s.StartDir,
				Dest:     curr,
				DestDir:  currDir,
				Len:      len,
				Path:     traced,
			}, true
		}
	}
	return Step{}, false
}

// Returns the position at the end of this step
func (s Step) Pos() Vec2 {
	return s.Dest
}

// Returns the number of forward movements of this step
func (s Step) Cost() int {
	return s.Len
}

// Moves from the end of this step to the end of the next one
func (s Step) Move(nextStep Step) (next Step, cost int) {
	if s.Dest.Row != nextStep.Start.Row || s.Dest.Col != nextStep.Start.Col {
		return Step{}, -1
	}
	path := s.Path
	switch nextStep.StartDir {
	case s.DestDir:
	case s.DestDir.TurnCW():
		path = path + string(Clockwise)
	case s.DestDir.TurnCCW():
		path = path + string(CounterClockwise)
	case s.DestDir.TurnCW().TurnCW():
		path = path + string(Clockwise) + string(Clockwise)
	}
	path += nextStep.Path
	next = Step{
		Start:    s.Start,
		StartDir: s.StartDir,
		Dest:     nextStep.Dest,
		DestDir:  nextStep.DestDir,
		Len:      s.Len + nextStep.Len,
		Path:     path,
	}
	return next, next.Len
}

// Needed for BFS interface
func (s Step) Finish() (finished Step, cost int) {
	return s, s.Cost()
}

// Needed for BFS interface
func (s Step) Finished() bool {
	return true
}
