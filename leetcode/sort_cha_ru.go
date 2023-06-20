/**
 * @author:伯约
 * @date:2023/6/6
 * @note: 插入排序 时间复杂度 O(N^2)，空间复杂度 O(1)，稳定排序
**/

package main

import "fmt"

func sort_charu(arr []int64) {
	if len(arr) <= 1 {
		return
	}
	// 从第1号元素开始，往前对比
	for i := 1; i < len(arr); i++ {
		// 当前值
		cur := arr[i]
		// 对比尾值
		j := i - 1
		for {
			// 依次对比，直到碰到比当前值小的，退出对比； 否则依次交换
			if j >= 0 && arr[j] > cur {
				arr[j+1] = arr[j]
				j--
			} else {
				break
			}
		}
		// 确定j的位置，将本轮元素放置进来
		arr[j+1] = cur
		fmt.Println(arr)
	}
}

func main() {
	arr := []int64{1, 3, 2, 5, 4}
	sort_charu(arr)
	fmt.Println("result: ", arr)
}
