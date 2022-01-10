package main

import (
	"fmt"
	"math"
	"sort"
)

//  3 <= nums.length <= 1000
//  -1000 <= nums[i] <= 1000
//  -104 <= target <= 104

func threeSumClosest(nums []int, target int) int {
	var res int = nums[0] + nums[1] + nums[2]
	sort.Ints(nums)
	lg := len(nums)
	for i := 0; i < lg-2; i++ {
		l := i + 1
		r := lg - 1
		for l < r {
			s := nums[i] + nums[l] + nums[r]
			if s == target {
				return target
			} else if s > target {
				r--
			} else {
				l++
			}
			if math.Abs(float64(s-target)) < math.Abs(float64(res-target)) {
				res = s
			}
		}
	}
	return res
}

func main(){    
	fmt.Println(threeSumClosest([]int{-1,0,1,2,-1,-4}, 1))
	fmt.Println(threeSumClosest([]int{-1,2,1,-4}, 1)) 
	fmt.Println(threeSumClosest([]int{0, 0, 0}, 1))
}
