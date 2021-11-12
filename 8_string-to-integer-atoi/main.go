package main

import (
	"fmt"
)

// 32-bit integer range [-2^31, 2^31 - 1]  [-2147483648, 2147483647]

func myAtoi(s string) int {
	var n, digit int

	start := false
	sign := 1

	for x:=0;x<len(s);x++ {
		i := s[x]

		if start {
			if i<'0' || i>'9' {
				break
			}
			n = n*10 + int(i-'0')

			if n>0 { digit++ }

			if digit>10 { break } // 防止64位溢出
		} else {
			if i=='-' {
				sign = -1
				start = true
			} else if i=='+' {
				start = true
			} else if i >= '0' && i <= '9' {
				n = int(i-'0')
				start = true
			} else if i!=' ' {
				break
			}
		}
	}

	n = n*sign

	if n > 2147483647 { n = 2147483647 } else if n < -2147483648 { n = -2147483648 } 

	return n
}

func main(){	
	fmt.Println(myAtoi("10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000522545459"))
	fmt.Println(myAtoi("18446744073709551617"))
}
