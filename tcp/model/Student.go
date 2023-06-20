package model

import "fmt"

type Student struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
}

func init()  {
	fmt.Println("model.Student init")
	go func() {
		fmt.Println("model.Student init go routine")
	}()
}