// 28.实现strstr
/*
给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串出现的第一个位置（下标从 0 开始）。如果不存在，则返回  -1 。
输入：haystack = "hello", needle = "ll"
输出：2
*/
package main

import "fmt"

func main(){
	res := strStr("aaa","aaaa")
	fmt.Println(res)
}

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return -1
	}
	// 找到第一个，进入二级循环
	for i:=0;i<len(haystack);i++{
		if haystack[i] == needle[0] {
			for j:=i;j<i+len(needle);j++{
				if j >= len(haystack) {
					return -1
				}
				if haystack[j] != needle[j-i] {
					break
				}
				if j+1 == i+len(needle) {
					return i
				}
			}
		}
	}
	return -1
}