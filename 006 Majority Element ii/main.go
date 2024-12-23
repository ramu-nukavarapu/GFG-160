package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	fmt.Println("Output:=", majorityElement2Optimal(arr))

	arr = []int{2, 1, 5, 5, 5, 5, 6, 6, 6, 6, 6}
	fmt.Println(arr)
	fmt.Println("Output:=", majorityElement2Optimal(arr))

	arr = []int{2, 2, 1, 3, 1}
	fmt.Println(arr)
	fmt.Println("Output:=", majorityElement2Optimal(arr))

	arr = []int{1, 1, 2, 3, 4}
	fmt.Println(arr)
	fmt.Println("Output:=", majorityElement2Optimal(arr))

}

func majorityElement2Bruteforce(arr []int) (result []int) {
	//use nested loops to check whether a element occurs more than n/3 times
	return arr
}

func majorityElement2Optimal(arr []int) []int {
	n := len(arr)
	element1, element2 := -1, -1
	count1, count2 := 0, 0

	for i := 0; i < n; i++ {
		if element1 == arr[i] {
			count1++
		} else if element2 == arr[i] {
			count2++
		} else if count1 == 0 {
			element1 = arr[i]
			count1++
		} else if count2 == 0 {
			element2 = arr[i]
			count2++
		} else {
			count1--
			count2--
		}
	}

	count1 = 0
	count2 = 0

	for i := 0; i < n; i++ {
		if element1 == arr[i] {
			count1++
		} else if element2 == arr[i] {
			count2++
		}
	}

	result := make([]int, 2)

	if count1 > n/3 && count2 > n/3 {
		if element1 < element2 {
			result[0] = element1
			result[1] = element2
		} else {
			result[0] = element2
			result[1] = element1
		}
	} else {
		if count1 > n/3 {
			result[0] = element1
		}

		if count2 > n/3 {
			result[0] = element2
		}
	}
	return result
}
