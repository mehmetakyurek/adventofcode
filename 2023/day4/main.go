package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var lines []string

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines = strings.Split(string(file), "\n")
	file = nil
	fmt.Printf("Part1(): %v\n", Part1())
	fmt.Printf("Part2(): %v\n", Part2())
}
