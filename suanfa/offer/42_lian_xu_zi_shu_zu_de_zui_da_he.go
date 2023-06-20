/*
连续子数组的最大和
输入一个整型数组，数组中的一个或连续多个整数组成一个子数组。求所有子数组的和的最大值。
要求时间复杂度为O(n)。

输入: nums = [-2,1,-3,4,-1,2,1,-5,4]
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
*/
package main

import "fmt"

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	//res := maxSubArray2(nums)
	//_, _, _, res := get2(nums, 0, len(nums)-1)
	res := get3(nums)
	fmt.Println(res)
}

// 暴力解法， 复杂度 O(n^2)
func maxSubArray(nums []int) int {
	cnt := len(nums)
	maxVal := 0
	for i := 0; i < cnt; i++ {
		tmpVal := 0
		for j := i; j < cnt; j++ {
			tmpVal += nums[j]
			if maxVal < tmpVal {
				maxVal = tmpVal
			}
		}
	}
	return maxVal
}

// 分治法，归并思想
// 一分为二，递归地求解子列和
// lSum 表示以l为开始下标的最大和， left_sum or left_isum + right_lsum
// rSum 表示右区间和  right_sum or right_isum + left_rsum
// iSum 表示全区间和  left_isum + right_isum
// mSum 表示最大和	max(lSum,rSum,iSum)

func pushUp(l, r Status) Status {
	iSum := l.iSum + r.iSum
	lSum := max(l.lSum, l.iSum+r.lSum)
	rSum := max(r.rSum, r.iSum+l.rSum)
	mSum := max(max(l.mSum, r.mSum), l.rSum+r.lSum)
	return Status{lSum, rSum, mSum, iSum}
}

func get(nums []int, l, r int) Status {
	if l == r {
		return Status{nums[l], nums[l], nums[l], nums[l]}
	}
	// 分
	m := (l + r) >> 1
	lSub := get(nums, l, m)
	rSub := get(nums, m+1, r)
	// 治
	return pushUp(lSub, rSub)
}

type Status struct {
	// 左区间和、右区间和、最大和、全区间和
	lSum, rSum, mSum, iSum int
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func get2(nums []int, l, r int) (left, right, sumVal, maxVal int) {
	if l == r {
		return nums[l], nums[l], nums[l], nums[l]
	}
	// 分
	middle := (l + r) / 2
	left1, right1, sum1, max1 := get2(nums, l, middle)
	left2, right2, sum2, max2 := get2(nums, middle+1, r)
	// 合
	sumVal = sum1 + sum2
	left = max(left1, sum1+left2)
	right = max(right2, sum2+right1)
	maxVal = max(max(max1, max2), right1+left2)
	return
}

// 动态规划
// f(i) = max(f(i-1)+nums[i],nums[i])
func get3(nums []int) int {
	maxVal := nums[0]
	fn := make([]int, len(nums))
	for i, num := range nums {
		if i == 0 {
			fn[i] = num
		} else {
			fn[i] = max(fn[i-1]+nums[i], nums[i])
			if fn[i] > maxVal {
				maxVal = fn[i]
			}
		}
	}
	return maxVal
}
