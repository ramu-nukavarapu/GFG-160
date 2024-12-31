package main

import "fmt"

func main() {
	arr := []int{-1, 0, 2, -3, 1, 3}
	fmt.Println(smallestPositiveMissingOptimal(arr))

	arr = []int{1, 11, 2, 8, 9, 8, 9, 14, 8, 17, 0, 2, -5, 9, -16, 8, 9, 16, 0}
	fmt.Println(smallestPositiveMissingOptimal(arr))

	arr = []int{2, -3, 4, 1, 1, 7}
	fmt.Println(smallestPositiveMissingOptimal(arr))

	arr = []int{5, 3, 2, 5, 1}
	fmt.Println(smallestPositiveMissingOptimal(arr))

	arr = []int{1, 2, 3, 4, 5}
	fmt.Println(smallestPositiveMissingOptimal(arr))
}

func smallestPositiveMissingBruteforce(arr []int) int {
	// sort the array first
	// start iteration if the element is negative continue the iteration and move to next element
	// if the difference b/w element and the next is zero or one continue the iteration
	// else return the current element + 1
	return 0
}

func smallestPositiveMissingOptimal(arr []int) int {
	start, end, n := 0, 1, len(arr)
	var temp int

	for end < n {
		if arr[start] == start+1 {
			start++
			end = start + 1
		} else if arr[end] == start+1 {
			temp = arr[start]
			arr[start] = arr[end]
			arr[end] = temp
			start++
			end = start + 1
		} else {
			end++
		}
	}

	for i := 0; i < n; i++ {
		if arr[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}
