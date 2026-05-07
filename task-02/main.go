package main

import (
	"fmt"
)

/*
массив 10 чисел, заполнить его знач от 1 до 10, вывести на экран сумму чисел массива
*/

func main() {
	var arr = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var sum = 0
	for i := 0; i < len(arr); i++ {
		sum = arr[i] + sum
	}

	fmt.Println(sum)
	fmt.Println(arr[len(arr)-1])
}
