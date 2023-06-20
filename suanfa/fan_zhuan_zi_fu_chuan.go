// 反转字符串
package main

import "fmt"

func main(){
	str := "abcdefg"
	str = reverse(str)
	fmt.Println(str)
}

func reverse(str string)(res string){
	arr := []rune(str)
	for i,j := 0,len(arr)-1;i<j;i,j=i+1,j-1 {
		arr[i],arr[j] = arr[j],arr[i]
	}
	res = string(arr)
	return
}
