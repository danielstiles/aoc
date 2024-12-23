package parts

import (
	"maps"
	"regexp"
	"slices"
)

var nodeRegex = regexp.MustCompile("[^-]+")

func Process1(lines []string) (total int) {
	edges := make(map[string][]string)
	for _, line := range lines {
		nodes := nodeRegex.FindAllString(line, -1)
		edges[nodes[0]] = append(edges[nodes[0]], nodes[1])
		edges[nodes[1]] = append(edges[nodes[1]], nodes[0])
	}
	for n, e := range edges {
		if n[0] != 't' {
			continue
		}
		for _, n2 := range e {
			if n2[0] == 't' && n2 > n {
				continue
			}
			for _, n3 := range edges[n2] {
				if n3 > n2 || (n3[0] == 't' && n3 > n) {
					continue
				}
				if slices.Contains(edges[n3], n) {
					total += 1
				}
			}
		}
	}
	return
}

func Process2(lines []string) (total string) {
	edges := make(map[string][]string)
	for _, line := range lines {
		nodes := nodeRegex.FindAllString(line, -1)
		edges[nodes[0]] = append(edges[nodes[0]], nodes[1])
		edges[nodes[1]] = append(edges[nodes[1]], nodes[0])
	}
	nodes := slices.Sorted(maps.Keys(edges))
	parties := make(map[string][]string)
	for _, n := range nodes {
		parties[n] = edges[n]
	}
	canContinue := true
	for canContinue {
		canContinue = false
		nextParties := make(map[string][]string)
		for ns, p := range parties {
			for _, n := range p {
				if n < ns {
					continue
				}
				nextParty := ns + "," + n
				nextParties[nextParty] = []string{}
				for _, n2 := range p {
					if n2 < n {
						continue
					}
					if slices.Contains(edges[n], n2) {
						canContinue = true
						nextParties[nextParty] = append(nextParties[nextParty], n2)
					}
				}
			}
		}
		parties = nextParties
	}
	for ns := range parties {
		if len(ns) > len(total) {
			total = ns
		}
	}
	return
}
