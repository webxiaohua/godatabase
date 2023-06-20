package main

import "fmt"

func main(){
	source := []int{7,3,4,2,8,1,9,0,2,5,8,1}
	res := split(source)
	fmt.Println(res)
}

// 不断分解,直到只有一个元素
func split(arr []int) []int{
	length := len(arr)
	if length < 2 {
		return arr
	}
	middle := length / 2
	left := arr[0:middle]
	right := arr[middle:]
	//
	return sort(split(left),split(right))
}

// left right 各自有序，进行合并
func sort(left,right []int)(result []int) {
	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			result = append(result,left[0])
			left = left[1:]
			continue
		}else {
			result = append(result,right[0])
			right = right[1:]
			continue
		}
	}
	for len(left) > 0 {
		result = append(result,left[0])
		left = left[1:]
	}
	for len(right) > 0 {
		result = append(result,right[0])
		right = right[1:]
	}
	return
}
