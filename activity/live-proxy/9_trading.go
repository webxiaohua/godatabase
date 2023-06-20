package main

import (
	"encoding/json"
	"fmt"
)

// 获取数值记录
type GetLogByBizIdReq struct {
	// 数值id
	ScoreId int64 `protobuf:"varint,1,opt,name=score_id,json=scoreId,proto3" json:"score_id" form:"score_id"`
	// uid
	Uid int64 `protobuf:"varint,2,opt,name=uid,proto3" json:"uid" form:"uid"`
	// 业务Id类型 1:订单 2: 抽奖 3: 系统任务 4:退款 https://info.bilibili.co/pages/viewpage.action?pageId=498625680
	BusinessType int64 `protobuf:"varint,3,opt,name=business_type,json=businessType,proto3" json:"business_type" form:"business_type"`
	// 对应的id
	BusinessId string `protobuf:"bytes,4,opt,name=business_id,json=businessId,proto3" json:"business_id" form:"business_id"`
}
type LogList struct {
	// 数值id scoreid可能是string
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	// 类型 1内部数值 2外部数值 3价值数值
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type"`
	// 名称
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name"`
	// 时间
	UpdateTime int64 `protobuf:"varint,4,opt,name=update_time,json=updateTime,proto3" json:"update_time"`
	// 类型 1增加 2减少 3 过期 4 退款
	UpdateType int64 `protobuf:"varint,5,opt,name=update_type,json=updateType,proto3" json:"update_type"`
	// 数量
	Num int64 `protobuf:"varint,6,opt,name=num,proto3" json:"num"`
	// 变更前数值(参考值)
	BeforeScore int64 `protobuf:"varint,7,opt,name=before_score,json=beforeScore,proto3" json:"before_score"`
	// 来源或用处
	Source string `protobuf:"bytes,8,opt,name=source,proto3" json:"source"`
	//发放扣除时 上游传递的保存信息 (可存储发放理由,或json字符串)
	ExtraData string `protobuf:"bytes,9,opt,name=extra_data,json=extraData,proto3" json:"extra_data"`
	// 通用 数值id 扩展为string
	CommonId string `protobuf:"bytes,10,opt,name=common_id,json=commonId,proto3" json:"common_id"`
	// 数据库中的msgid
	MsgId string `protobuf:"bytes,11,opt,name=msg_id,json=msgId,proto3" json:"msg_id"`
	// 数据创建时间 秒
	Ctime int64 `protobuf:"varint,12,opt,name=ctime,proto3" json:"ctime"`
}
type GetLogByBizIdResp struct {
	LogList []*LogList `protobuf:"bytes,1,rep,name=logList,proto3" json:"logList,omitempty"`
}

func GetLogByBizId(req *GetLogByBizIdReq, isPre bool) {
	apiUrl := "https://manager.bilibili.co/x/admin/live-proxy/unsafeproxy"
	if isPre {
		apiUrl = "http://ff-pre.manager.bilibili.co/x/admin/live-proxy/unsafeproxy"
	}
	reqByte, _ := json.Marshal(req)
	apiProxyParams := &ApiProxyParams{
		Params:   string(reqByte),
		AppId:    "live.trading",
		Module:   "1020001",
		GrpcPath: "/live.trading.Score/GetLogByBizId",
	}
	apiProxyParamsBody := GetFormParams(apiProxyParams)
	resp := HttpPost(apiUrl, apiProxyParamsBody)
	fmt.Println(resp)
	return
}
