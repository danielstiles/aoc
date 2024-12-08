package safety

import (
	"log/slog"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type rule struct {
	first int
	last  int
}

var ruleRegex = regexp.MustCompile("(\\d+)\\|(\\d+)")

func Process1(lines []string) (total int) {
	rules, count := parseRules(lines)
	for _, line := range lines[count+1:] {
		update := parseUpdate(line)
		pass := true
		for _, rule := range rules {
			if !rule.Validate(update) {
				pass = false
				break
			}
		}
		if pass {
			total += update[(len(update)-1)/2]
		}
	}
	return
}

func Process2(lines []string) (total int) {
	rules, count := parseRules(lines)
	for _, line := range lines[count+1:] {
		update := parseUpdate(line)
		pass := true
		for _, rule := range rules {
			if !rule.Validate(update) {
				pass = false
				break
			}
		}
		if pass {
			continue
		}
		ordering := getOrdering(update, rules)
		total += ordering[(len(ordering)-1)/2]
	}
	return
}

func parseRules(lines []string) (rules []rule, count int) {
	for i, line := range lines {
		match := ruleRegex.FindStringSubmatch(line)
		if match == nil {
			count = i
			return
		}
		first, _ := strconv.Atoi(match[1])
		last, _ := strconv.Atoi(match[2])
		rules = append(rules, rule{first: first, last: last})
	}
	return
}

func parseUpdate(line string) (update []int) {
	nums := strings.Split(line, ",")
	for _, num := range nums {
		updatePart, _ := strconv.Atoi(num)
		update = append(update, updatePart)
	}
	return update
}

func (r rule) Validate(update []int) bool {
	firstI := slices.Index(update, r.first)
	lastI := slices.Index(update, r.last)
	return firstI == -1 || lastI == -1 || firstI < lastI
}

func getOrdering(update []int, rules []rule) (order []int) {
	after := make(map[int][]int)
	for _, u := range update {
		after[u] = []int{}
	}
	for _, rule := range rules {
		_, fFound := after[rule.first]
		_, lFound := after[rule.last]
		if fFound && lFound {
			after[rule.last] = append(after[rule.last], rule.first)
		}
	}
	for len(after) > 0 {
		new := -1
		for a, l := range after {
			if isFirst(l, order) {
				new = a
				break
			}
		}
		if new == -1 {
			slog.Error("inf loop", slog.Any("order", order), slog.Any("after", after))
			break
		}
		order = append(order, new)
		delete(after, new)
	}
	return
}

func isFirst(l, order []int) bool {
	for _, b := range l {
		if !slices.Contains(order, b) {
			return false
		}
	}
	return true
}
