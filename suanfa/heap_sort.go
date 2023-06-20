package main

func main()  {
	
}

func heapSort(arr []int) []int {
	arrLen := len(arr)
	// 构建大顶堆
	buildMaxHeap(arr, arrLen)
	// 遍历堆
	for i := arrLen - 1; i >= 0; i-- {
		swap(arr, 0, i)
		arrLen -= 1
		heapify(arr, 0, arrLen)
	}
	return arr
}

// 构建大顶堆
func buildMaxHeap(arr []int,arrLen int){
	for i := arrLen / 2; i >= 0; i-- {
		heapify(arr, i, arrLen)
	}
}

// 数组堆化
func heapify(arr []int, i, arrLen int) {
	// 左右节点
	left := 2*i + 1
	right := 2*i + 2
	// 最大节点下标
	largest := i
	// 左节点比较
	if left < arrLen && arr[left] > arr[largest] {
		largest = left
	}
	// 右节点比较
	if right < arrLen && arr[right] > arr[largest] {
		largest = right
	}
	if largest != i {
		swap(arr, i, largest)
		heapify(arr, largest, arrLen)
	}
}

// 数组内元素交换位置
func swap(arr []int,i,j int)  {
	arr[i], arr[j] = arr[j], arr[i]
}