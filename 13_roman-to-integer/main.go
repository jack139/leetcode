package main

import (
	"fmt"
)

// 1 <= num <= 3999
// 1 <= s.length <= 15

/*
	I             1
	V             5
	X             10
	L             50
	C             100
	D             500
	M             1000
*/

func romanToInt(s string) int {
	var alpha = map[byte]int{'I':1, 'V':5, 'X':10, 'L':50, 'C':100, 'D':500, 'M':1000}
	var num, last int
	p := len(s)-1

	for p>=0 {
		d := alpha[s[p]]
		if d<last {
			num -= d
		} else {
			num += d
		}
		last = d
		p--
	}

	return num
}


func main(){	
	fmt.Println(romanToInt("III"))// (3)) 
	fmt.Println(romanToInt("IV"))// (4))
	fmt.Println(romanToInt("IX"))// (9))
	fmt.Println(romanToInt("LVIII"))// (58))
	fmt.Println(romanToInt("MCMXCIV"))// (1994))
}
