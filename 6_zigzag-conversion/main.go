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

pos = 

PINALSIGYAHRPI
PINALSIGYAHRPI
*/

func convert(s string, numRows int) string {
	var hash [1000]string
	var way, p int
	way = -1
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
}