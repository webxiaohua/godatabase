// 2.两数相加
/*
给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[7,0,8]
解释：342 + 465 = 807.
*/
package main

import "fmt"

func main(){
	l1 := &ListNode2{
		Val:  2,
		Next: &ListNode2{
			Val:  4,
			Next: &ListNode2{
				Val:  3,
				Next: nil,
			},
		},
	}
	l2 := &ListNode2{
		Val:  5,
		Next: &ListNode2{
			Val:  6,
			Next: &ListNode2{
				Val:  4,
				Next: nil,
			},
		},
	}
	res := addTwoNumbers(l1,l2)
	fmt.Println(res)
}

type ListNode2 struct {
	 Val int
	 Next *ListNode2
 }

func addTwoNumbers(l1 *ListNode2, l2 *ListNode2) *ListNode2 {
	res := &ListNode2{}
	cursor := res
	carry := 0
	for {
		if l1 == nil && l2 == nil && carry == 0 {
			break
		}
		var l1Val,l2Val int
		if l1 != nil {
			l1Val = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			l2Val = l2.Val
			l2 = l2.Next
		}
		sumVal := l1Val + l2Val + carry
		carry = sumVal / 10
		sumNode := &ListNode2{
			Val:  sumVal % 10,
			Next: nil,
		}
		cursor.Next = sumNode
		cursor = sumNode
	}
	return res.Next
}
