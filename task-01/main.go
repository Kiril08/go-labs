package main

import "fmt"

/*
	1. Create 4 variables of different types (string, int64, bool, float32)
	2. Assign values to these variables
	3. Print the following sentence using these variables:
		"My name is __, I am __ years old and it's __ that I can drive a car, my pet weights __ kilograms"

	Use this: https://pkg.go.dev/fmt#hdr-Printing
*/

func main() {
	var name string
	var age int64
	var legal bool
	var weight float32

	name = "Anna"
	age = 29
	legal = false
	weight = 70.12

	// Here goes your code
	fmt.Printf("My name is __, I am __ years old and it's __ that I can drive a car, my pet weights __ kilograms")
}
