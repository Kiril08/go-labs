package main

import (
	"fmt"
)

/*
	1. Create 4 variables of different types (string, int64, bool, float32)
	2. Assign values to these variables
	3. Print the following sentence using these variables:
		"My name is __, I am __ years old and it's __ that I can drive a car, my pet weights __ kilograms"

	Use this: https://pkg.go.dev/fmt#hdr-Printing
*/

func main() {
	var name string = "Kiril"
	var age int64 = 20
	var legal bool = true
	var weight float32 = 7.34

	fmt.Printf("My name is %s, I am %d years old and it's %t that I can drive a car, my pet weights %.2git f kilograms\n",
		name, age, legal, weight)
}
