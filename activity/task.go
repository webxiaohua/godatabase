// 主播任务加分
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main(){

	/*var uid,taskId,timestamp,score int64
	uid=433693710
	taskId=763
	timestamp = time.Now().Unix()
	score=1
	msgId:=fmt.Sprintf("fix_sxh:20220819190001")
	incrAnchorTaskScore(uid,taskId,timestamp,score,msgId)*/
	//incrAnchorTaskScore(28008918,786,time.Now().Unix(),200,"sxh:20220824150006")
}

// 主播任务加分，执行时需要绑定host: [172.18.21.10 grpc-proxy.bilibili.co]
func incrAnchorTaskScore(uid,taskId,timestamp,score int64,msgId string)  {
	type IncrTaskProDataReq struct {
		//主任务id
		TaskId int64 `protobuf:"varint,1,opt,name=task_id,json=taskId,proto3" json:"task_id" validate:"required"`
		//时间戳
		Timestamp int64 `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp"`
		//增加多少进度 (由于一个任务有多个子任务， 如果不传子任务结构体则默认加所有子任务的进度,否则只加指定的任务类型进度)
		Score int64 `protobuf:"varint,3,opt,name=score,proto3" json:"score"  validate:"required"`
		//子任务结构体 -  指定某个类型任务增加多少进度 taskTYpe -> score
		SubTaskProData map[int64]int64 `protobuf:"bytes,4,rep,name=sub_task_pro_data,json=subTaskProData,proto3" json:"sub_task_pro_data" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
		//用户id
		Uid int64 `protobuf:"varint,5,opt,name=uid,proto3" json:"uid" validate:"required", gt=0`
		//幂等id
		MsgId                string   `protobuf:"bytes,6,opt,name=msg_id,json=msgId,proto3" json:"msg_id" validate:"required"`
	}
	apiUrl := "http://grpc-proxy.bilibili.co/live.activitytask/live.activitytask.AnchorTaskInfo/IncrTaskProData"
	incrReq := &IncrTaskProDataReq{
		TaskId:         taskId,
		Timestamp:      timestamp,
		Score:          score,
		SubTaskProData: nil,
		Uid:            uid,
		MsgId:          msgId,
	}
	jsonByte,_ := json.Marshal(incrReq)
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