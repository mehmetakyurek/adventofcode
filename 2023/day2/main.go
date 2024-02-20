package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const MaxRed = 12
const MaxBlue = 14
const MaxGreen = 13

var ids = []int{}
var sum = 0
func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	filter(string(file))

}
func filter(input string) {
	powerSum := 0
	lines := strings.Split(input, "\n")
	for i := range lines {
		lines[i] = strings.Replace(lines[i], "Game "+fmt.Sprint(i+1)+":", "", -1)
		games := strings.Split(lines[i], ";")
		red, green, blue := 0,0,0;
		for g := range games {
			cubes := strings.Split(games[g], ",")
			for c := range cubes {
				num, color := getNum(cubes[c]), getColor(cubes[c])
				if color == "green" && num > green {
					green = num
				}
				if color == "blue" && num > blue {
					blue = num
				}
				if color == "red" && num > red {
					red = num
				}
			}
		}
		fmt.Printf("red: %v %v %v\n", red, green, blue);
		power := red * green * blue
		powerSum += power;
		fmt.Printf("%v power: %v\n", i + 1 , power)

	}
	fmt.Println(powerSum)
}

var replacer = strings.NewReplacer(" green", "", " blue", "", " red", "", " ", "")

func getColor(cube string) string {
	if strings.Index(cube, "green") >=0 {return "green"}
	if strings.Index(cube, "blue") >=0 {return "blue"}
	if strings.Index(cube, "red") >=0 {return "red"}
	return ""
}

func getNum(cube string) int {
	num, err := strconv.Atoi(replacer.Replace(cube))
	if err != nil{
		return 0
	}
	return num
}
func CheckCube(cube string) bool {
	num := getNum(cube)
	if strings.Index(cube, "green") >= 0 && num > MaxGreen {
		return false
	}
	if strings.Index(cube, "blue") >= 0 && num > MaxBlue {
		return false
	}
	if strings.Index(cube, "red") >= 0 && num > MaxRed {
		return false
	}
	return true
}
func getPower(line string) int {
	return 0
}
