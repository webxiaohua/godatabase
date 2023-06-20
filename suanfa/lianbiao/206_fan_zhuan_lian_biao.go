// 206 反转链表
/*
给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
输入：head = [1,2,3,4,5]
输出：[5,4,3,2,1]
*/
package main

import "fmt"

type Node206 struct {
	Val int
	Next *Node206
}

func main(){
	node := &Node206{
		Val:  1,
		Next: &Node206{
			Val:  2,
			Next: &Node206{
				Val:  3,
				Next: &Node206{
					Val:  4,
					Next: nil,
				},
			},
		},
	}
	res := reverseDigui(node)
	fmt.Println(res)
}

// 递归处理
func reverseDigui(head *Node206) *Node206 {
	// 退出条件
	if head == nil || head.Next == nil {
		return head
	}
	ret := reverseDigui(head.Next)
	head.Next.Next = head // ret.Next = head
	head.Next = nil
	return ret
}

// 双指针
func reverseDoublePointer(head *Node206) *Node206 {

	return nil
}