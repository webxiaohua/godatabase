// 94.二叉树的中序遍历
/*
输入：root = [1,null,2,3]
输出：[1,3,2]
*/
package main

import "fmt"

type TreeNode94 struct {
	 Val int
	 Left *TreeNode94
	 Right *TreeNode94
 }

 func main(){
 	tree := &TreeNode94{
		Val:   1,
		Left:  nil,
		Right: &TreeNode94{
			Val:   2,
			Left:  &TreeNode94{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}
	res := inorderTraversal(tree)
	fmt.Println(res)
 }
 // 递归
func inorderTraversal(root *TreeNode94) []int {
	res := make([]int,0)
	if root == nil {
		return res
	}
	res = append(res,inorderTraversal(root.Left)...)
	res = append(res,root.Val)
	res = append(res,inorderTraversal(root.Right)...)
	return res
}
