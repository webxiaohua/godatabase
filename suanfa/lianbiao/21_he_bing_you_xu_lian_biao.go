// 21.合并两个有序链表
/*
输入：l1 = [1,2,4], l2 = [1,3,4]
输出：[1,1,2,3,4,4]
*/
package main

import "fmt"

func main()  {
	list1 := &ListNode21{
		Val:  1,
		Next: &ListNode21{
			Val:  4,
			Next: &ListNode21{
				Val:  7,
				Next: nil,
			},
		},
	}
	list2 := &ListNode21{
		Val:  2,
		Next: &ListNode21{
			Val:  3,
			Next: &ListNode21{
				Val:  8,
				Next: nil,
			},
		},
	}
	res := mergeTwoLists(list1,list2)
	fmt.Println(res)
}

type ListNode21 struct {
	 Val int
	 Next *ListNode21
}

func mergeTwoLists(list1 *ListNode21, list2 *ListNode21) *ListNode21 {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val <= list2.Val {
		list1.Next = mergeTwoLists(list1.Next,list2)
		return list1
	}else{
		list2.Next = mergeTwoLists(list2.Next,list1)
		return list2
	}
}

