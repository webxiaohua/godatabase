package main

import (
	"fmt"
	"log"
	"net"
)

func main(){
	sockP := "/Users/shenxinhua/workspace/tmp/test.sock"
	conn,err := net.Dial("unix",sockP)
	if err != nil {
		panic(err)
	}
	if _, err := conn.Write([]byte("hello server 12345678901234567890")); err != nil {
		log.Print(err)
		return
	}
	var buf = make([]byte, 1024)
	if _, err := conn.Read(buf); err != nil {
		panic(err)
	}
	fmt.Println("client recv: ", string(buf))
}
