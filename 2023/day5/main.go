package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	maps := strings.Split(string(input), "\n\n")
	seeds := getSeeds(maps[0])
	fmt.Printf("Part1(string(input)): %v\n", Part1(*seeds, &maps))
	fmt.Printf("Part2: %v\n", Part2(*seeds, &maps))
}

func Part2(seeds []int, maps *[]string) int {
	closest := -1
	for i := 0; i < len(seeds); i++ {
		for n := range seeds[i+1] {
			location := readMap(n, maps)
			if closest < 0 || location < closest {
				closest = location
			}
		}
		i++
	}
	return closest
}

func Part1(seeds []int, maps *[]string) int {
	closest := -1
	for _, seed := range seeds {
		location := readMap(seed, maps)
		if closest < 0 || location < closest {
			closest = location
		}
	}
	return closest
}

func readMap(seed int, maps *[]string) int {
	soil := findCode(seed, decodeMap((*maps)[1]))
	fertilizer := findCode(soil, decodeMap((*maps)[2]))
	water := findCode(fertilizer, decodeMap((*maps)[3]))
	light := findCode(water, decodeMap((*maps)[4]))
	temperature := findCode(light, decodeMap((*maps)[5]))
	humidity := findCode(temperature, decodeMap((*maps)[6]))
	location := findCode(humidity, decodeMap((*maps)[7]))
	return location
}

func getSeeds(text string) *[]int {
	seeds := strings.Split(strings.Split(text, ": ")[1], " ")
	var res []int
	for _, seed := range seeds {
		i, err := strconv.Atoi(seed)
		if err != nil {
			log.Fatal(err)
		}
		res = append(res, i)
	}
	return &res
}

func decodeMap(text string) *[][]int {
	var res [][]int
	maps := strings.Split(text, "\n")[1:]
	for _, m := range maps {
		if len(m) < 1 {
			continue
		}
		numbers := strings.Split(m, " ")
		res = append(res, make([]int, 3))
		for number := range numbers {
			n, err := strconv.Atoi(numbers[number])
			if err != nil {
				log.Fatal(err)
			}
			res[len(res)-1][number] = n
		}

	}
	return &res
}
func findCode(code int, codeMap *[][]int) int {
	for _, m := range *codeMap {
		if code >= m[1] && m[1]+m[2] >= code {
			return m[0] + (code - m[1])
		}
	}
	return code
}
