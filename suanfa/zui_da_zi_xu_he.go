package main

import "fmt"

func main()  {
	arr := []int64{-2,1,-3,4,-1,2,1,-5,4}
	//res := slotion(arr)
	res := maxSum(arr)
	fmt.Println(res)
}

func slotion(arr []int64) int64 {
	if len(arr) < 1 {
		return 0
	}
	dp := make([]int64,len(arr))
	dp[0] = arr[0]
	result := dp[0]
	for i:=1;i<len(arr);i++{
		if dp[i-1] < 0 {
			dp[i] = arr[i]
		}else{
			dp[i] = dp[i-1] + arr[i]
		}
		if result < dp[i]{
			result = dp[i]
		}
	}
	return result
}

func maxSum(nums []int64) int64{
	var temp int64
	var max int64
	max = -100

	for _, v := range nums {
		if temp < 0 {
			temp = 0
		}

		cur := temp + v
		if cur > max {
			max = cur
		}
		temp = cur
	}

	return max
}