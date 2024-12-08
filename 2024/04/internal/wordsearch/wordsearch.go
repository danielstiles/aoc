package wordsearch

func Process1(lines []string) (total int) {
	search := parse(lines)
	for i, row := range search {
		for j := range row {
			total += findSeqs(search, i, j)
		}
	}
	return
}

func Process2(lines []string) (total int) {
	search := parse(lines)
	for i, row := range search {
		for j := range row {
			total += findX(search, i, j)
		}
	}
	return
}

var letters = map[rune]int{
	'X': 1,
	'M': 2,
	'A': 3,
	'S': 4,
}

func parse(lines []string) (search [][]int) {
	for _, line := range lines {
		search = append(search, parseLine(line))
	}
	return
}

func parseLine(line string) (search []int) {
	for _, c := range line {
		if val, ok := letters[c]; ok {
			search = append(search, val)
		} else {
			search = append(search, 0)
		}
	}
	return
}

func findSeqs(search [][]int, i, j int) (total int) {
	total += checkDir(search, i, j, 1, 0)
	total += checkDir(search, i, j, 1, 1)
	total += checkDir(search, i, j, 0, 1)
	total += checkDir(search, i, j, -1, 1)
	total += checkDir(search, i, j, -1, 0)
	total += checkDir(search, i, j, -1, -1)
	total += checkDir(search, i, j, 0, -1)
	total += checkDir(search, i, j, 1, -1)
	return
}

func checkDir(search [][]int, i, j, iDir, jDir int) int {
	if len(search) == 0 || i+3*iDir < 0 || i+3*iDir >= len(search) ||
		j+3*jDir < 0 || j+3*jDir >= len(search[0]) {
		return 0
	}
	if search[i][j] == 1 && search[i+iDir][j+jDir] == 2 &&
		search[i+2*iDir][j+2*jDir] == 3 && search[i+3*iDir][j+3*jDir] == 4 {
		return 1
	}
	return 0
}

func findX(search [][]int, i, j int) int {
	if len(search) == 0 || i == 0 || i >= len(search)-1 ||
		j == 0 || j >= len(search[0])-1 {
		return 0
	}
	if search[i][j] != 3 {
		return 0
	}
	diags := [][]int{
		{search[i+1][j+1], search[i-1][j-1]},
		{search[i+1][j-1], search[i-1][j+1]},
	}
	for _, diag := range diags {
		if diag[0] == 2 && diag[1] == 4 {
			continue
		}
		if diag[0] == 4 && diag[1] == 2 {
			continue
		}
		return 0
	}
	return 1
}
