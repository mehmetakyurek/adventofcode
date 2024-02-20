package main

import (
	"fmt"
	"testing"
)






func TestGetStrNum (t *testing.T) {
	getFirstStrNum("fivesixthreeptcqjnkzgdfgzspmlvmmhn3")
}
func TestFindLastStrNum (t *testing.T) {
	if findFirstStrNum("fivesixthreeptcqjnkzgdfgzspmlvmmhn3") != 5 {
		fmt.Println(findFirstStrNum("fivesixthreeptcqjnkzgdfgzspmlvmmhn3"))
		t.Fail()
	}
}
