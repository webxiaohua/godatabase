package main

import "fmt"

type Test8 struct {
	Id int
}

func main() {
	m := map[string]*Test8{"a": &Test8{Id: 1}}
	p := m["a"]
	fmt.Println(p)
	p = &Test8{Id: 2}
	fmt.Println(m["a"])

	b := make(map[int64]int64)
	b1 := b[0]
	fmt.Println(b1)
}
