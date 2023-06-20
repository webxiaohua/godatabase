package main

import (
	"bufio"
	"fmt"
	"github.com/dlclark/regexp2"
	"godatabase/gobuild/json"
	"os"
	"strings"
)

func main() {
	logFile := "/Users/shenxinhua/workspace/code/go/src/godatabase/activity/pickUpLog/rankdb_20230410001.txt"
	msgList, totalLine := GetLogMstList(logFile, `(?<=req = ).+(?=","log.observed_time)`)
	fmt.Println(msgList)
	fmt.Printf("totalLine: %d \nhitLine: %d", totalLine, len(msgList))
	dataList := PickUpKeywords(msgList)
	fmt.Println(dataList)
	for _, dataMap := range dataList {
		dataByte, _ := json.Marshal(dataMap)
		tmp := &RankParams{}
		json.Unmarshal(dataByte, &tmp)
		fmt.Println(tmp)
	}
}

// load log file
/*type LogMsgStruct struct {
	LogMsg string `json:"log.msg"`
}*/

// 解析日志字段
func GetLogMstList(filePath string, expression string) (msgList []string, total int) {
	msgList = make([]string, 0)
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("[error1]", err)
		return
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		total++
		// 读行
		line := scanner.Text()
		logMsgReg, _ := regexp2.Compile(expression, 0)
		logMsgMatch, _ := logMsgReg.FindStringMatch(line)
		if logMsgMatch != nil {
			msgList = append(msgList, logMsgMatch.String())
		}

		// build log
		/*var logStruct LogMsgStruct
		err := json.Unmarshal([]byte(line), &logStruct)
		if err != nil {
			fmt.Println("[error2]", err)
			break
		}
		msgList = append(msgList, logStruct.LogMsg)*/

		/*// build content
		item := &APCenterSendPoolRewardReq{}
		// 字段1
		poolIdReg, _ := regexp.Compile(`(pool_id:)\d+`)
		poolIdStr := poolIdReg.FindString(line)
		poolId, _ := strconv.ParseInt(strings.Split(poolIdStr, ":")[1], 10, 64)
		item.PoolId = poolId
		// 字段2
		//extBodyReg, err := regexp.Compile(`(?<=extBody:\\\").+(?=\\" third_id)`) // 这种模式不支持

		fmt.Println(item)*/
	}
	return
}

// 提取关键字段
func PickUpKeywords(msgList []string) (dataList []map[string]interface{}) {
	dataList = make([]map[string]interface{}, 0)
	for _, msg := range msgList {
		tmp := make(map[string]interface{})
		msg = strings.ReplaceAll(msg, "\\", "")
		if err := json.Unmarshal([]byte(msg), &tmp); err != nil {
			fmt.Println("error:", err)
		}
		dataList = append(dataList, tmp)
	}
	return
}

type RankParams struct {
	Ruid     int64  `json:"ruid"`
	Platform string `json:"platform"`
}
