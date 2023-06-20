package main

import (
	"encoding/json"
)

// 通用web端道具左侧栏入口接口
type WebLeftEntryReq struct {
	ActIds       []int64 `protobuf:"varint,1,rep,packed,name=act_ids,json=actIds,proto3" json:"act_ids"`
	RoomId       int64   `protobuf:"varint,2,opt,name=room_id,json=roomId,proto3" json:"room_id"`
	Ruid         int64   `protobuf:"varint,3,opt,name=ruid,proto3" json:"ruid"`
	AreaId       int64   `protobuf:"varint,4,opt,name=area_id,json=areaId,proto3" json:"area_id"`
	ParentAreaId int64   `protobuf:"varint,5,opt,name=parent_area_id,json=parentAreaId,proto3" json:"parent_area_id"`
	Uid          int64   `protobuf:"varint,6,opt,name=uid,proto3" json:"uid"`
}

type WebLeftEntryResp_Record struct {
	IsExist   bool   `protobuf:"varint,1,opt,name=is_exist,json=isExist,proto3" json:"is_exist"`
	Icon      string `protobuf:"bytes,2,opt,name=icon,proto3" json:"icon"`
	Title     string `protobuf:"bytes,3,opt,name=title,proto3" json:"title"`
	RedDot    bool   `protobuf:"varint,4,opt,name=red_dot,json=redDot,proto3" json:"red_dot"`
	RedDotTag string `protobuf:"bytes,5,opt,name=red_dot_tag,json=redDotTag,proto3" json:"red_dot_tag"`
	ActId     int64  `protobuf:"varint,6,opt,name=act_id,json=actId,proto3" json:"act_id"`
	JumpUrl   string `protobuf:"bytes,7,opt,name=jump_url,json=jumpUrl,proto3" json:"jump_url"`
}

type WebLeftEntryResp struct {
	List map[int64]*WebLeftEntryResp_Record `protobuf:"bytes,1,rep,name=list,proto3" json:"list" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

// web左侧栏入口
func WebLeftEntry(req *WebLeftEntryReq, isPre bool) (resp string) {
	apiUrl := "https://manager.bilibili.co/x/admin/live-proxy/unsafeproxy"
	if isPre {
		apiUrl = "http://ff-pre.manager.bilibili.co/x/admin/live-proxy/unsafeproxy"
	}
	reqByte, _ := json.Marshal(req)
	apiProxyParams := &ApiProxyParams{
		Params:   string(reqByte),
		AppId:    "live.xactivity",
		Module:   "1020001",
		GrpcPath: "/live.xactivity.v1.generalRankFrontend/WebLeftEntry",
	}
	apiProxyParamsBody := GetFormParams(apiProxyParams)
	//fmt.Println(apiProxyParamsBody)
	resp = HttpPost(apiUrl, apiProxyParamsBody)
	return
}

// 注入送分池
type AddPoolBalanceReq struct {
	MsgId string `protobuf:"bytes,1,opt,name=msg_id,json=msgId,proto3" json:"msg_id,omitempty" form:"msg_id"`
	Cnt   int64  `protobuf:"varint,2,opt,name=cnt,proto3" json:"cnt,omitempty" form:"cnt"`
}

func AddPoolBalance(req *AddPoolBalanceReq, isPre bool) (resp string) {
	apiUrl := "https://manager.bilibili.co/x/admin/live-proxy/unsafeproxy"
	if isPre {
		apiUrl = "http://ff-pre.manager.bilibili.co/x/admin/live-proxy/unsafeproxy"
	}
	reqByte, _ := json.Marshal(req)
	apiProxyParams := &ApiProxyParams{
		Params:   string(reqByte),
		AppId:    "live.xactivity",
		Module:   "1020001",
		GrpcPath: "/live.xactivity.v1.specApi/AddPoolBalance",
	}
	apiProxyParamsBody := GetFormParams(apiProxyParams)
	//fmt.Println(apiProxyParamsBody)
	resp = HttpPost(apiUrl, apiProxyParamsBody)
	return
}

// 挂件
type GetMultipleWidgetBannerReq struct {
	// 活动ids
	ActIds []int64 `protobuf:"varint,1,rep,packed,name=act_ids,json=actIds,proto3" json:"act_ids,omitempty" form:"act_ids`
	// 房间id
	RoomId int64 `protobuf:"varint,2,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty" form:"room_id" validate:"required,gt=0"`
	// 主播uid
	Ruid int64 `protobuf:"varint,3,opt,name=ruid,proto3" json:"ruid,omitempty" form:"ruid" validate:"required,gt=0"`
	// 投放终端 1 粉版直播间 2 blink 3 粉版播端 4 web 5 pc_link
	Source int64 `protobuf:"varint,4,opt,name=source,proto3" json:"source,omitempty" form:"source"`
	// uid 登录态提供，非登录态传0
	Uid int64 `protobuf:"varint,5,opt,name=uid,proto3" json:"uid,omitempty" form:"uid"`
	// 挂件类型map
	WidgetTypeMap map[int64]int64 `protobuf:"bytes,6,rep,name=widget_type_map,json=widgetTypeMap,proto3" json:"widget_type_map,omitempty" form:"widget_type_map" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

// 挂件列表
func GetMultipleWidgetBanner(req *GetMultipleWidgetBannerReq, isPre bool) (resp string) {
	apiUrl := "https://manager.bilibili.co/x/admin/live-proxy/unsafeproxy"
	if isPre {
		apiUrl = "http://ff-pre.manager.bilibili.co/x/admin/live-proxy/unsafeproxy"
	}
	reqByte, _ := json.Marshal(req)
	apiProxyParams := &ApiProxyParams{
		Params:   string(reqByte),
		AppId:    "live.xactivity",
		Module:   "1020001",
		GrpcPath: "/live.xactivity.rank_pendant/getMultipleWidgetBanner",
	}
	apiProxyParamsBody := GetFormParams(apiProxyParams)
	//fmt.Println(apiProxyParamsBody)
	resp = HttpPost(apiUrl, apiProxyParamsBody)
	return
}

// 半屏初始化
type ActivityHalfInitReq struct {
	// 活动id
	ActId int64 `protobuf:"varint,1,opt,name=act_id,json=actId,proto3" json:"act_id,omitempty" form:"act_id" validate:"required,gt=0"`
	// 房间id
	RoomId int64 `protobuf:"varint,2,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty" form:"room_id" validate:"required,gt=0"`
	//  登录用户id
	Uid int64 `protobuf:"varint,3,opt,name=uid,proto3" json:"uid,omitempty" form:"uid"`
	// 主播ruid
	Ruid int64 `protobuf:"varint,4,opt,name=ruid,proto3" json:"ruid,omitempty" form:"ruid"`
}

func ActivityHalfInit(req *ActivityHalfInitReq, isPre bool) (resp string) {
	apiUrl := "https://manager.bilibili.co/x/admin/live-proxy/unsafeproxy"
	if isPre {
		apiUrl = "http://ff-pre.manager.bilibili.co/x/admin/live-proxy/unsafeproxy"
	}
	reqByte, _ := json.Marshal(req)
	apiProxyParams := &ApiProxyParams{
		Params:   string(reqByte),
		AppId:    "live.xactivity",
		Module:   "1020001",
		GrpcPath: "/live.xactivity.v1.generalRankFrontend/ActivityHalfInit",
	}
	apiProxyParamsBody := GetFormParams(apiProxyParams)
	//fmt.Println(apiProxyParamsBody)
	resp = HttpPost(apiUrl, apiProxyParamsBody)
	return
}

// JobApi
type JobApiReq struct {
	// 活动id
	ActId int64 `protobuf:"varint,1,opt,name=act_id,json=actId,proto3" json:"act_id,omitempty" form:"act_id" validate:"required,gt=0"`
	// 脚本名 用于表示具体调用哪个脚本
	ActionName string `protobuf:"bytes,2,opt,name=action_name,json=actionName,proto3" json:"action_name"`
	// 执行时间 yyyy-MM-dd HH:mm:ss
	ExecTime string `protobuf:"bytes,3,opt,name=exec_time,json=execTime,proto3" json:"exec_time"`
}

func JobApi(req *JobApiReq, isPre bool) (resp string) {
	apiUrl := "https://manager.bilibili.co/x/admin/live-proxy/unsafeproxy"
	if isPre {
		apiUrl = "http://ff-pre.manager.bilibili.co/x/admin/live-proxy/unsafeproxy"
	}
	reqByte, _ := json.Marshal(req)
	apiProxyParams := &ApiProxyParams{
		Params:   string(reqByte),
		AppId:    "live.xactivity",
		Module:   "1020001",
		GrpcPath: "/live.xactivity.v1.generalRankFrontend/JobApi",
	}
	apiProxyParamsBody := GetFormParams(apiProxyParams)
	//fmt.Println(apiProxyParamsBody)
	resp = HttpPost(apiUrl, apiProxyParamsBody)
	return
}
