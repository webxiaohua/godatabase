package main

import "fmt"

type Student struct {
	Age int64 `json:"age"`
	Name string `json:"name"`
}

func main()  {
	a := &Student{
		Age:  11,
		Name: "testa",
	}
	b := new(Student)
	fmt.Println(*a)
	*b = *a
	fmt.Printf("a addr: %p, b addr: %p, a val: %+v, b val: %+v",a,b,a,b)
}