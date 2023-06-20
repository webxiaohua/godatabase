// 46.全排列
/*
输入：nums = [1,2,3] 不重复，不为空
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
*/
package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}
	r := permute(arr)
	fmt.Println(r)
}

var res [][]int

func permute(nums []int) [][]int {
	res = make([][]int, 0)
	path := make([]int, 0)
	backtrack(nums, path)
	return res
}

// nums:选择列表 path:路径
func backtrack(nums []int, path []int) {
	// 退出条件
	if len(path) == len(nums) {
		// 注意这里不能直接引用path
		newPath := make([]int, 0)
		newPath = append(newPath, path...)
		res = append(res, newPath)
		return
	}
	// 处理决策
	for _, num := range nums {
		// 去除不合法的选择后继续进行决策
		if isDiff(num, path) {
			continue
		}
		// 做选择
		path = append(path, num)
		// 进入下一层决策树
		backtrack(nums, path)
		// 撤销选择
		path = path[0 : len(path)-1]
	}
}

// 判断是否重复
func isDiff(target int, nums []int) bool {
	for _, num := range nums {
		if target == num {
			return true
		}
	}
	return false
}
