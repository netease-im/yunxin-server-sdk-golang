
## 快速开始（使用服务封装）

```go
package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/base"
	"github.com/netease-im/yunxin-server-sdk-golang/src/rtc"
	"github.com/netease-im/yunxin-server-sdk-golang/src/rtc/room"
)

func main() {
	// 初始化
	appkey := "xxxx"
	appsecret := "xxx"
	timeoutMillis := 5000 * time.Millisecond

	// 创建客户端
	client := core.NewYunxinApiHttpClientBuilder(base.BizRTC, appkey, appsecret).
		TimeoutMillis(timeoutMillis).
		Build()

	// 创建服务聚合器
	services := rtc.NewYunxinRtcApiServices(client)
	rtcRoomService := services.GetRtcRoomService()

	// 创建房间
	{
		request := &room.RtcCreateRoomRequest{
			ChannelName: "xxx",
			Mode:        2,
			Uid:         123,
		}

		result, _ := rtcRoomService.CreateRoom(request)
		if result.IsSuccess() {
			responseJson, _ := json.Marshal(result.GetResponse())
			fmt.Println(string(responseJson))
		} else {
			resultJson, _ := json.Marshal(result)
			fmt.Println(string(resultJson))
		}
	}

	// 根据cname查询房间信息
	{
		request := &room.RtcGetRoomByCnameRequest{
			Cname: "xxx",
		}

		result, _ := rtcRoomService.GetRoomByCname(request)
		if result.IsSuccess() {
			responseJson, _ := json.Marshal(result.GetResponse())
			fmt.Println(string(responseJson))
		} else {
			resultJson, _ := json.Marshal(result)
			fmt.Println(string(resultJson))
		}
	}

	// 根据cid查询房间信息
	{
		request := &room.RtcGetRoomByCidRequest{
			Cid: 123,
		}

		result, _ := rtcRoomService.GetRoomByCid(request)
		if result.IsSuccess() {
			responseJson, _ := json.Marshal(result.GetResponse())
			fmt.Println(string(responseJson))
		} else {
			resultJson, _ := json.Marshal(result)
			fmt.Println(string(resultJson))
		}
	}
}
```

* rtc相关api，http.code和业务code均可能非200，需要同时处理
