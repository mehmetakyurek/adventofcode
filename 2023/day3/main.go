package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	nonSymbol = "0123456789."
	digits    = "0123456789"
	symbols   = "%*#&$@/=+-"
)

func read() string {
	if file, err := os.ReadFile("input.txt"); err != nil {
		log.Fatal(err)
	} else {
		return string(file)
	}
	return ""
}

var lines = []string{}
var sum = 0

func main() {
	file := read()
	lines = strings.Split(file, "\n")
	lines = lines[:(len(lines) - 1)]
	fmt.Printf("Part1: %v\n", Part1())
	fmt.Printf("Part2: %v\n", Part2())
}

func getNumber(line, pos int) string {
	c := pos
	num := ""
	for c < len(lines[line]) && isDigit(string(lines[line][c])) {
		num += string(lines[line][c])
		c++
	}
	return num
}

func scanAdjacent(line, pos, length int) string {
	res := ""
	sl := line - 1
	if sl < 0 {
		sl = 0
	}
	for ; sl < len(lines) && sl < line+2; sl++ {
		ps := pos - 1
		if ps < 0 {
			ps = 0
		}
		for ; ps < len(lines[sl]) && ps < pos+length+1; ps++ {
			res += string(lines[sl][ps])
		}
		res += "\n"
	}
	return res
}

func isSymbol(r rune) bool {
	return strings.IndexByte(nonSymbol, byte(r)) < 0
}

func isDigit(r string) bool {
	return strings.Index(digits, r) >= 0
}
