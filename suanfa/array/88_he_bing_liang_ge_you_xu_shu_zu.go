/*
合并两个有序数组
【示例1】
输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
输出：[1,2,2,3,5,6]
解释：需要合并 [1,2,3] 和 [2,5,6] 。
【示例2】
输入：nums1 = [0], m = 0, nums2 = [1], n = 1
输出：[1]
解释：需要合并的数组是 [] 和 [1] 。
合并结果是 [1] 。
*/
package main

import "fmt"

func main() {
	// 不使用额外空间
	nums1 := []int{1,2,4,5,6,0}
	nums2 := []int{3}
	merge2(nums1,5,nums2,1)
	fmt.Println(nums1)
}
// 不使用额外存储空间
func merge1(nums1 []int, m int, nums2 []int, n int)  {
	if m == 0 {
		for i:=0;i<n;i++ {
			nums1[i]=nums2[i]
		}
		return
	}
	if n == 0 {
		return
	}
	front := m - 1
	back := len(nums1) - 1
	for i:=0;i<m;i++ {
		nums1[back] = nums1[front]
		front--
		back--
	}
	a,b,i:=n,0,0
	for a < len(nums1) && b<n {
		if nums1[a] > nums2[b]{
			nums1[i] = nums2[b]
			b++
		}else{
			nums1[i] = nums1[a]
			a++
		}
		i++
	}
	for b < n {
		nums1[i] = nums2[b]
		i++
		b++
	}
}

// 使用额外存储空间
func merge2(nums1 []int, m int, nums2 []int, n int)  {
	tmp := make([]int,m+n)
	  i,a,b := 0,0,0
	  for a<m && b<n {
	      if nums1[a] > nums2[b]{
	          tmp[i]=nums2[b]
	          b++
	      }else{
	          tmp[i]=nums1[a]
	          a++
	      }
	      i++
	  }
	  for a<m{
	      tmp[i] = nums1[a]
	      a++
	      i++
	  }
	  for b<n{
	      tmp[i] = nums2[b]
	      b++
	      i++
	  }
	  for j:=0;j<m+n;j++ {
	      nums1[j] = tmp[j]
	  }
}