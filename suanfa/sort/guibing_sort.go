// 归并排序思想：
// 1.将数组从中间拆分为2个数组，左边可能比右边多一个
// 2.将步骤1拆分下来的数组继续拆分，直到每个数组都只剩一个元素
// 3.从底层开始逐步合并两个排好序的数组

package main

import "fmt"

func main()  {
	source := []int{4,3,2,7,9,2,1,6,0,1}
	source = splitTwo(source)
	fmt.Println(source)
}

func mergeSort(arr []int) []int {
	length := len(arr)
	if length < 2 {
		return arr
	}
	middle := length / 2
	left := arr[0:middle]
	right := arr[middle:]
	return merge(mergeSort(left), mergeSort(right))
}

func merge(left []int, right []int) []int {
	// 两个有序数组的排序
	var result []int
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	for len(left) != 0 {
		result = append(result, left[0])
		left = left[1:]
	}
	for len(right) != 0 {
		result = append(result, right[0])
		right = right[1:]
	}
	return result
}






func splitTwo(arr []int) []int {
	// 拆分为2个数组
	length := len(arr)
	if length < 2 {
		return arr
	}
	middle := length / 2
	left := arr[0:middle]
	right := arr[middle:]
	// 递归拆分并归并排序
	return mergeTwo(splitTwo(left),splitTwo(right))
}
// 合并两个数组并排序
func mergeTwo(left []int,right []int) (result []int) {
	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			result = append(result,left[0])
			left = left[1:]
		}else{
			result = append(result,right[0])
			right = right[1:]
		}
	}
	if len(left) > 0 {
		result = append(result,left...)
	}
	if len(right) > 0 {
		result = append(result,right...)
	}
	return
}