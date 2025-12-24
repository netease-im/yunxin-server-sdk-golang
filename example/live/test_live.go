package main

import (
	"encoding/json"
	"fmt"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/base"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/http"
	"github.com/netease-im/yunxin-server-sdk-golang/src/live"
	"github.com/netease-im/yunxin-server-sdk-golang/src/live/manage"
	"github.com/sirupsen/logrus"
)

func main() {
	// 初始化配置
	appkey := "b882128a2a30430fd3269b2a272565df"
	appsecret := "4eb5f9177a7b"
	timeoutMillis := 5000

	fmt.Println("=== 云信 Live 测试 ===")

	// 创建 Live HTTP 客户端
	client := core.NewYunxinApiHttpClientBuilder(base.BizLIVE, appkey, appsecret).
		TimeoutMillis(timeoutMillis).
		Build()

	// 测试1: 使用 SDK 创建直播频道
	testCreateLiveChannelWithSDK(client)

	// 测试2: 使用原始 HTTP 调用创建直播频道
	testCreateLiveChannelWithRawHTTP(client)

	// 关闭客户端
	client.Shutdown()
	fmt.Println("\n=== 所有测试完成 ===")
}

// testCreateLiveChannelWithSDK 测试使用 SDK 创建直播频道
func testCreateLiveChannelWithSDK(client core.YunxinApiHttpClient) {
	fmt.Println("\n========================================")
	fmt.Println("测试1: 使用 SDK 创建直播频道")
	fmt.Println("========================================")

	// 创建 Live API 服务
	liveApiServices := live.NewYunxinLiveApiServices(client)
	liveManageService := liveApiServices.GetLiveManageService()

	request := &manage.LiveCreateChannelRequest{
		Name: "xxxx",
		Type: 0,
	}

	result, err := liveManageService.CreateChannel(request)
	if err != nil {
		logrus.Errorf("创建直播频道失败: %v", err)
		return
	}

	if result.GetCode() != 200 {
		// 失败时打印错误码
		fmt.Printf("✗ 创建失败，错误码: %d\n", result.GetCode())
		// 打印完整结果
		jsonData, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			logrus.Errorf("JSON序列化失败: %v", err)
			fmt.Printf("原始结果: %+v\n", result)
		} else {
			fmt.Printf("错误信息: %s\n", string(jsonData))
		}
	} else {
		// 成功时打印响应数据
		response := result.GetResponse()
		jsonData, err := json.MarshalIndent(response, "", "  ")
		if err != nil {
			logrus.Errorf("JSON序列化失败: %v", err)
			fmt.Printf("原始结果: %+v\n", response)
		} else {
			fmt.Printf("✓ 创建成功，响应数据: %s\n", string(jsonData))
		}
	}
}

// testCreateLiveChannelWithRawHTTP 测试使用原始 HTTP 调用创建直播频道
func testCreateLiveChannelWithRawHTTP(client core.YunxinApiHttpClient) {
	fmt.Println("\n========================================")
	fmt.Println("测试2: 使用原始 HTTP 调用创建直播频道")
	fmt.Println("========================================")

	// 构造请求数据
	requestData := map[string]interface{}{
		"name": "yyyy",
		"type": 0,
	}

	// 将请求数据转换为 JSON 字符串
	requestJSON, err := json.Marshal(requestData)
	if err != nil {
		logrus.Errorf("请求数据序列化失败: %v", err)
		return
	}

	// 使用 HTTP 客户端发送 POST 请求
	apiResponse, err := client.ExecuteJson(http.POST, "/app/channel/create", nil, nil, string(requestJSON))
	if err != nil {
		logrus.Errorf("请求失败: %v", err)
		return
	}

	// 解析响应数据
	var responseData map[string]interface{}
	if err := json.Unmarshal([]byte(apiResponse.GetData()), &responseData); err != nil {
		logrus.Errorf("响应数据解析失败: %v", err)
		return
	}

	// 检查 code 字段
	code := 0
	if codeValue, exists := responseData["code"]; exists {
		if codeFloat, ok := codeValue.(float64); ok {
			code = int(codeFloat)
		}
	}

	if code != 200 {
		// 请求失败，打印完整响应
		jsonData, err := json.MarshalIndent(apiResponse, "", "  ")
		if err != nil {
			logrus.Errorf("JSON序列化失败: %v", err)
			fmt.Printf("原始响应: %+v\n", apiResponse)
		} else {
			fmt.Printf("✗ 请求失败: %s\n", string(jsonData))
		}
	} else {
		// 请求成功，打印 ret 字段
		if retValue, exists := responseData["ret"]; exists {
			jsonData, err := json.MarshalIndent(retValue, "", "  ")
			if err != nil {
				logrus.Errorf("JSON序列化失败: %v", err)
				fmt.Printf("原始数据: %+v\n", retValue)
			} else {
				fmt.Printf("✓ 请求成功: %s\n", string(jsonData))
			}
		} else {
			fmt.Println("✓ 请求成功，但未返回 ret 字段")
		}
	}
}
