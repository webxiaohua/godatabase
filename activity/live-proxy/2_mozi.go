package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// 新增报名
func SignUp(uid, signId, startTime, endTime, teamId int64, otherMessage string) {
	type SignUpReq struct {
		// 用户uid
		Uid int64 `json:"uid"`
		// 报名id
		SignId int64 `json:"sign_id"`
		// 报名类型 1:特征报名2:录入
		SignType int64 `json:"sign_type"`
		// 报名开始时间
		StartTime int64 `json:"start_time"`
		// 报名结束时间
		EndTime int64 `json:"end_time"`
		// 额外信息
		OtherMessage string `json:"other_message"`
		// 赛道
		TeamId int64 `json:"team_id"`
	}
	apiUrl := "https://manager.bilibili.co/x/admin/live-proxy/unsafeproxy"
	req := &SignUpReq{
		Uid:          uid,
		SignId:       signId,
		SignType:     2,
		StartTime:    startTime,
		EndTime:      endTime,
		OtherMessage: otherMessage,
		TeamId:       teamId,
	}
	reqByte, _ := json.Marshal(req)
	apiProxyParams := &ApiProxyParams{
		Params:   string(reqByte),
		AppId:    "live.mozi",
		Module:   "1020001",
		GrpcPath: "/live.mozi.signup/SignUp",
	}
	apiProxyParamsBody := GetFormParams(apiProxyParams)
	fmt.Println(apiProxyParamsBody)
	HttpPost(apiUrl, apiProxyParamsBody)
}

// 更新报名
func UpdateUserSignUpInfo(uid, signId, teamId, startTime, endTime int64, otherMessage string) {
	type UpdateUserSignUpInfoReq struct {
		// 用户uid
		Uid int64 `json:"uid"`
		// 报名id
		SignId int64 `json:"sign_id"`
		// 修改数据是哪个 0 所有参数全部覆盖 1只修改other_message 2只修改时间 3. 只修改赛道team_id
		UpdateType int64 `json:"update_type"`
		// 其他信息
		OtherMessage string `json:"other_message"`
		// 开始时间
		StartTime int64 `json:"start_time"`
		// 结束时间
		EndTime int64 `json:"end_time"`
		// 赛道
		TeamId int64 `json:"team_id"`
		// sign state
		SignState int64 `json:"sign_state"`
	}
	apiUrl := "https://manager.bilibili.co/x/admin/live-proxy/unsafeproxy"
	req := &UpdateUserSignUpInfoReq{
		Uid:          uid,
		SignId:       signId,
		UpdateType:   0,
		OtherMessage: otherMessage,
		StartTime:    startTime,
		EndTime:      endTime,
		TeamId:       teamId,
		SignState:    1,
	}
	reqByte, _ := json.Marshal(req)
	apiProxyParams := &ApiProxyParams{
		Params:   string(reqByte),
		AppId:    "live.mozi",
		Module:   "1020001",
		GrpcPath: "/live.mozi.signup/UpdateUserSignUpInfo",
	}
	apiProxyParamsBody := GetFormParams(apiProxyParams)
	fmt.Println(apiProxyParamsBody)
	HttpPost(apiUrl, apiProxyParamsBody)
}

// 获取生效中的活动列表
func GetEffectActList() {
	type GetEffectActReq struct {
		// 主活动开始时间戳
		Timestamp int64 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp"`
	}
	apiUrl := "https://manager.bilibili.co/x/admin/live-proxy/unsafeproxy"
	req := &GetEffectActReq{
		Timestamp: time.Now().AddDate(-2, 0, 0).Unix(),
	}
	reqByte, _ := json.Marshal(req)
	apiProxyParams := &ApiProxyParams{
		Params:   string(reqByte),
		AppId:    "live.mozi",
		Module:   "1020001",
		GrpcPath: "/live.mozi.actCenterBackend/GetEffectActList",
	}
	apiProxyParamsBody := GetFormParams(apiProxyParams)
	fmt.Println(apiProxyParamsBody)
	res := HttpPost(apiUrl, apiProxyParamsBody)
	fmt.Println(res)
}
