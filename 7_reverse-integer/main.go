package main

import (
	"fmt"
)

// 32-bit integer range [-2^31, 2^31 - 1]  [-2147483648, 2147483647]

func reverse(x int) int {
	var n int

	for x!=0 {
		n = n*10 + x%10
		x = x/10
	}

	if n > 2147483647 || n < -2147483648 {
		n = 0
	} 
	
	return n
}

func main(){
	
	fmt.Println(reverse(-21477))
}
