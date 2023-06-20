package main

import (
	"encoding/json"
	"fmt"
)

type AreUidsInRoomReq struct {
	RoomId int64   `json:"roomid"`
	Uids   []int64 `json:"uids"`
	Group  string  `json:"group"`
}

func AreUidsInRoom(req *AreUidsInRoomReq, isPre bool) {
	apiUrl := "https://manager.bilibili.co/x/admin/live-proxy/unsafeproxy"
	if isPre {
		apiUrl = "http://ff-pre.manager.bilibili.co/x/admin/live-proxy/unsafeproxy"
	}
	reqByte, _ := json.Marshal(req)
	apiProxyParams := &ApiProxyParams{
		Params:   string(reqByte),
		AppId:    "live.live.online",
		Module:   "1020001",
		GrpcPath: "/live.live.online.v1.Online/AreUidsInRoom",
	}
	apiProxyParamsBody := GetFormParams(apiProxyParams)
	resp := HttpPost(apiUrl, apiProxyParamsBody)
	fmt.Println(resp)
	return
}
