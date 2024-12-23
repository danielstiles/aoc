package parts

import (
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/danielstiles/aoc/pkg/display"
	"github.com/danielstiles/aoc/pkg/grid"
)

func Process1(lines []string) (total int) {
	rows := 103
	cols := 101
	qs := make([]int, 4)
	for _, line := range lines {
		vals := findAllInt(line)
		pos := grid.Vec2{
			Row: vals[1],
			Col: vals[0],
		}
		vel := grid.Vec2{
			Row: vals[3],
			Col: vals[2],
		}
		endPos := pos.Add(vel.Mul(100))
		endPos.Row = endPos.Row % rows
		if endPos.Row < 0 {
			endPos.Row += rows
		}
		endPos.Col = endPos.Col % cols
		if endPos.Col < 0 {
			endPos.Col += cols
		}
		if endPos.Row == rows/2 || endPos.Col == cols/2 {
			continue
		}
		q := 0
		if endPos.Row > rows/2 {
			q += 2
		}
		if endPos.Col > cols/2 {
			q += 1
		}
		qs[q] += 1
	}
	total = qs[0] * qs[1] * qs[2] * qs[3]
	return
}

func Process2(lines []string) (total int) {
	rows := 103
	cols := 101
	var robots []grid.Vec2
	var robotVels []grid.Vec2
	for _, line := range lines {
		vals := findAllInt(line)
		pos := grid.Vec2{
			Row: rows - 1 - vals[1],
			Col: vals[0],
		}
		vel := grid.Vec2{
			Row: -vals[3],
			Col: vals[2],
		}
		robots = append(robots, pos)
		robotVels = append(robotVels, vel)
	}
	for {
		g := grid.New(grid.Vec2{Row: rows, Col: cols})
		total += 1
		for i := range robots {
			pos := robots[i]
			vel := robotVels[i]
			endPos := pos.Add(vel)
			endPos.Row = endPos.Row % rows
			if endPos.Row < 0 {
				endPos.Row += rows
			}
			endPos.Col = endPos.Col % cols
			if endPos.Col < 0 {
				endPos.Col += cols
			}
			robots[i] = endPos
			g.Set(endPos, 1)
		}
		c := strings.Join(display.PrintCondensedGrid(g), "\n")
		if strings.Count(c, "█") > len(robots)/16 {
			os.Stdout.WriteString("Round: " + strconv.Itoa(total) + "\n")
			os.Stdout.WriteString(strings.Join(display.PrintCondensedGrid(g), "\n") + "\n")
			break
		}
	}
	return
}

var numRegex = regexp.MustCompile("-?\\d+")

func findAllInt(line string) (nums []int) {
	matches := numRegex.FindAllString(line, -1)
	for _, m := range matches {
		num, _ := strconv.Atoi(m)
		nums = append(nums, num)
	}
	return
}

func show(robots []grid.Vec2, rows, cols, round int) {
	g := &grid.Grid{
		Grid: make([]int, rows*cols),
		Size: grid.Vec2{
			Row: rows,
			Col: cols,
		},
	}
	for _, r := range robots {
		g.Set(r, 1)
	}
	os.Stdout.WriteString("Round: " + strconv.Itoa(round) + "\n")
	os.Stdout.WriteString(strings.Join(display.PrintCondensedGrid(g), "\n") + "\n")
}
