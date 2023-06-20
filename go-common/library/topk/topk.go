package topk

type Item struct {
	Key   string
	Count uint32
}

type Topk interface {
	// 新增元素，如果存在则返回true
	Add(item string, incr uint32) bool
	// 返回所有元素
	List() []Item
	// 移除元素
	Expelled() <-chan Item
}
