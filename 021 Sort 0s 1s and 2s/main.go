package main

import "fmt"

func main() {
	arr := []int{1, 2, 0, 1, 2, 0}
	sor0s1sand2stOptimal(arr)
	fmt.Println(arr)

	arr = []int{0, 1, 2, 0, 1, 2}
	sor0s1sand2stOptimal(arr)
	fmt.Println(arr)

	arr = []int{0, 1, 1, 0, 1, 2, 1, 2, 0, 0, 0, 1}
	sor0s1sand2stOptimal(arr)
	fmt.Println(arr)

}

func sort0s1sand2sBruteforce(arr []int) {
	// sort the array any sorting technique
}

func sor0s1sand2stOptimal(arr []int) {
	low, mid, high := 0, 0, len(arr)-1

	for mid <= high {
		if arr[mid] == 0 {
			swap(low, mid, arr)
			low++
			mid++
		} else if arr[mid] == 2 {
			swap(mid, high, arr)
			high--
		} else {
			mid++
		}
	}
}

func swap(low, high int, arr []int) {
	temp := arr[low]
	arr[low] = arr[high]
	arr[high] = temp
}
