/**
 * @Desc: 榜单
 */

package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Dimensions struct {
	// 主播id
	Ruid int64 `protobuf:"varint,1,opt,name=ruid,proto3" json:"ruid"`
	// 用户id
	Uid int64 `protobuf:"varint,2,opt,name=uid,proto3" json:"uid"`
	// 房间id
	RoomId int64 `protobuf:"varint,3,opt,name=room_id,json=roomId,proto3" json:"room_id"`
	// 一级分区id
	ParentAreaId int64 `protobuf:"varint,4,opt,name=parent_area_id,json=parentAreaId,proto3" json:"parent_area_id"`
	// 二级分区id
	AreaId int64 `protobuf:"varint,5,opt,name=area_id,json=areaId,proto3" json:"area_id"`
	// 道具id
	GiftId int64 `protobuf:"varint,6,opt,name=gift_id,json=giftId,proto3" json:"gift_id"`
	// 时间戳
	Timestamp int64 `protobuf:"varint,7,opt,name=timestamp,proto3" json:"timestamp"`
	// 大乱斗赛季id
	SeasonId int64 `protobuf:"varint,8,opt,name=season_id,json=seasonId,proto3" json:"season_id"`
	// 大乱斗pk id
	PkId int64 `protobuf:"varint,9,opt,name=pk_id,json=pkId,proto3" json:"pk_id"`
	// 扩展1
	Other string `protobuf:"bytes,10,opt,name=other,proto3" json:"other"`
}

type GeneralRankGetResp struct {
	// 榜单项id
	ItemId int64 `json:"item_id"` // int
	// 积分
	Score int64 `json:"score"` // int
	// 排名
	Rank int64 `json:"rank"` // int
	// 是否是灰度数据
	IsGray int64 `json:"is_gray"` // int
	// 是否在动态黑名单
	IsIgnore int64 `json:"is_ignore"` // int
	// 执行状态  0成功  大于0 失败  【10：未获取到榜单配置】
	Status int64 `json:"status"` // int
	// 提示信息 方便定位问题
	Desc string `json:"desc"` // string
}
type GeneralRankDelReq struct {
	// source 来源，由榜单系统提供
	Source int64 `json:"source"`
	// 幂等ID,同一个source下的msg_id由业务方保证不重复   redis24小时幂等
	MsgId string `json:"msg_id"`
	// 榜单ID 如果为0则自动扣分，否则为指定榜单扣分
	RankId int64 `json:"rank_id"`
	// 成员id
	ItemId int64 `json:"item_id"`
	// 维度信息值
	Dimensions *Dimensions `json:"dimensions"`
}
type GeneralRankDelResp struct {
	// 状态 0成功（包含幂等） 1失败
	Status int64 `json:"status"`
}

// 增加榜单积分
func GeneralRankIncr(rankId, itemId, score, subScore, ruid, uid, timestamp int64, msgId string, dimensionOther string) {
	type GeneralRankIncrReq struct {
		// source 来源，由榜单系统提供
		Source int64 `json:"source"`
		// 幂等ID,同一个source下的msg_id由业务方保证不重复   redis24小时幂等
		MsgId string `json:"msg_id"`
		// 榜单ID 如果为0则自动加分，否则为指定榜单加分
		RankId int64 `json:"rank_id"`
		// 加分值（自动记分传0）
		Score int64 `json:"score"`
		// 二级积分值（自动记分传0）
		SubScore int64 `json:"sub_score"`
		// 成员id
		ItemId int64 `json:"item_id"`
		// 是否是动态黑名单 (更新item信息和db,但是不放入榜单列表中)
		IsIgnore int64 `json:"is_ignore"`
		// 维度信息值
		Dimensions *Dimensions `json:"dimensions"`
	}
	apiUrl := "https://manager.bilibili.co/x/admin/live-proxy/unsafeproxy"
	rankReq := &GeneralRankIncrReq{
		Source:   1000,
		MsgId:    msgId,
		RankId:   rankId,
		Score:    score,
		SubScore: subScore,
		ItemId:   itemId,
		IsIgnore: 0,
		Dimensions: &Dimensions{
			Ruid:         ruid,
			Uid:          uid,
			RoomId:       0,
			ParentAreaId: 0,
			AreaId:       0,
			GiftId:       0,
			Timestamp:    timestamp,
			SeasonId:     0,
			PkId:         0,
			Other:        dimensionOther,
		},
	}
	rankReqByte, _ := json.Marshal(rankReq)
	apiProxyParams := &ApiProxyParams{
		Params:   string(rankReqByte),
		AppId:    "live.live.rankdb",
		Module:   "1020001",
		GrpcPath: "/live.rankdb.v2.generalRankFrontend/GeneralRankIncr",
	}
	apiProxyParamsBody := GetFormParams(apiProxyParams)
	fmt.Println(apiProxyParamsBody)
	HttpPost(apiUrl, apiProxyParamsBody)
}

// 榜单移除
func GeneralRankDel(req *GeneralRankDelReq) (resp *GeneralRankDelResp) {
	return
}

// 查询榜单列表
func GeneralRankTop(rankId, start, end int64, dimensions *Dimensions) {
	type GeneralRankTopReq struct {
		// 榜单ID 必须传
		RankId int64 `protobuf:"varint,1,opt,name=rank_id,json=rankId,proto3" json:"rank_id"`
		// 维度信息值
		Dimensions *Dimensions `protobuf:"bytes,2,opt,name=dimensions,proto3" json:"dimensions"`
		// 榜单开始排名 从0开始
		Start int64 `protobuf:"varint,3,opt,name=start,proto3" json:"start"`
		// 榜单结束排名 按0开始计名次
		End int64 `protobuf:"varint,4,opt,name=end,proto3" json:"end"`
		// 屏蔽 ture、不屏蔽 false  屏蔽后会将错误信息放到 Resp
		MaskError bool `protobuf:"varint,5,opt,name=mask_error,json=maskError,proto3" json:"mask_error"`
	}
	apiUrl := "https://manager.bilibili.co/x/admin/live-proxy/unsafeproxy"
	params := &GeneralRankTopReq{
		RankId:     rankId,
		Dimensions: dimensions,
		Start:      start,
		End:        end,
		MaskError:  false,
	}
	jsonByte, _ := json.Marshal(params)
	apiProxyParams := &ApiProxyParams{
		Params:   string(jsonByte),
		AppId:    "live.live.rankdb",
		Module:   "1020001",
		GrpcPath: "/live.rankdb.v2.generalRankFrontend/GeneralRankTop",
	}
	apiProxyParamsBody := GetFormParams(apiProxyParams)
	fmt.Println(apiProxyParamsBody)
	HttpPost(apiUrl, apiProxyParamsBody)
}

// 查询单个榜单积分
func GeneralRankGet(rankId, itemId int64, dimensions *Dimensions) (resp *GeneralRankGetResp) {
	type GeneralRankGetReq struct {
		// 榜单ID 必须传
		RankId int64 `protobuf:"varint,1,opt,name=rank_id,json=rankId,proto3" json:"rank_id"`
		// 榜单对象ID
		ItemId int64 `protobuf:"varint,2,opt,name=item_id,json=itemId,proto3" json:"item_id"`
		// 维度信息值
		Dimensions *Dimensions `protobuf:"bytes,3,opt,name=dimensions,proto3" json:"dimensions"`
		// 屏蔽 ture、不屏蔽 false  屏蔽后会将错误信息放到 Resp
		MaskError bool `protobuf:"varint,4,opt,name=mask_error,json=maskError,proto3" json:"mask_error"`
	}
	type GeneralRankGetRespDataTmp struct {
		// 榜单项id
		ItemId string `json:"item_id"` // int
		// 积分
		Score string `json:"score"` // int
		// 排名
		Rank string `json:"rank"` // int
		// 是否是灰度数据
		IsGray string `json:"is_gray"` // int
		// 是否在动态黑名单
		IsIgnore string `json:"is_ignore"` // int
		// 执行状态  0成功  大于0 失败  【10：未获取到榜单配置】
		Status string `json:"status"` // int
		// 提示信息 方便定位问题
		Desc string `json:"desc"` // string
	}
	type GeneralRankGetRespTmp struct {
		Data *GeneralRankGetRespDataTmp
	}
	apiUrl := "https://manager.bilibili.co/x/admin/live-proxy/unsafeproxy"
	rankReq := &GeneralRankGetReq{
		RankId:     rankId,
		ItemId:     itemId,
		Dimensions: dimensions,
		MaskError:  false,
	}
	rankReqByte, _ := json.Marshal(rankReq)
	apiProxyParams := &ApiProxyParams{
		Params:   string(rankReqByte),
		AppId:    "live.live.rankdb",
		Module:   "1020001",
		GrpcPath: "/live.rankdb.v2.generalRankFrontend/GeneralRankGet",
	}
	apiProxyParamsBody := GetFormParams(apiProxyParams)
	fmt.Println(apiProxyParamsBody)
	respBody := HttpPost(apiUrl, apiProxyParamsBody)
	fmt.Println(respBody)
	var tmpResp *GeneralRankGetRespTmp
	json.Unmarshal([]byte(respBody), &tmpResp)
	score, _ := strconv.ParseInt(tmpResp.Data.Score, 10, 64)
	rank, _ := strconv.ParseInt(tmpResp.Data.Rank, 10, 64)
	resp = &GeneralRankGetResp{
		ItemId: itemId,
		Score:  score,
		Rank:   rank,
	}
	return
}

// web & 看端查高能榜接口
type GetHighEnergyRankReq struct {
	Uid      int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid"`
	Ruid     int64 `protobuf:"varint,2,opt,name=ruid,proto3" json:"ruid"`
	RoomId   int64 `protobuf:"varint,3,opt,name=roomId,proto3" json:"roomId"`
	PageSize int64 `protobuf:"varint,4,opt,name=pageSize,proto3" json:"pageSize"`
	Page     int64 `protobuf:"varint,5,opt,name=page,proto3" json:"page"`
	//这个优先级大于 pageSize 不要轻易用，
	ForcePageSize int64 `protobuf:"varint,6,opt,name=force_pageSize,json=forcePageSize,proto3" json:"forcePageSize"`
}

func GetHighEnergyRank(req *GetHighEnergyRankReq, isPre bool) {
	apiUrl := "https://manager.bilibili.co/x/admin/live-proxy/unsafeproxy"
	if isPre {
		apiUrl = "http://ff-pre.manager.bilibili.co/x/admin/live-proxy/unsafeproxy"
	}
	reqByte, _ := json.Marshal(req)
	apiProxyParams := &ApiProxyParams{
		Params:   string(reqByte),
		AppId:    "live.live.rankdb",
		Module:   "1020001",
		GrpcPath: "/live.rankdb.v1.OnlineGoldRank/GetHighEnergyRank",
	}
	apiProxyParamsBody := GetFormParams(apiProxyParams)
	resp := HttpPost(apiUrl, apiProxyParamsBody)
	fmt.Println(resp)
	return
}

// 播端 查在线榜接口
type UserGetOnlineRankReq struct {
	Uid      int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid"`
	Ruid     int64 `protobuf:"varint,2,opt,name=ruid,proto3" json:"ruid"`
	RoomId   int64 `protobuf:"varint,3,opt,name=roomId,proto3" json:"roomId"`
	Page     int64 `protobuf:"varint,4,opt,name=page,proto3" json:"page"`
	PageSize int64 `protobuf:"varint,5,opt,name=pageSize,proto3" json:"pageSize"`
}

func UserGetOnlineRank(req *UserGetOnlineRankReq, isPre bool) {
	apiUrl := "https://manager.bilibili.co/x/admin/live-proxy/unsafeproxy"
	if isPre {
		apiUrl = "http://ff-pre.manager.bilibili.co/x/admin/live-proxy/unsafeproxy"
	}
	reqByte, _ := json.Marshal(req)
	apiProxyParams := &ApiProxyParams{
		Params:   string(reqByte),
		AppId:    "live.live.rankdb",
		Module:   "1020001",
		GrpcPath: "/live.rankdb.v1.OnlineRank/UserGetOnlineRank",
	}
	apiProxyParamsBody := GetFormParams(apiProxyParams)
	resp := HttpPost(apiUrl, apiProxyParamsBody)
	fmt.Println(resp)
	return
}
