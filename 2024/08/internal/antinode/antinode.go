package antinode

type point struct {
	r int
	c int
}

func Process1(lines []string) (total int) {
	size, nodes := parse(lines)
	antinodes := make(map[int]struct{})
	for _, n := range nodes {
		new := getAntinodes(n, size, false)
		for k, v := range new {
			antinodes[k] = v
		}
	}
	total = len(antinodes)
	return
}

func Process2(lines []string) (total int) {
	size, nodes := parse(lines)
	antinodes := make(map[int]struct{})
	for _, n := range nodes {
		new := getAntinodes(n, size, true)
		for k, v := range new {
			antinodes[k] = v
		}
	}
	total = len(antinodes)
	return
}

func parse(lines []string) (size point, nodes map[rune][]point) {
	if len(lines) == 0 {
		return
	}
	size.r = len(lines)
	size.c = len(lines[0])
	nodes = make(map[rune][]point)
	for r, line := range lines {
		for c, val := range line {
			if val != '.' {
				nodes[val] = append(nodes[val], point{r: r, c: c})
			}
		}
	}
	return
}

func getAntinodes(nodes []point, size point, resonance bool) (antinodes map[int]struct{}) {
	antinodes = make(map[int]struct{})
	for i := range nodes {
		for j := i + 1; j < len(nodes); j++ {
			diff := nodes[i].sub(nodes[j])
			if !resonance {
				p := nodes[i].add(diff)
				if p.checkBounds(size) {
					loc := p.r*size.c + p.c
					antinodes[loc] = struct{}{}
				}
				p = nodes[j].sub(diff)
				if p.checkBounds(size) {
					loc := p.r*size.c + p.c
					antinodes[loc] = struct{}{}
				}
			} else {
				for p := nodes[i]; p.checkBounds(size); p = p.add(diff) {
					loc := p.r*size.c + p.c
					antinodes[loc] = struct{}{}
				}
				for p := nodes[j]; p.checkBounds(size); p = p.sub(diff) {
					loc := p.r*size.c + p.c
					antinodes[loc] = struct{}{}
				}
			}
		}
	}
	return
}

func (p point) checkBounds(size point) bool {
	if p.r < 0 || p.r >= size.r {
		return false
	}
	if p.c < 0 || p.c >= size.c {
		return false
	}
	return true
}

func (p point) add(q point) (new point) {
	new.r = p.r + q.r
	new.c = p.c + q.c
	return
}

func (p point) sub(q point) (new point) {
	new.r = p.r - q.r
	new.c = p.c - q.c
	return
}
