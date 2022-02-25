package main

import (
	"fmt"
)


/**
1 <= s.length <= 10000
s consists of parentheses only '()[]{}'.
*/
func isValid(s string) bool {
	var stack []byte
	for i:=0;i<len(s);i++ {
		if s[i]=='(' || s[i]=='[' || s[i]=='{' {
			stack = append(stack, s[i])
		} else {
			l := len(stack)
			if l==0 {
				return false
			}
			pop := stack[l-1]
			switch pop {
			case '(':
				pop = ')'
			case '[':
				pop = ']'
			case '{':
				pop = '}'
			}
			if pop!=s[i] {
				return false
			}
			stack = stack[:l-1]
		}
	}
	if len(stack)==0 {
		return true
	} else {
		return false
	}
}

func main(){
	fmt.Println(isValid("([{}])[]{}"))
	fmt.Println(isValid("([{}]))[]{}"))
	fmt.Println(isValid("([{}])[]{}("))
}
