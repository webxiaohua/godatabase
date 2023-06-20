package main

import "fmt"
// go build --gcflags="-l -N" -o pointer
func main()  {
	testa := 1000
	testb := 2000
	fmt.Printf("a addr: %p \n",&testa)
	fmt.Printf("b addr: %p \n",&testb)
}
