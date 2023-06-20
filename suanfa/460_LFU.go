/*
实现 LFUCache 类：

LFUCache(int capacity) - 用数据结构的容量 capacity 初始化对象
int get(int key) - 如果键 key 存在于缓存中，则获取键的值，否则返回 -1 。
void put(int key, int value) - 如果键 key 已存在，则变更其值；如果键不存在，请插入键值对。当缓存达到其容量 capacity 时，则应该在插入新项之前，移除最不经常使用的项。在此问题中，当存在平局（即两个或更多个键具有相同使用频率）时，应该去除 最近最久未使用 的键。
为了确定最不常使用的键，可以为缓存中的每个键维护一个 使用计数器 。使用计数最小的键是最久未使用的键。

当一个键首次插入到缓存中时，它的使用计数器被设置为 1 (由于 put 操作)。对缓存中的键执行 get 或 put 操作，使用计数器的值将会递增。

函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。

*/
// https://blog.csdn.net/qq_43756091/article/details/122426202
package main

import "fmt"

// 双向链表
type LFUNode struct {
	key, val, cnt int // 键,值,访问次数
	prev          *LFUNode
	next          *LFUNode
}

type LFUCache struct {
	capacity, size int              // 容量，数量
	data           map[int]*LFUNode // 缓存数据 key:数据节点  此处注意不支持并发场景
	head, tail     *LFUNode         // 头节点，尾节点
}

func NewLFUCache(capacity int) *LFUCache {
	res := &LFUCache{
		capacity: capacity,
		size:     0,
		data:     make(map[int]*LFUNode),
		head:     nil,
		tail:     nil,
	}
	//res.head.next = res.tail
	//res.tail.prev = res.head
	return res
}

// 移除节点
func (l *LFUCache) remove(node *LFUNode) {
	if node.prev != nil {
		if node.next != nil {
			// 中间节点
			node.prev.next = node.next
			node.next.prev = node.prev
		} else {
			// 尾节点
			node.prev.next = nil
			l.tail = node.prev
		}
	} else {
		// 头节点
		if node.next != nil {
			l.head = node.next
			l.head.prev = nil
		} else {
			// 单节点删除
			l.head = nil
			l.tail = nil
		}
	}
	l.size--
	delete(l.data, node.key)
}

// 链表排序，访问次数越多链表排序越靠前，如果节点已经存在则先删除节点
func (l *LFUCache) sort(node *LFUNode, isUpdate bool) {
	if isUpdate {
		// 定位到第一个不满足条件的节点
		front := node.prev
		for front != nil {
			if front.cnt > node.cnt {
				break
			} else {
				front = front.prev
			}
		}
		// 在此节点之后插入node
		if front == nil {
			// 插到链表首位置
			tmp := l.head
			l.head = node
			node.next = tmp
		} else {
			//front->tmp => front->node->tmp
			tmp := front.next
			front.next = node
			node.prev = front
			node.next = tmp
			tmp.prev = node
		}
	} else {
		// 新增元素直接排入队尾
		l.tail.next = node
		node.prev = l.tail
		l.tail = node
		l.size++
	}
}

func (l *LFUCache) Get(key int) int {
	if l.capacity == 0 {
		return -1
	}
	if node, ok := l.data[key]; !ok {
		return -1
	} else {
		node.cnt++
		l.data[key] = node
		l.sort(node, true)
		return node.val
	}
}

func (l *LFUCache) Put(key int, value int) {
	if l.capacity == 0 {
		return
	}
	node, ok := l.data[key]
	if !ok {
		// 容量已满
		if l.capacity == l.size {
			l.remove(l.tail)
		}
		new := &LFUNode{
			key: key,
			val: value,
			cnt: 1,
		}
		node = new
		l.sort(node, false)
		l.size++
	} else {
		node.val = value
		node.cnt++
		l.sort(node, true)
	}
	l.data[key] = node
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	lfu := NewLFUCache(2)
	lfu.Put(1, 1)
	lfu.Put(2, 2)
	lfu.Put(3, 3)
	fmt.Println("get 1: ", lfu.Get(1))
	fmt.Println("get 2: ", lfu.Get(2))
	fmt.Println("get 3: ", lfu.Get(3))
	fmt.Println("get 4: ", lfu.Get(4))
}
