// 4.获取两个有序数组的中位数
package main

import (
	"fmt"
	"sort"
)

func main(){
	arr1 := []int{1,2,3,4,5}
	arr2 := []int{6,7,8,9,10}
	//fmt.Println(MergeSort_4(arr1,arr2))
	fmt.Println(GuiBingSort_4(arr1,arr2))

}

// 合并排序解法
func MergeSort_4(arr1,arr2 []int) float64 {
	arr1 = append(arr1,arr2...)
	sort.Ints(arr1)
	if len(arr1) == 0 {
		return 0
	}
	if len(arr1) == 1 {
		return float64(arr1[0])
	}
	half := len(arr1) / 2
	if len(arr1) % 2 == 0 {
		return float64(arr1[half-1] + arr1[half]) / 2
	}else{
		return float64(arr1[half])
	}
}

// 归并排序解法 O(nlogn)
func GuiBingSort_4(arr1,arr2 []int)float64 {
	arr := make([]int,0)
	for len(arr1) > 0 && len(arr2) > 0 {
		if arr1[0] <= arr2[0] {
			arr = append(arr,arr1[0])
			arr1 = arr1[1:]
		}else{
			arr = append(arr,arr2[0])
			arr2 = arr2[1:]
		}
	}
	if len(arr1) != 0 {
		arr = append(arr,arr1...)
	}else if len(arr2) != 0 {
		arr = append(arr,arr2...)
	}
	// 取中位数
	if len(arr) == 0 {
		return 0
	}
	if len(arr) == 1 {
		return float64(arr[0])
	}
	half := len(arr) / 2
	if len(arr) % 2 == 0 {
		return float64(arr[half-1] + arr[half]) / 2
	}else{
		return float64(arr[half])
	}
}
