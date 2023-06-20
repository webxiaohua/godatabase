// 387.字符串中第一个唯一字符
/*
给定一个字符串 s ，找到 它的第一个不重复的字符，并返回它的索引 。如果不存在，则返回 -1
输入: s = "loveleetcode"
输出: 2
*/
package main

import "fmt"

func main(){
	s := "loveleetcode"
	res := firstUniqChar(s)
	fmt.Println(res)
}

func firstUniqChar(s string) int {
	// 定义一个数组，考虑到小写字母只有26个，a 表示 arr[0] ，b 表示 arr[1],数组记录每个元素最后一次出现的下标
	// 再遍历一次，如果发现当前下标和最后下标不一致，则标记为-1表示为重复，一直找到第一个相同的下标
	arr := [26]int{}
	for i,c := range s {
		arr[c - 'a'] = i
	}
	for i,c := range s{
		if arr[c - 'a'] != i {
			arr[c - 'a'] = -1
		}else{
			return i
		}
	}
	return -1
}