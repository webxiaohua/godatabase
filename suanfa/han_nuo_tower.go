package main

import "fmt"

var count int

func main()  {
	hannuo(3,"A","B","C")
}

func hannuo(n int, a,b,c string) {
	if n == 1 {
		// 仅有一个时，将盘子从A移动到C
		count++
		fmt.Printf("第 %d 次移动，把 %d 号圆盘从 %s -> 移到 -> %s \n", count,n, a, c)
	} else {
		// 把 n-1 个盘子从A移动到B 以C为辅助塔
		hannuo(n-1, a, c, b)
		// 把 n 号盘子从A移动到C
		count++
		fmt.Printf("第 %d 次移动，把 %d 号圆盘从 %s -> 移到 -> %s \n", count,n, a, c)
		// 把 n-1 个盘子从B移动到C 以A为辅助塔
		hannuo(n-1, b, a, c)
	}
}
