package main

import (
	"fmt"
	"io"
	"strings"
)

func ReadFrom(reader io.Reader,nums int) ([]byte,error) {
	b := make([]byte,nums)
	n,err := reader.Read(b)
	if n > 0 {
		return b[:n],nil
	}
	return b,err
}

func main(){
	data,err := ReadFrom(strings.NewReader("hello world"),11)
	if err != nil {
		fmt.Println("[error]:",err)
	}else{
		fmt.Println(string(data))
	}
}