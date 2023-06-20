// 最长上升子序列
package main

import "fmt"

// 给定一个无序的整数数组，找到其中最长上升子序列的长度。
func main()  {
	arr := []int64{10,9,2,5,3,7,101,18}
	result := sloution(arr)
	fmt.Println(result)
}

func sloution(arr []int64) int64 {
	if len(arr) == 0 {
		return 0
	}
	if len(arr) == 1 {
		return 1
	}
	dp := make([]int64,len(arr)) // 用来记录下标对应的最长序列
	var result int64
	// 向前找到所有比自己小的元素，在对应下标的dp上加1，一轮比较下来取最大值
	result = 1
	dp[0] = 1
	for i:=1;i<len(arr);i++ {
		dp[i] = 1
		for j:=0;j<i;j++ {
			if arr[i] > arr[j] {
				if dp[i] < (dp[j] + 1) {
					dp[i] = dp[j] + 1
				}
			}
		}
		if dp[i] > result {
			result = dp[i]
		}
	}
	return result
}