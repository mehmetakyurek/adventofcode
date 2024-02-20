package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

var cards []int

func Part2() int {
	cardsCount := 0
	cards = make([]int, len(lines))

	for li := range lines {
		if len(lines[li]) > 0 {

			cards[li]++
			match := getmatchingNumber(li)
			for ci := 0; ci < cards[li]; ci++ {
				for n := li + 1; n <= li+match; n++ {
					cards[n]++
				}
			}
		}
	}

	for c := range cards {
		cardsCount += cards[c]
	}
	fmt.Printf("cards: %v\n", cards)
	return cardsCount
}

func getmatchingNumber(line int) int {
	w, m := decodeLine(line)
	count := 0
	for wi := range w {
		if slices.Index(m, w[wi]) >= 0 {
			count++
		}
	}
	return count
}

func decodeLine(line int) ([]int, []int) {
	i := strings.Index(lines[line], ":")
	if i >= 0 {
		l := strings.TrimSpace(lines[line][i:])
		winning := make([]int, 0)
		mine := make([]int, 0)
		currentlyWorking := &winning
		for c := 0; c < len(l); c++ {
			if l[c] == '|' {
				currentlyWorking = &mine
				continue
			}
			if isDigit(l[c]) {
				n := ""
				for c < len(l) && isDigit(l[c]) {
					n += string(l[c])
					c++
				}
				if num, err := strconv.Atoi(n); err != nil {
					panic(err)
				} else {
					*currentlyWorking = append(*currentlyWorking, num)
				}
			}
		}
		return winning, mine
	}
	return make([]int, 0), make([]int, 0)
}

func isDigit(c byte) bool {
	return c <= '9' && c >= '0'
}
