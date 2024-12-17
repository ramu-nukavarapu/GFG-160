package main

func main() {

}

func reverseArrayBruteforce(arr []int) {
	// take a new array of same size as input array
	// traverse the input array in reverse order and copy the elements to the new array
}

func reverseArrayOptimal(nums []int) {
	start, end := 0, len(nums)-1
	var temp int

	for start < end {
		temp = nums[start]
		nums[start] = nums[end]
		nums[end] = temp
		start++
		end--
	}
}
