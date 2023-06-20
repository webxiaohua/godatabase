package main

import "fmt"

type TreeNode struct {
	Val int `json:"val"`
	Left *TreeNode `json:"left"`
	Right *TreeNode `json:"right"`
}

// 构建满二叉树
func buildTree(deep int,cur *TreeNode) {
	if deep == 0 {
		return
	}else{
		cur.Left = &TreeNode{
			Val:   deep+1,
			Left:  nil,
			Right: nil,
		}
		cur.Right = &TreeNode{
			Val:   deep+2,
			Left:  nil,
			Right: nil,
		}
		deep--
		buildTree(deep,cur.Left)
		buildTree(deep,cur.Right)
	}
}

// 前序遍历
func frontEach(root *TreeNode) {
	if root.Left == nil {
		return
	}

}

func main()  {
	root := &TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	}
	buildTree(3,root)
	fmt.Println("前序遍历：")

}