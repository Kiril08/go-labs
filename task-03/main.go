package main

import (
	"fmt"
)

func main() {
	goal_01()
	goal_02()
}

/*
1. Создайте массив из 4 чисел с плавающей запятой (например, [4]float64).
Напишите цикл, который вычисляет сумму всех элементов массива
, а затем выводит среднее арифметическое.
*/

func goal_01() {
	var arr = [4]float64{6.5, 16.8, 8.2, 9.3}
	var sum float64 = 0
	for i := 0; i < len(arr); i++ {
		sum = arr[i] + sum
	}
	fmt.Printf("Sum of elements = %.1f\n", sum)
	fmt.Printf("Average = %.1f\n", float64(sum)/float64(len(arr)))
}

/*
2. Нахождение максимального значения
Определите массив целых чисел с 10 случайными элементами (определите его самостоятельно).
Используйте цикл for для поиска наибольшего числа в этом массиве и выведите его на консоль.
*/

func goal_02() {
	var arr1 = [10]int{50, 67, 30, 10, 3, 54, 105, 56, 75, 25}
	var max = arr1[0]
	for _, v := range arr1 {
		if v > max {
			max = v
		}
	}
	fmt.Println("Max in array =", max)
}
