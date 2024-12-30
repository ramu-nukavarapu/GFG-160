package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{8, -8, 9, -9, 10, -11, 12}
	fmt.Println(maximumSubarraySumOptimal(arr))

	arr = []int{-1, 40, -14, 7, 6, 5, -4, -1}
	fmt.Println(maximumSubarraySumOptimal(arr))

	arr = []int{-2}
	fmt.Println(maximumSubarraySumOptimal(arr))

	arr = []int{10, -3, -4, 7, 6, 5, -4, -1}
	fmt.Println(maximumSubarraySumOptimal(arr))

}

func maximumSubarraySumBruteforce(arr []int) int {
	// use nested loops start with first one and in the inner loop from that element iterate over and calculate maximum sum
	// check if the currentsum exceeds the maximum value if exceeds change the maximum value to the current sum
	return 0
}

func maximumSubarraySumOptimal(arr []int) int {
	max := math.MinInt
	min := math.MaxInt
	sum := 0
	currentMax := 0
	currentMin := 0

	if len(arr) == 1 {
		return arr[0]
	}

	for i := 0; i < len(arr); i++ {
		sum += arr[i]

		currentMax += arr[i]
		currentMin += arr[i]

		if currentMax > max {
			max = currentMax
		}

		if currentMax < 0 {
			currentMax = 0
		}

		if currentMin < min {
			min = currentMin
		}

		if currentMin > 0 {
			currentMin = 0
		}
	}

	if max > (sum - min) {
		return max
	} else {
		return sum - min
	}
}
