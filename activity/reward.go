package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

/************************ models ************************/
type SendRewardReq struct {
	Source       int64   `json:"source"`
	MsgId        string  `json:"msg_id"`
	Uids         []int64 `json:"uids"`
	PackageId    string  `json:"package_id"`
	MsgTime      int64   `json:"msg_time"`
	ExtraData    string  `json:"extra_data"`
	IsApply      int64   `json:"is_apply"`
	BusinessType string  `json:"business_type"`
	BusinessId   string  `json:"business_id"`
}
type SendRewardReq_ExtraData_Lottery struct {
	Count int64 `json:"count"`
}
type SendRewardReq_ExtraData struct {
	Lottery *SendRewardReq_ExtraData_Lottery `json:"lottery"`
}

func main() {
	arr := [][]string{}
	for i, item := range arr {
		msgId := "fix2022081114001:" + item[0]
		uid, _ := strconv.ParseInt(item[0], 10, 64)
		count, _ := strconv.ParseInt(item[1], 10, 64)
		packageId := "1642_0_5"
		//count,_ := strconv.ParseInt(item[3],10,64)
		fmt.Printf("第 %d 行： msgId:%s, uid:%d, packageId:%s \n", i, msgId, uid, packageId)
		normalSendReward(&SendRewardReq{
			Source:       10000,
			MsgId:        msgId,
			Uids:         []int64{uid},
			PackageId:    packageId,
			MsgTime:      0,
			ExtraData:    fmt.Sprintf("{\"lottery\":{\"count\":%d}}", count),
			IsApply:      0,
			BusinessType: "",
			BusinessId:   "",
		})
	}
}

// 发放活动奖励 执行时需要绑定host: [172.18.21.10 grpc-proxy.bilibili.co]
func simpleSendReward(uid, lotteryNum int64, packageId, msgId string) {
	apiUrl := "http://grpc-proxy.bilibili.co/live.reward/live.reward.rewardConfig/sendReward"
	extraData := &SendRewardReq_ExtraData{Lottery: &SendRewardReq_ExtraData_Lottery{Count: lotteryNum}}
	extraDataByte, _ := json.Marshal(extraData)
	params := &SendRewardReq{
		Source:       10000,
		MsgId:        msgId,
		Uids:         []int64{uid},
		PackageId:    packageId,
		MsgTime:      0,
		ExtraData:    string(extraDataByte),
		IsApply:      0,
		BusinessType: "",
		BusinessId:   "",
	}
	jsonByte, _ := json.Marshal(params)
	req, _ := http.NewRequest("POST", apiUrl, bytes.NewBuffer(jsonByte))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error: %+v", err)
		return
	}
	defer resp.Body.Close()
	fmt.Sprintf("status:%d\n", resp.Status)
	fmt.Sprintf("response:%+v\n", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

// 标准发奖
func normalSendReward(sendReward *SendRewardReq) {
	apiUrl := "http://grpc-proxy.bilibili.co/live.reward/live.reward.rewardConfig/sendReward"
	jsonByte, _ := json.Marshal(sendReward)
	req, _ := http.NewRequest("POST", apiUrl, bytes.NewBuffer(jsonByte))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error: %+v", err)
		return
	}
	defer resp.Body.Close()
	fmt.Sprintf("status:%d\n", resp.Status)
	fmt.Sprintf("response:%+v\n", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
