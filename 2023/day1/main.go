package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)


var numbers = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
func main() {

	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := string(file)
	lines := strings.Split(input, "\n")
	sum := 0
	for _, line := range lines {
		first, last := getFirstNLastStrNum(line)
		fmt.Printf("%v %v %v\n", line, first, last)
		sum += (first * 10) + last;
	}
	println(sum)
	println(len(lines))

}
type StrSearchResult struct {
	index int
	num int
}
func getFirstNLastStrNum(text string) (fisrt int, last int) {
	return findFirstStrNum(text), findLastStrNum(text);
}

func findFirstStrNum(text string) int {
	if len(text) ==0 {
		return 0
	}
	num, index := -1, -1
	for numInt,numStr := range numbers {
		if i := strings.Index(text, numStr); (i > -1 && i < index) || index < 0 {
			index = i
			num = numInt 
		}
	}
	
	for i := 0; (index < 0 && i < len(text)) || i < index; i++ {
		if unicode.IsDigit(rune(text[i])) {
			return int(text[i] - '0')
		}

	}
	return num
}
func findLastStrNum(text string) int {
	if len(text) ==0 {
		return 0
	}
	num, index := -1, -1
	for numInt,numStr := range numbers {
		if i := strings.LastIndex(text, numStr); (i> -1 && i > index) || index < 0 {
			index = i
			num = numInt 
		}
	}
	for i:=len(text) - 1; i > index; i-- {
		if unicode.IsDigit(rune(text[i])) {
			return int(text[i] - '0')
		}

	}
	return num
}
