package main

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"time"
)

func main(){
	node1,_ := snowflake.NewNode(1)
	node2,_ := snowflake.NewNode(2)
	go func() {
		for i:=0;i<10;i++{
			time.Sleep(time.Millisecond * 10)
			snowId := node1.Generate().Int64()
			fmt.Println("node1:",snowId)
		}
	}()
	go func() {
		for i:=0;i<10;i++{
			time.Sleep(time.Millisecond * 10)
			snowId := node2.Generate().Int64()
			fmt.Println("node2:",snowId)
		}
	}()
	time.Sleep(time.Second * 5)
}
