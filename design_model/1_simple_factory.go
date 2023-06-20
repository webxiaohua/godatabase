/*
创建型-简单工厂：
	抽象产品接口
    具体产品实现
	具体工厂实现
具体思路：
	通过具体工厂创建产品，再使用产品，方便了同一种功能实现多样性
常见应用应用：
	发奖（发放不同的奖品，基于奖励类型屏蔽不同奖品的发放差异）、活动（基于活动ID屏蔽各个活动差异，执行不同的活动逻辑）
*/
package main

import "fmt"

// 定义奖励抽象
type IReward interface {
	Send()
}

// 定义具体奖励 - 头像框
type HeadPicReward struct{}

func (*HeadPicReward) Send() {
	fmt.Println("发送头像框奖励")
}

// 定义具体奖励 - 实物奖励
type PhysicalReward struct{}

func (*PhysicalReward) Send() {
	fmt.Println("发送实物奖励")
}

// 奖励工厂
type RewardFactory struct {
}

// 创建奖品实例
func (*RewardFactory) CreateReward(rewardType string) IReward {
	if rewardType == "headpic" {
		return new(HeadPicReward)
	} else {
		return new(PhysicalReward)
	}
}

func main() {
	factory := new(RewardFactory)
	factory.CreateReward("headpic").Send()
	factory.CreateReward("physical").Send()
}
