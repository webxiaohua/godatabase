package main

import (
	"fmt"
	sort1 "sort"
)

func main()  {
	res := intersect([]int{1,1,2,4,6,10},[]int{1,2,3,5,6})
	fmt.Println(res)
}

func intersect(nums1 []int, nums2 []int) []int {
	i, j, k := 0, 0, 0
	sort1.Ints(nums1)
	sort1.Ints(nums2)
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] > nums2[j] {
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			nums1[k] = nums1[i]
			i++
			j++
			k++
		}
	}
	return nums1[:k]
}