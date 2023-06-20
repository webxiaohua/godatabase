package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	// todo
	//for _,uid := range uids {
	//buildSignUp(uid,987,1647878400,1648137600,10,"{\"room_id\": 0, \"task_id\": 10}")
	//updateSignInfo(uid,987,6,1647878400,1648137600,"{\"room_id\": 0, \"task_id\": 6}")
	//}
	//buildSignUp(28008860,710,1648137600,1648396799,2,"")
	// {"room_id": 0, "task_id": 429}
	//var teamId int64
	// todo
	//buildSignUp(28008918, 987, 1661270400, 1661356800, 2, "")

	getEffectActList()
}

// 新增报名，执行时需要绑定host: [172.18.21.10 grpc-proxy.bilibili.co]
func buildSignUp(uid, signId, startTime, endTime, teamId int64, otherMessage string) {
	type SignUpReq struct {
		// 用户uid
		Uid int64 `json:"uid"`
		// 报名id
		SignId int64 `json:"sign_id"`
		// 报名类型 1:特征报名2:录入
		SignType int64 `json:"sign_type"`
		// 报名开始时间
		StartTime int64 `json:"start_time"`
		// 报名结束时间
		EndTime int64 `json:"end_time"`
		// 额外信息
		OtherMessage string `json:"other_message"`
		// 赛道
		TeamId int64 `json:"team_id"`
	}
	apiUrl := "http://grpc-proxy.bilibili.co/live.mozi/live.mozi.signup/SignUp"
	signUpReq := &SignUpReq{
		Uid:          uid,
		SignId:       signId,
		SignType:     2,
		StartTime:    startTime,
		EndTime:      endTime,
		OtherMessage: otherMessage,
		TeamId:       teamId,
	}
	jsonByte, _ := json.Marshal(signUpReq)
	req, _ := http.NewRequest("POST", apiUrl, bytes.NewBuffer(jsonByte))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error: %+v", err)
		return
	}
	defer resp.Body.Close()
	fmt.Sprintf("status:%d\n", resp.Status)
	fmt.Sprintf("response:%+v\n", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

// 更新报名队伍，执行时需要绑定host: [172.18.21.10 grpc-proxy.bilibili.co]
func updateSignInfo(uid, signId, teamId, startTime, endTime int64, otherMessage string) {
	type UpdateUserSignUpInfoReq struct {
		// 用户uid
		Uid int64 `json:"uid"`
		// 报名id
		SignId int64 `json:"sign_id"`
		// 修改数据是哪个 0 所有参数全部覆盖 1只修改other_message 2只修改时间 3. 只修改赛道team_id
		UpdateType int64 `json:"update_type"`
		// 其他信息
		OtherMessage string `json:"other_message"`
		// 开始时间
		StartTime int64 `json:"start_time"`
		// 结束时间
		EndTime int64 `json:"end_time"`
		// 赛道
		TeamId int64 `json:"team_id"`
		// sign state
		SignState int64 `json:"sign_state"`
	}
	apiUrl := "http://grpc-proxy.bilibili.co/live.mozi/live.mozi.signup/UpdateUserSignUpInfo"
	signUpReq := &UpdateUserSignUpInfoReq{
		Uid:          uid,
		SignId:       signId,
		UpdateType:   0,
		OtherMessage: otherMessage,
		StartTime:    startTime,
		EndTime:      endTime,
		TeamId:       teamId,
		SignState:    1,
	}
	jsonByte, _ := json.Marshal(signUpReq)
	req, _ := http.NewRequest("POST", apiUrl, bytes.NewBuffer(jsonByte))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error: %+v", err)
		return
	}
	defer resp.Body.Close()
	fmt.Sprintf("status:%d\n", resp.Status)
	fmt.Sprintf("response:%+v\n", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
