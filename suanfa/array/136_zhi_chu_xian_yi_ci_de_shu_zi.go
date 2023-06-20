// 136.只出现一次的数字
/*
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
输入: [2,2,1]
输出: 1
*/
package main

import "fmt"

func main(){
	arr := []int{2,2,1}
	res := singleNumber2(arr)
	fmt.Println(res)
}

// 引入额外空间解决，时间复杂度 O(n)
func singleNumber(nums []int) int {
	tmpMap := make(map[int]int)
	for _,num := range nums {
		if _,ok := tmpMap[num];ok {
			delete(tmpMap,num)
		}else{
			tmpMap[num] = num
		}
	}
	for key,_ := range tmpMap{
		return key
	}
	return 0
}

// 不引入额外空间计算，采用位运算（0^n=n; n^n=0）
func singleNumber2(nums []int) int {
	res := nums[0]
	for i:=1;i<len(nums);i++{
		res = res ^ nums[i]
	}
	return res
}