package main

import (
	"fmt"
)


func isMatch(s string, p string) bool {
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
