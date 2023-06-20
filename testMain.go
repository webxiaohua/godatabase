package main

import (
	"fmt"
	"time"
)

type Day int

func (d *Day) toString() {
	fmt.Printf("today is %d", d)
}

func doIt2(a *int) {
	*a = 123
}

func main() {
	actBeginTime, _ := time.ParseInLocation("2006-01-02 15:04:05", "2023-01-01 00:00:00", time.Local)
	actEndTime, _ := time.ParseInLocation("2006-01-02 15:04:05", "2023-01-01 23:59:59", time.Local)
	now, _ := time.ParseInLocation("2006-01-02 15:04:05", "2023-01-01 23:59:59", time.Local)
	roundId, roundStartTime := calcCurrentRoundId(now, actBeginTime, actEndTime)
	roundStartTimeObj := time.Unix(roundStartTime, 0)
	fmt.Println(roundId, roundStartTimeObj.Format("2006-01-02 15:04:05"))
}

// 根据时间计算当前轮次ID
func calcCurrentRoundId(now, actBeginTime, actEndTime time.Time) (roundId, roundStartTime int64) {
	if now.Before(actBeginTime) || now.After(actEndTime) {
		return
	}
	diffSeconds := now.Unix() - actBeginTime.Unix()
	roundId = diffSeconds/(60+20) + 1
	roundStartTime = actBeginTime.Unix() + (60+20)*(roundId-1)
	return
}
