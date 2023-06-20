package main

import (
	"encoding/json"
	"fmt"
)

type CostIConfItem struct {
	// 成本项目id
	CostObjId int64 `protobuf:"varint,1,opt,name=cost_obj_id,json=costObjId,proto3" json:"cost_obj_id" form:"cost_obj_id"`
	// 成本项目名字
	CostObjName string `protobuf:"bytes,2,opt,name=cost_obj_name,json=costObjName,proto3" json:"cost_obj_name" form:"cost_obj_name"`
	// 预期产出概率 概率*10000
	Chance int64 `protobuf:"varint,3,opt,name=chance,proto3" json:"chance" form:"chance"`
	// 成本（一般以电池为单位）
	Cost int64 `protobuf:"varint,4,opt,name=cost,proto3" json:"cost" form:"cost"`
}

type UpdateCostIncomeReq struct {
	// 监控规则id
	RuleId int64 `protobuf:"varint,1,opt,name=rule_id,json=ruleId,proto3" json:"rule_id" form:"rule_id"`
	//成本项目
	CostItems []*CostIConfItem `protobuf:"bytes,2,rep,name=cost_items,json=costItems,proto3" json:"cost_items" form:"cost_items"`
	// 单次行为收益
	OnceIncome int64 `protobuf:"varint,3,opt,name=once_income,json=onceIncome,proto3" json:"once_income" form:"once_income"`
}

type UpdateCostIncomeResp struct {
}

// 更新成本项
func UpdateCostIncome(req *UpdateCostIncomeReq) (resp *UpdateCostIncomeResp) {
	apiUrl := "https://manager.bilibili.co/x/admin/live-proxy/unsafeproxy"
	reqByte, _ := json.Marshal(req)
	apiProxyParams := &ApiProxyParams{
		Params:   string(reqByte),
		AppId:    "live.live.roi-monitor",
		Module:   "1020001",
		GrpcPath: "/live.live.roimonitor.v1.RoiWarring/UpdateCostIncome",
	}
	apiProxyParamsBody := GetFormParams(apiProxyParams)
	//fmt.Println(apiProxyParamsBody)
	respJson := HttpPost(apiUrl, apiProxyParamsBody)
	resp = &UpdateCostIncomeResp{}
	err := json.Unmarshal([]byte(respJson), &resp)
	if err != nil {
		fmt.Println(err)
	}
	return resp
}
