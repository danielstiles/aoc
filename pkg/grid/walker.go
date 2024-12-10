package grid

type Walker struct {
	Grid    *Grid
	Blocker int

	moveMap map[int]map[int]int
}

// CalcMoveMap determines how far from each space the walker needs to move in
// each direction before they hit a blocker, then stores the results for later use.
func (w *Walker) CalcMoveMap() {
	if w.moveMap == nil {
		w.moveMap = make(map[int]map[int]int)
	}
	for i := 0; i < w.Grid.Size.Row; i += 1 {
		w.CalcMoveMapRow(i)
	}
	for j := 0; j < w.Grid.Size.Row; j += 1 {
		w.CalcMoveMapCol(j)
	}
}

// CalcMoveMapRow updates the move map for a given row for the Left and Right directions.
func (w *Walker) CalcMoveMapRow(row int) {
	prevBlocker := -1
	for j := 0; j < w.Grid.Size.Col; j += 1 {
		pos := Vec2{Row: row, Col: j}
		loc := pos.Loc(w.Grid.Size)
		if _, ok := w.moveMap[loc]; !ok {
			w.moveMap[loc] = make(map[int]int)
		}
		if w.Grid.Get(pos) == w.Blocker {
			for j2 := prevBlocker + 1; j2 < j; j2 += 1 {
				loc2 := Vec2{Row: row, Col: j}.Loc(w.Grid.Size)
				w.moveMap[loc2][int(Left)] = j2 - prevBlocker
				if prevBlocker == -1 {
					w.moveMap[loc2][int(Left)] += 1
				}
				w.moveMap[loc2][int(Right)] = j - j2
			}
			prevBlocker = j
		}
	}
	for j2 := prevBlocker + 1; j2 < w.Grid.Size.Col; j2 += 1 {
		loc2 := Vec2{Row: row, Col: j2}.Loc(w.Grid.Size)
		w.moveMap[loc2][int(Left)] = j2 - prevBlocker
		if prevBlocker == -1 {
			w.moveMap[loc2][int(Left)] += 1
		}
		w.moveMap[loc2][int(Right)] = w.Grid.Size.Col - j2 + 1
	}
}

// CalcMoveMapCol updates the move map for a given column for the Up and Down directions.
func (w *Walker) CalcMoveMapCol(col int) {
	prevBlocker := -1
	for i := 0; i < w.Grid.Size.Row; i += 1 {
		pos := Vec2{Row: i, Col: col}
		loc := pos.Loc(w.Grid.Size)
		if _, ok := w.moveMap[loc]; !ok {
			w.moveMap[loc] = make(map[int]int)
		}
		if w.Grid.Get(pos) == w.Blocker {
			for i2 := prevBlocker + 1; i2 < i; i2 += 1 {
				loc2 := Vec2{Row: i2, Col: col}.Loc(w.Grid.Size)
				w.moveMap[loc2][int(Down)] = i2 - prevBlocker - 1
				if prevBlocker == -1 {
					w.moveMap[loc2][int(Down)] += 1
				}
				w.moveMap[loc2][int(Up)] = i - i2
			}
			prevBlocker = i
		}
	}
	for i2 := prevBlocker + 1; i2 < w.Grid.Size.Row; i2 += 1 {
		loc2 := Vec2{Row: i2, Col: col}.Loc(w.Grid.Size)
		w.moveMap[loc2][int(Down)] = i2 - prevBlocker - 1
		if prevBlocker == -1 {
			w.moveMap[loc2][int(Down)] += 1
		}
		w.moveMap[loc2][int(Up)] = w.Grid.Size.Row - i2 + 1
	}
}

// Next determines where the walker will go next, using a calculated move map if allowed.
// Will move one square off the board if it can.
func (w *Walker) Next(start Vec2, dir Dir, useMoveMap bool) Vec2 {
	if useMoveMap && w.moveMap == nil {
		w.CalcMoveMap()
	}
	len := 1
	if useMoveMap {
		len = w.moveMap[start.Loc(w.Grid.Size)][int(dir)]
	}
	return start.Move(dir, len)
}
