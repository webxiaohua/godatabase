/*
【题目83】删除排序链表中的重复元素
给定一个已排序的链表的头 head ， 删除所有重复的元素，使每个元素只出现一次 。返回 已排序的链表 。

示例1：
输入：head = [1,1,2]
输出：[1,2]

示例2：
输入：head = [1,1,2,3,3]
输出：[1,2,3]

提示：
链表中节点数目在范围 [0, 300] 内
-100 <= Node.val <= 100
题目数据保证链表已经按升序 排列

*/

package main

import "fmt"

func main() {
	head := &ListNode83{Val: 1, Next: &ListNode83{Val: 1, Next: &ListNode83{Val: 2, Next: &ListNode83{Val: 3, Next: &ListNode83{Val: 3, Next: nil}}}}}
	res := sol83_2(head)
	for {
		if res == nil {
			return
		}
		fmt.Println(res.Val)
		res = res.Next
	}
}

type ListNode83 struct {
	Val  int
	Next *ListNode83
}

// 递归法
func sol83_1(head *ListNode83) *ListNode83 {
	// 假设最多只有1个值
	if head == nil || head.Next == nil {
		return head
	}
	// 如果相同,则往后选择
	if head.Val == head.Next.Val {
		head.Next = sol83_1(head.Next.Next)
	} else {
		head.Next = sol83_1(head.Next)
	}
	return head
}

// 循环法
func sol83_2(head *ListNode83) *ListNode83 {
	if head == nil {
		return head
	}
	cur := head
	for {
		if cur.Next == nil {
			break
		}
		if cur.Next.Val == cur.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}
