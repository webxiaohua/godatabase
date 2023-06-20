// 19.删除链表倒数第N个节点
/*
输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]
*/
package main

import "fmt"

func main(){
	node := &ListNode19{
		Val:  1,
		Next: &ListNode19{
			Val:  2,
			/*Next: &ListNode19{
				Val:  3,
				Next: &ListNode19{
					Val:  4,
					Next: &ListNode19{
						Val:  5,
						Next: nil,
					},
				},
			},*/
		},
	}
	res := removeNthFromEnd(node,1)
	fmt.Println(res)
}

type ListNode19 struct {
	 Val int
	 Next *ListNode19
 }
func removeNthFromEnd(head *ListNode19, n int) *ListNode19 {
	// 双指针解法
	fast,slow := head,head
	for i:=0;i<n;i++ {
		if fast.Next != nil {
			fast = fast.Next
		}else{
			return head.Next
		}
	}
	for {
		if fast.Next != nil {
			fast = fast.Next
			slow = slow.Next
		}else{
			break
		}
	}
	slow.Next = slow.Next.Next
	return head
}