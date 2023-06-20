// 27.移除元素 你必须仅使用 O(1) 额外空间并 原地 修改输入数组。
/*
输入：nums = [3,2,2,3], val = 3
输出：2, nums = [2,2]
*/
package main

import "fmt"

func main(){
	arr := []int{0,1,2,2,3,0,4,2}
	res := removeElement(arr,2)
	fmt.Println(res)
}

// 切片法
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	cursor := 0
	for {
		num := nums[cursor]
		if num == val {
			if cursor == len(nums) - 1 {
				nums = nums[0:cursor]
				break
			}else {
				nums = append(nums[0:cursor], nums[cursor+1:]...)
			}
		}else{
			if cursor == len(nums) - 1 {
				break
			}
			cursor++
		}
	}
	return len(nums)
}
