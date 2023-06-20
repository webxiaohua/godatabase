/**
【题目42】接雨水
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
示例1:
输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。

示例2:
输入：height = [4,2,0,3,2,5]
输出：9

提示：
n == height.length
1 <= n <= 2 * 104
0 <= height[i] <= 105
**/

package main

import "fmt"

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	res := sol_42_2(height)
	fmt.Println(res)
}

// 按行求解，时间复杂度 O(m*n)
func sol_42_1(height []int) int {
	// 求出最高值，按行遍历，
	// 1.当前列高度 < 行，且当前未开始累计，则跳过；若已经开始累计，则累积雨水 tmp
	// 2.当前列高度 >= 行高，若未开始累计，则标记开始累计；若已经开始累计，则把结果追加，同时标记不开始累计 & tmp 置为0
	maxHeight := sol_42_1_max(height)
	var ans, tmp int
	var isStart bool
	for i := 1; i <= maxHeight; i++ {
		tmp = 0
		isStart = false
		for j := 0; j < len(height); j++ {
			if height[j] < i {
				if isStart {
					tmp++
				} else {
					continue
				}
			} else {
				if isStart {
					ans += tmp
					tmp = 0
				} else {
					isStart = true
				}
			}
		}
	}
	return ans
}
func sol_42_1_max(arr []int) int {
	max := arr[0]
	for _, i := range arr {
		if max < i {
			max = i
		}
	}
	return max
}

// 按列求解,时间复杂度 O(N^2)
func sol_42_2(height []int) int {
	// 分为3个指针，分别代表 左边最高的墙、正在求解的列、右边最高的墙
	// 根据短板效应，找到两边最矮的墙，跟正在求解的列进行比对
	// 如果 最矮的墙 > 正在求解的列，则能够存储  最矮的墙 - 正在求解的列 的雨水
	// 如果 最矮的墙 = 正在求解的列，不存储雨水
	// 如果 最矮的墙 < 正在求解的列，不存储雨水
	sum := 0
	for i := 1; i < len(height)-1; i++ {
		// 找到左边最高的墙
		maxLeft := 0
		for j := i - 1; j >= 0; j-- {
			if height[j] > maxLeft {
				maxLeft = height[j]
			}
		}
		// 找到右边最高的墙
		maxRight := 0
		for k := i + 1; k < len(height); k++ {
			if height[k] > maxRight {
				maxRight = height[k]
			}
		}
		// 获取最矮的墙
		min := maxLeft
		if maxRight < maxLeft {
			min = maxRight
		}
		// 求雨水
		if min > height[i] {
			sum = sum + (min - height[i])
		}
	}
	return sum
}
