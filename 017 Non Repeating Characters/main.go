package main

import "fmt"

func main() {
	s := "geekforeeks"
	fmt.Printf("%c\n", nonRepeatingCharactersOptimal(s))

	s = "geekforgeeks"
	fmt.Printf("%c\n", nonRepeatingCharactersOptimal(s))

	s = "aabbccc"
	fmt.Printf("%c\n", nonRepeatingCharactersOptimal(s))

	s = "racecar"
	fmt.Printf("%c\n", nonRepeatingCharactersOptimal(s))

}

func nonRepeatingCharactersBruteforce(s string) rune {
	// using nested loops, start with one element and checks whether the element is matched to any other element
	//if matches then move to next one else return the element
	return '$'
}

func nonRepeatingCharactersOptimal(s string) rune {
	arr := make([]int, 26)

	for i := 0; i < len(s); i++ {
		arr[s[i]-'a']++
	}

	for i := 0; i < len(s); i++ {
		if arr[s[i]-'a'] == 1 {
			return rune(s[i])
		}
	}
	return '$'
}
