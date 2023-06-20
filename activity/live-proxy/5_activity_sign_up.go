package main

import (
	"encoding/json"
	"fmt"
)

type GetUserIdentityReq struct {
	// 用户uid
	Uid int64 `json:"uid"`
}

type GetUserIdentityResp struct {
	Code    int64                    `json:"code"`
	Message string                   `json:"message"`
	Data    *GetUserIdentityRespData `json:"data"`
}
type GetUserIdentityRespData struct {
	// 当前用户身份 //直签: 1, 素人: 2, 超电: 3, 官签: 4, 普通公会: 5
	UidIdentity string `json:"uid_identity"`
}

// 更新活动数值
func GetUserIdentity(uid int64) (resp *GetUserIdentityResp) {
	apiUrl := "https://manager.bilibili.co/x/admin/live-proxy/unsafeproxy"
	req := &GetUserIdentityReq{
		Uid: uid,
	}
	reqByte, _ := json.Marshal(req)
	apiProxyParams := &ApiProxyParams{
		Params:   string(reqByte),
		AppId:    "live.live.activitysignup",
		Module:   "1020001",
		GrpcPath: "/live.activitysignup.SignupFront/GetUserIdentity",
	}
	apiProxyParamsBody := GetFormParams(apiProxyParams)
	//fmt.Println(apiProxyParamsBody)
	respJson := HttpPost(apiUrl, apiProxyParamsBody)
	resp = &GetUserIdentityResp{}
	err := json.Unmarshal([]byte(respJson), &resp)
	if err != nil {
		fmt.Println(err)
	}
	return resp
}
