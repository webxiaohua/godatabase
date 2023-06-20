package main

import (
	"fmt"
	"time"
)

func main()  {

	i:=0
	go func() {
		for a:=0;a<10;a++ {
			time.Sleep(time.Second)
			i++
		}
	}()

	go func() {
		for a:=0;a<15;a++ {
			time.Sleep(time.Second)
			i++
		}
	}()

	time.Sleep(time.Second * 15)
	fmt.Printf("i value : %d",i)

	var err error
	doIt(err)

}

func doIt(err error)  {
	fmt.Printf("err: %v",err)
	return
}