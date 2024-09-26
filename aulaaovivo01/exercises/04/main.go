package main

import (
	"fmt"
)

// 11111111 = 2 elevado a 8 = 256
// int8 = -127 => +127
// uint8 = 0 255

func main() {
	var lastName string = "Smith"
	var age int = 35
	var boolean1 bool
	boolean := false
	var salary float64 = 45857.90
	var firstName string = "Mary"

	fmt.Println(
		firstName,
		lastName,
		age,
		boolean,
		boolean1,
		salary,
	)

	// Alphanumeric to integer
	// strconv.Atoi
	// Integer to alphanumeric
	// strconv.Itoa
}
