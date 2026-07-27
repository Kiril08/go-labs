package main

import (
	"fmt"
)


/*the condition for the second task
   [1,2,6,8]
   [0,5,6,10]
*/

// 1) Missing numbers
func main() {
	// data for the first task
	// replace it intead of the other data into the func main()
    //                  |
    //                  |
	//				   \|/
	/*var arr = []int{10,9,6,4,8,3,7,5,1,0}
	sorted := BubbleSort(arr)
	missing := missingNumber(sorted)
	fmt.Printf("Sorted array: %v\nMissing number: %d\n", sorted, missing)
	*/
	var arr1 = []int{1,2,6,8}
	var arr2 = []int{0,5,6,10}
	var f = intersection(arr1,arr2)
	fmt.Println(f)
}

func missingNumber(nums []int) int {
	if nums[0] != 0 {
		return 0
	}
	
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1]-nums[i] != 1 {
			return nums[i] + 1 
		}
	}
	return len(nums)
}

func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

// 2) Intersection of elements of the arrays
func intersection(nums1 []int, nums2 []int) []int {
	for i:=0; i<len(nums1); i++{
		for j:=i; j==i; {
			if nums1[i] == nums2[j] {
				return nums1
			}
		}
	}
	return nums1
}






