package main

import "fmt"

func main() {
	s := "abc"
	fmt.Println(minCharstoAddforPalindromeOptimal(s))

	s = "aacecaaaa"
	fmt.Println(minCharstoAddforPalindromeOptimal(s))

	s = "madam"
	fmt.Println(minCharstoAddforPalindromeOptimal(s))

}

func minCharstoAddforPalindromeBruteforce(s string) int {
	// check all possible prefixes whether they are palindrome or not
	// if there is longest possible prefix palindrome return string length - prefix length
	return 0
}

func minCharstoAddforPalindromeOptimal(s string) int {
	n := len(s)
	rev := reverse(s)
	s = s + "$" + rev
	lps := make([]int, len(s))
	createLPS(s, lps)

	return n - lps[len(s)-1]
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

func reverse(s string) string {
	rev := ""
	for i := 0; i < len(s); i++ {
		rev = string(s[i]) + rev
	}
	return rev
}
