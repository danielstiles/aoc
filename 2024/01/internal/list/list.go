package list

import (
	"slices"
	"strconv"
	"strings"
)

func Process1(lines []string) (total int) {
	l1, l2 := getLists(lines)
	slices.Sort(l1)
	slices.Sort(l2)
	for i := range l1 {
		diff := l1[i] - l2[i]
		if diff > 0 {
			total += diff
		} else {
			total -= diff
		}
	}
	return
}

func Process2(lines []string) (total int) {
	l1, l2 := getLists(lines)
	slices.Sort(l2)
	for _, num := range l1 {
		i := slices.Index(l2, num)
		if i == -1 {
			continue
		}
		count := 1
		for ; i+count < len(l2) && l2[i+count] == num; count += 1 {
		}
		total += num * count
	}
	return
}

func getLists(lines []string) (list1, list2 []int) {
	for _, line := range lines {
		parts := strings.Split(line, " ")
		n1, _ := strconv.Atoi(parts[0])
		n2, _ := strconv.Atoi(parts[len(parts)-1])
		list1 = append(list1, n1)
		list2 = append(list2, n2)
	}
	return
}
