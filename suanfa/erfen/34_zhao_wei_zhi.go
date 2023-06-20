// 34.在排序数组中查找元素的第一个和最后一个位置
package main

import "fmt"

func main()  {
	arr :=[]int{1,1,1}
	fmt.Println(searchRange(arr,1))
}

func searchRange(nums []int, target int) []int {
	idx := doIt(nums,target,0)
	if idx == -1 {
		return []int{-1,-1}
	}
	// 左右探测
	left := idx
	right := idx
	for left > 0 {
		if nums[left-1] == target {
			left--
		}else{
			break
		}
	}
	for right < len(nums)-1 {
		if nums[right+1] == target {
			right++
		}else{
			break
		}
	}
	return []int{left,right}
}
// 找到目标值的位置
func doIt(arr []int, target,tmp int) int {
	if len(arr) == 0 {
		return -1
	}
	half := len(arr) / 2
	if arr[half] > target {
		if half == 0 {
			return -1
		}
		arr = arr[0:half]
		return doIt(arr,target,tmp)
	}else if arr[half] < target {
		if half == 0 {
			return -1
		}
		arr = arr[half:]
		tmp+=half
		return doIt(arr,target,tmp)
	}else{
		return tmp+half
	}
}

