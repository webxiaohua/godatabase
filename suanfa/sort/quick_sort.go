// 快速排序
package main

import "fmt"

func main(){
	arr := []int64{9,4,2,7,5}
	res := quickSortInt64(arr)
	fmt.Println(res)
}


func quickSortInt64(arr []int64) []int64 {
	return _quickSortInt64(arr, 0, len(arr)-1)
}
func _quickSortInt64(arr []int64, left, right int) []int64 {
	if left < right {
		partitionIndex := partition(arr, left, right)
		_quickSortInt64(arr, left, partitionIndex-1)
		_quickSortInt64(arr, partitionIndex+1, right)
	}
	return arr
}
func partition(arr []int64, left, right int) int {
	pivot := left
	index := pivot + 1
	for i := index; i <= right; i++ {
		if arr[i] < arr[pivot] {
			arr[i],arr[index] = arr[index],arr[i]
			index += 1
		}
	}
	arr[pivot],arr[index-1] = arr[index-1],arr[pivot]
	return index - 1
}