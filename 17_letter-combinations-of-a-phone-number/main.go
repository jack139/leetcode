package main

import (
	"fmt"
)

// 0 <= digits.length <= 4
// digits[i] is a digit in the range ['2', '9'].
var (
	res []string
	letters = [][]byte{
		{},
		{},
		{'a','b','c'},
		{'d','e','f'},
		{'g','h','i'},
		{'j','k','l'},
		{'m','n','o'},
		{'p','q','r','s'},
		{'t','u','v'},
		{'w','x','y','z'},
	}
)

func letterCombinations(digits string) []string {
	res = make([]string, 0)
	if len(digits) == 0 {
		return res
	}
	find(digits, 0, []byte{})
	return res
}

func find(d string, index int, p []byte) {
	if index == len(d) {
		res = append(res, string(p))
		return
	}
	
	num := d[index]-'0'
	for i := 0; i < len(letters[num]); i++ {
		find(d, index+1, append(p, letters[num][i]))
	}
}

func main(){    
	fmt.Println(letterCombinations("23"))
	fmt.Println(letterCombinations("")) 
	fmt.Println(letterCombinations("2"))
}
