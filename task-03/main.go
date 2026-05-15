package main

import (
	"fmt"
	"math/rand"
)

func main() {
	m()
}

func goal_01() {
	var arr = [4]float64{6.5, 16.8, 8.2, 9.3}
	var sum float64 = 0
	for i := 0; i < len(arr); i++ {
		sum = arr[i] + sum
	}
	fmt.Printf("Sum of elements = %.1f\n", sum)
	fmt.Printf("Average = %.1f\n", float64(sum)/float64(len(arr)))
}

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

func goal_03() {
	arr2 := [30]int{}
	for i := 0; i < len(arr2); i++ {
		arr2[i] = rand.Intn(30)
	}
	fmt.Println(arr2)

	for i, v := range arr2 {
		if v%2 == 1 {
			arr2[i] = 0
		} else {
			arr2[i] = 1
		}
	}
	fmt.Println(arr2)

}

func goal_04() {
	var sum int = 0
	arr3 := [8]int{52, 23, 90, 45, 20, 94, 41, 39}
	fmt.Println(arr3)
	for i := 0; i < len(arr3); i++ {
		if i%2 == 0 {
			sum += arr3[i]
		}
	}
	fmt.Println(sum)
}

func goal_05() {
	arr4 := [5]int{1, 2, 3, 4, 5}
	arr5 := [5]int{}
	for i, j := 0, len(arr4)-1; i < len(arr4) && j >= 0; i++ {
		fmt.Println(j)
		arr5[j] = arr4[i]
		j--
	}
	fmt.Println(arr5)
}

/*
Дан массив из 6 целых чисел. Напишите цикл, который сдвигает все элементы на одну позицию влево.
При этом самый первый элемент должен переместиться на место самого последнего (например, массив [10, 20, 30, 40]
должен превратиться в [20, 30, 40, 10]).Что тренируем: работу с относительными индексами (arr[i] = arr[i+1]) и
контроль границ массива, чтобы не вызвать ошибку panic: index out of range.
*/

func goal_07() {
	arr := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	var s = moveArray(arr)
	fmt.Println(s)
}

func moveArray(arr []int) []int {
	fmt.Println("До:    ", arr) // [60 10 20 30 40 50]

	// ── Шаг 1: сохраняем первый элемент ────────────────
	temp := arr[len(arr)-1] // temp = 60

	// ── Шаг 2: сдвигаем всё влево ──────────────────────
	// i < len(arr)-1, а не len(arr)! иначе panic при i=5
	for i := len(arr) - 1; i > 0; i-- {
		arr[i] = arr[i-1] // берём значение правого соседа
	}
	// Достаём из «кармана» на последнее место
	arr[0] = temp                  // arr[5] = 10
	fmt.Println("После:    ", arr) // [10 20 30 40 50 60]

	return arr
}

/*
Зеркальный переворот (Реверс массива)Создайте массив из 6 целых чисел.
Напишите цикл, который полностью развернет массив задом наперед прямо «на месте» (in-place),
не используя второй вспомогательный массив (например, [1, 2, 3, 4, 5, 6] превращается в [6, 5, 4, 3, 2, 1]).
Что тренируем: одновременную работу с двумя индексами (начало и конец)
и понимание того, что цикл должен дойти только до середины массива (len(arr) / 2), иначе элементы поменяются местами обратно.
*/

func f() {
	var arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	fmt.Println(arr)
	// we need: [10, 9, 8, 7, 6, 5, 4, 3 ,2 ,1]
	// we'll get

	// last = len(arr)-1
	// next = i+1
	// prev = i-1
	// first = 0

	for i, j := 0, len(arr)-1; i != j; {
		var temp = arr[i]
		arr[i] = arr[j]
		arr[j] = temp

		i++
		j--
	}
	/*
	   we'll get

	      Iterations:

	      1) arr[0] = arr[10  (11-1)]
	      arr[10] = arr(0)                  [11,2,3,4,5,6,7,8,9,10,1]

	   i++
	   j--

	      2) arr[1] = arr[9]
	      arr[9] = arr(1)                  [11,10,3,4,5,6,7,8,9,2,1]

	      3) arr[2] = arr[8]
	      arr[8] = arr(2)                  [11,10,9,4,5,6,7,8,3,2,1]

	      4) arr[3] = arr[7]
	      arr[7] = arr(3)                  [11,10,9,8,5,6,7,4,3,2,1]

	      5) arr[4] = arr[6]
	      arr[6] = arr(4)                  [11,10,9,8,7,6,5,4,3,2,1]

	      6) arr[5] = arr[5]
	      arr[5] = arr(5)          the cycle stops because our condition doesn't work       [11,10,9,8,7,6,5,4,3,2,1]  end
	*/
	fmt.Println(arr)
}

func m() {
	var arr = [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println("До: ", arr)
	// print it for the comparing with result
	// we need to change the places of indexes of the neighbouring elements
	// it takes [1 , 2 , 3 , 4 , 5 , 6] , [2 , 1 , 4 , 3 , 6 , 5]
	//           i  i+1  i  i+1  i  i+1
	//           step+2  step+2    no steps
	for i := 0; i < len(arr)-1; i += 2 {
		// create temporary variable for avoid overwriting of the elements
		var temp = arr[i]
		arr[i] = arr[i+1] // начальный = следующему, а следующий = нашей временной переменной
		arr[i+1] = temp
	}
	fmt.Println("После: ", arr)
}
