package main

import "fmt"

func main() {

	s1 := "aba"
	s2 := "aab"
	fmt.Println(stringRotationsOptimal(s1, s2))

	s1 = "abcd"
	s2 = "cdab"
	fmt.Println(stringRotationsOptimal(s1, s2))

	s1 = "abcd"
	s2 = "acbd"
	fmt.Println(stringRotationsOptimal(s1, s2))
}

func stringRotationsBruteforce(s1, s2 string) bool {
	// rotate the string and check whether it is matched or not if matched return true
	// else rotate till the length times
	return false
}

func stringRotationsOptimal(s1, s2 string) bool {
	s1 = s1 + s1
	lps := make([]int, len(s2))
	createLPS(s2, lps)

	i := 0
	j := 0
	for i < len(s1) {
		if s2[j] == s1[i] {
			i++
			j++
		} else {
			if j == 0 {
				i++
			} else {
				j = lps[j-1]
			}
		}

		if j == len(s2) {
			return true
		}
	}
	return false
}

func createLPS(pattern string, lps []int) {
	prevLPS, i := 0, 1
	for i < len(pattern) {
		if pattern[prevLPS] == pattern[i] {
			lps[i] = prevLPS + 1
			prevLPS++
			i++
		} else if prevLPS == 0 {
			lps[i] = 0
			i++
		} else {
			prevLPS = lps[prevLPS-1]
		}
	}
}
