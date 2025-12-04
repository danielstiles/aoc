package parts

import "github.com/danielstiles/aoc/pkg/grid"

var tiles = map[rune]int{
	'@': 1,
}

var dirs = []grid.Vec2{
	grid.DirVecs[grid.Up],
	grid.DirVecs[grid.Up].Add(grid.DirVecs[grid.Right]),
	grid.DirVecs[grid.Right],
	grid.DirVecs[grid.Right].Add(grid.DirVecs[grid.Down]),
	grid.DirVecs[grid.Down],
	grid.DirVecs[grid.Down].Add(grid.DirVecs[grid.Left]),
	grid.DirVecs[grid.Left],
	grid.DirVecs[grid.Left].Add(grid.DirVecs[grid.Up]),
}

func Process1(lines []string) (total int) {
	g, locsMap := grid.Load(lines, tiles)
	locs := locsMap['@']
	counts := grid.New(g.Size)
	for _, loc := range locs {
		for _, dir := range dirs {
			toCheck := loc.Add(dir)
			if g.CheckBounds(toCheck) {
				counts.Set(toCheck, counts.Get(toCheck)+1)
			}
		}
	}
	for _, loc := range locs {
		if counts.Get(loc) < 4 {
			total += 1
		}
	}
	return
}

func Process2(lines []string) (total int) {
	g, locsMap := grid.Load(lines, tiles)
	locs := locsMap['@']
	counts := grid.New(g.Size)
	for _, loc := range locs {
		for _, dir := range dirs {
			toCheck := loc.Add(dir)
			if g.CheckBounds(toCheck) {
				counts.Set(toCheck, counts.Get(toCheck)+1)
			}
		}
	}
	removed := 1 // Force one execution of the loop
	for removed > 0 {
		removed = 0
		var newLocs []grid.Vec2
		for _, loc := range locs {
			if counts.Get(loc) < 4 {
				removed += 1
				for _, dir := range dirs {
					toCheck := loc.Add(dir)
					if g.CheckBounds(toCheck) {
						counts.Set(toCheck, counts.Get(toCheck)-1)
					}
				}
				continue
			}
			newLocs = append(newLocs, loc)
		}
		total += removed
		locs = newLocs
	}
	return
}
