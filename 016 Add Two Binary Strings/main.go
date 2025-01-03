package main

import "fmt"

func main() {
	s1 := "1101"
	s2 := "111"
	fmt.Println(addTwoBinaryStrings(s1, s2))
}

func addTwoBinaryStrings(s1, s2 string) string {
	index := 0
	firstLength := len(s1)
	secondLength := len(s2)

	for index < firstLength && s1[index] == '0' {
		index++
	}
	s1 = s1[index:firstLength]
	firstLength = len(s1) - 1

	index = 0
	for index < secondLength && s2[index] == '0' {
		index++
	}
	s2 = s2[index:firstLength]
	secondLength = len(s2) - 1

	str := ""
	carry := 0

	for firstLength >= 0 && secondLength >= 0 {
		if s1[firstLength] == s2[secondLength] {
			if s1[firstLength] == '1' {
				if carry == 1 {
					str = "1" + str
				} else {
					str = "0" + str
				}
				carry = 1
			} else {
				if carry == 1 {
					str = "1" + str
				} else {
					str = "0" + str
				}
				carry = 0
			}
		} else {
			if carry == 1 {
				str = "0" + str
				carry = 1
			} else {
				str = "1" + str
			}
		}
		firstLength--
		secondLength--
	}

	for firstLength >= 0 {
		if carry == 1 {
			if s1[firstLength] == '1' {
				str = "0" + str
				carry = 1
			} else {
				str = "1" + str
			}
		} else {
			str = string(s1[firstLength]) + str
		}
		firstLength--
	}

	for secondLength >= 0 {
		if carry == 1 {
			if s2[secondLength] == '1' {
				str = "0" + str
				carry = 1
			} else {
				str = "1" + str
			}
		} else {
			str = string(s2[secondLength]) + str
		}
		secondLength--
	}

	if carry == 1 {
		str = "1" + str
	}
	return str
}
