// 9.回文数
/*
回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
输入：x = 121
输出：true
输入：x = -121
输出：false
*/
package main

import (
	"fmt"
	"strconv"
)

func main(){
	res := isPalindrome(-121)
	fmt.Println(res)
}

func isPalindrome(x int) bool {
	// 转string，再利用双指针比较
	str := strconv.Itoa(x)
	if len(str) == 1 {
		return true
	}
	left,right := 0,len(str)-1
	for {
		if left >= right {
			return true
		}
		if str[left] != str[right] {
			return false
		}
		left++
		right--
	}
}