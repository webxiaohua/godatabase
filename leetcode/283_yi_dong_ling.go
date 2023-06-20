/*
【题目283】移动零
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
请注意 ，必须在不复制数组的情况下原地对数组进行操作。

示例 1:
输入: nums = [0,1,0,3,12]
输出: [1,3,12,0,0]

示例 2:
输入: nums = [0]
输出: [0]

提示:
1 <= nums.length <= 104
-231 <= nums[i] <= 231 - 1

进阶：你能尽量减少完成的操作次数吗？

*/
package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 3, 12}
	sol_283_1(nums)
	fmt.Println(nums)
}

// 遍历数组，如果值不为0，则将指针前移动1位，并且将指针前一位设置位当前值，如果为0，指针不动，如此遍历完成后 [0,j) 区间都是不为0的元素，[j:]都是0
func sol_283_1(nums []int) {
	j := 0
	for _, num := range nums {
		if num != 0 {
			j += 1
			nums[j-1] = num
		}
	}
	for j < len(nums) {
		nums[j] = 0
		j++
	}
}
