package main

import (
	"encoding/json"
	"fmt"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/base"
	"github.com/netease-im/yunxin-server-sdk-golang/src/rtc"
	"github.com/netease-im/yunxin-server-sdk-golang/src/rtc/room"
	"github.com/sirupsen/logrus"
)

func main() {
	// 请替换为你的实际 appkey 和 appsecret
	appkey := "xxx"
	appsecret := "yyy"

	fmt.Println("=== 云信 RTC 房间服务测试 ===")

	// 创建 RTC HTTP 客户端
	httpClient := core.NewYunxinApiHttpClientBuilder(base.BizRTC, appkey, appsecret).
		TimeoutMillis(10000). // 设置超时时间为10秒
		RetryPolicy(base.NewDefaultRetryPolicy(2, false)).
		Build()

	// 创建 RTC API 服务
	rtcApiServices := rtc.NewYunxinRtcApiServices(httpClient)
	rtcRoomService := rtcApiServices.GetRtcRoomService()

	// 测试创建房间
	//testCreateRoom(rtcRoomService)

	// 测试获取房间信息（通过CID）
	testGetRoomByCid(rtcRoomService)

	// 测试获取房间信息（通过CNAME）
	//testGetRoomByCname(rtcRoomService)

	// 测试列出房间成员（V2版本）
	//testListRoomMembersV2(rtcRoomService)

	// 测试列出房间成员（V3版本）
	//testListRoomMembersV3(rtcRoomService)

	// 测试添加成员到踢出列表（V2版本）
	//testAddMemberToKicklistV2(rtcRoomService)

	// 测试添加成员到踢出列表（V3版本）
	//testAddMemberToKicklistV3(rtcRoomService)

	// 测试删除房间（V2版本）
	//testDeleteRoomV2(rtcRoomService)

	// 测试删除房间（V3版本）
	//testDeleteRoomV3(rtcRoomService)

	// 关闭客户端
	httpClient.Shutdown()
	fmt.Println("\n=== 测试完成 ===")
}

// 测试创建房间
func testCreateRoom(service room.RtcRoomService) {
	fmt.Println("\n1. 测试创建房间")

	request := &room.RtcCreateRoomRequest{
		ChannelName: "test_room_001",
		Mode:        1, // 通话模式
		Uid:         123456789,
	}

	result, err := service.CreateRoom(request)
	if err != nil {
		logrus.Infof("创建房间失败: %v", err)
		return
	}

	printResult("创建房间", result)
}

// 测试根据CID获取房间信息
func testGetRoomByCid(service room.RtcRoomService) {
	fmt.Println("\n2. 测试根据CID获取房间信息")

	request := &room.RtcGetRoomByCidRequest{
		Cid: 1349545895743435, // 替换为实际的房间ID
	}

	result, err := service.GetRoomByCid(request)
	if err != nil {
		logrus.Infof("根据CID获取房间信息失败: %v", err)
		return
	}

	printResult("根据CID获取房间信息", result)
}

// 测试根据CNAME获取房间信息
func testGetRoomByCname(service room.RtcRoomService) {
	fmt.Println("\n3. 测试根据CNAME获取房间信息")

	request := &room.RtcGetRoomByCnameRequest{
		Cname: "csh1218",
	}

	result, err := service.GetRoomByCname(request)
	if err != nil {
		logrus.Infof("根据CNAME获取房间信息失败: %v", err)
		return
	}

	printResult("根据CNAME获取房间信息", result)
}

// 测试列出房间成员（V2版本）
func testListRoomMembersV2(service room.RtcRoomService) {
	fmt.Println("\n4. 测试列出房间成员（V2版本）")

	request := &room.RtcListRoomMembersRequestV2{
		Cid: 1349545916223392, // 替换为实际的房间ID
	}

	result, err := service.ListRoomMembersV2(request)
	if err != nil {
		logrus.Infof("列出房间成员（V2版本）失败: %v", err)
		return
	}

	printResult("列出房间成员（V2版本）", result)
}

// 测试列出房间成员（V3版本）
func testListRoomMembersV3(service room.RtcRoomService) {
	fmt.Println("\n5. 测试列出房间成员（V3版本）")

	role := 0
	request := &room.RtcListRoomMembersRequestV3{
		Cname:    "80191043",
		UserRole: &role,
	}

	result, err := service.ListRoomMembersV3(request)
	if err != nil {
		logrus.Infof("列出房间成员（V3版本）失败: %v", err)
		return
	}

	printResult("列出房间成员（V3版本）", result)
}

// 测试添加成员到踢出列表（V2版本）
func testAddMemberToKicklistV2(service room.RtcRoomService) {
	fmt.Println("\n6. 测试添加成员到踢出列表（V2版本）")

	request := &room.RtcAddMemberToKicklistRequestV2{
		Cid:      1349545916223392, // 替换为实际的房间ID
		Uid:      1764309371,
		Duration: 3600, // 踢出时长1小时（秒）
	}

	result, err := service.AddMemberToKicklistV2(request)
	if err != nil {
		logrus.Infof("添加成员到踢出列表（V2版本）失败: %v", err)
		return
	}

	printResult("添加成员到踢出列表（V2版本）", result)
}

// 测试添加成员到踢出列表（V3版本）
func testAddMemberToKicklistV3(service room.RtcRoomService) {
	fmt.Println("\n7. 测试添加成员到踢出列表（V3版本）")

	request := &room.RtcAddMemberToKicklistRequestV3{
		Cname:    "80191043", // 替换为实际的房间名
		Uid:      1764309371,
		Duration: 3600, // 踢出时长1小时（秒）
	}

	result, err := service.AddMemberToKicklistV3(request)
	if err != nil {
		logrus.Infof("添加成员到踢出列表（V3版本）失败: %v", err)
		return
	}

	printResult("添加成员到踢出列表（V3版本）", result)
}

// 测试删除房间（V2版本）
func testDeleteRoomV2(service room.RtcRoomService) {
	fmt.Println("\n8. 测试删除房间（V2版本）")

	request := &room.RtcDeleteRoomRequestV2{
		Cid: 1349545895743435, // 替换为实际的房间ID
	}

	result, err := service.DeleteRoomV2(request)
	if err != nil {
		logrus.Infof("删除房间（V2版本）失败: %v", err)
		return
	}

	printResult("删除房间（V2版本）", result)
}

// 测试删除房间（V3版本）
func testDeleteRoomV3(service room.RtcRoomService) {
	fmt.Println("\n9. 测试删除房间（V3版本）")

	request := &room.RtcDeleteRoomRequestV3{
		Cname: "test_room_001", // 替换为实际的房间名称
	}

	result, err := service.DeleteRoomV3(request)
	if err != nil {
		logrus.Infof("删除房间（V3版本）失败: %v", err)
		return
	}

	printResult("删除房间（V3版本）", result)
}

// 通用的结果打印函数
func printResult(operation string, result interface{}) {
	fmt.Printf("操作: %s\n", operation)

	// 将结果转换为JSON格式进行美化打印
	jsonData, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		logrus.Infof("JSON序列化失败: %v", err)
		fmt.Printf("原始结果: %+v\n", result)
	} else {
		fmt.Printf("结果: %s\n", string(jsonData))
	}
}
