// 27.二叉树的镜像
/*
输入：root = [4,2,7,1,3,6,9]
输出：[4,7,2,9,6,3,1]
*/
package main

import "fmt"

type TreeNode27 struct {
	Val int
	Left *TreeNode27
	Right *TreeNode27
}

func main (){
	tree := &TreeNode27{
		Val:   1,
		Left:  &TreeNode27{
			Val:   2,
			Left:  &TreeNode27{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode27{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode27{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}
	//tree = mirrorTree(tree)
	tree = mirrorTreeWithLoop(tree)
	fmt.Println(tree)
}
// 递归法
func mirrorTree(root *TreeNode27) *TreeNode27 {
	// 当前不具备递归条件，退出
	if root != nil {
		// 交换左右子节点
		root.Left,root.Right = root.Right,root.Left
		// 继续处理左子树
		mirrorTree(root.Left)
		// 继续处理右子树
		mirrorTree(root.Right)
	}
	return root
}

// 循环法 栈
//
// 只要left和right有一个不为空，加入数组
// 直到数组为空退出循环
func mirrorTreeWithLoop(root *TreeNode27) *TreeNode27 {
	stack := &Stack{
		Count: 0,
		Data:  make([]interface{},0),
	}
	stack.Push(root)
	for !stack.Empty() {
		node := stack.Pop()
		if node == nil {
			continue
		}
		tmp := node.(*TreeNode27)
		tmp.Left,tmp.Right = tmp.Right,tmp.Left
		if tmp.Right!=nil {
			stack.Push(tmp.Right)
		}
		if tmp.Left!=nil {
			stack.Push(tmp.Left)
		}
	}
	return root
}

type Stack struct {
	Count int
	Data []interface{}
}
func (s *Stack) Empty() bool{
	return s.Count == 0
}
func (s *Stack) Pop() interface{}{
	if s.Count == 0 {
		return nil
	}
	s.Count--
	res := s.Data[s.Count]
	s.Data = s.Data[0:s.Count]
	return res
}
func (s *Stack) Push(p interface{}){
	s.Count++
	s.Data = append(s.Data,p)
}