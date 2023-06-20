package main

import (
	"fmt"
	"godatabase/rpc/service"
	"log"
	"net/rpc"
)

func main()  {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	args := &service.Args{7,8}
	var reply int
	err = client.Call("Service.Multiply",args,&reply)
	if err != nil {
		log.Fatal("client.Call() error:",err)
	}
	fmt.Println(reply)
}
