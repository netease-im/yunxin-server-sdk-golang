package main

import (
	"encoding/json"
	"fmt"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/base"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/http"
	"github.com/netease-im/yunxin-server-sdk-golang/src/sms"
	"github.com/sirupsen/logrus"
)

func main() {
	// 初始化配置
	appkey := "b882128a2a30430fd3269b2a272565df"
	appsecret := "4eb5f9177a7b"
	timeoutMillis := 5000

	fmt.Println("=== 云信 SMS 测试 ===")

	// 创建 SMS HTTP 客户端
	client := core.NewYunxinApiHttpClientBuilder(base.BizSMS, appkey, appsecret).
		TimeoutMillis(timeoutMillis).
		Build()

	// 测试1: 使用 SDK 发送短信验证码
	testSendCodeWithSDK(client)

	// 测试2: 使用原始 HTTP 调用发送短信验证码
	testSendCodeWithRawHTTP(client)

	// 关闭客户端
	client.Shutdown()
	fmt.Println("\n=== 所有测试完成 ===")
}

// testSendCodeWithSDK 测试使用 SDK 发送短信验证码
func testSendCodeWithSDK(client core.YunxinApiHttpClient) {
	fmt.Println("\n========================================")
	fmt.Println("测试1: 使用 SDK 发送短信验证码")
	fmt.Println("========================================")

	// 创建 SMS API 服务
	smsApiServices := sms.NewYunxinSmsApiServices(client)
	smsApiService := smsApiServices.GetSmsApiService()

	request := &sms.SmsSendCodeRequest{
		Mobile: "13012340000",
	}

	result, err := smsApiService.SendCode(request)
	if err != nil {
		logrus.Errorf("发送短信验证码失败: %v", err)
		return
	}

	if result.GetCode() != 200 {
		// 失败时打印错误码
		fmt.Printf("✗ 发送失败，错误码: %d\n", result.GetCode())
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
			fmt.Printf("✓ 发送成功，响应数据: %s\n", string(jsonData))
		}
	}
}

// testSendCodeWithRawHTTP 测试使用原始 HTTP 调用发送短信验证码
func testSendCodeWithRawHTTP(client core.YunxinApiHttpClient) {
	fmt.Println("\n========================================")
	fmt.Println("测试2: 使用原始 HTTP 调用发送短信验证码")
	fmt.Println("========================================")

	// 构建表单参数
	paramBuilder := http.NewParamBuilder()
	paramBuilder.AddParam("mobile", "13012340000")

	// 使用 HTTP 客户端发送 POST 请求（Form格式）
	apiResponse, err := client.ExecuteForm(http.POST, "/sendcode.action", nil, nil, paramBuilder.Build())
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
		// 请求成功，提取 sendid 和 authCode
		var sendid int64
		var authCode string

		if msgValue, exists := responseData["msg"]; exists {
			if msgFloat, ok := msgValue.(float64); ok {
				sendid = int64(msgFloat)
			}
		}

		if objValue, exists := responseData["obj"]; exists {
			if objStr, ok := objValue.(string); ok {
				authCode = objStr
			}
		}

		fmt.Printf("✓ 请求成功: sendid=%d, authcode=%s\n", sendid, authCode)
	}
}
