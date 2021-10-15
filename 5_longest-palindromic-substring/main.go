package main

import (
	"fmt"
)

func longestPalindrome3(s string) string {
    
    res := ""
    for i := range s {
        odd := helper(s, i, i)
        even := helper(s, i, i+1)
        
        if len(odd) > len(res) {
            res = odd
        }

        if len(even) > len(res) {
            res = even
        }
    }
    
    return res
}

func helper(s string, l, r int) string {
    for l >= 0 && r < len(s) && s[l] == s[r] {
        l--
        r++
    }
    
    aa := s[l+1:r]
    
    return aa
}

// babad
// b 		i=0 j=0
// ba           j=0
//  a       i=1 j=0
//	ab          j=0
// bab          j=1
// baba         j=1
//   b      i=2
//	 ba
//  aba
//  abad        j=1
// babad        j=2
func longestPalindrome(s string) string {
	var maxSize, start, end int
	l := len(s)
	for i := range s {
		for j:=maxSize/2; j<i+1; j++ {
			if j+j+1>maxSize && i+j+1<=l{
				fmt.Println(s[i-j:i+j+1], i, j)
				if checkPalindrome(s[i-j:i+j+1], j+j+1) {
					maxSize = j+j+1
					start = i-j
					end = i+j
				}
			}
			if j+j+2>maxSize && i+j+2<=l{
				fmt.Println(s[i-j:i+j+2], i, j)
				if checkPalindrome(s[i-j:i+j+2], j+j+2) {
					maxSize = j+j+2
					start = i-j
					end = i+j+1
				}
			}
		}
	}
	return s[start:end+1]
}

func longestPalindrome2(s string) string {
	var (
		alphb [128][]int
		maxSize, maxstart int
	)

	if len(s)==0 {
		return ""
	}

	maxSize = 1

	for i:=0; i<len(s); i++ {
		alphb[s[i]] = append(alphb[s[i]], i)

		l := len(alphb[s[i]])

		if l<2 {
			continue
		}

		for k:=0; k<l; k++ {
			start := alphb[s[i]][k]
			size := i-start+1
			if size<=maxSize {
				continue
			}
			fmt.Printf("%c %s\n", s[i], s[start:i+1])
			if checkPalindrome(s[start:i+1], size) {
				maxSize = size
				maxstart = start
			}
		}
	}

	return s[maxstart:maxstart+maxSize]
}


func checkPalindrome(s string, l int) bool {
	for i:=0;i<l/2;i++ {
		if s[i]!=s[l-i-1] {
			return false
		}
	}
	return true
}

func main(){
	
	//fmt.Println(longestPalindrome("babad"))
	//fmt.Println(longestPalindrome("cbbd"))
	fmt.Println(longestPalindrome("eabcb"))
	//fmt.Println(longestPalindrome("aacabdkacaa"))
	//fmt.Println(longestPalindrome("aacabdkakdbaca"))
	
	//fmt.Println(checkPalindrome("bb", 2))
}
