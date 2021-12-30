package main

import (
	"fmt"
	"sort"
)

// 0 <= nums.length <= 3000
// -105 <= nums[i] <= 105

func threeSum(nums []int) [][]int {
	var r [][]int
	
	if len(nums) == 0 || len(nums) == 1 { return r }
	
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] { continue }
		
		j, k := i + 1, len(nums) - 1

		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			
			if sum < 0 {
				j++
			} else if sum > 0 {
				k--
			} else {
				r = append(r, []int{nums[i], nums[j], nums[k]})
				
				for j < k && nums[j] == nums[j + 1] { j++ }
				for j < k && nums[k] == nums[k - 1] { k-- }
				
				j++
				k--
			}
		}
	}
	
	return r
}

func main(){    
	//fmt.Println(threeSum([]int{-1,0,1,2,-1,-4})) // [[-1,-1,2],[-1,0,1]]
	fmt.Println(threeSum([]int{-2,0,0,2,2})) // [[-2,0,2]]
	fmt.Println(threeSum([]int{0}))
}

/*
C3,6 = P3

-1 0 1
-1 0   2
-1 0     -1
-1 0        -4
-1   1 2
-1   1   -1
-1   1      -4
-1     2 -1
-1     2    -4
-1       -1 -4
   0 1 2
   0 1   -1
   0 1      -4
   0   2 -1
   0   2    -4
   0     -1 -4
	 1 2 -1
	 1 2    -4
	 1   -1 -4
	   2 -1 -4
*/