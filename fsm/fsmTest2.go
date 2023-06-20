package main

import (
	"fmt"
	"github.com/looplab/fsm"
)

func main(){
	fmt.Print("xxx")
	fsm := fsm.NewFSM("close",fsm.Events{
		// 开： close -> open
		{Name:"open",Src:[]string{"close"},Dst:"open"},
		// 关： open -> close
		{Name:"close",Src:[]string{"open"},Dst:"close"},
	},fsm.Callbacks{})
	fmt.Println("初始状态：",fsm.Current())
	// 关
	err := fsm.Event("close")
	if err != nil {
		fmt.Println("已经关闭，无需重复关闭。",err)
	}
	fmt.Println("当前状态：",fsm.Current())
	// 开
	err = fsm.Event("open")
	if err != nil {
		fmt.Println("已经打开，无需重复打开。",err)
	}
	fmt.Println("当前状态：",fsm.Current())
	// 关
	err = fsm.Event("close")
	if err != nil {
		fmt.Println("已经关闭，无需重复关闭。",err)
	}
	fmt.Println("当前状态：",fsm.Current())
}