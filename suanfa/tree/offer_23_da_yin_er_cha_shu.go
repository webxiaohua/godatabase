// offer 23.从上往下打印二叉树
package main

import "fmt"

func main(){
	tree := &TreeNode23{
		Val:   1,
		Left:  &TreeNode23{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode23{
			Val:   3,
			Left:  &TreeNode23{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}
	res := levelOrder(tree)
	fmt.Println(res)
}

type TreeNode23 struct {
	 Val int
	 Left *TreeNode23
	 Right *TreeNode23
}

func levelOrder(root *TreeNode23) []int {
	// 从跟节点开始打印，如果存在子节点，插入队列，后续处理队列，直到队列为空
	res := make([]int,0)
	queue := make([]*TreeNode23,0)
	if root == nil {
		return res
	}
	queue = append(queue,root)
	for {
		if len(queue) == 0 {
			break
		}
		item := queue[0]
		queue = queue[1:]
		res = append(res,item.Val)
		if item.Left != nil {
			queue = append(queue,item.Left)
		}
		if item.Right != nil {
			queue = append(queue,item.Right)
		}
	}
	return res
}