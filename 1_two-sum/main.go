package main

import (
	"fmt"
)

func twoSum_v1(nums []int, target int) []int {
	for i, x := range nums {
		for j, y := range nums[:i] {
			if (x+y) == target {
				return []int{j, i}
			}
		}
	}
	return []int{}
}


func twoSum_v2(nums []int, target int) []int {
	type Item struct {
		val int
		pos int
	}

	const L=20
	var hash map[int][]Item
	hash = make(map[int][]Item)

	for i, x := range nums {
		y := target - x
		for _, item := range hash[y%L] {
			if y==item.val {
				return []int{item.pos, i}
			}
		}

		hash[x%L] = append(hash[x%L], Item{x, i})
	}
	return nil
}

func main(){
	var list = []int{2,7,11,15}
	var target = int(26)

	r := twoSum_v2(list, target)
	fmt.Println(r)
}