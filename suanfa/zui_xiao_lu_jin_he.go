// 最小路径和
// 给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
package main

import "fmt"

func main()  {
	arr := [][]int{
		{1,3,1},
		{1,5,1},
		{4,2,1},
	}
	result := solution20220403(arr)
	fmt.Println(result)
}

func solution20220403(grid [][]int) int{
	// 状态方程： dp[i][j] =  min((dp[i-1][j] + arr[i][j]),(dp[i][j-1] + arr[i][j]))
	// 特殊情况： 空数组、只有1行、最左侧只能从上方元素累加
	if len(grid) < 0 {
		return 0
	}
	for i:=0;i<len(grid);i++ {
		for j:=0;j<len(grid[i]);j++{
			if i == 0 && j != 0{
				grid[i][j] = grid[i][j-1] + grid[i][j]
			}else if j == 0 && i != 0 {
				grid[i][j] = grid[i-1][j]+grid[i][j]
			}else if i!=0 && j!=0 {
				if grid[i][j-1] <= grid[i-1][j]{
					grid[i][j] = grid[i][j-1] + grid[i][j]
				}else{
					grid[i][j] = grid[i-1][j]+grid[i][j]
				}
			}
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}
