package main

import (
	"fmt"
)

/*

PAYPALISHIRING

P  I  N
A LS IG
YA HR
P  I

n 字符个数
r 总行数
l 行号
lnum 一行个数
i 原始位置 0 开始
pos 新位置  0 开始

0     6     12      3 = round((n-1)/6)+1
1   5 7  11 13      5 = round((n-2)/3)+1
2 4   8 10          4 = round((n-3)/3)+1
3     9             2 = round((n-4)/6)+1

0   4   8    12     4 = round((n-1)/4)+1
1 3 5 7 9 11 13     7 = round((n-2)/2)+1
2   6   10          3 = round((n-3)/4)+1

lnum = round((n-l)/(r-1))+1
lnum = round((n-l)/((r-1)*2))+1  首尾行

0   1   2    3
4 5 6 7 8 9  10
11  12  13

0       8		0 0	 4  k=4	4=4 - 0
1     7 9		0 3     k=3	3=4 - 1
2   6   10 		0 2     k=2	2=4 - 2
3 5     11 13	0 1     k=1 1=4 - 3
4		12		0 0     k=0	0=4 - 4

pos = 

PINALSIGYAHRPI
02468AC13579BD
048C13579BD26A
06C157BD248A39
0817926A35BD4C

*/

func convert(s string, numRows int) string {
	var s2 [1000]byte

	if numRows==1 {
		return s
	}

	l := len(s)
	r := 0
	k := numRows - 1 - r
	n := 0
	for i:=0;i<l;i++ {
		//fmt.Println("\n",i,r,k,n)
		if n*(numRows-1)*2+r > l-1 {
			r++
			k = numRows - 1 - r
			n = 0
		}
		if n*(numRows-1)*2+r < l {
			s2[i] = s[n*(numRows-1)*2+r]
			//fmt.Printf("\n; i=%d, r=%d, k=%d, n=%d, s[]=%c", i, r, k, n, s[n*(numRows-1)*2+r])
		}
		if k%(numRows-1)!=0 {
			if (n*(numRows-1)+k)*2+r < l {
				i++
				s2[i] = s[(n*(numRows-1)+k)*2+r]
				//fmt.Printf("\n, i=%d, r=%d, k=%d, n=%d, s[]=%c", i, r, k, n, s[(n*(numRows-1)+k)*2+r])
			}
		}
		n++
	}

	return string(s2[:l])
}

func convert2(s string, numRows int) string {
	var hash [1000]string
	var way, p int
	way = -1

	if numRows==1 {
		return s
	}

	for n, v := range s {
		hash[p] = hash[p] + string(v)
		if n % (numRows-1) == 0 {
			way = -way
		}
		if way==1 {
			p += 1
		} else {
			p -= 1
		}
	}

	var ss string
	for i:=0;i<numRows;i++ {
		ss = ss + hash[i]
	}

	return ss
}

func main(){
	fmt.Println(convert("PAYPALISHIRING", 4))
	fmt.Println(convert("0123456789ABCD", 2))
	fmt.Println(convert("0123456789ABCD", 3))
	fmt.Println(convert("0123456789ABCD", 4))
	fmt.Println(convert("0123456789ABCD", 5))
}