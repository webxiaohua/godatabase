/*
【题目142】环形链表2
给定一个链表的头节点  head ，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。
不允许修改 链表。


示例 1:
3 -> 2 -> 0 -> -4
     |__________|
输入：head = [3,2,0,-4], pos = 1
输出：返回索引为 1 的链表节点
解释：链表中有一个环，其尾部连接到第二个节点。


提示：
链表中节点的数目范围在范围 [0, 104] 内
-105 <= Node.val <= 105
pos 的值为 -1 或者链表中的一个有效索引

*/
package main

import "fmt"

type ListNode142 struct {
	Val  int
	Next *ListNode142
}

func main() {
	node1 := &ListNode142{Val: 1}
	node2 := &ListNode142{Val: 2}
	node3 := &ListNode142{Val: 3}
	node4 := &ListNode142{Val: 4}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node2
	res := sol142_2(node1)
	if res == nil {
		fmt.Println("not exist.")
	} else {
		fmt.Println("exist:", res.Val)
	}
}

/*
1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> (3)
2 3
3 5
4 7
5 4
6 6
采用快慢指针可以判断出来是否存在环;
如果存在环，则整个环形链表可以分为3块： [1_2_3  4_交汇点  交汇点_3
这三块分别用 a,b,c 表示，代数公示：
2(a + b) = a + n*(b+c) + b
消除后得到 a+b = n*(b+c) 即： a = (n-1) * b + n*c
如果：
n=1 则 a=c
n=2 则 a=b+c+c 1圈+c
n=3 则 a=b+c+b+c+c 2圈+c
实际上随便多少圈，最终的a一定等于c
因此我们只要将slow指针移到头节点，fast保持交汇位置，然后步长都设置为1开始移动，找到交汇点即表示环入口

*/

// 快慢指针 先判断出是否有环，然后将慢指针移动回原点，快慢指针步长都改为1，等到下次相交处便是入口
func sol142_2(head *ListNode142) *ListNode142 {
	if head == nil || head.Next == nil {
		return nil
	}
	existCycle := false
	i, j := head, head
	for {
		if i == nil || j == nil || j.Next == nil {
			break
		}
		i = i.Next
		j = j.Next.Next
		if i == j {
			existCycle = true
			break
		}
	}
	if !existCycle {
		return nil
	}
	i = head
	for {
		if i == j {
			return i
		}
		i = i.Next
		j = j.Next
	}
}
