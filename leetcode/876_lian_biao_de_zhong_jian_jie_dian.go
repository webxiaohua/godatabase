/*
【题目876】链表的中间节点
给定一个头结点为 head 的非空单链表，返回链表的中间结点。
如果有两个中间结点，则返回第二个中间结点。

示例 1：
输入：[1,2,3,4,5]
输出：此列表中的结点 3 (序列化形式：[3,4,5])
返回的结点值为 3 。 (测评系统对该结点序列化表述是 [3,4,5])。
注意，我们返回了一个 ListNode 类型的对象 ans，这样：
ans.val = 3, ans.next.val = 4, ans.next.next.val = 5, 以及 ans.next.next.next = NULL.

示例 2：
输入：[1,2,3,4,5,6]
输出：此列表中的结点 4 (序列化形式：[4,5,6])
由于该列表有两个中间结点，值分别为 3 和 4，我们返回第二个结点。

提示：
给定链表的结点数介于 1 和 100 之间。
*/
package main

import "fmt"

type ListNode876 struct {
	Val  int
	Next *ListNode876
}

func main() {
	node := &ListNode876{Val: 1, Next: &ListNode876{Val: 2, Next: &ListNode876{Val: 3, Next: &ListNode876{Val: 4, Next: &ListNode876{Val: 5}}}}}
	res := sol876_1(node)
	fmt.Println(res)
}

// 快慢指针
func sol876_1(head *ListNode876) *ListNode876 {
	fast, slow := head, head
	for {
		if fast == nil || fast.Next == nil {
			break
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
