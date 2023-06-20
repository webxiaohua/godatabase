package main

import "fmt"

type ListNode struct {
  	 Val int
     Next *ListNode
}

// 循环法
func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	// 定义预先指针，用于返回结果
	root := &ListNode{Val:0}
	// 移动指针
	cursor := root
	carry := 0
	for l1 != nil || l2 != nil || carry != 0 {
		var l1Val,l2Val int
		if l1 != nil {
			l1Val = l1.Val
		}
		if l2 != nil {
			l2Val = l2.Val
		}
		sumVal := l1Val + l2Val + carry
		carry = sumVal / 10
		sumNode := &ListNode{Val:sumVal % 10}
		// 第1轮相当于是 root.Next 第2轮相当于 root.Next.Next
		cursor.Next = sumNode
		// 第1轮相当于 cursor = root.Next  这里理解成地址比较好
		cursor = sumNode
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	return root.Next
}

func main()  {
	l1 := &ListNode{
		Val:  9,
		Next: &ListNode{
			Val:  9,
			Next: &ListNode{
				Val:  9,
				Next: nil,
			},
		},
	}
	l2 := &ListNode{
		Val:  9,
		Next: &ListNode{
			Val:  9,
			Next: &ListNode{
				Val:  9,
				Next: &ListNode{
					Val:  9,
					Next: nil,
				},
			},
		},
	}
	l3 := addTwoNumbers1(l1,l2)
	fmt.Println(l3)
	//l4:=addTwoNumbers2(l1,l2)
	//fmt.Println(l4)
	fmt.Println(`{"A":"aaa"}`)
}