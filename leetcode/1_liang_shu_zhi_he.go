/*
【题目1】两数之和

给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
你可以按任意顺序返回答案。

输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。

提示：
2 <= nums.length <= 104
-109 <= nums[i] <= 109
-109 <= target <= 109
只会存在一个有效答案

进阶：你可以想出一个时间复杂度小于 O(n2) 的算法吗？


*/

package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	//res := sol1_1(nums, target)
	res := sol1_2(nums, target)
	fmt.Println(res)
}

// 1.暴力穷举，时间复杂度 O(N2)
func sol1_1(nums []int, target int) []int {
	res := []int{0, 0}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				res = []int{i, j}
			}
		}
	}
	return res
}

// 2.引入map，反向匹配
func sol1_2(nums []int, target int) []int {
	// val : index
	_map := make(map[int]int)
	for i, v := range nums {
		if i2, ok := _map[target-v]; ok {
			return []int{i, i2}
		}
		_map[v] = i
	}
	return nil
}
