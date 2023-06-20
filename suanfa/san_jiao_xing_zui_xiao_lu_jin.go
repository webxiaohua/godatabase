// 三角形最小路径之和
// 给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。
package main

import "fmt"

func main() {
	arr := [][]int{
		{2},
		{3,4},
		{6,5,7},
		{4,1,8,3},
	}
	result := solution(arr)
	fmt.Println(result)
}

func solution(arr [][]int) int {
	// 如果arr长度为0，返回0
	// 如果arr长度为1，返回1
	// 状态方程：dp[i][j] = min(dp[i-1][j-1],dp[i-1][j]) + arr[i][j]，每个数组第一位 dp[i][0] = dp[i-1][0] + arr[i][0]; 每个数组最末位 dp[i][j] = dp[i-1][j-1] + arr[i][j]
	// 特殊情况1：dp[0][0] = arr[0][0]
	// 特殊情况2：dp[1][0] = arr[0][0] + arr[1][0]
	// 特殊情况3：dp[1][1] = arr[0][0] + arr[1][1]
	result := 0
	length := len(arr)
	if length == 0 {
		return result
	}else if length == 1 {
		return arr[0][0]
	}
	dp := make([][]int,length)
	dp[0] = make([]int,1)
	dp[0][0] = arr[0][0]
	dp[1] = make([]int,2)
	dp[1][0] = arr[1][0] + dp[0][0]
	dp[1][1] = arr[1][1] + dp[0][0]
	for i:=2;i<length;i++ {
		subLength := len(arr[i])
		dp[i] = make([]int,subLength)
		for j:=0;j<subLength;j++ {
			if j == 0 {
				dp[i][j] = dp[i-1][j] + arr[i][j]
			}else if j == (subLength - 1) {
				dp[i][j] = dp[i-1][j-1] + arr[i][j]
			}else{
				if dp[i-1][j-1] <= dp[i-1][j] {
					dp[i][j] = dp[i-1][j-1] + arr[i][j]
				}else{
					dp[i][j] = dp[i-1][j] + arr[i][j]
				}
			}
		}
	}
	// 求最小值
	result = dp[length-1][0]
	for i:=1;i<len(dp[length-1]);i++{
		if result > dp[length-1][i] {
			result = dp[length-1][i]
		}
	}
	return result
}
