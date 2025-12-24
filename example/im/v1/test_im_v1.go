package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/base"
	imv1 "github.com/netease-im/yunxin-server-sdk-golang/src/im/v1"
	accountv1 "github.com/netease-im/yunxin-server-sdk-golang/src/im/v1/account"
	"github.com/sirupsen/logrus"
)

func main() {
	// 初始化配置
	appkey := "b882128a2a30430fd3269b2a272565df"
	appsecret := "4eb5f9177a7b"
	timeoutMillis := 5000

	fmt.Println("=== 云信 IM V1 API 功能测试 ===")

	// 测试使用原始 V1 API
	testRawV1Api(appkey, appsecret, timeoutMillis)

	// 测试使用 SDK 封装的 V1 API
	testV1ApiWithSDK(appkey, appsecret, timeoutMillis)

	fmt.Println("\n=== 所有测试完成 ===")
}

// testRawV1Api 测试使用原始 V1 API
func testRawV1Api(appkey, appsecret string, timeoutMillis int) {
	fmt.Println("\n========================================")
	fmt.Println("测试1: 测试使用原始 V1 API")
	fmt.Println("========================================")

	client := core.NewYunxinApiHttpClientBuilder(base.BizIM, appkey, appsecret).
		TimeoutMillis(timeoutMillis).
		Build()

	// 创建账号请求
	queryParams := make(map[string]string)
	queryParams["accid"] = "zhangsan"
	response, err := client.ExecuteV1Api("/user/create.action", queryParams)
	if err != nil {
		// 判断是否是 SDK 异常
		var sdkErr *base.YunxinSdkError
		if errors.As(err, &sdkErr) {
			fmt.Printf("✗ 注册失败（异常）, endpoint=%s, traceId=%s, err=%v\n",
				sdkErr.Endpoint, sdkErr.TraceId, sdkErr.Err)
		} else {
			logrus.Errorf("注册失败: %v", err)
		}
		client.Shutdown()
		return
	}

	var responseData map[string]interface{}
	_ = json.Unmarshal([]byte(response.GetData()), &responseData)

	code := int(responseData["code"].(float64))
	if code != 200 {
		fmt.Printf("✗ 注册失败, response=%s, traceId=%s\n",
			response.GetData(), response.GetTraceId())
	} else {
		data := responseData["info"].(map[string]interface{})
		fmt.Printf("✓ 注册成功, accid=%s, token=%s, traceId=%s\n",
			data["accid"], data["token"], response.GetTraceId())
	}

	client.Shutdown()
}

// testV1ApiWithSDK 测试使用 SDK 封装的 V1 API
func testV1ApiWithSDK(appkey, appsecret string, timeoutMillis int) {
	fmt.Println("\n========================================")
	fmt.Println("测试2: 测试使用 SDK 封装的 V1 API")
	fmt.Println("========================================")

	client := core.NewYunxinApiHttpClientBuilder(base.BizIM, appkey, appsecret).
		TimeoutMillis(timeoutMillis).
		Build()

	// 创建 V1 API 服务
	services := imv1.NewYunxinV1ApiServices(client)

	// 创建账号请求
	request := &accountv1.CreateAccountRequestV1{
		Accid: "zhangsan2",
	}

	result, err := services.GetAccountService().CreateAccount(request)
	if err != nil {
		// 判断是否是 SDK 异常
		var sdkErr *base.YunxinSdkError
		if errors.As(err, &sdkErr) {
			fmt.Printf("✗ 注册失败（异常）, endpoint=%s, traceId=%s, err=%v\n",
				sdkErr.Endpoint, sdkErr.TraceId, sdkErr.Err)
		} else {
			logrus.Errorf("注册失败: %v", err)
		}
		client.Shutdown()
		return
	}

	if result.IsSuccess() {
		response := result.GetResponse()
		fmt.Printf("✓ 注册成功, accid=%s, token=%s, traceId=%s\n",
			response.Accid, response.Token, result.GetTraceId())
	} else {
		fmt.Printf("✗ 注册失败, code=%d, msg=%s, traceId=%s\n",
			result.GetCode(), result.GetMsg(), result.GetTraceId())
	}

	client.Shutdown()
}
