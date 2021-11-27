package main

import (
	"fmt"
)


func isMatch1(s string, p string) bool {
	var (
		star, fail bool
		ps, pp, plen, slen int
		char, last_char byte
	)

	fmt.Printf("%s\t%s\t", s, p)

	slen = len(s)
	plen = len(p)

	for pp < plen {
		char = p[pp]	
		pp++
		if pp < plen && p[pp] == '*' {
			star = true
			pp++
		}
		if ps==slen { 
			//fmt.Println(ps, slen, pp, plen)
			fail = true
		}
		for ps < slen {
			if char=='.' { 
				last_char=s[ps]
				ps++ 
			} else if s[ps]==char { 
				last_char=s[ps]
				ps++ 
			} else { break }
			if !star { break }
		}
		//fmt.Println(star, ps, slen, pp, plen)
		if star {
			last_pp := pp
			for pp<plen {
				fmt.Printf("%c %c %c\n", last_char, char, p[pp])
				fmt.Println(star, ps, slen, pp, plen)
				if p[pp]==char ||char=='.' { pp++ } else { break }
			}
			if last_pp!=pp {
				if ps-(pp-last_pp)<0 { return false }
				//fmt.Println(p[last_pp:pp], s[ps-(pp-last_pp):ps])
				if p[last_pp:pp]!=s[ps-(pp-last_pp):ps] { return false }
			}
			fail = false
		}
		star = false
	}

	//fmt.Println(ps, slen, pp, plen)

	return ps==slen && pp==plen && !fail
}

func isMatch(s string, p string) bool {
	var pp, sp int
	var char byte

	dp := make([][]bool, len(p)+1)
	for i := range dp {
		dp[i] = make([]bool, len(s)+1)
	}
	dp[0][0] = true
	
	for i := range p {
		if p[i] == '*' {
			pp = i+1
			dp[pp][0] = dp[pp-2][0]
		}
	}
	
	for j := 0; j < len(p); j++ {
		pp = j+1
		char = p[j]
		
		for i := 0; i < len(s); i++ {
			sp = i+1
			
			if char == '*' {
				dp[pp][sp] = dp[pp-2][sp] || (s[i] == p[j-1] || p[j-1] == '.') && dp[pp][sp-1]
			} else if char == '.' {
				dp[pp][sp] = dp[pp-1][sp-1]
			} else {
				dp[pp][sp] = s[i] == p[j] && dp[pp-1][sp-1]
			}
		}
	}   
	
	return dp[len(p)][len(s)]
}

func main(){	
	fmt.Println(isMatch("a", "ab*"))
	fmt.Println(isMatch("aaa", "aaaa"))
	fmt.Println(isMatch("aaa", "ab*a*c*a"))
	fmt.Println(isMatch("aaa", ".*a"))
	fmt.Println(isMatch("aaab", ".*ab"))
	fmt.Println(isMatch("aaaa", "a*aaaaa"))
	fmt.Println(isMatch("aaa", "a*a"))
	fmt.Println(isMatch("aaab", "a*ab"))
	fmt.Println(isMatch("ab", ".*c"))
	fmt.Println(isMatch("aa", "a"))
	fmt.Println(isMatch("aa", "a*"))
	fmt.Println(isMatch("ab", ".*"))
	fmt.Println(isMatch("aab", "c*a*b"))
	fmt.Println(isMatch("mississippi", "mis*is*p*."))
}
