package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var arr = make([]int, 100)
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(10)
	}

	fmt.Println(arr)

	var c = 0

	for i := 0; i < len(arr); i++ {
		if arr[i] == 5 {
			c++
		}

	}
	fmt.Println(c)
}
