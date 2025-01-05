package main

import "fmt"

func main() {
	text := "abcab"
	pattern := "ab"
	fmt.Println(patternMatchingOptimal(pattern, text))

	text = "abesdu"
	pattern = "edu"
	fmt.Println(patternMatchingOptimal(pattern, text))

	text = "aabaacaadaabaaba"
	pattern = "aaba"
	fmt.Println(patternMatchingOptimal(pattern, text))

	text = "efdsvsdvsdfeiresdifdffddssdde"
	pattern = "vs"
	fmt.Println(patternMatchingOptimal(pattern, text))
}

func patternMatchingBruteforce(pattern, text string) []int {
	// using nested loops, check the pattern is matched in the text or not
	// if matched return the currentIndex - len(pattern) and start with next character checking else
	// repeat the process with the next Index
	return []int{}
}

func patternMatchingOptimal(pattern, text string) []int {
	result := make([]int, 0)

	lps := make([]int, len(pattern))
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

	i = 0
	j := 0
	for i < len(text) {
		if pattern[j] == text[i] {
			i++
			j++
		} else {
			if j == 0 {
				i++
			} else {
				j = lps[j-1]
			}
		}

		if j == len(pattern) {
			result = append(result, i-j)
			j = lps[j-1]
		}
	}

	return result
}
