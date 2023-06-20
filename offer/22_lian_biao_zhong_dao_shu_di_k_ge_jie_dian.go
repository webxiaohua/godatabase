/*
【题目22】链表中倒数第K个节点
输入一个链表，输出该链表中倒数第K个节点。

示例 1：
输入：[1,2,3,4,5] k=2
输出：4->5

提示：
给定链表的结点数介于 1 和 100 之间。
*/
package main

import "fmt"

type ListNode22 struct {
	Val  int
	Next *ListNode22
}

func main() {
	node := &ListNode22{Val: 1, Next: &ListNode22{Val: 2, Next: &ListNode22{Val: 3, Next: &ListNode22{Val: 4, Next: &ListNode22{Val: 5}}}}}
	res := sol22_2(node, 3)
	fmt.Println(res)
}

func sol22_1(node *ListNode22, k int) *ListNode22 {
	// 倒数第K个节点相当于正数第 len(node) - k 个节点，求出链表长度，从头开始遍历即可
	total := 0
	first := node
	for {
		if first == nil {
			break
		}
		total++
		first = first.Next
	}
	target := total - k
	i := 0
	for {
		if i == target {
			break
		}
		i++
		node = node.Next
	}
	return node
}

func sol22_2(node *ListNode22, k int) *ListNode22 {
	/*
		公式推导，设总长为n，求解为x，则有
		x = n-k+1
		n = x+k-1
		所以当快指针走了 k-1 步时，剩余x步，而慢指针这个时候出发正好走了x步后快指针到达重点，慢指针即为求解
	*/
	fast, slow := node, node
	counter := 0
	for {
		if counter > k-1 {
			slow = slow.Next
		}
		if fast.Next == nil {
			return slow
		}
		counter++
		fast = fast.Next
	}
}
