package main

import "fmt"

func main() {
	s1 := "eef"
	s2 := "eedctv"
	fmt.Println(anagramsOptimal(s1, s2))

	s1 = "cake"
	s2 = "caek"
	fmt.Println(anagramsOptimal(s1, s2))

	s1 = "guava"
	s2 = "vuaga"
	fmt.Println(anagramsOptimal(s1, s2))

	s1 = "cat"
	s2 = "eat"
	fmt.Println(anagramsOptimal(s1, s2))
}

func anagramsBruteforce(s1, s2 string) bool {
	// sort both strings and check one element after other
	// if any element is mismatched return false else return true
	return true
}

func anagramsOptimal(s1, s2 string) bool {

	if len(s1) != len(s2) {
		return false
	}
	arr := make([]int, 26)

	for i := 0; i < len(s1); i++ {
		val := int(s1[i] - 'a')
		arr[val]++
	}
	for i := 0; i < len(s2); i++ {
		val := int(s2[i] - 'a')
		arr[val]--
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			return false
		}
	}
	return true
}
