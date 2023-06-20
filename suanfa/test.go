package main

import (
	"fmt"
	"strings"
)

func main() {
	a:=[]int64{1,2,3,4,6}
	res := strings.Replace(strings.Trim(fmt.Sprint(a), "[]"), " ", ",", -1)
	fmt.Println(res)
}

