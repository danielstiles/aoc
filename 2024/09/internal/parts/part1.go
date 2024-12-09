package parts

import (
	"strconv"
)

func Process1(lines []string) (total int) {
	line := lines[0]
	layout := parse(line)
	filled := defrag(layout)
	total = checksum(filled)
	return
}

func parse(line string) (layout []*int) {
	fileno := 0
	isFile := true
	for i := range line {
		num, _ := strconv.Atoi(line[i : i+1])
		for j := 0; j < num; j++ {
			if isFile {
				file := fileno
				layout = append(layout, &file)
			} else {
				layout = append(layout, nil)
			}
		}
		if isFile {
			fileno += 1
		}
		isFile = !isFile
	}
	return
}

func defrag(layout []*int) (new []*int) {
	end := len(layout) - 1
	for i, val := range layout {
		if i > end {
			return
		}
		if val != nil {
			new = append(new, val)
			continue
		}
		for ; end > i && layout[end] == nil; end -= 1 {
		}
		if i != end {
			new = append(new, layout[end])
			end -= 1
		}
	}
	return
}

func checksum(layout []*int) (total int) {
	for i, val := range layout {
		if val != nil {
			total += i * (*val)
		}
	}
	return
}

func printLayout(layout []*int) (res string) {
	for _, val := range layout {
		if val == nil {
			res += "."
		} else {
			res += strconv.Itoa(*val)
		}
	}
	return
}
