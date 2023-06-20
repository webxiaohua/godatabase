// 70.爬楼梯
/*
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

输入：n = 3
输出：3
解释：有三种方法可以爬到楼顶。
1. 1 阶 + 1 阶 + 1 阶
2. 1 阶 + 2 阶
3. 2 阶 + 1 阶

1 <= n <= 45
*/
package main

import "fmt"

/*
【思路】
数列推导：
1	1
2	11
3	111	12	21
4	1111	121		211		22		112
5	11111	1112	1121	1211	122		2111	212		221

1:1
2:2
3:3
4:5
5:8


f(1) = 1
f(2) = 2
f(n) = f(n-1) + f(n-2)
*/
var tmpMap70 map[int]int

func main() {
	//res := sol70_1(45)
	/*tmpMap = make(map[int]int)
	res := sol70_2(45)*/
	//res := sol70_3(45)
	res := sol70_4(45)
	fmt.Println(res)

}

// 1.递归法 (复杂度太高，O(N2)，存在大量重复计算)
func sol70_1(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return sol70_1(n-1) + sol70_1(n-2)
}

// 2.递归借助map，减少重复计算，复杂度降低为 O(N)
func sol70_2(n int) int {
	if n == 1 {
		tmpMap70[1] = 1
		return 1
	}
	if n == 2 {
		tmpMap70[2] = 2
		return 2
	}
	if _, ok := tmpMap70[n]; ok {
		return tmpMap70[n]
	} else {
		tmpMap70[n] = sol70_1(n-1) + sol70_1(n-2)
		return tmpMap70[n]
	}
}

// 3.循环法
func sol70_3(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	prev := 2
	prevPrev := 1
	var res int
	for i := 3; i <= n; i++ {
		res = prev + prevPrev
		prev = res
		prevPrev = res - prevPrev
	}
	return res
}

// 4.动态规划
func sol70_4(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
