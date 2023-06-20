package main

import (
	"encoding/json"
	"fmt"
)

type UpdateUserItemReq struct {
	// 类型id
	ItemId int64 `protobuf:"varint,1,opt,name=item_id,json=itemId,proto3" json:"item_id" form:"item_id" validate:"required"`
	// 增加或减少的数量 正增负减少
	Num int64 `protobuf:"varint,2,opt,name=num,proto3" json:"num" form:"num" validate:"required"`
	// uid 用户uid
	Uid int64 `protobuf:"varint,3,opt,name=uid,proto3" json:"uid" form:"uid" validate:"required"`
	// 有效日期 秒级时间戳
	Expire             int64  `protobuf:"varint,4,opt,name=expire,proto3" json:"expire" form:"expire"`
	TargetUid          int64  `protobuf:"varint,5,opt,name=target_uid,json=targetUid,proto3" json:"target_uid" form:"target_uid"`
	TargetParentAreaId int64  `protobuf:"varint,6,opt,name=target_parent_area_id,json=targetParentAreaId,proto3" json:"target_parent_area_id" form:"target_parent_area_id"`
	TargetAreaId       int64  `protobuf:"varint,7,opt,name=target_area_id,json=targetAreaId,proto3" json:"target_area_id" form:"target_area_id"`
	TargetOther        string `protobuf:"bytes,8,opt,name=target_other,json=targetOther,proto3" json:"target_other" form:"target_other"`
	// 事务id 用于幂等
	Tid string `protobuf:"bytes,9,opt,name=tid,proto3" json:"tid" form:"tid" validate:"required"`
	// 来源或用处
	Source string `protobuf:"bytes,10,opt,name=source,proto3" json:"source" form:"source" validate:"required"`
}

// 更新活动数值
func UpdateUserItem(req *UpdateUserItemReq, isPre bool) {
	apiUrl := "https://manager.bilibili.co/x/admin/live-proxy/unsafeproxy"
	if isPre {
		apiUrl = "http://ff-pre.manager.bilibili.co/x/admin/live-proxy/unsafeproxy"
	}
	reqByte, _ := json.Marshal(req)
	apiProxyParams := &ApiProxyParams{
		Params:   string(reqByte),
		AppId:    "live.reward",
		Module:   "1020001",
		GrpcPath: "/live.reward.userItem/UpdateUserItem",
	}
	apiProxyParamsBody := GetFormParams(apiProxyParams)
	fmt.Println(apiProxyParamsBody)
	resp := HttpPost(apiUrl, apiProxyParamsBody)
	fmt.Println(resp)
}

// 查询用户数值
type GetUserItemReq struct {
	// 类型id
	ItemId int64 `protobuf:"varint,1,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty" form:"item_id" validate:"required"`
	// 用户uid
	Uid                int64  `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty" form:"uid" validate:"required"`
	TargetUid          int64  `protobuf:"varint,3,opt,name=target_uid,json=targetUid,proto3" json:"target_uid,omitempty" form:"target_uid"`
	TargetParentAreaId int64  `protobuf:"varint,4,opt,name=target_parent_area_id,json=targetParentAreaId,proto3" json:"target_parent_area_id,omitempty" form:"target_parent_area_id"`
	TargetAreaId       int64  `protobuf:"varint,5,opt,name=target_area_id,json=targetAreaId,proto3" json:"target_area_id,omitempty" form:"target_area_id"`
	TargetOther        string `protobuf:"bytes,6,opt,name=target_other,json=targetOther,proto3" json:"target_other,omitempty" form:"target_other"`
}
type GetUserItemResp_UserItemData_ItemDetail struct {
	// 累计获取积分
	Num int64 `protobuf:"varint,1,opt,name=num,proto3" json:"num"`
	// 有效期 秒级时间戳
	Expire int64 `protobuf:"varint,2,opt,name=expire,proto3" json:"expire"`
	// 用户item记录id
	Id int64 `protobuf:"varint,3,opt,name=id,proto3" json:"id"`
}
type GetUserItemResp_UserItemData struct {
	//用户拥有的总共的item数量
	Num int64 `protobuf:"varint,1,opt,name=num,proto3" json:"num"`
	// 列表
	List []*GetUserItemResp_UserItemData_ItemDetail `protobuf:"bytes,2,rep,name=list,proto3" json:"list"`
}
type GetUserItemResp struct {
	//状态 0失败 1成功
	Status int64 `protobuf:"varint,1,opt,name=status,proto3" json:"status"`
	//具体信息
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	//用户物品信息
	UserItemData *GetUserItemResp_UserItemData `protobuf:"bytes,3,opt,name=user_item_data,json=userItemData,proto3" json:"user_item_data,omitempty"`
}

func GetUserItem(req *GetUserItemReq, isPre bool) {
	apiUrl := "https://manager.bilibili.co/x/admin/live-proxy/unsafeproxy"
	if isPre {
		apiUrl = "http://ff-pre.manager.bilibili.co/x/admin/live-proxy/unsafeproxy"
	}
	reqByte, _ := json.Marshal(req)
	apiProxyParams := &ApiProxyParams{
		Params:   string(reqByte),
		AppId:    "live.reward",
		Module:   "1020001",
		GrpcPath: "/live.reward.userItem/GetUserItem",
	}
	apiProxyParamsBody := GetFormParams(apiProxyParams)
	fmt.Println(apiProxyParamsBody)
	resp := HttpPost(apiUrl, apiProxyParamsBody)
	fmt.Println(resp)
}
