/*
【题目145】二叉树的后序遍历
给定一个二叉树的根节点 root ，返回 它的 后序 遍历 。

示例 1：
输入：root = [1,null,2,3]
输出：[3,2,1]

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
	node := &TreeNode145{Val: 1, Right: &TreeNode145{Val: 2, Left: &TreeNode145{Val: 3}}}
	res := sol145_1(node)
	fmt.Println(res)
}

type TreeNode145 struct {
	Val   int
	Left  *TreeNode145
	Right *TreeNode145
}

// 1.中序遍历：左子树 - 右子树 - 根节点，时间复杂度 O(N)，空间复杂度 O(N)
func sol145_1(root *TreeNode145) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	res = append(res, sol145_1(root.Left)...)
	res = append(res, sol145_1(root.Right)...)
	res = append(res, root.Val)
	return res
}

// 2.迭代方式，使用栈，判断
func sol145_2(root *TreeNode145) []int {
	if root == nil {
		return []int{}
	}
	res, stack := make([]int, 0), make([]*TreeNode145, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
	}
	return res
}
