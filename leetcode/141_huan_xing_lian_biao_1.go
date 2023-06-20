/*
【题目141】环形链表
给你一个链表的头节点 head ，判断链表中是否有环。
如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。注意：pos 不作为参数进行传递 。仅仅是为了标识链表的实际情况。
如果链表中存在环 ，则返回 true 。 否则，返回 false 。

示例 1:
3 -> 2 -> 0 -> -4
     |__________|
输入：head = [3,2,0,-4], pos = 1
输出：true
解释：链表中有一个环，其尾部连接到第二个节点。


提示：
链表中节点的数目范围是 [0, 104]
-105 <= Node.val <= 105
pos 为 -1 或者链表中的一个 有效索引 。

*/
package main

import "fmt"

type ListNode141 struct {
	Val  int
	Next *ListNode141
}

func main() {
	node1 := &ListNode141{Val: 1}
	node2 := &ListNode141{Val: 2}
	node3 := &ListNode141{Val: 3}
	node4 := &ListNode141{Val: 4}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node2
	res := sol141_2(node1)
	fmt.Println(res)
}

// 哈希法 额外O(N)的空间
func sol141_1(head *ListNode141) bool {
	_map := make(map[*ListNode141]int)
	for {
		if head == nil {
			return false
		}
		if _, ok := _map[head]; ok {
			return true
		} else {
			_map[head] = 1
			head = head.Next
		}
	}
}

// 快慢指针 无需额外的空间占用
func sol141_2(head *ListNode141) bool {
	if head == nil || head.Next == nil {
		return false
	}
	i, j := head, head
	for {
		if i == nil || j == nil || j.Next == nil {
			return false
		}
		i = i.Next
		j = j.Next.Next
		if i == j {
			return true
		}
	}
}
