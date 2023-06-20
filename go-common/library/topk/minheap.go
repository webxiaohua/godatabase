// https://github.com/go-kratos/aegis/blob/main/pkg/minheap/minheap.go
package topk

type Node struct {
	Key   string
	Count uint32
}

type Nodes []*Node

func (n Nodes) Len() int {
	return len(n)
}

func (n Nodes) Less(i, j int) bool {
	return (n[i].Count < n[j].Count) || (n[i].Count == n[j].Count && n[i].Key > n[j].Key)
}

func (n Nodes) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n *Nodes) Push(val interface{}) {
	*n = append(*n, val.(*Node))
}

func (n *Nodes) Pop() interface{} {
	var val *Node
	val, *n = (*n)[len(*n)-1], (*n)[:len(*n)-1]
	return val
}

type Heap struct {
}
