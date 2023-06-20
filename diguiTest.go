package main

import "fmt"

func main(){
	processing(0)
}

func processing(i int64)  {
	fmt.Println(i)
	i++
	if i > 10000000{
		return
	}
	processing(i)
}