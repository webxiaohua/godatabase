package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 5}}}
	l2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
	res := sol_21_3(l1, l2)
	fmt.Println(res)
}

func sol_21_3(l1 *ListNode, l2 *ListNode) *ListNode {
	// 合并两个有序链表
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val <= l2.Val {
		l1.Next = sol_21_3(l1.Next, l2)
		return l1
	} else {
		l2.Next = sol_21_3(l2.Next, l1)
		return l2
	}
}
