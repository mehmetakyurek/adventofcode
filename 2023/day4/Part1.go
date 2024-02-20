package main

import (
	"slices"
	"strconv"
	"strings"
)

func Part1() int {
	sum := 0
	for _, line := range lines {
		if len(line) < 10 {
			continue
		}
		line = line[10:]
		line = strings.Replace(line, "  ", " ", -1)
		s := strings.Split(line, " | ")

		winning, mine := strArrayToInt(strings.Split(strings.TrimSpace(s[0]), " ")), strArrayToInt(strings.Split(strings.TrimSpace(s[1]), " "))
		sum += findIntersectionCount(winning, mine)

	}
	return sum
}

func findIntersectionCount(a []int, b []int) int {
	count := 0
	for _, be := range b {
		if slices.Contains(a, be) {
			if count == 0 {
				count = 1
			} else {
				count *= 2
			}
		}
	}
	return count
}

func strArrayToInt(str []string) []int {
	res := make([]int, len(str))

	for s, st := range str {
		if n, err := strconv.Atoi(strings.TrimSpace(st)); err != nil {

		} else {
			res[s] = n
		}
	}
	return res
}
