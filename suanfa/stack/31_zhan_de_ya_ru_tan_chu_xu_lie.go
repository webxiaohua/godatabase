// 剑指offer 31.栈的压入、弹出序列
/*
输入：pushed = [1,2,3,4,5], popped = [4,5,3,2,1]
输出：true
解释：我们可以按以下顺序执行：
push(1), push(2), push(3), push(4), pop() -> 4,
push(5), pop() -> 5, pop() -> 3, pop() -> 2, pop() -> 1
*/
package main

import "fmt"

func main(){
	pushed := []int{2,1,0}
	popped := []int{1,2,0}
	res := validateStackSequences(pushed,popped)
	fmt.Println(res)
}

func validateStackSequences(pushed []int, popped []int) bool {
	// 定义一个栈，pushed进栈时判断是否为popped当前指针元素，如果是，直接跳过压栈，popped指针指向后一个元素，最后当所有栈内元素弹出后若 popped 指针指向 len(popped)，则成功，否则失败
	stack := make([]int,0)
	l,r := 0,0
	for ;l<len(pushed);l++{
		stack = append(stack,pushed[l])
		REDO:
		if stack[len(stack)-1] == popped[r] {
			stack = stack[0:len(stack)-1]
			r++
			if len(stack) != 0 {
				goto REDO
			}
		}
	}
	return r == len(popped)
	/*for {
		if len(stack) == 0 {
			break
		}
		popVal := stack[len(stack)-1]
		stack = stack[0:len(stack)-1]
		if popVal != popped[r] {
			return false
		}else{
			r++
		}
	}
	return true*/
}