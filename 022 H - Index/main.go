package main

import "fmt"

func main() {
	arr := []int{1, 0, 3, 5, 3}
	fmt.Println(hIndexOptimal(arr))

	arr = []int{5, 1, 2, 4, 1}
	fmt.Println(hIndexOptimal(arr))

	arr = []int{1, 0, 0}
	fmt.Println(hIndexOptimal(arr))
}

func hIndexBruteforce(arr []int) int {
	return 1
}

func hIndexOptimal(arr []int) int {
	n := len(arr)
	frequency := make([]int, n+1)

	for i := 0; i < n; i++ {
		if arr[i] >= n {
			frequency[n]++
		} else {
			frequency[arr[i]]++
		}
	}
	index := n
	s := frequency[index]
	for s < index {
		index--
		s += frequency[index]
	}
	return index
}
