package parts

import (
	"slices"

	"github.com/danielstiles/aoc/pkg/parse"
)

type Vec3 struct {
	X int
	Y int
	Z int
}

func (v Vec3) DistanceSquared(o Vec3) int {
	return (v.X-o.X)*(v.X-o.X) +
		(v.Y-o.Y)*(v.Y-o.Y) +
		(v.Z-o.Z)*(v.Z-o.Z)
}

type Distance struct {
	Dsquared int
	Box1     int
	Box2     int
}

func Process1(lines []string) (total int) {
	process := 1000
	topN := 3
	boxes, dists := loadBoxes(lines)
	circuits, maxCircuit, _ := makeCircuits(boxes, dists, process)
	fullCircuits := make([][]int, maxCircuit+1)
	for i, circuit := range circuits {
		if circuit != 0 {
			fullCircuits[circuit] = append(fullCircuits[circuit], i)
		}
	}
	slices.SortFunc(fullCircuits, func(a, b []int) int {
		if len(a) < len(b) {
			return 1
		} else if len(a) == len(b) {
			return 0
		}
		return -1
	})
	total = 1
	for i := range topN {
		if len(fullCircuits[i]) > 0 {
			total *= len(fullCircuits[i])
		}
	}
	return
}

func Process2(lines []string) (total int) {
	boxes, dists := loadBoxes(lines)
	_, _, conn := makeCircuits(boxes, dists, -1)
	total = boxes[conn.Box1].X * boxes[conn.Box2].X
	return
}

func loadBoxes(lines []string) ([]Vec3, []Distance) {
	var boxes []Vec3
	var dists []Distance
	for _, line := range lines {
		coords := parse.FindAllInt(line)
		newBox := Vec3{
			X: coords[0],
			Y: coords[1],
			Z: coords[2],
		}
		for i, oldBox := range boxes {
			dists = append(dists, Distance{
				Dsquared: newBox.DistanceSquared(oldBox),
				Box1:     i,
				Box2:     len(boxes),
			})
		}
		boxes = append(boxes, newBox)
	}
	slices.SortFunc(dists, func(a, b Distance) int {
		if a.Dsquared < b.Dsquared {
			return -1
		} else if a.Dsquared == b.Dsquared {
			return 0
		}
		return 1
	})
	return boxes, dists
}

func makeCircuits(boxes []Vec3, dists []Distance, maxProcess int) (circuits []int, maxCircuit int, conn Distance) {
	circuits = make([]int, len(boxes)) // 0 indicates a circuit of length 1
	boxesFound := 0
	numCircuits := 0
	for pos := 0; boxesFound != len(boxes) || numCircuits != 1; pos++ {
		if maxProcess != -1 && pos >= maxProcess {
			break
		}
		conn = dists[pos]
		c1 := circuits[conn.Box1]
		c2 := circuits[conn.Box2]
		if c1 == 0 {
			if c2 == 0 {
				maxCircuit += 1
				boxesFound += 2
				numCircuits += 1
				circuits[conn.Box1] = maxCircuit
				circuits[conn.Box2] = maxCircuit
				continue
			}
			boxesFound += 1
			circuits[conn.Box1] = c2
			continue
		}
		if c2 == 0 {
			boxesFound += 1
			circuits[conn.Box2] = c1
			continue
		}
		if c1 == c2 {
			continue
		}
		numCircuits -= 1
		for i, circuit := range circuits {
			if circuit == c2 {
				circuits[i] = c1
			}
		}
	}
	return
}
