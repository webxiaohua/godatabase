// 合并链表排序
package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func main(){
	arr := []*ListNode{
		&ListNode{Val:1,Next:&ListNode{Val:4,Next:&ListNode{Val:7}}},
		&ListNode{Val:2,Next:&ListNode{Val:5,Next:&ListNode{Val:8}}},
		&ListNode{Val:3,Next:&ListNode{Val:6,Next:&ListNode{Val:9}}},
	}
	res := mergeKLists(arr)
	fmt.Println(res)
}

func mergeKLists(lists []*ListNode) *ListNode {
	return mergeN(lists,0,len(lists)-1)
}

// 分治法
func mergeN(lists []*ListNode,l,r int) *ListNode {
	if l == r {
		return lists[l]
	}else if l>r {
		return nil
	}
	mid := (l+r)>>1
	return mergeTwo(mergeN(lists,l,mid),mergeN(lists,mid+1,r))
}

// 假设只有两个
func mergeTwo(a *ListNode,b *ListNode) *ListNode{
	if a == nil {
		return b
	}else if b==nil {
		return a
	}
	// 先找到头节点
	var head *ListNode
	if a.Val < b.Val {
		head = a
		head.Next = mergeTwo(a.Next,b)
	}else{
		head = b
		head.Next = mergeTwo(b.Next,a)
	}
	return head
}