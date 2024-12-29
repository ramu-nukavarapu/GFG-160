package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{2, 3, 4}
	fmt.Println(maximumSubarrayProductOptimal(arr))

	arr = []int{0}
	fmt.Println(maximumSubarrayProductOptimal(arr))

	arr = []int{-1, 3, 2, 1, -5}
	fmt.Println(maximumSubarrayProductOptimal(arr))

	arr = []int{2, 3, -4}
	fmt.Println(maximumSubarrayProductOptimal(arr))
}

func maximumSubarrayProductOptimal(arr []int) int {
	max := math.MinInt
	prefix, suffix := 1, 1
	n := len(arr)

	for i := 0; i < n; i++ {
		if prefix == 0 {
			prefix = 1
		}

		if suffix == 0 {
			suffix = 1
		}

		prefix *= arr[i]
		suffix *= arr[n-1-i]

		if prefix > max {
			max = prefix
		}

		if suffix > max {
			max = suffix
		}
	}
	return max
}
