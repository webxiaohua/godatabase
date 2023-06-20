/*
创建型-抽象工厂：
	抽象产品接口族
    具体产品实现
	抽象工厂接口族
	具体工厂实现
	客户端调用
具体思路：
	客户端无需关心具体产品，只需要对接具体的产品工厂即可，可以利用抽象工厂进行替换
常见应用应用：

*/
package main

import (
	"fmt"
)

/****************** 按钮 ******************/
type IButton interface {
	Display()
}

type SpringButton struct {
}

func (*SpringButton) Display() {
	fmt.Println("显示浅绿色按钮")
}

type SummerButton struct {
}

func (*SummerButton) Display() {
	fmt.Println("显示浅蓝色按钮")
}

/****************** 文本框 ******************/
type ITextBox interface {
	Display()
}

type SpringTextBox struct {
}

func (*SpringTextBox) Display() {
	fmt.Println("显示浅绿色文本框")
}

type SummerTextBox struct {
}

func (*SummerTextBox) Display() {
	fmt.Println("显示浅蓝色文本框")
}

/****************** 皮肤界面抽象工厂 ******************/

type SkinFactory interface {
	CreateButton() IButton
	CreateTextBox() ITextBox
}

/****************** 春季皮肤界面具体工厂 ******************/
type SpringSkinFactory struct {
}

func (*SpringSkinFactory) CreateButton() IButton {
	return new(SpringButton)
}

func (*SpringSkinFactory) CreateTextBox() ITextBox {
	return new(SpringTextBox)
}

/****************** 夏季皮肤界面具体工厂 ******************/
type SummerSkinFactory struct {
}

func (*SummerSkinFactory) CreateButton() IButton {
	return new(SummerButton)
}

func (*SummerSkinFactory) CreateTextBox() ITextBox {
	return new(SummerTextBox)
}

/****************** 客户端 ******************/
func main() {
	skinFactory := new(SpringSkinFactory)
	skinFactory.CreateButton().Display()
	skinFactory.CreateTextBox().Display()
}
