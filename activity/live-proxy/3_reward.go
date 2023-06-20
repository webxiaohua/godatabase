package main

import (
	"encoding/json"
	"fmt"
)

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

// 标准发奖
func SendReward(msgId, packageId string, uid, lotteryCount int64) {
	apiUrl := "https://manager.bilibili.co/x/admin/live-proxy/unsafeproxy"
	req := &SendRewardReq{
		Source:       10000,
		MsgId:        msgId,
		Uids:         []int64{uid},
		PackageId:    packageId,
		MsgTime:      0,
		ExtraData:    fmt.Sprintf("{\"lottery\":{\"count\":%d}}", lotteryCount),
		IsApply:      0,
		BusinessType: "",
		BusinessId:   "",
	}
	reqByte, _ := json.Marshal(req)
	apiProxyParams := &ApiProxyParams{
		Params:   string(reqByte),
		AppId:    "live.reward",
		Module:   "1020001",
		GrpcPath: "/live.reward.rewardConfig/sendReward",
	}
	apiProxyParamsBody := GetFormParams(apiProxyParams)
	fmt.Println(apiProxyParamsBody)
	HttpPost(apiUrl, apiProxyParamsBody)
}
