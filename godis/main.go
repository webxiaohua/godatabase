package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func ListenAndServe(address string)  {
	listener,err := net.Listen("tcp",address)
	if err != nil {
		log.Fatal("net.Listen() err(%v)",err)
	}
	defer listener.Close()
	log.Println(fmt.Sprintf("bind: %s, start listening...",address))
	for{
		conn,err := listener.Accept()
		if err != nil {
			log.Fatal(fmt.Sprintf("accept err: %v",err))
		}
		// 开启新的处理协程
		go Handler(conn)
	}
}

func Handler(conn net.Conn)  {
	reader := bufio.NewReader(conn)
	for{
		// ReadString 会一直阻塞到遇到分隔符 \n
		msg,err := reader.ReadString('\n')
		// 遇到分隔符后 ReadString 会返回上次遇到分隔符到现在收到的所有数据
		// 若在遇到分隔符之前发生异常，ReadString会返回已收到的数据和错误信息
		if err != nil {
			if err == io.EOF {
				log.Println("connection close")
			}else{
				log.Println(err)
			}
			return
		}
		b:=[]byte(msg)
		// 将收到的消息发送给客户端
		conn.Write(b)
	}
}

func main()  {
	ListenAndServe(":7900")
}