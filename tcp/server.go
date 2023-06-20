package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"time"
)

func init()  {
	fmt.Println("main.main init")
}
func main()  {
	// tcp 通信地址
	var tcpAddr *net.TCPAddr
	tcpAddr,_ = net.ResolveTCPAddr("tcp","127.0.0.1:9999")
	// 监听地址
	tcpListener,_:=net.ListenTCP("tcp",tcpAddr)
	defer tcpListener.Close()
	fmt.Println("Server ready to read...")
	// 循环接收客户端连接，创建协程处理具体链接
	for{
		tcpConn,err := tcpListener.AcceptTCP()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("A client connected :"+tcpConn.RemoteAddr().String())
		go tcpPipe(tcpConn)
	}
}

// 具体处理连接过程
func tcpPipe(conn *net.TCPConn) {
	ipStr := conn.RemoteAddr().String()
	defer func() {
		fmt.Println(" Disconnected :"+ipStr)
		conn.Close()
	}()
	// 获取一个连接的reader读取流
	reader := bufio.NewReader(conn)
	i:=0
	// 接收并返回信息
	for{
		message,err := reader.ReadString('\n')
		if err != nil || err == io.EOF {
			break
		}
		fmt.Println(message)
		time.Sleep(time.Second * 3)
		msg := time.Now().String() + conn.RemoteAddr().String() + " Server say hello! \n"
		b :=[]byte(msg)
		conn.Write(b)
		i++
		if i>10{
			break
		}
	}
}
