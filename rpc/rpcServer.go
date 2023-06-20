package main

import (
	"fmt"
	"godatabase/rpc/service"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)


func main()  {
	stop := make(chan int)
	service := new(service.Service)
	rpc.Register(service)
	rpc.HandleHTTP()
	l,e := net.Listen("tcp",":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l,nil)
	go func() {
		fmt.Println("server shutdown after 5 seconds...")
		time.Sleep(5 * time.Second)
		stop<- 1
	}()
	select {
		case <-stop:
			fmt.Println("stop")
	}
}