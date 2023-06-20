// 选择排序 时间复杂度 O(n2)； 空间复杂度 O(1)； 稳定排序
package main

import "fmt"

func main()  {
	arr := []int{3,2,7,1,9,4,5}
	xuanze(arr)
	fmt.Println(arr)
}

func xuanze(arr []int) {
	// 基本思想：借助变量下标标记索引，每一轮循环过后会找出本轮应该交换交换的下标
	for i:= 0;i<len(arr);i++ {
		minIndex:=0
		for j:= i+1;i<len(arr);j++ {
			if arr[minIndex] < arr[j]{
				minIndex = j
			}
		}
		arr[i],arr[minIndex] = arr[minIndex],arr[i]
	}
}