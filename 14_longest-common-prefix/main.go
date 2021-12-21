package main

import (
	"fmt"
)

// 1 <= strs.length <= 200
// 0 <= strs[i].length <= 200
// strs[i] consists of only lower-case English letters.

func longestCommonPrefix(strs []string) string {
	var p int
	var end bool

	if len(strs)==0 { return "" }

	for {
		for _, v := range strs {
			if p>=len(v) || v[p]!=strs[0][p] { 
				end = true 
				break
			}
		}
		if end { break } else { p++ }
	}

	return strs[0][:p]
}


func main(){    
	fmt.Println(longestCommonPrefix([]string{"flower","flow","flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog","racecar","car"}))
	fmt.Println(longestCommonPrefix([]string{"a"}))
}
