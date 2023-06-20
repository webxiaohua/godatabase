// 插入排序
package main

import "fmt"

func main(){
	arr := []int{3,2,7,1,9,4,5}
	charu(arr)
	fmt.Println(arr)
}

func charu(arr []int)  {
	// 基本思想： 取出任意一个元素，往前对比，找到第一个小于该元素的位置插入，再将之后的元素后移
	// 第一轮 3,2,7,1,9,4,5
	// 第二轮 2,3,7,1,9,4,5
	// 第三轮 2,3,7,1,9,4,5
	// 第四轮 1,2,3,7,9,4,5
	// 第五轮 1,2,3,7,9,4,5
	// 第六轮 1,2,3,4,7,9,5
	// 第七轮 1,2,3,4,5,7,9
	if len(arr) == 1 {
		return
	}
	for i:=1;i<len(arr);i++ {
		// 取出当前值
		current := arr[i]
		// 可移动指针
		j := i-1
		for j>=0 && current < arr[j]{
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = current
	}
}