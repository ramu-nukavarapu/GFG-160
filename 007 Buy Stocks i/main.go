package main

import "fmt"

func main() {
	arr := []int{5, 4, 3, 2, 1}
	fmt.Println(arr)
	fmt.Println("Profit:=", buyStocks1Optimal(arr))

	arr = []int{7, 10, 1, 3, 6, 9, 2}
	fmt.Println(arr)
	fmt.Println("Profit:=", buyStocks1Optimal(arr))

	arr = []int{1, 3, 6, 9, 11}
	fmt.Println(arr)
	fmt.Println("Profit:=", buyStocks1Optimal(arr))

	arr = []int{4, 5}
	fmt.Println(arr)
	fmt.Println("Profit:=", buyStocks1Optimal(arr))

}

func buyStocks1Bruteforce(arr []int) int {
	// use nested loops to check the maximum profit
	// by calculating the difference b/w current element and the remaining ones
	return 0
}

func buyStocks1Optimal(arr []int) int {
	profit, current := 0, 0
	first, second := 0, 1

	for second < len(arr) {
		if arr[first] < arr[second] {
			current = arr[second] - arr[first]
			if profit < current {
				profit = current
			}
			second++
		} else {
			first = second
			second++
		}
	}

	return profit
}
