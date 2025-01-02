package main

import (
	"fmt"
	"math"
)

func main() {
	s := "    -123"
	fmt.Println(convertAtoI(s))

	s = "  -"
	fmt.Println(convertAtoI(s))

	s = " 1231231231311133"
	fmt.Println(convertAtoI(s))

	s = "-999999999999"
	fmt.Println(convertAtoI(s))

	s = "  -0012gfg4"
	fmt.Println(convertAtoI(s))
}

func convertAtoI(s string) int {
	index := 0
	for index < len(s) && s[index] == ' ' {
		index++
	}

	sign := 1
	if index < len(s) && s[index] == '+' || s[index] == '-' {
		if s[index] == '-' {
			sign = -1
		}
		index++
	}

	res := 0
	for index < len(s) && s[index] >= '0' && s[index] <= '9' {
		if res > math.MaxInt32/10 || (res == math.MaxInt32/10 && s[index] > '7') {
			if sign == 1 {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}
		res = res*10 + int(s[index]-'0')
		index++
	}

	return sign * res
}
