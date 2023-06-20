// 29.顺时针打印矩阵 【通过】
/*
输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[1,2,3,6,9,8,7,4,5]
*/
package main

import "fmt"

func main() {
	arr := [][]int{
		{1},
		//{1,2,3,4,5},
		//{11,12,13,14,15},
		//{21,22,23,24,25},
	}
	res := spiralOrder(arr)
	fmt.Println(res)
}
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	res := make([]int,len(matrix) * len(matrix[0]))
	l := 0
	t := 0
	r := len(matrix[0]) - 1
	b := len(matrix) - 1
	x := 0
	for {
		// 从左往右
		for i:=l;i<=r;i++ {
			res[x] = matrix[t][i]
			x++
		}
		if t == b {
			break
		}
		t++
		// 从上往下
		for i:=t;i<=b;i++{
			res[x] = matrix[i][r]
			x++
		}
		if l==r {
			break
		}
		r--
		// 从右往左
		for i:=r;i>=l;i-- {
			res[x] = matrix[b][i]
			x++
		}
		if t==b {
			break
		}
		b--
		// 从下往上
		for i:=b;i>=t;i-- {
			res[x] = matrix[i][l]
			x++
		}
		if l==r {
			break
		}
		l++
	}
	return res
}
