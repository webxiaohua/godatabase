// offer 33.二叉树的后序遍历序列
/*
输入: [1,3,2,6,5]
输出: true
*/
package main

func verifyPostorder(postorder []int) bool {
	// 左右中 -> 中右左


	// 二叉树中左子树的节点小于根节点

	// 二叉树中右子树的节点大于根节点

	// 判断左子树是不是二叉搜索树

	// 判断右子树是不是二叉搜索树

}

func verifyPostorderDigui(postorder []int, left,right int) bool {
	// 当前区域不合法的时候直接返回true就好
	if left >= right {
		return true
	}
	// 找到根节点
	root := postorder[right]
	// 从当前区间找到第一个大于跟节点的，说明后续元素都在右子树
	k:=left
	for {
		if k < right && postorder[k] < root {
			continue
		}else{
			k++
		}
	}
	// 判断后续区域是否全部大于跟节点，如果出现小于的情况则返回false

	// 当前树没问题就检查左右子树

	// 最终都没问题就返回true
	return true
}