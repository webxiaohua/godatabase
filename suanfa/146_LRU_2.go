package main

import "container/list"

type CacheNode struct {
	Key int
	Value int
}

type LRUCache2 struct {
	cap int // 容量
	dList *list.List // 链表
	cacheMap map[int]*list.Element // hash
}


func Constructor2(capacity int) *LRUCache2 {
	ret := &LRUCache2{
		cap:capacity,
		dList : list.New(),
		cacheMap:make(map[int]*list.Element),
	}
	return ret
}

func (l *LRUCache2) Set(key int,val int) {
	if l.dList == nil {
		return
	}
	if ele,ok := l.cacheMap[key];ok {
		// 移到最前面
		l.dList.MoveToFront(ele)
		ele.Value = val
		return
	}
	if l.dList.Len() == l.cap {
		// 删除表尾元素
		lastEle := l.dList.Back()
		if lastEle == nil {
			return
		}
		l.dList.Remove(lastEle)
		delete(l.cacheMap,lastEle.Value.(CacheNode).Key)
	}
	node := &CacheNode{
		Key:   key,
		Value: val,
	}
	l.dList.MoveToFront()
}

func (l *LRUCache2) Get(key int) int {

}