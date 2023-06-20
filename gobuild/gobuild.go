
package main

import (
	"fmt"
	"godatabase/gobuild/json"
)

func main(){
	human := &Human{
		Name: "Robot",
		Age:  10,
	}
	fmt.Println(json.Marshal(human))
}