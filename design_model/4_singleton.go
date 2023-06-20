/*
创建型-单例：

具体思路：
	仅提供一个全局访问入口
常见应用应用：

*/
package main

import (
	"fmt"
	"sync"
	"time"
)

type Singleton struct {
}

var singletonInstance *Singleton
var lock sync.Mutex

func getSingletonInstance() *Singleton {
	if singletonInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singletonInstance == nil {
			singletonInstance = new(Singleton)
			fmt.Println("实例创建成功")
		} else {
			fmt.Println("实例已经创建2")
		}
	} else {
		fmt.Println("实例已经创建1")
	}
	return singletonInstance
}

func main() {
	for i := 0; i < 5; i++ {
		go getSingletonInstance()
	}
	time.Sleep(time.Second * 3)
}
