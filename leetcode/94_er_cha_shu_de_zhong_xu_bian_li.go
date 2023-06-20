/*
【题目94】二叉树de中序遍历
给定一个二叉树的根节点 root ，返回 它的 中序 遍历 。

示例 1：
输入：root = [1,null,2,3]
输出：[1,3,2]

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
	node := &TreeNode94{Val: 1, Right: &TreeNode94{Val: 2, Left: &TreeNode94{Val: 3}}}
	res := sol94_3(node)
	fmt.Println(res)
}

type TreeNode94 struct {
	Val   int
	Left  *TreeNode94
	Right *TreeNode94
}

// 1.中序遍历：左子树 - 根节点 -右子树，时间复杂度 O(N)，空间复杂度 O(N)
func sol94_1(root *TreeNode94) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	res = append(res, sol94_1(root.Left)...)
	res = append(res, root.Val)
	res = append(res, sol94_1(root.Right)...)
	return res
}

// 2.迭代方式，使用栈，判断

type Stack94 struct {
	data []*TreeNode94
}

func NewStack94() *Stack94 {
	return &Stack94{data: make([]*TreeNode94, 0)}
}
func (s *Stack94) Push(node94 *TreeNode94) {
	s.data = append(s.data, node94)
}
func (s *Stack94) Pop() *TreeNode94 {
	if len(s.data) == 0 {
		return nil
	}
	res := s.data[len(s.data)-1]
	s.data = s.data[0 : len(s.data)-1]
	return res
}
func (s *Stack94) Empty() bool {
	return len(s.data) == 0
}
func (s *Stack94) Top() *TreeNode94 {
	if len(s.data) == 0 {
		return nil
	} else {
		return s.data[len(s.data)-1]
	}
}

func sol94_2(root *TreeNode94) []int {
	// 左-根-右
	res := make([]int, 0)
	stack := NewStack94()
	for {
		if root == nil && stack.Empty() {
			break
		}
		for {
			if root != nil {
				stack.Push(root)
				root = root.Left
			} else {
				break
			}
		}
		root = stack.Pop()
		res = append(res, root.Val)
		root = root.Right
	}
	return res
}
