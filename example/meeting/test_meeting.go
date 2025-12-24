package main

import (
	"encoding/json"
	"fmt"

	"github.com/netease-im/yunxin-server-sdk-golang/example"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/base"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/http"
	"github.com/netease-im/yunxin-server-sdk-golang/src/meeting"
	"github.com/netease-im/yunxin-server-sdk-golang/src/meeting/user"
	"github.com/sirupsen/logrus"
)

func main() {
	// 初始化配置
	appkey := "b882128a2a30430fd3269b2a272565df"
	appsecret := "4eb5f9177a7b"
	timeoutMillis := 5000

	fmt.Println("=== 云信 Meeting 测试 ===")

	// 创建 Meeting HTTP 客户端
	client := core.NewYunxinApiHttpClientBuilder(base.BizMEETING, appkey, appsecret).
		TimeoutMillis(timeoutMillis).
		Build()

	// 测试1: 使用 SDK 创建会议账号
	testCreateMeetingAccountWithSDK(client)

	// 测试2: 使用原始 HTTP 调用创建会议账号
	testCreateMeetingAccountWithRawHTTP(client)

	// 关闭客户端
	client.Shutdown()
	fmt.Println("\n=== 所有测试完成 ===")
}

// testCreateMeetingAccountWithSDK 测试使用 SDK 创建会议账号
func testCreateMeetingAccountWithSDK(client core.YunxinApiHttpClient) {
	fmt.Println("\n========================================")
	fmt.Println("测试1: 使用 SDK 创建会议账号")
	fmt.Println("========================================")

	// 创建 Meeting API 服务
	meetingApiServices := meeting.NewYunxinMeetingApiServices(client)
	meetingUserService := meetingApiServices.GetMeetingUserService()

	userUuid := example.GenerateUuidWithoutDash()
	request := &user.CreateMeetingAccountRequest{
		UserUuid: userUuid,
		Username: "zhangsan",
	}

	result, err := meetingUserService.CreateAccount(request)
	if err != nil {
		logrus.Errorf("创建会议账号失败: %v", err)
		return
	}

	if result.IsSuccess() {
		// 成功时打印响应数据
		jsonData, err := json.MarshalIndent(result.GetResponse(), "", "  ")
		if err != nil {
			logrus.Errorf("JSON序列化失败: %v", err)
			fmt.Printf("原始结果: %+v\n", result.GetResponse())
		} else {
			fmt.Printf("✓ 创建成功，响应数据: %s\n", string(jsonData))
		}
	} else {
		// 失败时打印完整结果
		jsonData, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			logrus.Errorf("JSON序列化失败: %v", err)
			fmt.Printf("原始结果: %+v\n", result)
		} else {
			fmt.Printf("✗ 创建失败，错误信息: %s\n", string(jsonData))
		}
	}
}

// testCreateMeetingAccountWithRawHTTP 测试使用原始 HTTP 调用创建会议账号
func testCreateMeetingAccountWithRawHTTP(client core.YunxinApiHttpClient) {
	fmt.Println("\n========================================")
	fmt.Println("测试2: 使用原始 HTTP 调用创建会议账号")
	fmt.Println("========================================")

	// 构造请求数据
	requestData := map[string]interface{}{
		"userUuid": example.GenerateUuidWithoutDash(),
		"name":     "zhangsan",
	}

	// 将请求数据转换为 JSON 字符串
	requestJSON, err := json.Marshal(requestData)
	if err != nil {
		logrus.Errorf("请求数据序列化失败: %v", err)
		return
	}

	// 使用 HTTP 客户端发送 POST 请求
	response, err := client.ExecuteJson(http.POST, "/scene/meeting/api/v2/add-user", nil, nil, string(requestJSON))
	if err != nil {
		logrus.Errorf("请求失败: %v", err)
		return
	}

	// 解析响应数据
	var responseData map[string]interface{}
	if err := json.Unmarshal([]byte(response.GetData()), &responseData); err != nil {
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

	if code != 0 {
		// 请求失败，打印完整响应
		jsonData, err := json.MarshalIndent(responseData, "", "  ")
		if err != nil {
			logrus.Errorf("JSON序列化失败: %v", err)
			fmt.Printf("原始响应: %+v\n", responseData)
		} else {
			fmt.Printf("✗ 请求失败: %s\n", string(jsonData))
		}
	} else {
		// 请求成功，打印 data 字段
		if dataValue, exists := responseData["data"]; exists {
			jsonData, err := json.MarshalIndent(dataValue, "", "  ")
			if err != nil {
				logrus.Errorf("JSON序列化失败: %v", err)
				fmt.Printf("原始数据: %+v\n", dataValue)
			} else {
				fmt.Printf("✓ 请求成功: %s\n", string(jsonData))
			}
		} else {
			fmt.Println("✓ 请求成功，但未返回 data 字段")
		}
	}
}
