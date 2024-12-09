package parts

import (
	"log/slog"
	"strconv"
)

type disc struct {
	prev *disc
	next *disc
	f    file
}

type file struct {
	fileno int
	size   int
}

func Process2(lines []string) (total int) {
	line := lines[0]
	layout := parseToDisc(line)
	defragDisc(layout)
	total = checksumDisc(layout)
	return
}

func parseToDisc(line string) (layout *disc) {
	num, _ := strconv.Atoi(line[0:1])
	layout = &disc{
		f: file{
			fileno: 0,
			size:   num,
		},
	}
	curr := layout
	fileno := 1
	isFile := false
	for i := range line {
		if i == 0 {
			continue
		}
		num, _ := strconv.Atoi(line[i : i+1])
		newFile := file{
			fileno: fileno,
			size:   num,
		}
		if isFile {
			fileno += 1
		} else {
			newFile.fileno = -1
		}
		isFile = !isFile
		if num == 0 {
			continue
		}
		next := &disc{
			prev: curr,
			f:    newFile,
		}
		curr.next = next
		curr = next
	}
	curr.next = layout
	layout.prev = curr
	return
}

func defragDisc(layout *disc) {
	for curr := layout.prev; curr != layout; curr = curr.prev {
		if curr.f.fileno == -1 {
			continue
		}
		for check := layout; check != curr; check = check.next {
			if check.f.fileno != -1 || check.f.size < curr.f.size {
				continue
			}
			swapTo(curr, check)
			break
		}
	}
}

func swapTo(from, to *disc) {
	if to.f.size > from.f.size {
		insertAfter(to, file{
			fileno: -1,
			size:   to.f.size - from.f.size,
		})
	}
	to.f.fileno = from.f.fileno
	to.f.size = from.f.size
	from.f.fileno = -1
	merge(from)
}

func insertAfter(curr *disc, new file) {
	next := &disc{
		prev: curr,
		next: curr.next,
		f:    new,
	}
	curr.next.prev = next
	curr.next = next
}

func merge(curr *disc) {
	if curr.prev.f.fileno == curr.f.fileno {
		curr.f.size += curr.prev.f.size
		curr.prev.prev.next = curr
		curr.prev = curr.prev.prev
	}
	if curr.next.f.fileno == curr.f.fileno {
		curr.f.size += curr.next.f.size
		curr.next.next.prev = curr
		curr.next = curr.next.next
	}
}

func checksumDisc(layout *disc) (total int) {
	pos := layout.f.size
	for curr := layout.next; curr != layout; curr = curr.next {
		if curr.f.fileno == -1 {
			pos += curr.f.size
			continue
		}
		for i := 0; i < curr.f.size; i += 1 {
			total += (pos + i) * curr.f.fileno
		}
		pos += curr.f.size
	}
	return
}

func printDisc(layout *disc) {
	str := "|" + strconv.Itoa(layout.f.fileno) + "*" + strconv.Itoa(layout.f.size)
	for curr := layout.next; curr != layout; curr = curr.next {
		if curr.f.fileno == -1 {
			str += "|.*" + strconv.Itoa(curr.f.size)
		} else {
			str += "|" + strconv.Itoa(curr.f.fileno) + "*" + strconv.Itoa(curr.f.size)
		}
	}
	str += "|"
	slog.Info(str)
}
