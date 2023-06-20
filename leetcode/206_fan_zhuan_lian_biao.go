/*
【题目206】反转链表
给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。


示例 1：
输入：head = [1,2,3,4,5]
输出：[5,4,3,2,1]

提示：
链表中节点的数目范围是 [0, 5000]
-5000 <= Node.val <= 5000

进阶：链表可以选用迭代或递归方式完成反转。你能否用两种方法解决这道题？
*/
package main

import "fmt"

type ListNode206 struct {
	Val  int
	Next *ListNode206
}

func main() {
	node := &ListNode206{Val: 1, Next: &ListNode206{Val: 2, Next: &ListNode206{Val: 3, Next: &ListNode206{Val: 4}}}}
	res := sol206_2(node)
	fmt.Println(res)
}

// 双指针循环法，记住上一个节点和下一个节点，需要借助临时变量处理断链问题，两两串联
func sol206_1(head *ListNode206) *ListNode206 {
	/*
		1->2->3->4->5
		nil<-1 2->3->4->5 (cur:1 next:2)
		nil<-1<-2 3->4->5 (cur:2 next:3)
		nil<-1<-2<-3 4->5 (cur:3 next:4)
		nil<-1<-2<-3<-4 5 (cur:4 next:5)
		nil<-1<-2<-3<-4<-5 (cur:5 next:nil)
	*/
	var current *ListNode206
	next := head
	for {
		if next == nil {
			break
		}
		tmp := next
		next = next.Next
		tmp.Next = current
		current = tmp
	}
	return current
}

// 递归法
func sol206_2(head *ListNode206) *ListNode206 {
	/*
			1.先一路到头，找到链表尾节点
			2.返回的过程中，将当前节点的下一个节点的next指针指向当前节点  (2->3 | res:3 head:2 | 2->3->2)
		    3.当前节点的next指向nil，实现局部反转 (2->3 | res:3 head:2 | 3->2->nil)
			4.当递归函数完成全部出栈后，得出解
			1 ->
				2 ->
					3
				2 <-
			1 <-
	*/
	if head == nil || head.Next == nil {
		return head
	}
	ret := sol206_2(head.Next) // ret = head.Next
	head.Next.Next = head
	head.Next = nil
	return ret
}
