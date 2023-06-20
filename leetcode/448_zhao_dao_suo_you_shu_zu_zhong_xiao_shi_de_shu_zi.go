/*
【题目448】找到所有数组中消失的数字
给你一个含 n 个整数的数组 nums ，其中 nums[i] 在区间 [1, n] 内。请你找出所有在 [1, n] 范围内但没有出现在 nums 中的数字，并以数组的形式返回结果。

示例 1：
输入：nums = [4,3,2,7,8,2,3,1]
输出：[5,6]

示例 2：
输入：nums = [1,1]
输出：[2]


提示：
n == nums.length
1 <= n <= 105
1 <= nums[i] <= n

进阶：你能在不使用额外空间且时间复杂度为 O(n) 的情况下解决这个问题吗? 你可以假定返回的数组不算在额外空间内。
*/
package main

import "fmt"

func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	res := sol_448_1(nums)
	fmt.Println(res)
}

func sol_448_1(nums []int) []int {
	n := len(nums)
	for _, num := range nums {
		x := (num - 1) % n // 取模后还是原来的下标： val - 1
		nums[x] += n
	}
	res := make([]int, 0)
	// 如果值不大于总长度，则表示此下标不存在
	for i, num := range nums {
		if num <= n {
			res = append(res, i+1)
		}
	}
	return res
}
