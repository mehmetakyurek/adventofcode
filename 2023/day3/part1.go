package main

import (
	"log"
	"strconv"
	"strings"
)

func Part1() int {
	for li, line := range lines {
		for c := 0; c < len(line); c++ {

			if isDigit(string(line[c])) {
				num := getNumber(li, c)
				c += len(num)
				adj := scanAdjacent(li, c-len(num), len(num))
				nRemoved := strings.Replace(adj, num, "", -1)
				counts := strings.IndexAny(nRemoved, symbols) >= 0
				if counts {
					if n, err := strconv.Atoi(num); err != nil {
						log.Fatal(err)
					} else {
						sum += n
					}
				}
			}

		}
	}
	return sum
}
