package main

import (
	"encoding/json"
	"fmt"
)

// 根据uid获取房间信息

type UIDsReq struct {
	// 主播uids
	Uids []int64 `protobuf:"varint,1,rep,packed,name=uids,proto3" json:"uids,omitempty" form:"uids" validate:"required,gt=0,dive,gt=0"`
	// 要获取的房间信息维度 status:状态相关 show:展示相关 area:分区相关 anchor:主播相关
	Attrs []string `protobuf:"bytes,2,rep,name=attrs,proto3" json:"attrs,omitempty" form:"attrs" validate:"required,gt=0,dive,gt=0"`
}

func GetMultipleByUids(req *UIDsReq, isPre bool) {
	apiUrl := "https://manager.bilibili.co/x/admin/live-proxy/unsafeproxy"
	if isPre {
		apiUrl = "http://ff-pre.manager.bilibili.co/x/admin/live-proxy/unsafeproxy"
	}
	reqByte, _ := json.Marshal(req)
	apiProxyParams := &ApiProxyParams{
		Params:   string(reqByte),
		AppId:    "live.xroom",
		Module:   "1020001",
		GrpcPath: "/live.xroom.v1.Room/getMultipleByUids",
	}
	apiProxyParamsBody := GetFormParams(apiProxyParams)
	resp := HttpPost(apiUrl, apiProxyParamsBody)
	fmt.Println(resp)
	return
}
