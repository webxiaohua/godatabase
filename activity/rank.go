package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main()  {
	//addRankScore(508,25237642,1469,0,25237642,25237642,time.Now().Unix(),"fix:20220619010001","1")
	//addRankScore(508,470748582,1961,0,470748582,470748582,time.Now().Unix(),"fix:20220619010002","35")
	/*params := [][]int64{
		{2045003,1,200},
	}
	for _,item := range params{
		itemId := item[0]
		score := item[2]
		msgId:=fmt.Sprintf("manual_:%d:%d",itemId,score)
		addRankScore(370,itemId,score,0,0,0,0,msgId,"")
	}*/
}

// 增加榜单积分
func addRankScore(rankId,itemId,score,subScore,ruid,uid,timestamp int64,msgId string,dimensionOther string)  {
	// 通用榜单加分接口
	type Dimensions struct {
		// 主播id
		Ruid int64 `json:"ruid"`
		// 用户id
		Uid int64 `json:"uid"`
		// 房间id
		RoomId int64 `json:"room_id"`
		// 一级分区id
		ParentAreaId int64 `json:"parent_area_id"`
		// 二级分区id
		AreaId int64 `json:"area_id"`
		// 道具id
		GiftId int64 `json:"gift_id"`
		// 时间戳
		Timestamp int64 `json:"timestamp"`
		// 大乱斗赛季id
		SeasonId int64 `json:"season_id"`
		// 大乱斗pk id
		PkId int64 `json:"pk_id"`
		// 扩展1
		Other                string   `json:"other"`
	}
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
		Dimensions           *Dimensions `json:"dimensions"`
	}

	apiUrl := "http://grpc-proxy.bilibili.co/live.live.rankdb/live.rankdb.v2.generalRankFrontend/GeneralRankIncr"
	signUpReq := &GeneralRankIncrReq{
		Source:     1000,
		MsgId:      msgId,
		RankId:     rankId,
		Score:      score,
		SubScore:   subScore,
		ItemId:     itemId,
		IsIgnore:   0,
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
	jsonByte,_ := json.Marshal(signUpReq)
	req,_ := http.NewRequest("POST",apiUrl,bytes.NewBuffer(jsonByte))
	req.Header.Set("Content-Type","application/json")
	client := &http.Client{}
	resp,err := client.Do(req)
	if err != nil {
		fmt.Println("error: %+v",err)
		return
	}
	defer resp.Body.Close()
	fmt.Sprintf("status:%d\n",resp.Status)
	fmt.Sprintf("response:%+v\n",resp.Header)
	body,_ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

// 设置榜单积分
func setRankScore(rankId,itemId,score,subScore,ruid,uid,timestamp int64, msgId string)  {
	// 通用榜单加分接口
	type Dimensions struct {
		// 主播id
		Ruid int64 `json:"ruid"`
		// 用户id
		Uid int64 `json:"uid"`
		// 房间id
		RoomId int64 `json:"room_id"`
		// 一级分区id
		ParentAreaId int64 `json:"parent_area_id"`
		// 二级分区id
		AreaId int64 `json:"area_id"`
		// 道具id
		GiftId int64 `json:"gift_id"`
		// 时间戳
		Timestamp int64 `json:"timestamp"`
		// 大乱斗赛季id
		SeasonId int64 `json:"season_id"`
		// 大乱斗pk id
		PkId int64 `json:"pk_id"`
		// 扩展1
		Other                string   `json:"other"`
	}
	type GeneralRankSetReq struct {
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
		Dimensions           *Dimensions `json:"dimensions"`
	}

	apiUrl := "http://grpc-proxy.bilibili.co/live.live.rankdb/live.rankdb.v2.generalRankFrontend/GeneralRankSet"
	signUpReq := &GeneralRankSetReq{
		Source:     1000,
		MsgId:      msgId,
		RankId:     rankId,
		Score:      score,
		SubScore:   subScore,
		ItemId:     itemId,
		IsIgnore:   0,
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
			Other:        "",
		},
	}
	jsonByte,_ := json.Marshal(signUpReq)
	req,_ := http.NewRequest("POST",apiUrl,bytes.NewBuffer(jsonByte))
	req.Header.Set("Content-Type","application/json")
	client := &http.Client{}
	resp,err := client.Do(req)
	if err != nil {
		fmt.Println("error: %+v",err)
		return
	}
	defer resp.Body.Close()
	fmt.Sprintf("status:%d\n",resp.Status)
	fmt.Sprintf("response:%+v\n",resp.Header)
	body,_ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

// 删除榜单积分
func delRankScore(rankId,itemId int64,msgId string)  {
	// 通用榜单加分接口
	type Dimensions struct {
		// 主播id
		Ruid int64 `json:"ruid"`
		// 用户id
		Uid int64 `json:"uid"`
		// 房间id
		RoomId int64 `json:"room_id"`
		// 一级分区id
		ParentAreaId int64 `json:"parent_area_id"`
		// 二级分区id
		AreaId int64 `json:"area_id"`
		// 道具id
		GiftId int64 `json:"gift_id"`
		// 时间戳
		Timestamp int64 `json:"timestamp"`
		// 大乱斗赛季id
		SeasonId int64 `json:"season_id"`
		// 大乱斗pk id
		PkId int64 `json:"pk_id"`
		// 扩展1
		Other                string   `json:"other"`
	}
	type GeneralRankDelReq struct {
		// source 来源，由榜单系统提供
		Source int64 `json:"source"`
		// 幂等ID,同一个source下的msg_id由业务方保证不重复   redis24小时幂等
		MsgId string `json:"msg_id"`
		// 榜单ID 如果为0则自动加分，否则为指定榜单加分
		RankId int64 `json:"rank_id"`
		// 成员id
		ItemId int64 `json:"item_id"`
		// 维度信息值
		Dimensions           *Dimensions `json:"dimensions"`
	}

	apiUrl := "http://grpc-proxy.bilibili.co/live.live.rankdb/live.rankdb.v2.generalRankFrontend/GeneralRankDel"
	signUpReq := &GeneralRankDelReq{
		Source:     1000,
		MsgId:      msgId,
		RankId:     rankId,
		ItemId:     itemId,
		Dimensions: &Dimensions{
			Ruid:         0,
			Uid:          0,
			RoomId:       0,
			ParentAreaId: 0,
			AreaId:       0,
			GiftId:       0,
			Timestamp:    0,
			SeasonId:     0,
			PkId:         0,
			Other:        "",
		},
	}
	jsonByte,_ := json.Marshal(signUpReq)
	req,_ := http.NewRequest("POST",apiUrl,bytes.NewBuffer(jsonByte))
	req.Header.Set("Content-Type","application/json")
	client := &http.Client{}
	resp,err := client.Do(req)
	if err != nil {
		fmt.Println("error: %+v",err)
		return
	}
	defer resp.Body.Close()
	fmt.Sprintf("status:%d\n",resp.Status)
	fmt.Sprintf("response:%+v\n",resp.Header)
	body,_ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

// 榜单自动加分
func generalRankAutoIncr(){
	type GeneralRankAutoIncrReq struct {
		// Uid
		Uid int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid"`
		// ruid
		Ruid int64 `protobuf:"varint,2,opt,name=ruid,proto3" json:"ruid"`
		// 房间号
		RoomId int64 `protobuf:"varint,3,opt,name=room_id,json=roomId,proto3" json:"room_id"`
		// 父分区id
		ParentAreaId int64 `protobuf:"varint,4,opt,name=parent_area_id,json=parentAreaId,proto3" json:"parent_area_id"`
		// 分区id
		AreaId int64 `protobuf:"varint,5,opt,name=area_id,json=areaId,proto3" json:"area_id"`
		// 平台
		Platform string `protobuf:"bytes,6,opt,name=platform,proto3" json:"platform"`
		// app类型
		MobileApp string `protobuf:"bytes,7,opt,name=mobile_app,json=mobileApp,proto3" json:"mobile_app"`
		// 礼物类型 1付费道具 2大航海 3付费留言 4免费道具
		GoodsType int64 `protobuf:"varint,8,opt,name=goods_type,json=goodsType,proto3" json:"goods_type"`
		// 商品id
		GoodsId int64 `protobuf:"varint,9,opt,name=goods_id,json=goodsId,proto3" json:"goods_id"`
		// 币种 1金瓜子 2银瓜子
		CoinType int64 `protobuf:"varint,10,opt,name=coin_type,json=coinType,proto3" json:"coin_type"`
		// 礼物数量 | 大航海开通月数
		GoodsNum int64 `protobuf:"varint,11,opt,name=goods_num,json=goodsNum,proto3" json:"goods_num"`
		// 实际支付货币数量
		PayCoin int64 `protobuf:"varint,12,opt,name=pay_coin,json=payCoin,proto3" json:"pay_coin"`
		// 货币应支付数量
		TotalCoin int64 `protobuf:"varint,13,opt,name=total_coin,json=totalCoin,proto3" json:"total_coin"`
		// 发生时间戳
		Timestamp int64 `protobuf:"varint,14,opt,name=timestamp,proto3" json:"timestamp"`
		// 幂等id
		MsgId string `protobuf:"bytes,15,opt,name=msg_id,json=msgId,proto3" json:"msg_id"`
		// 请求来源 应用名
		Source int64 `protobuf:"varint,16,opt,name=source,proto3" json:"source"`
		// 需要更新的榜单ids,不传默认匹配当前生效的所有榜单(主要用于上游重试,或活动更新指定榜单) 逗号分隔,
		RankIds string `protobuf:"bytes,17,opt,name=rank_ids,json=rankIds,proto3" json:"rank_ids"`
		// 道具 id
		GiftId               int64    `protobuf:"varint,18,opt,name=gift_id,json=giftId,proto3" json:"gift_id"`
	}
	apiUrl := "http://grpc-proxy.bilibili.co/live.live.rankdb/live.rankdb.v2.generalRankFrontend/GeneralRankAutoIncr"
	signUpReq := &GeneralRankAutoIncrReq{
		Uid:          28007796,
		Ruid:         28007843,
		RoomId:       460710,
		ParentAreaId: 1,
		AreaId:       145,
		Platform:     "android",
		MobileApp:    "android",
		GoodsType:    1,
		GoodsId:      15,
		GiftId:       20004,
		CoinType:     1,
		GoodsNum:     1,
		PayCoin:      110,
		TotalCoin:    110,
		Timestamp:    time.Now().Unix(),
		MsgId:        "pre_test:20220418140011",
		Source:       1,
		RankIds:      "1",
	}
	jsonByte,_ := json.Marshal(signUpReq)
	req,_ := http.NewRequest("POST",apiUrl,bytes.NewBuffer(jsonByte))
	req.Header.Set("Content-Type","application/json")
	client := &http.Client{}
	resp,err := client.Do(req)
	if err != nil {
		fmt.Println("error: %+v",err)
		return
	}
	defer resp.Body.Close()
	fmt.Sprintf("status:%d\n",resp.Status)
	fmt.Sprintf("response:%+v\n",resp.Header)
	body,_ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}