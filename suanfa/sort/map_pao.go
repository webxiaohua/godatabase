// 冒泡排序 时间复杂度 O(n2)，最坏逆序，最好顺序； 空间复杂度 O(1)； 稳定排序
package main

import "fmt"

func main()  {
	arr := []int{3,2,7,1,9,4,5}
	maopao(arr)
	fmt.Println(arr)
}

func maopao(arr []int) {
	// 基本思想：外层循环控制轮次，每一次会找到一个最值；内层循环用于比较替换，一旦发现比外层值大或者小的值会进行替换
	for i:=0;i<len(arr);i++ {
		for j:=i+1;j<len(arr);j++ {
			if arr[i] > arr[j] {
				arr[i],arr[j] = arr[j],arr[i]
			}
		}
	}
}
