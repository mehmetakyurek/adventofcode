package main

import (
	"fmt"
	"testing"

)


func TestIsDigit(t *testing.T) {
	for _,digit := range digits {
		if !isDigit(digit) {
			fmt.Printf("%q %v", digit, isDigit(digit))
			t.Fail()
		}
	}
}
