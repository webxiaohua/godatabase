// 倒数第K个节点
package main

import "fmt"

type Node struct {
	Val int
	Next *Node
}
func main(){
	node := &Node{
		Val:  1,
		Next: &Node{
			Val:  2,
			Next: &Node{
				Val:  3,
				Next: &Node{
					Val:  4,
					Next: nil,
				},
			},
		},
	}
	//res := daoShuDiKGeJieDian_Bianli(node,1)
	res := daoShuDiKGeJieDian_ShuangZhiZhen(node,1)
	fmt.Println(res)
}
// 遍历法
func daoShuDiKGeJieDian_Bianli(head *Node,k int) (res *Node) {
	// 先取得链表总长度
	var cnt int
	for node:=head;node!=nil;node=node.Next {
		cnt++
	}
	if cnt == 0 || k > cnt{
		return
	}
	// 循环遍历直到取到目标值
	res = head
	for i:=0;i<cnt-k;i++ {
		res=res.Next
	}
	return 
}

// 双指针法
func daoShuDiKGeJieDian_ShuangZhiZhen(head *Node,k int)(res *Node){
	// 本质上就是一个方程式, 正数第K个表示法：k + ? = cnt; 所以倒数第K个的表示法即为：? = cnt - k，最终求出 ? 即可
	// 第一个指针先走k步，定位到第k+1个节点上
	// 第二个指针定位到第1个节点
	// 同步开始前进，直到第一个指针结束，此时第二个指针指向的正好是倒数第K个节点
	fast,slow := head,head
	for fast!=nil && k>0 {
		fast = fast.Next
		k--
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}


