package main

import "fmt"

func main() {
	//sample inputs
	arr := []int{12, 35, 1, 10, 34, 1}
	fmt.Println(secondLargestOptimal(arr)) //34
	arr = []int{10, 5, 10}
	fmt.Println(secondLargestOptimal(arr)) //5
	arr = []int{10, 10, 10}
	fmt.Println(secondLargestOptimal(arr)) //-1
}

// brute force solution
func secondLargestBruteForce(arr []int) int {
	// answer will be added
	return -1
}

// optimal solution
func secondLargestOptimal(arr []int) int {
	firstMax, secondMax := -1, -1
	for i := 0; i < len(arr); i++ {
		if firstMax < arr[i] {
			secondMax = firstMax
			firstMax = arr[i]
		} else if secondMax < arr[i] && firstMax != arr[i] {
			secondMax = arr[i]
		}
	}

	return secondMax
}
