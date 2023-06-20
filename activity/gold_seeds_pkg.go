// 金瓜子包裹补发
package main

import (
	"bufio"
	"fmt"
	"godatabase/gobuild/json"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main(){
	logFile := "/Users/shenxinhua/workspace/code/go/src/godatabase/activity/reward0623.txt"
	loadLogFile(logFile)
}

// load log file
func loadLogFile(filePath string)(list []*APCenterSendPoolRewardReq){
	list = make([]*APCenterSendPoolRewardReq,0)
	f,err := os.Open(filePath)
	if err != nil {
		fmt.Println("[error1]",err)
		return
	}
	defer f.Close()
	scanner:=bufio.NewScanner(f)
	for scanner.Scan(){
		// 读行
		line:=scanner.Text()
		// build log
		var logStruct LogStruct
		err := json.Unmarshal([]byte(line),&logStruct)
		if err != nil {
			fmt.Println("[error2]",err)
			break
		}
		// build content
		item := &APCenterSendPoolRewardReq{}
		// 字段1
		poolIdReg,_ := regexp.Compile(`(pool_id:)\d+`)
		poolIdStr := poolIdReg.FindString(line)
		poolId,_ := strconv.ParseInt(strings.Split(poolIdStr,":")[1],10,64)
		item.PoolId = poolId
		// 字段2
		//extBodyReg, err := regexp.Compile(`(?<=extBody:\\\").+(?=\\" third_id)`) // 这种模式不支持

		fmt.Println(item)
	}
	return
}

func SendPoolReward(req *APCenterSendPoolRewardReq){

}

type APCenterSendPoolRewardReq struct {
	// 预算池id
	PoolId int64 `protobuf:"varint,1,opt,name=pool_id,json=poolId,proto3" json:"pool_id"`
	// 货币类型 3：金瓜子包裹
	PoolType int64 `protobuf:"varint,2,opt,name=pool_type,json=poolType,proto3" json:"pool_type"`
	// 订单ID (幂等处理：transaction + source)
	Transaction string `protobuf:"bytes,3,opt,name=transaction,proto3" json:"transaction"`
	// 订单来源 (幂等处理：transaction + source)
	Source int64 `protobuf:"varint,4,opt,name=source,proto3" json:"source"`
	// 用户uid
	Uid int64 `protobuf:"varint,5,opt,name=uid,proto3" json:"uid"`
	// 货币数量
	Num int64 `protobuf:"varint,6,opt,name=num,proto3" json:"num"`
	// 奖励原因标志
	Reason int64 `protobuf:"varint,7,opt,name=reason,proto3" json:"reason"`
	// 主播侧显示奖励备注
	Remark string `protobuf:"bytes,8,opt,name=remark,proto3" json:"remark"`
	// 调用方时间
	Timestamp int64 `protobuf:"varint,9,opt,name=timestamp,proto3" json:"timestamp"`
	// 额外信息字段 金仓鼠 {int account_type,int delay_day} 金瓜子 {int account_type} 金瓜子包裹道具 {int gift_id,int expired_day}
	ExtBody string `protobuf:"bytes,10,opt,name=extBody,proto3" json:"extBody"`
	// 第三方关联订单号
	ThirdId              string   `protobuf:"bytes,11,opt,name=third_id,json=thirdId,proto3" json:"third_id"`
}

type LogStruct struct {
	Index string `json:"_index"`
	Type string `json:"_type"`
	Source LogStructSource `json:"_source"`
}
type LogStructSource struct {
	Log string `json:"log"`
}