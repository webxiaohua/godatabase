// 15.三数之和
/*
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。
输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]
*/
package main

import (
	"fmt"
	"sort"
)

func main(){
	nums:=[]int{-1,0,1,2,-1,-4}
	res := threeSum2(nums)
	fmt.Println(res)
}

func threeSum(nums []int) [][]int {
	res := make([][]int,0)
	if len(nums) < 3 {
		return res
	}
	// 先排序
	sort.Ints(nums)
	// 三重循环
	for i:=0;i<len(nums)-2;i++ {
		if i==0 || nums[i] != nums[i-1] {
			for j := i + 1; j < len(nums)-1; j++ {
				if j==i+1 || nums[j] != nums[j-1] {
					for k := j + 1; k < len(nums); k++ {
						if k==j+1 || nums[k] != nums[k-1] {
							if nums[i]+nums[j]+nums[k] == 0 {
								res = append(res, []int{nums[i], nums[j], nums[k]})
							}
						}
					}
				}
			}
		}
	}
	return res
}

func threeSum2(nums []int) [][]int {
	// 先排序，再左右指针， L=nums[i+1] R=nums[len(nums)-1]
	// 当 nums[i] + nums[L] + nums[len()] = 0
	// 如果 nums[i] == nums[L] 继续下一轮循环
	// 如果和为正数，则R向左移动
	// 如果和为负数，则L向右移动
	res := make([][]int,0)
	if len(nums) < 3 {
		return res
	}
	sort.Ints(nums)
	for i,_ := range nums{
		if nums[i] > 0 {
			break
		}
		if i>0 && nums[i] == nums[i-1] {
			continue
		}
		l := i+1
		r := len(nums)-1
		for {
			if l >= r {
				break
			}
			if nums[i] + nums[l] + nums[r] == 0 {
				res = append(res,[]int{nums[i],nums[l],nums[r]})
				// 如果存在一样的结果则继续往下遍历
				for{
					if l<r && nums[l]==nums[l+1]{
						l=l+1
					}else{
						break
					}
				}
				l=l+1
				for{
					if l<r && nums[r] == nums[r-1] {
						r = r-1
					}else{
						break
					}
				}
				r=r-1
			}else if nums[i] + nums[l] + nums[r] > 0 {
				r--
			}else{
				l++
			}
		}
	}
	return res
}

