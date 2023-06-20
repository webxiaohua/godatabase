// 300.最长递增子序列
/*
输入：nums = [10,9,2,5,3,7,101,18]
输出：4
解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。

1 <= nums.length <= 2500
*/
package main

import "fmt"

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println(getMaxSubList(nums))
	//fmt.Println(test(nums))
}

func test(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	global := make([]int, len(nums))
	local := make([]int, len(nums))
	global[0] = 1
	local[0] = 1
	for i := 1; i < len(nums); i++ {
		if nums[i] >= nums[i-1] {
			local[i] = local[i-1] + 1
			if local[i] > global[i-1] {
				global[i] = local[i]
			} else {
				global[i] = global[i-1]
			}
		} else {
			local[i] = 1
			global[i] = global[i-1]
		}
	}
	return global[len(nums)-1]
}

func getMaxSubList(nums []int) int {
	// 状态方程 dp[n] 表示下标n的最长子序列值
	// dp[n] = max(dp[a]+1,dp[b]+1,...)
	// nums[n] > nums[a] && nums[n] > nums[b] && ...
	if len(nums) == 1 {
		return 1
	}
	dp := make([]int, len(nums))
	maxVal := 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := i - 1; j >= 0; j-- {
			if nums[j] < nums[i] {
				dp[i] = getMax(dp[j]+1, dp[i])
			}
		}
		maxVal = getMax(maxVal, dp[i])
	}
	return maxVal
}

func getMax(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

// 栈求解 由低到高
