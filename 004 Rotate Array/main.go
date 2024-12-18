package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	d := 2
	fmt.Println(arr)
	rotateArrayOptimal(arr, d)
	fmt.Println(arr)

	arr = []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	d = 3
	fmt.Println(arr)
	rotateArrayOptimal(arr, d)
	fmt.Println(arr)
	arr = []int{7, 3, 9, 1}
	d = 9
	fmt.Println(arr)
	rotateArrayOptimal(arr, d)
	fmt.Println(arr)

}

func rotateArrayBruteforce(arr []int, d int) {
	// Iterate d times the array
	// add the first element to the temp variable then swap all elements to the left and add the value in temp at last
}

func rotateArrayOptimal(nums []int, d int) {
	length := len(nums)
	d = d % length
	reverseArray(nums, 0, d-1)
	reverseArray(nums, d, length-1)
	reverseArray(nums, 0, length-1)
}

func reverseArray(nums []int, start, end int) {
	var temp int

	for start < end {
		temp = nums[start]
		nums[start] = nums[end]
		nums[end] = temp
		start++
		end--
	}
}
