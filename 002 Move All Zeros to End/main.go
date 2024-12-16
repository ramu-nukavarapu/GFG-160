package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 0, 3, 2, 0, 5}
	fmt.Println(arr)
	moveZerosOptimal(arr)
	fmt.Println(arr)

	arr = []int{5, 2, 1, 0, 0, 0}
	fmt.Println(arr)
	moveZerosOptimal(arr)
	fmt.Println(arr)

	arr = []int{1}
	fmt.Println(arr)
	moveZerosOptimal(arr)
	fmt.Println(arr)

	arr = []int{0, 0}
	fmt.Println(arr)
	moveZerosOptimal(arr)
	fmt.Println(arr)

	arr = []int{0, 0, 0, 1, 2, 3}
	fmt.Println(arr)
	moveZerosOptimal(arr)
	fmt.Println(arr)
}

func moveZerosBruteForce(arr []int) {
	// sort in decreasing order [ only applicable when order of non-zero elements are not important ]
}

func moveZerosOptimal(arr []int) {
	var temp int
	first := 0
	second := 1

	if len(arr) > 1 {
		for second < len(arr) {
			if arr[first] == 0 {
				temp = arr[second]
				arr[second] = arr[first]
				arr[first] = temp
				second++

				if arr[first] != 0 {
					first++
				}
			} else {
				first++
				second++
			}
		}
	}
}
