/*
创建型-工厂方法：
	抽象产品接口
    具体产品实现
	抽象工厂接口
	具体工厂实现
	客户端调用
具体思路：
	客户端无需关心具体产品，只需要对接具体的产品工厂即可，可以利用抽象工厂进行替换
常见应用应用：

*/
package main

import (
	"fmt"
)

type ILog interface {
	WriteLog()
}

type FileLog struct {
}

func (*FileLog) WriteLog() {
	fmt.Println("I'm file log")
}

type DBLog struct {
}

func (*DBLog) WriteLog() {
	fmt.Println("I'm db log")
}

type ILogFactory interface {
	CreateLog() ILog
}

type FileLogFactory struct {
}

func (*FileLogFactory) CreateLog() ILog {
	return new(FileLog)
}

type DBLogFactory struct {
}

func (*DBLogFactory) CreateLog() ILog {
	return new(DBLog)
}

func main() {
	factory := new(FileLogFactory)
	factory.CreateLog().WriteLog()
}
