package parts

import (
	"regexp"
	"strconv"
)

const (
	AND = 1
	OR  = 2
	XOR = 3
)

var opMap = map[string]int{
	"AND": AND,
	"OR":  OR,
	"XOR": XOR,
}

type gate struct {
	in1 string
	in2 string
	op  int
	out string
}

var wireRegex = regexp.MustCompile("(.+): (0|1)")
var lineRegex = regexp.MustCompile("(.+) (AND|OR|XOR) (.+) -> (.+)")

func Process1(lines []string) (total int) {
	wires, i := parseWires(lines)
	gates := parseGates(lines[i+1:])
	for len(gates) > 0 {
		var unsolved []*gate
		for _, g := range gates {
			val1, ok1 := wires[g.in1]
			val2, ok2 := wires[g.in2]
			if !ok1 || !ok2 {
				unsolved = append(unsolved, g)
				continue
			}
			out := eval(val1, val2, g.op)
			wires[g.out] = out
			if g.out[0] == 'z' {
				bit, _ := strconv.Atoi(g.out[1:])
				mask := out << bit
				total |= mask
			}
		}
		gates = unsolved
	}
	return
}

func Process2(lines []string) (total string) {
	return
}

func parseWires(lines []string) (map[string]int, int) {
	ret := make(map[string]int)
	for i, line := range lines {
		wire := wireRegex.FindStringSubmatch(line)
		if wire == nil {
			return ret, i
		}
		val, _ := strconv.Atoi(wire[2])
		ret[wire[1]] = val
		continue
	}
	return ret, len(lines)
}

func parseGates(lines []string) (gateSlice []*gate) {
	for _, line := range lines {
		vals := lineRegex.FindStringSubmatch(line)
		if vals == nil {
			continue
		}
		newGate := gate{
			in1: vals[1],
			in2: vals[3],
			op:  opMap[vals[2]],
			out: vals[4],
		}
		gateSlice = append(gateSlice, &newGate)
	}
	return
}

func eval(in1, in2, op int) int {
	switch op {
	case AND:
		return in1 & in2
	case OR:
		return in1 | in2
	case XOR:
		return in1 ^ in2
	}
	return -1
}
