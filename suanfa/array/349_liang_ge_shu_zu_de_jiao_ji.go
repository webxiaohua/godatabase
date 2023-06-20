// 349.两个数组的交集
/*
给定两个数组 nums1 和 nums2 ，返回 它们的交集 。输出结果中的每个元素一定是 唯一 的。我们可以 不考虑输出结果的顺序 。
输入：nums1 = [1,2,2,1], nums2 = [2,2]
输出：[2]
*/
package main

import "fmt"

func main(){
	arr1:=[]int{1,2,2,1}
	arr2:=[]int{2,2}
	res := intersection(arr1,arr2)
	fmt.Println(res)
}

func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	for _,num := range nums1 {
		if _,ok := m[num];!ok {
			m[num] = 1
		}
	}
	res := make([]int,0)
	for _,num := range nums2{
		if _,ok := m[num];ok {
			res = append(res,num)
			delete(m,num)
		}
	}
	return res
}
