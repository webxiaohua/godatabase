/*
【题目394】字符串编码
给定一个经过编码的字符串，返回它解码后的字符串。
编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。
你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。
此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。

示例 1：
输入：s = "3[a]2[bc]"
输出："aaabcbc"

示例 2：
输入：s = "3[a2[c]]"
输出："accaccacc"

示例 3：
输入：s = "2[abc]3[cd]ef"
输出："abcabccdcdcdef"

示例 4：
输入：s = "abc3[cd]xyz"
输出："abccdcdcdxyz"

提示：
1 <= s.length <= 30
s 由小写英文字母、数字和方括号 '[]' 组成
s 保证是一个 有效 的输入。
s 中所有整数的取值范围为 [1, 300]

*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "2[abc]3[cd]ef" // 3[a]2[bc]  3[a2[c]] 2[abc]3[cd]ef 100[leetcode]
	res := sol394_1(s)
	fmt.Println(res)
}

type Stack394 struct {
	data []string
}

func NewStack394() *Stack394 {
	return &Stack394{data: make([]string, 0)}
}
func (s *Stack394) Push(str string) {
	s.data = append(s.data, str)
}
func (s *Stack394) Pop() string {
	if len(s.data) == 0 {
		return ""
	}
	res := s.data[len(s.data)-1]
	s.data = s.data[0 : len(s.data)-1]
	return res
}
func (s *Stack394) Empty() bool {
	return len(s.data) == 0
}
func (s *Stack394) Top() string {
	if len(s.data) == 0 {
		return ""
	} else {
		return s.data[len(s.data)-1]
	}
}

// 栈求解
func sol394_1(s string) string {
	/*
			把字符串当中的字符依次循环入栈，直到找到右括号时开始出栈，
			出栈的元素遇到左括号时准备读取数量，这里注意数量的范围，
		    然后把元素拼接在一起重复N次，再压栈
			最后把栈内元素依次出栈拼接即可
	*/
	getCnt := func(stack *Stack394) int {
		cntStr := ""
		for {
			if stack.Top() == "" || stack.Top() == "[" {
				break
			}
			if stack.Top() < "0" || stack.Top() > "9" {
				break
			}
			cntStr = fmt.Sprintf("%s%s", stack.Pop(), cntStr)
		}
		res, _ := strconv.Atoi(cntStr)
		return res
	}
	out := func(stack *Stack394) string {
		res := ""
		for {
			val := stack.Pop()
			if val == "[" {
				times := getCnt(stack)
				res = strings.Repeat(res, times) // 重复N遍
				break
			} else {
				res = fmt.Sprintf("%s%s", val, res) // 前插法
			}
		}
		return res
	}
	in := func(stack *Stack394, s string) {
		res := ""
		for i, str := range s {
			if string(str) == "]" {
				res = out(stack)
				stack.Push(res)
			} else {
				stack.Push(string(str))
			}
			if i == len(s)-1 {
				break
			}
		}
		//return res
	}
	stack := NewStack394()
	// 入栈
	in(stack, s)
	// 出栈拼接
	res := ""
	for {
		if stack.Empty() {
			break
		}
		res = fmt.Sprintf("%s%s", stack.Pop(), res)
	}
	return res
}
