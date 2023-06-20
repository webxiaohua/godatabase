/*
【题目21】合并两个有序链表
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例 1：
输入：l1 = [1,2,4], l2 = [1,3,4]
输出：[1,1,2,3,4,4]

示例 2：
输入：l1 = [], l2 = []
输出：[]

示例 3：
输入：l1 = [], l2 = [0]
输出：[0]

提示：
两个链表的节点数目范围是 [0, 50]
-100 <= Node.val <= 100
l1 和 l2 均按 非递减顺序 排列
*/

package main

import "fmt"

func main() {
	l1 := &ListNode21{Val: 1, Next: &ListNode21{Val: 2, Next: &ListNode21{Val: 5}}}
	l2 := &ListNode21{Val: 1, Next: &ListNode21{Val: 3, Next: &ListNode21{Val: 4}}}
	res := sol_21_2(l1, l2)
	for {
		if res != nil {
			fmt.Println(res.Val)
			res = res.Next
		} else {
			break
		}
	}
}

type ListNode21 struct {
	Val  int
	Next *ListNode21
}

// 循环双指针，代码复杂，不易读
func sol_21_1(l1 *ListNode21, l2 *ListNode21) (res *ListNode21) {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	res = &ListNode21{}
	i, j, k := l1, l2, res
	for {
		if i.Val <= j.Val {
			k.Next = i
			k = k.Next
			i = i.Next
		} else {
			k.Next = j
			k = k.Next
			j = j.Next
		}
		if i == nil {
			k.Next = j
			break
		}
		if j == nil {
			k.Next = i
			break
		}
	}
	return res.Next
}

// 递归法 更优雅
func sol_21_2(l1 *ListNode21, l2 *ListNode21) *ListNode21 {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val <= l2.Val {
		l1.Next = sol_21_2(l1.Next, l2)
		return l1
	} else {
		l2.Next = sol_21_2(l1, l2.Next)
		return l2
	}
}
