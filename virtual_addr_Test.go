package main

import "fmt"

var global int

func main(){
	local := 5
	p := new(chan int)
	fmt.Printf("Address of main is %p\n",main)
	fmt.Printf("Address of global is %p\n",&global)
	fmt.Printf("Address of local is %p\n",&local)
	fmt.Printf("Address of p is %p\n",&p)
}
