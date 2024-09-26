package main

import "fmt"

/*
   batata frita

   batataFrita => camelCase
   BatataFrita => PascalCase
   batata_frita => snake_case
   batata-frita => kebab-case
*/

func main() {
	var firstName string
	var lastName string
	var age int
	lastName = "Da Silva"
	var driverLicense = true
	// var person height int
	// childsNumber := 2

	fmt.Println(firstName, lastName, driverLicense, age)
}
