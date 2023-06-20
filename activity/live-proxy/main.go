package main

import "fmt"

func main() {
	/*resp := AddPoolBalance(&AddPoolBalanceReq{MsgId: "sxh:20230331142701", Cnt: 500000}, false)
	fmt.Println(resp)*/

	/*resp := WebLeftEntry(&WebLeftEntryReq{
		ActIds:       []int64{101814},
		RoomId:       1250399,
		Ruid:         28009181,
		AreaId:       1,
		ParentAreaId: 1,
		Uid:          3493090564246435,
	}, true)
	fmt.Println(resp)*/

	/*resp := GetMultipleWidgetBanner(&GetMultipleWidgetBannerReq{
		ActIds:        []int64{101814},
		RoomId:        1250399,
		Ruid:          28009181,
		Source:        4,
		Uid:           3493090564246435,
		WidgetTypeMap: nil,
	}, true)
	fmt.Println(resp)*/

	/*resp := ActivityHalfInit(&ActivityHalfInitReq{
		ActId:  101814,
		RoomId: 1250399,
		Uid:    3493090564246435,
		Ruid:   28009181,
	}, false)
	fmt.Println(resp)*/

	/*resp := JobApi(&JobApiReq{
		ActId:      101814,
		ActionName: "JobSendMessageCard",
		ExecTime:   "2023-05-31 14:00:01",
	}, false)
	fmt.Println(resp)*/

	/*resp := GetWidgetBannerList(&GetWidgetBannerListReq{
		RoomId:       1250399,
		Ruid:         28009181,
		AreaId:       1,
		ParentAreaId: 1,
		Platform:     "ios",
		Build:        0,
		MobiApp:      "",
		Source:       "live",
		SpecialFlag:  0,
		Uid:          3493090564246435,
		PageSource:   0,
	}, true)
	fmt.Println(resp)*/

	/*resp := GetLeftEntryList(&GetLeftEntryListReq{
		RoomId:       1250399,
		Ruid:         28009181,
		AreaId:       1,
		ParentAreaId: 1,
		Uid:          3493090564246435,
	}, true)
	fmt.Println(resp)*/

	/*resp := RoomShowWidgetBanner(&RoomShowWidgetBannerReq{
		RoomId: 1250399,
	}, true)
	fmt.Println(resp)*/

	/*GetUserItem(&GetUserItemReq{
		ItemId:             1,
		Uid:                433693710,
		TargetUid:          69385121,
		TargetParentAreaId: 0,
		TargetAreaId:       0,
		TargetOther:        "",
	}, false)
	fmt.Println("success")*/

	/*UpdateUserItem(&UpdateUserItemReq{
		ItemId:             1,
		Num:                1,
		Uid:                433693710,
		Expire:             0,
		TargetUid:          69385121,
		TargetParentAreaId: 0,
		TargetAreaId:       0,
		TargetOther:        "",
		Tid:                "sxh:202304070001",
		Source:             "pre",
	}, true)*/

	/*GetMultipleByUids(&UIDsReq{
		Uids:  []int64{3493267610011759},
		Attrs: []string{"status", "area", "show", "live_info", "tags"},
	}, false)*/

	/*AreUidsInRoom(&AreUidsInRoomReq{
		RoomId: 27525672,
		Uids:   []int64{44087490},
		Group:  "",
	}, false)*/

	/*GetHighEnergyRank(&GetHighEnergyRankReq{
		Uid:           0,
		Ruid:          3493267610011759,
		RoomId:        27525672,
		PageSize:      20,
		Page:          1,
		ForcePageSize: 0,
	}, true)*/

	UserGetOnlineRank(&UserGetOnlineRankReq{
		Uid:      0,
		Ruid:     3493267610011759,
		RoomId:   27525672,
		Page:     1,
		PageSize: 20,
	}, false)

	fmt.Println("ok")
}
