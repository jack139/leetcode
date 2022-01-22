package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	var result [][]int
	l := len(nums)
	sort.Ints(nums)
	for i := 0; i < l-3; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i]+nums[i+1]+nums[i+2]+nums[i+3] > target {
			break
		}
		if nums[i]+nums[l-1]+nums[l-2]+nums[l-2] < target {
			continue
		}
		for j := i + 1; j < l-2; j++ {
			if j != i+1 && nums[j] == nums[j-1] {
				continue
			}
			if nums[i]+nums[j]+nums[j+1]+nums[j+2] > target {
				break
			}
			if nums[i]+nums[j]+nums[l-1]+nums[l-2] < target {
				continue
			}
			head, end := j+1, l-1
			for head < end {
				tmp := nums[i] + nums[j] + nums[head] + nums[end]
				if (tmp == target) {
					var t2 []int
					t2 = append(t2, nums[i])
					t2 = append(t2, nums[j])
					t2 = append(t2, nums[head])
					t2 = append(t2, nums[end])
					result = append(result, t2)
					head++;
					for head < end && nums[head] == nums[head-1] {
						head++;
					}
				} else if tmp > target {
					end--
				} else {
					head++
				}
			}
		}
	}
	return result
}

func main(){    
	fmt.Println(fourSum([]int{1,0,-1,0,-2,2}, 0))
	fmt.Println(fourSum([]int{2,2,2,2,2}, 8)) 
}
