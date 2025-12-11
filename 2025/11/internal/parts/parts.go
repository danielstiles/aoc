package parts

import (
	"strings"
)

func Process1(lines []string) (total int) {
	gates := loadGates(lines)
	total = countRoutes(gates, "you", "out")
	return
}

func Process2(lines []string) (total int) {
	gates := loadGates(lines)
	total = countRoutes2(gates, "svr", "out", []string{"dac", "fft"})
	return
}

func loadGates(lines []string) map[string][]string {
	gates := make(map[string][]string)
	for _, line := range lines {
		parts := strings.Split(line, " ")
		parts[0] = parts[0][:len(parts[0])-1]
		gates[parts[0]] = parts[1:]
	}
	return gates
}

func countRoutes(gates map[string][]string, start, end string) int {
	if start == end {
		return 1
	}
	dests, ok := gates[start]
	if !ok {
		return 0
	}
	total := 0
	for _, dest := range dests {
		total += countRoutes(gates, dest, end)
	}
	return total
}

func countRoutes2(gates map[string][]string, start, end string, toFind []string) (total int) {
	cache := make(map[string]map[int]int)
	routes := findRoutes(gates, start, end, toFind, cache)
	for found, count := range routes {
		foundAll := true
		for i := range toFind {
			if found&(1<<i) == 0 {
				foundAll = false
				break
			}
		}
		if foundAll {
			total += count
			break
		}
	}
	return
}

func findRoutes(gates map[string][]string, start, end string, toFind []string, cache map[string]map[int]int) map[int]int {
	found := 0
	toRet := make(map[int]int)
	defer func() {
		cache[start] = toRet
	}()
	for i, val := range toFind {
		if start == val {
			found |= (1 << i)
		}
	}
	if start == end {
		toRet[found] = 1
		return toRet
	}
	dests, ok := gates[start]
	if !ok {
		return toRet
	}
	for _, dest := range dests {
		var routes map[int]int
		if cached, ok := cache[dest]; ok {
			routes = cached
		} else {
			routes = findRoutes(gates, dest, end, toFind, cache)
		}
		for destFound, count := range routes {
			destFound |= found
			toRet[destFound] += count
		}
	}
	return toRet
}
