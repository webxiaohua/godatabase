// 3.无重复字符的最长子串
/*
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
*/
package main

import "fmt"

func main()  {
	s := "abccbaa"
	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	// 记录当前不重复队列的元素索引 元素:最新出现位置索引
	tmpMap := make(map[byte]int)
	max := 0
	// 滑动窗口偏移单位
	left := 0
	for i:=0;i<len(s);i++{
		if _,ok := tmpMap[s[i]];ok{
			// 出现重复元素，需要滑动，记录最大的滑动值
			left = maxFunc(tmpMap[s[i]]+1,left)
		}
		tmpMap[s[i]] = i
		max = maxFunc(max,i - left + 1)
	}
	return max
}
func maxFunc(x,y int) int{
	if x>=y{
		return x
	}
	return y
}

// abcabc
func wuchongfu3(a string) int {
	max := 0
	left := 0
	tmpMap := make(map[byte]int)
	for i:=0;i<len(a);i++ {
		if _,ok := tmpMap[a[i]];ok{
			left = maxFunc(left,tmpMap[a[i]] + 1)
		}
		tmpMap[a[i]] = i
		max = maxFunc(max,i-left+1)
	}
	return max
}