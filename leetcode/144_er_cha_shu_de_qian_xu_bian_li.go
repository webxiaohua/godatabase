/*
【题目144】二叉树的前序遍历
给定一个二叉树的根节点 root ，返回 它的 前序 遍历 。

示例 1：
输入：root = [1,null,2,3]
输出：[1,2,3]

示例 2：
输入：root = []
输出：[]

示例 3：
输入：root = [1]
输出：[1]

提示：
树中节点数目在范围 [0, 100] 内
-100 <= Node.val <= 100

进阶: 递归算法很简单，你可以通过迭代算法完成吗？
*/

package main

import "fmt"

func main() {
	node := &TreeNode144{Val: 1, Right: &TreeNode144{Val: 2, Left: &TreeNode144{Val: 3}}}
	res := sol144_2(node)
	fmt.Println(res)
}

type TreeNode144 struct {
	Val   int
	Left  *TreeNode144
	Right *TreeNode144
}

// 1.中序遍历：左子树 - 根节点 -右子树，时间复杂度 O(N)，空间复杂度 O(N)
func sol144_1(root *TreeNode144) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	res = append(res, root.Val)
	res = append(res, sol144_1(root.Left)...)
	res = append(res, sol144_1(root.Right)...)
	return res
}

// 2.迭代方式，使用栈，判断
func sol144_2(root *TreeNode144) []int {
	// 根-左-右
	if root == nil {
		return []int{}
	}
	res, stack := make([]int, 0), make([]*TreeNode144, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		res = append(res, curr.Val)
		stack = stack[:len(stack)-1]
		if curr.Right != nil {
			stack = append(stack, curr.Right)
		}
		if curr.Left != nil {
			stack = append(stack, curr.Left)
		}
	}
	return res
}
