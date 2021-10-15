package main

import (
	"fmt"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var i, j, v, m, n int

	m = len(nums1)
	n = len(nums2)

	for ;i+j<(m+n+1)/2; {
		if i<m && j<n {
			if nums1[i]<nums2[j] {
				v = nums1[i]
				i += 1
			} else {
				v = nums2[j]
				j += 1
			}
		} else {
			if i<m {
				v = nums1[i]
				i += 1
			} else {
				v = nums2[j]
				j += 1
			}
		}
	}

	if float32(m+n)/2.0 == float32((m+n)/2) { // 说明中间数有两个
		if i<m && j<n {
			if nums1[i]<nums2[j] {
				v += nums1[i]
			} else {
				v += nums2[j]
			}
		} else {
			if i<m {
				v += nums1[i]
			} else {
				v += nums2[j]
			}
		}

		return float64(v) / 2.0
	} else {
		return float64(v)
	}
}

func main(){
	n1 := []int{1,3}
	n2 := []int{2,4}
	r := findMedianSortedArrays(n1, n2)

	fmt.Println(n1, n2, r)
}