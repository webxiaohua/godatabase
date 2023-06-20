/*
【题目88】合并两个有序数组
给你两个按 非递减顺序 排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。
请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。
注意：最终，合并后数组不应由函数返回，而是存储在数组 nums1 中。为了应对这种情况，nums1 的初始长度为 m + n，其中前 m 个元素表示应合并的元素，后 n 个元素为 0 ，应忽略。nums2 的长度为 n 。

示例 1：
输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
输出：[1,2,2,3,5,6]
解释：需要合并 [1,2,3] 和 [2,5,6] 。
合并结果是 [1,2,2,3,5,6] ，其中斜体加粗标注的为 nums1 中的元素。

示例 2：
输入：nums1 = [1], m = 1, nums2 = [], n = 0
输出：[1]
解释：需要合并 [1] 和 [] 。
合并结果是 [1] 。

示例 3：
输入：nums1 = [0], m = 0, nums2 = [1], n = 1
输出：[1]
解释：需要合并的数组是 [] 和 [1] 。
合并结果是 [1] 。
注意，因为 m = 0 ，所以 nums1 中没有元素。nums1 中仅存的 0 仅仅是为了确保合并结果可以顺利存放到 nums1 中。

提示：
nums1.length == m + n
nums2.length == n
0 <= m, n <= 200
1 <= m + n <= 200
-109 <= nums1[i], nums2[j] <= 109

进阶：你可以设计实现一个时间复杂度为 O(m + n) 的算法解决此问题吗？

*/
package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3, 0, 0, 0}
	arr2 := []int{2, 5, 6}
	//res := sol88_1(arr1, 3, arr2, 3)
	//res := sol88_2(arr1, 3, arr2, 3)
	res := sol88_3(arr1, 3, arr2, 3)
	fmt.Println(res)

	/*nums := []int{4, 8, 3, 7, 2, 6, 1}
	nums = quick_sort2(nums)
	fmt.Println(nums)*/
}

// 合并数组 + 快速排序法
func sol88_1(nums1 []int, m int, nums2 []int, n int) []int {
	arr := make([]int, 0)
	arr = append(arr, nums1[0:m]...)
	arr = append(arr, nums2[0:n]...)
	return sol88_1_quick_sort(arr)
}

// 快速排序
func sol88_1_quick_sort(arr []int) []int {
	// 退出条件
	if len(arr) <= 1 {
		return arr
	}
	// 选择pivot中心轴
	pivot := arr[0]
	// 双指针
	l, r := 0, len(arr)-1
	// 是否选用左指针匹配
	isL := false
	for {
		if isL {
			// 左指针匹配
			if arr[l] >= pivot {
				// 大于基准，则将此值换到右指针指向的位置，右指针向左移动一步，控制权交给右侧指针
				arr[r] = arr[l]
				r--
				isL = false
			} else {
				// 小于等于基准，则无需交换元素，左指针向右移动一步，继续比对
				l++
			}
		} else {
			if arr[r] < pivot {
				// 将小于pivot的数字放在pivot的左侧 左侧指针加1
				arr[l] = arr[r]
				l++
				isL = true
			} else {
				// 比基准值大，右侧指针减1
				r--
			}
		}
		if l == r {
			// 所有数据处理完毕，插入基准值
			arr[l] = pivot
			break
		}
	}

	// 分别对左右子序列重复前三步操作
	sol88_1_quick_sort(arr[0:l])
	sol88_1_quick_sort(arr[l+1:])
	return arr
}

// 双指针法，额外O(N)空间
func sol88_2(nums1 []int, m int, nums2 []int, n int) []int {
	if m == 0 {
		if n == 0 {
			return []int{}
		} else {
			return nums2
		}
	} else {
		if n == 0 {
			return nums1
		}
	}
	l1, l2 := 0, 0
	res := make([]int, 0)
	for {
		if nums1[l1] <= nums2[l2] {
			res = append(res, nums1[l1])
			l1++
		} else {
			res = append(res, nums2[l2])
			l2++
		}
		if l1 == m-1 {
			res = append(res, nums2[l2:]...)
			break
		}
		if l2 == n-1 {
			res = append(res, nums1[l1:]...)
			break
		}
	}
	return res
}

// 双指针 & 内存最大利用
func sol88_3(nums1 []int, m int, nums2 []int, n int) []int {
	if m == 0 {
		for i := 0; i < len(nums2); i++ {
			nums1[i] = nums2[i]
		}
		return nums1
	}
	if n == 0 {
		return nums1
	}
	l1, l2, l3 := m-1, n-1, m+n-1
	for {
		// 从后往前放
		if nums1[l1] <= nums2[l2] {
			nums1[l3] = nums2[l2]
			l2--
			l3--
		} else {
			nums1[l3] = nums1[l1]
			l1--
			l3--
		}
		if l1 < 0 {
			for l2 >= 0 {
				nums1[l2] = nums2[l2]
				l2--
			}
			break
		}
		if l2 < 0 {
			for l1 >= 0 {
				nums1[l1] = nums1[l1]
				l1--
			}
			break
		}
	}
	return nums1
}
