package main

import "fmt"

func main() {
	arr := []int{5, 4, 3, 2, 1}
	fmt.Println(arr)
	fmt.Println(buyStocks2Optimal(arr))

	arr = []int{86, 92, 24, 5, 34, 72, 68, 52, 27, 95, 41, 28, 35}
	fmt.Println(arr)
	fmt.Println(buyStocks2Optimal(arr))

	arr = []int{100, 180, 260, 310, 40, 535, 695}
	fmt.Println(arr)
	fmt.Println(buyStocks2Optimal(arr))

	arr = []int{4, 2, 2, 2, 4}
	fmt.Println(arr)
	fmt.Println(buyStocks2Optimal(arr))

}

func buyStocks2Bruteforce(arr []int) int {
	// using nested loops check the difference b/w the elements and calculate the max profit
	return 0
}

func buyStocks2Optimal(arr []int) int {
	profit := 0

	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[i-1] {
			profit = profit + arr[i] - arr[i-1]
		}
	}
	return profit
}
