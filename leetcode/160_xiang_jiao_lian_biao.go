/*
【题目160】相交链表
给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表不存在相交节点，返回 null 。
图示两个链表在节点 c1 开始相交：

a1,a2,c1,c2,c3
b1,b2,b3,c1,c2,c3

题目数据 保证 整个链式结构中不存在环。

注意，函数返回结果后，链表必须 保持其原始结构 。

自定义评测：

评测系统 的输入如下（你设计的程序 不适用 此输入）：

intersectVal - 相交的起始节点的值。如果不存在相交节点，这一值为 0
listA - 第一个链表
listB - 第二个链表
skipA - 在 listA 中（从头节点开始）跳到交叉节点的节点数
skipB - 在 listB 中（从头节点开始）跳到交叉节点的节点数
评测系统将根据这些输入创建链式数据结构，并将两个头节点 headA 和 headB 传递给你的程序。如果程序能够正确返回相交节点，那么你的解决方案将被 视作正确答案 。


示例 1：
输入：intersectVal = 8, listA = [4,1,8,4,5], listB = [5,6,1,8,4,5], skipA = 2, skipB = 3
输出：Intersected at '8'
解释：相交节点的值为 8 （注意，如果两个链表相交则不能为 0）。
从各自的表头开始算起，链表 A 为 [4,1,8,4,5]，链表 B 为 [5,6,1,8,4,5]。
在 A 中，相交节点前有 2 个节点；在 B 中，相交节点前有 3 个节点。
— 请注意相交节点的值不为 1，因为在链表 A 和链表 B 之中值为 1 的节点 (A 中第二个节点和 B 中第三个节点) 是不同的节点。换句话说，它们在内存中指向两个不同的位置，而链表 A 和链表 B 中值为 8 的节点 (A 中第三个节点，B 中第四个节点) 在内存中指向相同的位置。

提示：
listA 中节点数目为 m
listB 中节点数目为 n
1 <= m, n <= 3 * 104
1 <= Node.val <= 105
0 <= skipA <= m
0 <= skipB <= n
如果 listA 和 listB 没有交点，intersectVal 为 0
如果 listA 和 listB 有交点，intersectVal == listA[skipA] == listB[skipB]


进阶：你能否设计一个时间复杂度 O(m + n) 、仅用 O(1) 内存的解决方案？


*/

package main

import "fmt"

type ListNode160 struct {
	Val  int
	Next *ListNode160
}

func main() {

	headA_5 := &ListNode160{Val: 5}
	headA_4 := &ListNode160{Val: 4, Next: headA_5}
	headA_3 := &ListNode160{Val: 3, Next: headA_4}
	headA_2 := &ListNode160{Val: 2, Next: headA_3}
	headA_1 := &ListNode160{Val: 1, Next: headA_2}

	headB_3 := &ListNode160{Val: 13, Next: headA_3}
	headB_2 := &ListNode160{Val: 12, Next: headB_3}
	headB_1 := &ListNode160{Val: 11, Next: headB_2}

	res := sol160_4(headA_1, headB_1)
	fmt.Println(res)
}

// 方案1:map法，先遍历完一条链表，计入map，再遍历另外一条链表，判断指针是否存在，时间复杂度 O(m+n)，需要额外 O(m) 空间复杂度，不满足要求
func sol160_1(headA, headB *ListNode160) *ListNode160 {
	tmpMap := make(map[*ListNode160]int)
	tmpA := headA
	for {
		if tmpA == nil {
			break
		}
		tmpMap[tmpA] = 1
		tmpA = tmpA.Next
	}
	for {
		if headB == nil {
			break
		}
		if _, ok := tmpMap[headB]; ok {
			return headB
		}
		headB = headB.Next
	}
	return nil
}

// 2.暴力循环法 时间复杂度 O(M*N)
func sol160_2(headA, headB *ListNode160) *ListNode160 {
	for {
		if headA == nil {
			break
		}
		tmpB := headB
		for {
			if tmpB == nil {
				break
			}
			if headA == tmpB {
				return headA
			}
			tmpB = tmpB.Next
		}
		headA = headA.Next
	}
	return nil
}

// 3.双指针交叉遍历 考虑到 len(A) + len(B) = len(B) + len(A) ，所以启动A、B两个指针，分别从AB两个链表的头节点开始遍历，当一个指针遍历到当前链表结尾的时候转头去遍历另一个链表的头节点，两个指针走到相同节点即可得出解
func sol160_3(headA, headB *ListNode160) *ListNode160 {
	if headA == nil || headB == nil {
		return nil
	}
	a, b := headA, headB
	for {
		if a == b {
			// 如果走到最后，ab都为nil，表示无相交
			return a
		}
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}

		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}
	return nil
}

// 4.求长度，然后长度大的减去长度小的，两个指针位于同一起跑线向后遍历
func sol160_4(headA, headB *ListNode160) *ListNode160 {
	lenA, lenB := 0, 0
	a, b := headA, headB
	for {
		if a == nil {
			break
		}
		lenA++
		a = a.Next
	}
	for {
		if b == nil {
			break
		}
		lenB++
		b = b.Next
	}
	a, b = headA, headB
	if lenA > lenB {
		_t := lenA - lenB
		for {
			if _t == 0 {
				break
			}
			_t--
			a = a.Next
		}
	} else if lenB > lenA {
		_t := lenB - lenA
		for {
			if _t == 0 {
				break
			}
			_t--
			b = b.Next
		}
	}
	for {
		if a == b {
			return a
		}
		a = a.Next
		b = b.Next
	}
	return nil
}
