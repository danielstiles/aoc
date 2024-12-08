package guard

func Process1(lines []string) (total int) {
	grid, guardR, guardC := getGrid(lines)
	_, total, _ = tracePath(grid, guardR, guardC, Up)
	return
}

func Process2(lines []string) (total int) {
	grid, guardR, guardC := getGrid(lines)
	path, _, _ := tracePath(grid, guardR, guardC, Up)
	for i, row := range grid {
		for j := range row {
			if i == guardR && j == guardC {
				continue
			}
			if path[i][j] > 2 {
				grid[i][j] = Blocker
				_, _, loop := tracePath(grid, guardR, guardC, Up)
				if loop {
					total += 1
				}
				grid[i][j] = Ground
			}
		}
	}
	return
}

const (
	Ground  = 0
	Guard   = 1
	Blocker = 2
	Up      = 4
	Right   = 8
	Down    = 16
	Left    = 32
)

var objects = map[rune]int{
	'^': Guard,
	'#': Blocker,
	'.': Ground,
}

func getGrid(lines []string) (grid [][]int, guardR, guardC int) {
	for i, line := range lines {
		gridLine, col := getGridLine(line)
		grid = append(grid, gridLine)
		if col != -1 {
			guardR = i
			guardC = col
		}
	}
	return
}

func getGridLine(line string) (gridLine []int, guardC int) {
	guardC = -1
	for i, r := range line {
		if objects[r] == Guard {
			guardC = i
			gridLine = append(gridLine, Ground)
		} else {
			gridLine = append(gridLine, objects[r])
		}
	}
	return
}

func tracePath(grid [][]int, guardR, guardC, dir int) (path [][]int, total int, loop bool) {
	if len(grid) == 0 {
		return
	}
	path = copyGrid(grid)
	for {
		if path[guardR][guardC]>>2 == 0 {
			total += 1
		}
		if path[guardR][guardC]&dir != 0 {
			return path, total, true
		}
		path[guardR][guardC] |= dir
		r, c := move(guardR, guardC, dir)
		if r < 0 || c < 0 || r >= len(grid) || c >= len(grid[0]) {
			return
		}
		for grid[r][c] == Blocker {
			dir = turn(dir)
			r, c = move(guardR, guardC, dir)
			if r < 0 || c < 0 || r >= len(grid) || c >= len(grid[0]) {
				return
			}
		}
		guardR = r
		guardC = c
	}
}

func turn(dir int) int {
	switch dir {
	case Up:
		return Right
	case Right:
		return Down
	case Down:
		return Left
	case Left:
		return Up
	}
	return Up
}

func move(r, c, dir int) (newR, newC int) {
	newR = r
	newC = c
	switch dir {
	case Up:
		newR -= 1
	case Right:
		newC += 1
	case Down:
		newR += 1
	case Left:
		newC -= 1
	}
	return
}

func copyGrid(grid [][]int) (newGrid [][]int) {
	for _, row := range grid {
		var gridLine []int
		for _, val := range row {
			gridLine = append(gridLine, val)
		}
		newGrid = append(newGrid, gridLine)
	}
	return
}
