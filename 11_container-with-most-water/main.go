package main

import (
	"fmt"
)


func maxArea0(height []int) int {
	var maxA, tmpA int
	var hash map[int]int
	hash = make(map[int]int)

	for i:=0;i<len(height);i++ {
		for k := range hash {
			left := hash[k]
			if height[i]>height[left] {
				tmpA = height[left] * (i - left)
			} else {
				tmpA = height[i] * (i - left)
			}
			if tmpA > maxA {
				maxA = tmpA
			}
		}

		if _, ok := hash[height[i]]; ok {
			//hash[height[i]] = height[i] * (i - left)
		} else {
			hash[height[i]] = i
		}
	}
	return maxA
}

func maxArea(height []int) int {
	max, s, e := 0, 0, len(height)-1
	for s < e {
		w := e - s
		h := 0
		if height[s] < height[e] {
			h = height[s]
			s++
		} else {
			h = height[e]
			e--
		}

		tmp := w * h
		if tmp > max {
			max = tmp
		}
	}
	return max
}

/*
		4
		4	3	3
		4	3	3
1	1	4	3	3	1
  1   1   3   3   1
*/
func main(){	
	fmt.Println(maxArea([]int{1,1}))
	fmt.Println(maxArea([]int{4,3,2,1,4}))
	fmt.Println(maxArea([]int{1,2,3,2,1,1}))
	fmt.Println(maxArea([]int{1,8,6,2,5,4,8,3,7}))
}
