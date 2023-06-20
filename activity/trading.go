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
	// todo 每次修改
	data:= [][]int64{

	}
	// todo 每次修改
	times:=1
	for idx,item := range data {
		uid := item[0]
		// todo 每次修改
		score := item[1]
		fmt.Printf("times:%d idx:%d uid:%d , score:%d \n",times,idx,uid,score)
		// todo 每次修改
		msgId := fmt.Sprintf("fix:20220401_20_%d:%d",times,idx)
		// todo 每次修改
		res := updateTradingScore(420011,score,uid,msgId)
		fmt.Println(res)
		fmt.Println("succeed:", uid)

	}
}

type UpdateUserScoreReq struct {
	// 数值id
	ScoreId int64 `json:"score_id"`
	// 增加或减少的数量
	Score int64 `json:"score"`
	// uid 用户uid
	Uid int64 `json:"uid"`
	// 有效日期 秒级时间戳
	Expire int64 `json:"expire"`
	// 事务id 用于幂等
	Tid string `json:"tid"`
	// 来源或用处
	Source string `json:"source"`
	// 平台
	Platform string `json:"platform"`
	//来自哪个app eg:blink,粉版
	MobiApp string `json:"mobi_app"`
	//版本号
	Build string `json:"build"`
	//设备
	Device string `json:"device"`
	//上下文类型，例如来自哪个页面
	ContextType string `json:"context_type"`
	//上下文id
	ContextId int64 `json:"context_id"`
	//用于保存上游额外数据信息
	ExtraData string `json:"extra_data"`
	//上游发放时间
	UpdateTime int64 `json:"update_time"`
	// 业务方id 用于对账 如不需要对账则填写0
	BizId int64 `json:"biz_id"`
	// 是否需要异步对账
	AsyncCheck bool `json:"async_check"`
	// 幂等重试无error 解决当前重复发放报错问题
	NoErrorRetry         bool     `json:"no_error_retry"`
}

// 增加 or 减少 数值
func updateTradingScore(scoreId,score,uid int64, tid string) (res string) {
	tradingReq := &UpdateUserScoreReq{
		ScoreId:     scoreId,
		Score:       score,
		Uid:         uid,
		Tid:         tid,
		Source:      "1100004",
		Platform:    "pc",
		MobiApp:     "internal",
		Build:       "internal",
		Device:      "internal",
		ContextType: "behavior",
		UpdateTime:  time.Now().Unix(),
	}
	jsonByte,_ := json.Marshal(tradingReq)
	apiUrl := "http://grpc-proxy.bilibili.co/live.trading/live.trading.Score/UpdateUserScore"
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
	res = string(body)
	return res
}