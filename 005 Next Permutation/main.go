package main

import "fmt"

func main() {
	arr := []int{2, 4, 1, 7, 5, 0}
	fmt.Println(arr)
	nextPermutationOptimal(arr)
	fmt.Println(arr)

	arr = []int{2, 4, 1, 0}
	fmt.Println(arr)
	nextPermutationOptimal(arr)
	fmt.Println(arr)

	arr = []int{3, 4, 2, 5, 1}
	fmt.Println(arr)
	nextPermutationOptimal(arr)
	fmt.Println(arr)

}

func nextPermutationBruteforce(arr []int) {
	//copy and sort the array. generate all permutations for  the sorted array
	//point out the array that is given and return the next one  as a permutation
}

func nextPermutationOptimal(arr []int) {
	pivot, length := -1, len(arr)

	for i := length - 2; i >= 0; i-- {
		if arr[i] < arr[i+1] {
			pivot = i
			break
		}
	}

	if pivot == -1 {
		reverseArray(arr, 0, length-1)
	} else {
		for i := length - 1; i > pivot; i-- {
			if arr[pivot] < arr[i] {
				swap(arr, pivot, i)
				break
			}
		}
		reverseArray(arr, pivot+1, length-1)
	}
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

func swap(num []int, one, two int) {
	temp := num[one]
	num[one] = num[two]
	num[two] = temp
}
