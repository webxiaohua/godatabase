package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){
 	test3()
}

func test1(){
	var lock sync.RWMutex
	lock.RLock()
	fmt.Println("RLock 1")
	lock.RLock()
	fmt.Println("RLock 2")
	lock.RUnlock()
	fmt.Println("RUnlock 2")
	lock.RUnlock()
	fmt.Println("RUnlock 1")
}

func test2(){
	// deadlock
	var wg sync.WaitGroup
	wg.Add(1)
	var lock sync.RWMutex
	lock.RLock()
	fmt.Println("rlock 1")
	go func() {
		lock.Lock()
		fmt.Println("lock...")
		time.Sleep(time.Second)
		lock.Unlock()
		fmt.Println("unlock...")
		wg.Done()
	}()
	time.Sleep(time.Second)
	lock.RLock()
	fmt.Println("rlock 2")
	lock.RUnlock()
	fmt.Println("runlock 2")
	lock.RUnlock()
	fmt.Println("runlock 1")
}

func test3(){
	// deadlock
	var wg sync.WaitGroup
	wg.Add(1)
	var lock sync.RWMutex
	lock.RLock()
	fmt.Println("rlock 1")
	lock.RUnlock()
	fmt.Println("runlock 1")
	go func() {
		lock.Lock()
		fmt.Println("lock...")
		time.Sleep(time.Second)
		lock.Unlock()
		fmt.Println("unlock...")
		wg.Done()
	}()
	time.Sleep(time.Second)
	lock.RLock()
	fmt.Println("rlock 2")
	lock.RUnlock()
	fmt.Println("runlock 2")
}
