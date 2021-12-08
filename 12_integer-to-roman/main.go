package main

import (
	"fmt"
	"strings"
)

// 1 <= num <= 3999

func intToRoman(num int) string {
	var px, pv int
	x := []byte{'I', 'X', 'C', 'M'}
	v := []byte{'V', 'L', 'D'}
	s := []byte("................")
	p := len(s)-1

	for num > 0 {
		d := num % 10
		switch d {
		case 1:
			s[p] = x[px]
			p--
		case 2:
			s[p-1] = x[px]
			s[p] = x[px]
			p -= 2
		case 3:
			s[p-2] = x[px]
			s[p-1] = x[px]
			s[p] = x[px]
			p -= 3
		case 4:
			s[p-1] = x[px]
			s[p] = v[pv]
			p -= 2
		case 5:
			s[p] = v[pv]
			p --
		case 6:
			s[p-1] = v[pv]
			s[p] = x[px]
			p -= 2
		case 7:
			s[p-2] = v[pv]
			s[p-1] = x[px]
			s[p] = x[px]
			p -= 3
		case 8:
			s[p-3] = v[pv]
			s[p-2] = x[px]
			s[p-1] = x[px]
			s[p] = x[px]
			p -= 4
		case 9:
			s[p-1] = x[px]
			s[p] = x[px+1]
			p -= 2
		}
		px++
		pv++
		num = num / 10
	}

	return string(s[p+1:])
}

func intToRoman2(num int) string {
	var px, pv int
	x := []string{"I", "X", "C", "M"}
	v := []string{"V", "L", "D"}
	s := ""

	for num > 0 {
		d := num % 10
		switch d {
		case 1:
			s = x[px] + s
		case 2:
			s = x[px] + x[px] + s
		case 3:
			s = x[px] + x[px]+ x[px] + s
		case 4:
			s = x[px] + v[pv] + s
		case 5:
			s = v[pv] + s
		case 6:
			s = v[pv] + x[px] + s
		case 7:
			s = v[pv] + x[px] + x[px] + s
		case 8:
			s = v[pv] + x[px] + x[px] + x[px] + s
		case 9:
			s = x[px] + x[px+1] + s
		}
		px++
		pv++
		num = num / 10
	}

	return s
}

func intToRoman1(num int) string {
	var px, pv int
	x := []string{"I", "X", "C", "M"}
	v := []string{"V", "L", "D"}
	s := ""

	for num > 0 {
		d := num % 10
		switch d {
		case 1,2,3:
			s = strings.Repeat(x[px], d) + s
		case 4:
			s = x[px] + v[pv] + s
		case 5, 6, 7, 8:
			s = v[pv] + strings.Repeat(x[px], d-5) + s
		case 9:
			s = x[px] + x[px+1] + s
		}
		px++
		pv++
		num = num / 10
	}

	return s
}

func main(){	
	fmt.Println(intToRoman(3)) //  III
	fmt.Println(intToRoman(4)) //  IV
	fmt.Println(intToRoman(9)) //  IX
	fmt.Println(intToRoman(58))  //  L VIII
	fmt.Println(intToRoman(1994)) // M CM XC IV
}
