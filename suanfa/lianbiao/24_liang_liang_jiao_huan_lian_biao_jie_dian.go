// 24.两两交换链表中的节点
/*
输入：head = [1,2,3,4]
输出：[2,1,4,3]
*/
package main

import "fmt"


type ListNode24 struct {
	Val int
	Next *ListNode24
}

func main(){
	head := &ListNode24{
		Val:  1,
		Next: &ListNode24{
			Val:  2,
			Next: &ListNode24{
				Val:  3,
				Next: &ListNode24{
					Val:  4,
					Next: nil,
				},
			},
		},
	}
	res := swapPairs(head)
	fmt.Println(res)
}

func swapPairs(head *ListNode24) *ListNode24 {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next // 第二个节点
	head.Next = swapPairs(newHead.Next) // 交换两个节点 head此时会作为第二个节点，
	newHead.Next = head // 第二个节点和第一个节点互换
	return newHead // 第二个节点作为头节点返回
}