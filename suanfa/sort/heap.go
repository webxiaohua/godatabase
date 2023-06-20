// 堆排序
package main

import (
	"fmt"
)

var count int //记录当前堆剩余位排序的长度
// 建立大顶堆
func buildMaxHeap(arr []int)  {
	count = len(arr)
	for i:= len(arr)/2;i>=0;i-- {
		maxHeapify(arr,i)
	}
}
// 堆调整
func maxHeapify(arr []int,i int)  {
	// i是跟节点，left是i的左子节点，right是i的右子节点
	left := 2*i+1
	right := left+1
	largest := i
	// 左子节点值和跟节点值判断
	if left < count && arr[left] > arr[largest] {
		largest = left
	}
	// 右子节点值和跟节点值判断
	if right < count && arr[right]>arr[largest] {
		largest = right
	}
	if largest != i {
		// 交换
		arr[i],arr[largest] = arr[largest],arr[i]
		// 只要发生变化，则需要递归地把整个树进行调整
		maxHeapify(arr,largest)
	}
}

func heapSortAsc(arr []int){
	// 先构建大顶堆
	buildMaxHeap(arr)
	for i:=len(arr)-1;i>0;i-- {
		// 将第一个元素和最后一个元素交换，得到最大值
		arr[i],arr[0] = arr[0],arr[i]
		// 减少数组长度
		count--
		// 重新构建大顶堆
		maxHeapify(arr,0)
	}
}

// 构建小顶堆
func buildMinHeap(arr []int){
	// arr[i] < arr[i*2+1] && arr[i] < arr[i*2+2]
	count = len(arr)
	for i:=len(arr)/2;i>=0;i--{
		minHeapify(arr,i)
	}
}

// 小顶堆调整
func minHeapify(arr []int,i int){
	smaller:=i
	left:=i*2+1
	right:=i*2+2
	if left < count && arr[left] < arr[smaller] {
		smaller = left
	}
	if right < count && arr[right] < arr[smaller] {
		smaller = right
	}
	if smaller != i {
		arr[smaller],arr[i] = arr[i],arr[smaller]
		minHeapify(arr,smaller)
	}
}

// 堆排序倒序
func heapSortDesc(arr []int){
	// 构建小顶堆
	// 循环，每次把堆顶元素和末位元素交换，数组长度计数器减1，重新调整堆，直到所有元素完成
	buildMinHeap(arr)
	for count > 1 {
		arr[0],arr[count-1] = arr[count-1],arr[0]
		count--
		minHeapify(arr,0)
	}
}

func main()  {
	arr := []int{3,5,2,7,4,9,0,5,6}
	//buildMaxHeap2(arr)
	//fmt.Println(arr)
	//heapSortDesc(arr)
	fmt.Println(arr)
}
