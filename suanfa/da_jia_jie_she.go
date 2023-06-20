// 打家劫舍
// 你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
package main

import "fmt"

func main()  {
	arr := []int{2,1,1,2}
	result := rob(arr)
	fmt.Println(result)
}

func rob(nums []int) int {
	// 定义状态转移方程：偷到第i个房子时所获得利益最大值  dp[i] = max(dp[i-2] + nums[i] , dp[i-1])
	// 特殊情况：不满足2时
	if len(nums) == 0 {
		return 0
	}else if len(nums) == 1 {
		return nums[0]
	}else if len(nums) == 2 {
		return max(nums[0],nums[1])
	}
	nums[1] = max(nums[0],nums[1])
	for i:=2;i<len(nums);i++{
		nums[i] = max(nums[i-2]+nums[i],nums[i-1])
	}
	return nums[len(nums)-1]
}

func max(a,b int) int {
	if a>= b {
		return a
	}
	return b
}