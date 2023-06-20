// 20.有效的括号
/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
输入：s = "()[]{}"
输出：true
*/
package main

import (
	"fmt"
	"strings"
)

func main(){
	a:="GO语言编程"
	fmt.Println([]byte(a))
	fmt.Println([]rune(a))

	str := "()[]"
	//res := isValid20(str)
	res := isValid20_2(str)
	fmt.Println(res)
}

func isValid20(s string) bool {
	// 思路：利用栈的特性，判断如果是左边括号就入栈，如果是右边括号就出栈并且和当前右括号匹配，最终栈不为空
	tmpMap := map[string]string{
		")":"(",
		"}":"{",
		"]":"[",
	}
	stack := make([]string,0)
	for i:=0;i<len(s);i++ {
		tmpChar := string(s[i])
		if val,ok := tmpMap[tmpChar];ok {
			if len(stack) == 0 {
				return false
			}
			// 弹出栈顶元素
			topChar := stack[len(stack)-1]
			if val != topChar {
				return false
			}
			stack = stack[0:len(stack)-1]
		}else{
			// 入栈
			stack = append(stack,tmpChar)
		}
	}
	if len(stack) > 0 {
		return false
	}
	return true
}

func isValid20_2(s string) bool {
	for {
		if strings.Contains(s,"{}") {
			s = strings.ReplaceAll(s,"{}","")
		}else if strings.Contains(s,"[]") {
			s = strings.ReplaceAll(s,"[]","")
		}else if strings.Contains(s,"()") {
			s = strings.ReplaceAll(s,"()","")
		}else{
			return len(s) == 0
		}
	}
}