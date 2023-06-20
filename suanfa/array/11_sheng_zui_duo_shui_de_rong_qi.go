// 11.盛最多水的容器
/*
给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。
找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
返回容器可以储存的最大水量。
输入：[1,8,6,2,5,4,8,3,7]
输出：49
*/
package main

import "fmt"

func main(){
	arr := []int{1,8,6,2,5,4,8,3,7}
	res := maxArea2(arr)
	fmt.Println(res)
}

// 双指针法
func maxArea2(height []int) int {
	max := 0
	left,right := 0,len(height)-1
	for {
		if left >= right {
			break
		}
		if height[left] < height[right] {
			max = getMax(max,height[left] * (right-left))
			left++
		}else{
			max = getMax(max,height[right]*(right-left))
			right--
		}
	}
	return max
}

func maxArea(height []int) int {
	max := 0
	for i:=0;i<len(height);i++ {
		for j:=1;j<len(height);j++{
			if height[i] < height[j] {
				max = getMax(max,height[i] * (j-i))
			}else{
				max = getMax(max,height[j]*(j-i))
			}
		}
	}
	return max
}

func getMax(a,b int) int {
	if a>b {
		return a
	}else{
		return b
	}
}