// 146. LRU缓存机制
package main

import "fmt"

func main(){
	lru := Constructor(3)
	lru.Put(1,1)
	lru.Put(2,2)
	lru.Put(3,3) // 3->2->1
	lru.Get(2)	// 2->3->1
	lru.Put(4,4) // 4->2->3
	lru.Get(0) // 4->2->3
	lru.Put(3,31) // 3->4->2
	fmt.Println("head:",lru.Head.Next.Val)
	fmt.Println("tail:",lru.Tail.Prev.Val)
}

type LRUCache struct {
	Capacity int
	Data map[int]*Node
	Head *Node
	Tail *Node
}

type Node struct{
	Val int
	Key int
	Prev *Node
	Next *Node
}

func Constructor(capacity int) LRUCache {
	ret := LRUCache{
		Capacity:capacity,
	}
	ret.Data = make(map[int]*Node)
	ret.Head = &Node{}
	ret.Tail = &Node{}
	ret.Head.Next = ret.Tail
	ret.Tail.Prev = ret.Head
	return ret
}


func (this *LRUCache) Get(key int) int {
	// 是否存在值
	if _,ok := this.Data[key];!ok {
		return -1
	}
	// 调整顺序
	cur := this.Data[key]
	// 后面指针的值绑定到前置节点 head->1->2->3->tail
	// 先断开
	cur.Next.Prev = cur.Prev
	cur.Prev.Next = cur.Next
	// 挂到head节点(将head的后置节点挂到当前节点下，再将当前节点挂到head节点后)
	cur.Next = this.Head.Next
	this.Head.Next.Prev = cur
	this.Head.Next = cur
	return this.Data[key].Val
}


func (this *LRUCache) Put(key int, value int)  {
	if _,ok := this.Data[key];ok {
		// 调整顺序 & 更新值
		cur := this.Data[key]
		// 后面指针的值绑定到前置节点 head->1->2->3->tail
		// 先断开
		cur.Next.Prev = cur.Prev
		cur.Prev.Next = cur.Next
		// 挂到head节点(将head的后置节点挂到当前节点下，再将当前节点挂到head节点后)
		cur.Next = this.Head.Next
		this.Head.Next.Prev = cur
		this.Head.Next = cur
		cur.Val = value
		return
	}
	if this.Capacity == 0 {
		return
	}
	if len(this.Data) > 0 && len(this.Data) == this.Capacity {
		// 先移除队尾元素
		last := this.Tail.Prev
		this.Tail.Prev = last.Prev
		last.Prev.Next = last.Next
		delete(this.Data,last.Key)
	}
	// 队首插入元素
	tmpNode := &Node{
		Val:value,
		Key:key,
		Prev:nil,
		Next:this.Head.Next,
	}
	tmpNode.Next = this.Head.Next
	tmpNode.Prev = this.Head
	this.Head.Next.Prev = tmpNode
	this.Head.Next = tmpNode
	// 设置缓存
	this.Data[key] = tmpNode
}

