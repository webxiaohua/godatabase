package main

import "fmt"

func main() {
	ids := []int{1}
	ids = ids[0:0]
	fmt.Println(ids)

	fmt.Printf("ids: %+v", ids)
}
