/**
 * @author:伯约
 * @date:2023/6/6
 * @note:
**/

package main

import (
	"fmt"
)

type Node struct {
	val  int64
	next *Node
}

func isHw(head *Node) bool {
	if head == nil || head.next == nil {
		return false
	}
	node1 := head
	node2 := head
	for node1 != nil && node2 != nil {
		node1 = node1.next
		if node2.next == nil {
			return false
		}
		node2 = node2.next.next
		if node1 == node2 {
			return true
		}

	}
	return false
}

func main() {
	node4 := &Node{val: 1}
	node3 := &Node{val: 2}
	node2 := &Node{val: 2}
	node1 := &Node{val: 1}

	node1.next = node2
	node2.next = node3
	node3.next = node4

	fmt.Println(isHw(node1))
}
