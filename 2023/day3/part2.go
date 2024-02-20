package main

import (
	"fmt"
	"log"
	"strconv"
)

func Part2() int {
	var sum = 0
	for li, line := range lines {
		for c, ch := range line {
			if ch == '*' {
				nums := make([]int, 0)

				if li > 0 {
					start, end := scanLineForStar(li-1, c)
					appendNumbers(lines[li-1][start:end], &nums)
				}

				start, end := scanLineForStar(li, c)
				appendNumbers(lines[li][start:end], &nums)
				if li < len(lines[li])-1 {
					start, end = scanLineForStar(li+1, c)
					appendNumbers(lines[li+1][start:end], &nums)
				}
				if len(nums) == 2 {
					sum += nums[0] * nums[1]
				}
			}
		}
	}
	return sum
}

func appendNumbers(input string, nums *[]int) {
	l := len(input)
	for i := 0; i < l; i++ {
		if isDigit(string(input[i])) {
			num := ""
			for i < l && isDigit(string(input[i])) {
				num += string(input[i])
				i++
			}
			if n, err := strconv.Atoi(num); err != nil {
				log.Fatal(err)
			} else {
				*nums = append(*nums, n)
			}
		}
	}
}

func printStartToEnd(line, start, end int) {
	for s := start; s < end; s++ {
		fmt.Print(string(lines[line][s]))
	}
	fmt.Print("\n")

}

func scanLineForStar(line, pos int) (start, end int) {
	l := pos
	for ; l > 0 && (l > pos-1 || (isDigit(string(lines[line][pos-1])) && isDigit(string(lines[line][l])))); l-- {
	}
	s := l
	for ; l < len(lines[line]) && (l < pos+1 || (isDigit(string(lines[line][pos+1])) && isDigit(string(lines[line][l])))); l++ {
	}
	return s, l
}
