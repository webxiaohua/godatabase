package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
)

var (
	a = "Hello world"
	b atomic.Value
)

func main() {
	randNum := rand.Int63n(0)
	fmt.Println(randNum)
}

func testA() {
	b.Store(true)
}

func test() (a int) {
	res := 0
	defer func() {
		a = 1
	}()
	return res
}
