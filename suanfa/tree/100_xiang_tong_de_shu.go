// 100.相同的树
/*
给你两棵二叉树的根节点 p 和 q ，编写一个函数来检验这两棵树是否相同。
输入：p = [1,2,3], q = [1,2,3]
输出：true

输入：p = [1,2], q = [1,null,2]
输出：false
*/
package main

import "fmt"

type TreeNode100 struct {
	 Val int
	 Left *TreeNode100
	 Right *TreeNode100
 }

func main(){
	p := &TreeNode100{
		Val:   1,
		Left:  &TreeNode100{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode100{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}
	q := &TreeNode100{
		Val:   1,
		Left:  &TreeNode100{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode100{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}
	res := isSameTree(p,q)
	fmt.Println(res)
}

func isSameTree(p *TreeNode100, q *TreeNode100) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left,q.Left) && isSameTree(p.Right,q.Right)
}