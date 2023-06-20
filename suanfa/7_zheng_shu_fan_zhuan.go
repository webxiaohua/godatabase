// 7.整数反转
/*
输入：x = 123
输出：321

-231 <= x <= 231 - 1
*/
package main

import (
	"fmt"
	"math"
)

func main(){
	res := reverse7(1534236469)
	fmt.Println(res)
}

func reverse7(x int) int {
	// 除10
	var res int
	for {
		if x == 0 {
			break
		}
		res = res*10 + x%10
		x = x/10
	}
	if res > math.MaxInt32 || res < math.MinInt32{
		return 0
	}
	return res
}