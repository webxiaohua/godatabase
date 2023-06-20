package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map
	wg := &sync.WaitGroup{}
	wg.Add(10)
	//Store
	for i:=0;i<10;i++ {
		go func(i int) {
			m.Store(i,"tt")
		}(i)
	}


	//LoadOrStore
	//若key不存在，则存入key和value，返回false和输入的value
	v,ok := m.LoadOrStore(1,"aaa")
	fmt.Println(ok,v) //false aaa

	//若key已存在，则返回true和key对应的value，不会修改原来的value
	v,ok = m.LoadOrStore(1,"aaa")
	fmt.Println(ok,v) //false aaa

	//Load
	v,ok = m.Load(1)
	if ok{
		fmt.Println("it's an existing key,value is ",v)
	} else {
		fmt.Println("it's an unknown key")
	}
}