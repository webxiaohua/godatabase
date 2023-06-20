// 141.环形链表
/*
如果链表中存在环 ，则返回 true 。 否则，返回 false 。
输入：head = [3,2,0,-4], pos = 1
输出：true
解释：链表中有一个环，其尾部连接到第二个节点。
*/
package main

import "fmt"

func main(){
	node1 := &ListNode141{
		Val:  1,
	}
	node2 := &ListNode141{
		Val:  2,
	}
	node1.Next = node2
	node2.Next = node1
	/*
	node3 = &ListNode141{
		Val:  1,
		Next: node4,
	}
	node4 = &ListNode141{
		Val:  1,
		Next: node2,
	}*/
	res := hasCycle2(node1)
	fmt.Println(res)
}

type ListNode141 struct {
	 Val int
	 Next *ListNode141
 }

func hasCycle(head *ListNode141) bool {
	// 哈希表
	tmpMap := make(map[*ListNode141]int)
	for {
		if head == nil {
			return false
		}
		if _,ok := tmpMap[head];ok{
			return true
		}
		tmpMap[head] = 1
		head = head.Next
	}
}


func hasCycle2(head *ListNode141) bool {
	// 快慢指针
	if head== nil || head.Next == nil {
		return false
	}
	fast,slow := head.Next,head
	for {
		if fast == slow{
			break
		}
		if fast == nil || fast.Next == nil {
			return false
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return true
}