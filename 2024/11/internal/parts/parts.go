package parts

import (
	"regexp"
	"strconv"
)

var numRegex = regexp.MustCompile("\\d+")

func Process1(lines []string) (total int) {
	stones := getNums(lines[0])
	for i := 0; i < 25; i += 1 {
		stones = iterate(stones)
	}
	total = len(stones)
	return
}

func Process2(lines []string) (total int) {
	stones := getNums(lines[0])
	total = calcIterations(stones, 75)
	return
}

func getNums(line string) (nums []int) {
	strs := numRegex.FindAllString(line, -1)
	for _, s := range strs {
		num, _ := strconv.Atoi(s)
		nums = append(nums, num)
	}
	return
}

func iterate(stones []int) (res []int) {
	for _, stone := range stones {
		if stone == 0 {
			res = append(res, 1)
		} else if stoneLen := intLen(stone); stoneLen%2 == 0 {
			num1, num2 := intSplit(stone, stoneLen/2)
			res = append(res, num1, num2)
		} else {
			res = append(res, stone*2024)
		}
	}
	return
}

func calcIterations(stones []int, depth int) (total int) {
	res := make(map[int]map[int]int)
	for _, stone := range stones {
		total += calcStoneTotal(stone, depth, res)
	}
	return
}

func calcStoneTotal(stone, depth int, res map[int]map[int]int) (total int) {
	if depth == 0 {
		return 1
	}
	if results, ok := res[stone]; ok {
		if total, ok = results[depth]; ok {
			return
		}
	} else {
		res[stone] = make(map[int]int)
	}
	if stone == 0 {
		total = calcStoneTotal(1, depth-1, res)
	} else if stoneLen := intLen(stone); stoneLen%2 == 0 {
		num1, num2 := intSplit(stone, stoneLen/2)
		total = calcStoneTotal(num1, depth-1, res)
		total += calcStoneTotal(num2, depth-1, res)
	} else {
		total = calcStoneTotal(stone*2024, depth-1, res)
	}
	res[stone][depth] = total
	return
}

func intLen(a int) (length int) {
	for ; a > 0; a /= 10 {
		length += 1
	}
	return
}

func pow(a, e int) (res int) {
	res = 1
	for ; e > 0; e -= 1 {
		res *= a
	}
	return
}

func intSplit(a, digits int) (front int, back int) {
	splitter := pow(10, digits)
	back = a % splitter
	front = a / splitter
	return
}
