/*
@Desc 斐波那契数列输出
*/

package main

import "fmt"

func main (){
	fmt.Println("循环法：")
	xunhuan()
	fmt.Println("\n递归法：")
	digui()
}
// 循环法
func xunhuan()  {
	a:=1
	b:=1
	for i:=0;i<10;i++ {
		fmt.Printf(" %d %d",a,b)
		a=a+b
		b=b+a
	}
}
// 递归法
func digui() {
	var myFun func(int)int
	myFun = func(n int)int{
		if n==1 || n == 2 {
			return 1
		}
		return myFun(n-1) + myFun(n-2)
	}
	for i:=1;i<=10;i++ {
		fmt.Printf("%d ",myFun(i))
	}
}