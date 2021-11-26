package main

import (
	"fmt"
)


func isMatch(s string, p string) bool {
	var (
		star, last_star bool
		ps, pp, plen, slen int
		char byte
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
		if ps==slen && !last_star { return false }
		for ps < slen {
			if char=='.' { ps++ } else if s[ps]==char { ps++ } else { break }
			if !star { break }
		}
		if star {
			last_pp := pp
			for pp<plen {
				if p[pp]==char || char=='.' { pp++ } else { break }
			}
			if last_pp!=pp {
				if ps-(pp-last_pp)<0 { return false }
				if p[last_pp:pp]!=s[ps-(pp-last_pp):ps] { return false }
			}
			last_star = true
		} else { last_star = false }
		star = false
	}

	//fmt.Println(ps, slen, pp, plen)

	return ps==slen && pp==plen
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
