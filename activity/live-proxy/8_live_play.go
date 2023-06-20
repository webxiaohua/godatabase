package main

import "encoding/json"

// web 左侧栏列表
type GetLeftEntryListReq struct {
	// 房间号
	RoomId int64 `protobuf:"varint,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty" form:"room_id" validate:"required,gt=0"`
	// 主播uid
	Ruid int64 `protobuf:"varint,2,opt,name=ruid,proto3" json:"ruid,omitempty" form:"ruid" validate:"required,gt=0"`
	// 二级分区id，若不存在二级分区id，则传0
	AreaId int64 `protobuf:"varint,3,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty" form:"area_id"`
	// 父分区id，若不存在父分区id，则传0
	ParentAreaId int64 `protobuf:"varint,4,opt,name=parent_area_id,json=parentAreaId,proto3" json:"parent_area_id,omitempty" form:"parent_area_id"`
	// 登陆用户ID, 0为未登录
	Uid int64 `protobuf:"varint,5,opt,name=uid,proto3" json:"uid" form:"uid"`
}

func GetLeftEntryList(req *GetLeftEntryListReq, isPre bool) (resp string) {
	apiUrl := "https://manager.bilibili.co/x/admin/live-proxy/unsafeproxy"
	if isPre {
		apiUrl = "http://ff-pre.manager.bilibili.co/x/admin/live-proxy/unsafeproxy"
	}
	reqByte, _ := json.Marshal(req)
	apiProxyParams := &ApiProxyParams{
		Params:   string(reqByte),
		AppId:    "live.liveplay",
		Module:   "1020001",
		GrpcPath: "/live.liveplay.WidgetService/GetLeftEntryList",
	}
	apiProxyParamsBody := GetFormParams(apiProxyParams)
	//fmt.Println(apiProxyParamsBody)
	resp = HttpPost(apiUrl, apiProxyParamsBody)
	return
}

// 挂件列表
type GetWidgetBannerListReq struct {
	// 房间号
	RoomId int64 `protobuf:"varint,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty" form:"room_id" validate:"required,gt=0"`
	// 主播uid
	Ruid int64 `protobuf:"varint,2,opt,name=ruid,proto3" json:"ruid,omitempty" form:"ruid" validate:"required,gt=0"`
	// 二级分区id，若不存在二级分区id，则传0
	AreaId int64 `protobuf:"varint,3,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty" form:"area_id"`
	// 父分区id，若不存在父分区id，则传0
	ParentAreaId int64 `protobuf:"varint,4,opt,name=parent_area_id,json=parentAreaId,proto3" json:"parent_area_id,omitempty" form:"parent_area_id"`
	// ios,android
	Platform string `protobuf:"bytes,5,opt,name=platform,proto3" json:"platform,omitempty" form:"platform" validate:"required,gt=0"`
	// 版本号
	Build int64 `protobuf:"varint,6,opt,name=build,proto3" json:"build,omitempty" form:"build" gt=0"`
	// 客户端类型 web不传这个参数
	MobiApp string `protobuf:"bytes,7,opt,name=mobi_app,json=mobiApp,proto3" json:"mobi_app" form:"mobi_app"`
	// 来源 （标识来源，房间页 ：live , 粉版播端：live_link , blink : blink , 直播姬：pc_link  ）
	Source string `protobuf:"bytes,8,opt,name=source,proto3" json:"source" form:"source"`
	// 特殊标记 0 表示挂件窗口判断 1 表示 获取挂件类容
	SpecialFlag int64 `protobuf:"varint,9,opt,name=special_flag,json=specialFlag,proto3" json:"special_flag" form:"special_flag"`
	// 登陆用户ID, 0为未登录
	Uid int64 `protobuf:"varint,10,opt,name=uid,proto3" json:"uid" form:"uid"`
	// 页面来源 1 挂件 2 半屏页
	PageSource int64 `protobuf:"varint,11,opt,name=page_source,json=pageSource,proto3" json:"page_source" form:"page_source"`
}

func GetWidgetBannerList(req *GetWidgetBannerListReq, isPre bool) (resp string) {
	apiUrl := "https://manager.bilibili.co/x/admin/live-proxy/unsafeproxy"
	if isPre {
		apiUrl = "http://ff-pre.manager.bilibili.co/x/admin/live-proxy/unsafeproxy"
	}
	reqByte, _ := json.Marshal(req)
	apiProxyParams := &ApiProxyParams{
		Params:   string(reqByte),
		AppId:    "live.liveplay",
		Module:   "1020001",
		GrpcPath: "/live.liveplay.WidgetService/GetWidgetBannerList",
	}
	apiProxyParamsBody := GetFormParams(apiProxyParams)
	//fmt.Println(apiProxyParamsBody)
	resp = HttpPost(apiUrl, apiProxyParamsBody)
	return
}

// 直播间是否展示挂件 & 左侧图标
type RoomShowWidgetBannerReq struct {
	// 房间号
	RoomId int64 `protobuf:"varint,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty" form:"room_id" validate:"required,gt=0"`
}

func RoomShowWidgetBanner(req *RoomShowWidgetBannerReq, isPre bool) (resp string) {
	apiUrl := "https://manager.bilibili.co/x/admin/live-proxy/unsafeproxy"
	if isPre {
		apiUrl = "http://ff-pre.manager.bilibili.co/x/admin/live-proxy/unsafeproxy"
	}
	reqByte, _ := json.Marshal(req)
	apiProxyParams := &ApiProxyParams{
		Params:   string(reqByte),
		AppId:    "live.liveplay",
		Module:   "1020001",
		GrpcPath: "/live.liveplay.WidgetService/RoomShowWidgetBanner",
	}
	apiProxyParamsBody := GetFormParams(apiProxyParams)
	//fmt.Println(apiProxyParamsBody)
	resp = HttpPost(apiUrl, apiProxyParamsBody)
	return
}
