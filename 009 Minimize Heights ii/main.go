package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	arr := []int{1, 5, 8, 10}
	k := 2
	fmt.Println(minimumHeights2Optimal(arr, k))

	arr = []int{5, 2, 6, 1, 5, 8, 10}
	k = 2
	fmt.Println(minimumHeights2Optimal(arr, k))

	arr = []int{3, 9, 12, 16, 20}
	k = 3
	fmt.Println(minimumHeights2Optimal(arr, k))
}

func minimumHeights2Optimal(arr []int, k int) int {
	sort.Ints(arr)

	n := len(arr)
	result := float64(arr[n-1] - arr[0])

	for i := 1; i < n; i++ {
		minH := math.Min(float64(arr[0]+k), float64(arr[i]-k))
		maxH := math.Max(float64(arr[i-1]+k), float64(arr[n-1]-k))

		result = math.Min(result, maxH-minH)
	}

	return int(result)
}
