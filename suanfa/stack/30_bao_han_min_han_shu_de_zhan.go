// 剑指offer 30.包含min函数的栈
/*
定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中，调用 min、push 及 pop 的时间复杂度都是 O(1)。


*/
package main

func main() {

}


type MinStack struct {
	MinArr []int // 利用数组存最小值，每次新来元素会跟当前最后一个元素进行比较，如果超过则继续压入当前元素
	Data []int // 栈内数据
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		MinArr:make([]int,0),
		Data:make([]int,0),
	}
}


func (this *MinStack) Push(x int)  {
	if len(this.MinArr) == 0 {
		this.MinArr = append(this.MinArr,x)
	}else{
		if(x >= this.MinArr[len(this.MinArr) - 1]) {
			this.MinArr = append(this.MinArr,this.MinArr[len(this.MinArr) - 1])
		}else{
			this.MinArr = append(this.MinArr,x)
		}
	}
	this.Data = append(this.Data,x)
}


func (this *MinStack) Pop()  {
	if len(this.Data) > 0 {
		this.MinArr = this.MinArr[0:len(this.MinArr)-1]
		this.Data = this.Data[0:len(this.Data)-1]
	}
}


func (this *MinStack) Top() int {
	if len(this.Data) > 0 {
		return this.Data[len(this.Data)-1]
	}else{
		return 0
	}
}


func (this *MinStack) Min() int {
	if len(this.MinArr) == 0 {
		return 0
	}else{
		return this.MinArr[len(this.MinArr)-1]
	}
}