package main

import (
	"fmt"
)

// 32-bit integer range [-2^31, 2^31 - 1]  [-2147483648, 2147483647]

func isPalindrome(x int) bool {
	if x < 0 || (x % 10 == 0 && x != 0) {
		return false
	}

	var n int

	for x>n {
		n = n*10 + x%10
		x /= 10
	}

	return x==n || x==n/10
}

func main(){	
	fmt.Println(isPalindrome(1211121))
	fmt.Println(isPalindrome(121121))
	fmt.Println(isPalindrome(-1201))
}
