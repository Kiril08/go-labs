package main

import (
	"fmt"
)

func task_02_1() {
	for i := 10; i >= 1; i-- {
		fmt.Printf("%d\n", i)
	}
	fmt.Println("Blast off")
}

func main() {
	task_02_3()
}

func task_02_2() {
	var sum = 0
	for t := -1; t >= -10; t-- {
		sum = sum + t
	}
	fmt.Println(sum)
}

func task_02_3() {
	var c int = 5
	var f int = 0
	for i := 1; i < 11; i++ {
		f = i * c
		fmt.Printf("%d * %d = %d\n", c, i, f)
	}
}
