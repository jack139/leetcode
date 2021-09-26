package main

import (
	"fmt"
)


func lengthOfLongestSubstring(s string) int {
	var alphb [128]byte
	var maxSize, start, end int

    if len(s)==0 {
		return 0
	}

	maxSize = 1
	start = 0
	end = 0
	alphb[byte(s[end])]=1
	for ; end<len(s)-1 ; {
		end += 1
		if alphb[byte(s[end])]==0 {

			if end-start+1>maxSize {
				maxSize = end - start + 1
			}

			// 没有重复，窗口变大
			alphb[byte(s[end])] = 1
		} else {
			for ;alphb[byte(s[end])]==1; {
				// 有重复，窗口进行后移
				if start!=end { alphb[byte(s[start])] = 0 }
				start += 1
			}
			alphb[byte(s[end])] = 1
		}
    }

	return maxSize
}

func main(){
	s := "pwwkew"
	//s := "abcabcbb"
	//s := "bbbbb"
	//s := ""
	r := lengthOfLongestSubstring(s)
	fmt.Println(r)
}
