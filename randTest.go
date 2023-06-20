package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	roundPoolMap := make(map[int64]*Act101814KvConf_PoolInfo)
	roundPoolMap[1] = &Act101814KvConf_PoolInfo{Rate: 1920}
	roundPoolMap[2] = &Act101814KvConf_PoolInfo{Rate: 990}
	roundPoolMap[3] = &Act101814KvConf_PoolInfo{Rate: 1920}
	roundPoolMap[4] = &Act101814KvConf_PoolInfo{Rate: 1920}
	roundPoolMap[5] = &Act101814KvConf_PoolInfo{Rate: 690}
	roundPoolMap[6] = &Act101814KvConf_PoolInfo{Rate: 400}
	roundPoolMap[7] = &Act101814KvConf_PoolInfo{Rate: 1920}
	roundPoolMap[8] = &Act101814KvConf_PoolInfo{Rate: 240}

	result := make(map[int64]int64)
	var i int64
	for i = 1; i <= 8; i++ {
		result[i] = 0
	}

	for j := 0; j < 1000; j++ {
		winPoolId := Lottery2(roundPoolMap)
		result[winPoolId]++
	}
	fmt.Println(result)
}

type Act101814KvConf_PoolInfo struct {
	Id       int64  `json:"id"`       // 奖池id
	Name     string `json:"name"`     // 奖池名
	Rate     int64  `json:"rate"`     // 中奖概率万分比
	Multiple int64  `json:"multiple"` // 获奖倍数
}

func Lottery(roundPoolMap map[int64]*Act101814KvConf_PoolInfo) (winPoolId int64) {
	var index int64
	lotteryPoolList := make([][]int64, 0) // index,poolId
	for poolId, _ := range roundPoolMap {
		index += roundPoolMap[poolId].Rate
		lotteryPoolList = append(lotteryPoolList, []int64{index, poolId})
	}
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Int63n(index)
	for _, item := range lotteryPoolList {
		if item[0] >= randNum {
			winPoolId = item[1]
			break
		}
	}
	return
}

func Lottery2(roundPoolMap map[int64]*Act101814KvConf_PoolInfo) (winPoolId int64) {
	var index int64
	lotteryPoolList := make([][]int64, 0) // index,poolId
	var i int64
	for i = 1; i <= 8; i++ {
		index += roundPoolMap[i].Rate
		lotteryPoolList = append(lotteryPoolList, []int64{index, i})
	}
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Int63n(index)
	for _, item := range lotteryPoolList {
		if item[0] >= randNum {
			winPoolId = item[1]
			break
		}
	}
	return
}
