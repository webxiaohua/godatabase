package main

import "fmt"

func main() {
	nums := []int{1, 9, 2, 8, 3, 7, 4, 6, 5}
	nums = quickSort(nums)
	fmt.Println(nums)
}

// 快速排序
func quickSort(nums []int) []int {
	// 取基准，然后左右指针分别进行判定，直到相遇，再递归子过程
	if len(nums) <= 1 {
		return nums
	}
	base := nums[0]
	l, r := 0, len(nums)-1
	isLeft := false
	for {
		if isLeft {
			if nums[l] > base {
				nums[r] = nums[l]
				r--
				isLeft = false
			} else {
				l++
			}
		} else {
			if nums[r] < base {
				nums[l] = nums[r]
				l++
				isLeft = true
			} else {
				r--
			}
		}
		if l == r {
			break
		}
	}
	nums[l] = base
	quickSort(nums[0:l])
	quickSort(nums[l+1:])
	return nums
}

// 归并排序

// 堆排序 [9,8,7,6,5,4]
func heapSort(nums []int) []int {
	// 从末尾开始进行比较
	// 父节点计算公式: （i-1)/2
	// 左子节点公式： 2*i+1
	// 右子节点公式： 2*i+2

}
