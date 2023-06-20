/*
【题目234】回文链表
给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。

示例 1：
输入：head = [1,2,2,1]
输出：true

提示：
链表中节点数目在范围[1, 105] 内
0 <= Node.val <= 9

进阶：你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
*/
package main

import "fmt"

type ListNode234 struct {
	Val  int
	Next *ListNode234
}

func main() {
	node := &ListNode234{Val: 1, Next: &ListNode234{Val: 2, Next: &ListNode234{Val: 2, Next: &ListNode234{Val: 1}}}}
	res := sol_234_2(node)
	fmt.Println(res)
}

// 链表转数组，再对数组求解，时间复杂度O(N)，需要额外的O(N)空间
func sol_234_1(head *ListNode234) bool {
	arr := make([]int, 0)
	for {
		if head == nil {
			break
		}
		arr = append(arr, head.Val)
		head = head.Next
	}
	i, j := 0, len(arr)-1
	for {
		if i >= j {
			return true
		}
		if arr[i] != arr[j] {
			return false
		}
		i++
		j--
	}
}

// 快慢指针 + 局部反转链表，时间复杂度 O(N)，空间复杂度O(1)
func sol_234_2(head *ListNode234) bool {
	// 快指针每次移动2步，慢指针每次移动1步
	fast, slow := head, head
	// 偶数个节点 123321  fast:nil,slow:3 ，将slow指针作为反转链表的头节点
	// 奇数个节点 12321   fast:1,slow:3 ，将slow指针作为反转链表的头节点，然后从头节点开始比较
	for {
		if fast == nil || fast.Next == nil {
			break
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	fast = reverse_234(slow)
	slow = head
	for {
		if fast == nil {
			break
		}
		if fast.Val != slow.Val {
			return false
		}
		fast = fast.Next
		slow = slow.Next
	}
	return true
}

// 递归求解
func reverse_234(head *ListNode234) (res *ListNode234) {
	// 找到最后一个节点
	if head == nil || head.Next == nil {
		return head
	}
	res = reverse_234(head.Next) // 1
	head.Next.Next = head        // 2->1 => 2->1->2 | 1->2
	head.Next = nil              // 2->1 | 2
	return res
}
