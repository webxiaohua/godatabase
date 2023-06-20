// 26.树的子结构
/*
输入两棵二叉树A和B，判断B是不是A的子结构。(约定空树不是任意一个树的子结构)
输入：A = [3,4,5,1,2], B = [4,1]
输出：true
*/
package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main(){
	A := &TreeNode{
		Val:   1,
		Left:  &TreeNode{
			Val:   2,
			Left:  &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}
	B := &TreeNode{
		Val:   2,
		Left:  &TreeNode{
			Val:   4,
			Left:  nil,
			Right: nil,
		},
		Right: nil,
	}
	result := haveSubTree(A,B)
	fmt.Println(result)
}

func haveSubTree(A *TreeNode,B *TreeNode) bool {
	result := false
	if A == nil || B == nil {
		return result
	}
	if A.Val == B.Val {
		return isAHaveB(A,B)
	}
	result = haveSubTree(A.Left,B)
	if !result {
		result = haveSubTree(A.Right,B)
	}
	return result
}

func isAHaveB(A,B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil {
		return false
	}
	if A.Val != B.Val {
		return false
	}
	return isAHaveB(A.Left,B.Left) && isAHaveB(A.Right,B.Right)
}