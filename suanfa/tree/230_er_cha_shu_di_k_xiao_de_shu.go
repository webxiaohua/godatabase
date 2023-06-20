/*
230. 二叉搜索树中第K小的元素
给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 个最小元素（从 1 开始计数）。
示例：

输入：root = [3,1,4,null,2], k = 1
输出：1
图例：
     3
    / \
   1   4
    \
     2

输入：root = [5,3,6,2,4,null,null,1], k = 3
输出：3
图例：
      5
     / \
    3   6
   / \
  2   4
 /
1
*/
package main

import "fmt"

type TreeNode230 struct {
	Val   int
	Left  *TreeNode230
	Right *TreeNode230
}

func kthSmallest(root *TreeNode230, k int) int {
	// 采用深度优先遍历方式，从最小的元素出发，每遍历一次k减去1，直到k为0时取出当前元素
	res := -1
	var dfs func(*TreeNode230)
	dfs = func(node *TreeNode230) {
		if node == nil {
			return
		}
		dfs(node.Left)
		k--
		if k == 0 {
			res = node.Val
			return
		}
		dfs(node.Right)
	}
	dfs(root)
	return res
}

func main() {
	data := &TreeNode230{
		Val: 3,
		Left: &TreeNode230{
			Val:  1,
			Left: nil,
			Right: &TreeNode230{
				Val:   20,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode230{
			Val:   4,
			Left:  nil,
			Right: nil,
		},
	}
	res := kthSmallest(data, 1)
	fmt.Println(res)
}
