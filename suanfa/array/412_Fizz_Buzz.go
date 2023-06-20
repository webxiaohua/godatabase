// 412
/*
给你一个整数 n ，找出从 1 到 n 各个整数的 Fizz Buzz 表示，并用字符串数组 answer（下标从 1 开始）返回结果，其中：
answer[i] == "FizzBuzz" 如果 i 同时是 3 和 5 的倍数。
answer[i] == "Fizz" 如果 i 是 3 的倍数。
answer[i] == "Buzz" 如果 i 是 5 的倍数。
answer[i] == i （以字符串形式）如果上述条件全不满足。

输入：n = 3
输出：["1","2","Fizz"]
*/
package main

import (
	"fmt"
	"strconv"
)

func main(){
	res := fizzBuzz(3)
	fmt.Println(res)
}

func fizzBuzz(n int) []string {
	res := make([]string,0)
	for i:=1;i<=n;i++ {
		if i % 15 == 0 {
			res = append(res,"FizzBuzz")
		}else if i % 3 == 0 {
			res = append(res,"Fizz")
		}else if i % 5 == 0 {
			res = append(res,"Buzz")
		}else{
			res = append(res,strconv.Itoa(i))
		}
	}
	return res
}