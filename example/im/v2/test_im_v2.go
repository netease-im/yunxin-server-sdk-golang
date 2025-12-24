package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/base"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/http"
	imv2 "github.com/netease-im/yunxin-server-sdk-golang/src/im/v2"
	accountv2 "github.com/netease-im/yunxin-server-sdk-golang/src/im/v2/account"
	"github.com/sirupsen/logrus"
)

func main() {
	// 初始化配置
	appkey := "b882128a2a30430fd3269b2a272565df"
	appsecret := "4eb5f9177a7b"
	timeoutMillis := 5000

	fmt.Println("=== 云信 IM V2 API 功能测试 ===")

	// 测试使用原始 V2 API
	testRawV2Api(appkey, appsecret, timeoutMillis)

	// 测试使用 SDK 封装的 V2 API
	testV2ApiWithSDK(appkey, appsecret, timeoutMillis)

	fmt.Println("\n=== 所有测试完成 ===")
}

// testRawV2Api 测试使用原始 V2 API
func testRawV2Api(appkey, appsecret string, timeoutMillis int) {
	fmt.Println("\n========================================")
	fmt.Println("测试1: 使用原始 V2 API")
	fmt.Println("========================================")

	client := core.NewYunxinApiHttpClientBuilder(base.BizIM, appkey, appsecret).
		TimeoutMillis(timeoutMillis).
		Build()

	// 1. 创建账号
	{
		fmt.Println("\n--- 1.1 创建账号 ---")
		uri := "/im/v2/accounts" // 仅用于统计
		requestData := map[string]interface{}{
			"account_id": "zhangsan",
		}
		requestJSON, _ := json.Marshal(requestData)

		response, err := client.ExecuteV2Api(http.POST, uri, nil, nil, string(requestJSON))
		if err != nil {
			var sdkErr *base.YunxinSdkError
			if errors.As(err, &sdkErr) {
				fmt.Printf("✗ 请求失败, traceId=%s\n", sdkErr.TraceId)
			} else {
				logrus.Errorf("请求失败: %v", err)
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
			data := responseData["data"].(map[string]interface{})
			fmt.Printf("✓ 注册成功, accid=%s, token=%s, traceId=%s\n",
				data["account_id"], data["token"], response.GetTraceId())
		}
	}

	// 2. 更新账号
	{
		fmt.Println("\n--- 1.2 更新账号 ---")
		uri := "/im/v2/accounts/{account_id}" // 仅用于统计
		pathParams := map[string]string{
			"account_id": "zhangsan",
		}
		requestData := map[string]interface{}{
			"token": "abc",
		}
		requestJSON, _ := json.Marshal(requestData)

		response, err := client.ExecuteV2Api(http.PATCH, uri, pathParams, nil, string(requestJSON))
		if err != nil {
			var sdkErr *base.YunxinSdkError
			if errors.As(err, &sdkErr) {
				fmt.Printf("✗ 请求失败, traceId=%s\n", sdkErr.TraceId)
			} else {
				logrus.Errorf("请求失败: %v", err)
			}
			client.Shutdown()
			return
		}

		var responseData map[string]interface{}
		_ = json.Unmarshal([]byte(response.GetData()), &responseData)

		code := int(responseData["code"].(float64))
		if code != 200 {
			fmt.Printf("✗ 更新失败, response=%s, traceId=%s\n",
				response.GetData(), response.GetTraceId())
		} else {
			data := responseData["data"].(map[string]interface{})
			fmt.Printf("✓ 更新成功, accid=%s, token=%s, traceId=%s\n",
				data["account_id"], data["token"], response.GetTraceId())
		}
	}

	// 3. 批量查询账号
	{
		fmt.Println("\n--- 1.3 批量查询账号 ---")
		uri := "/im/v2/accounts"
		queryString := map[string]string{
			"account_ids": "zhagnsan,lisi",
		}

		response, err := client.ExecuteV2Api(http.GET, uri, nil, queryString, "")
		if err != nil {
			var sdkErr *base.YunxinSdkError
			if errors.As(err, &sdkErr) {
				fmt.Printf("✗ 请求失败, traceId=%s\n", sdkErr.TraceId)
			} else {
				logrus.Errorf("请求失败: %v", err)
			}
			client.Shutdown()
			return
		}

		var responseData map[string]interface{}
		_ = json.Unmarshal([]byte(response.GetData()), &responseData)

		code := int(responseData["code"].(float64))
		if code != 200 {
			fmt.Printf("✗ 查询失败, response=%s, traceId=%s\n",
				response.GetData(), response.GetTraceId())
		} else {
			data := responseData["data"].(map[string]interface{})
			successList := data["success_list"]
			failedList := data["failed_list"]
			fmt.Printf("✓ 查询成功, success_list=%v, failed_list=%v\n",
				successList, failedList)
		}
	}

	client.Shutdown()
}

// testV2ApiWithSDK 测试使用 SDK 封装的 V2 API
func testV2ApiWithSDK(appkey, appsecret string, timeoutMillis int) {
	fmt.Println("\n========================================")
	fmt.Println("测试2: 使用 SDK 封装的 V2 API")
	fmt.Println("========================================")

	client := core.NewYunxinApiHttpClientBuilder(base.BizIM, appkey, appsecret).
		TimeoutMillis(timeoutMillis).
		Build()

	services := imv2.NewYunxinV2ApiServices(client)

	// 1. 创建账号
	{
		fmt.Println("\n--- 1. 创建账号 ---")
		request := &accountv2.CreateAccountRequestV2{
			AccountId: "zhangsan",
		}

		// 设置配置
		configuration := &accountv2.AccountConfiguration{
			ChatroomChatBanned: true,
		}
		request.Configuration = configuration

		// 设置用户信息
		userInformation := &accountv2.UserInformation{
			Email: "xxx@126.com",
		}
		request.UserInformation = userInformation

		result, err := services.GetAccountService().CreateAccount(request)
		if err != nil {
			var sdkErr *base.YunxinSdkError
			if errors.As(err, &sdkErr) {
				fmt.Printf("✗ 注册失败（异常）, endpoint=%s, traceId=%s, err=%v\n",
					sdkErr.Endpoint, sdkErr.TraceId, sdkErr.Err)
			} else {
				logrus.Errorf("注册失败: %v", err)
			}
		} else if result.IsSuccess() {
			response := result.GetResponse()
			jsonData, _ := json.MarshalIndent(response, "", "  ")
			fmt.Printf("✓ 注册成功, response=%s, traceId=%s\n",
				string(jsonData), result.GetTraceId())
		} else {
			fmt.Printf("✗ 注册失败, code=%d, msg=%s, traceId=%s\n",
				result.GetCode(), result.GetMsg(), result.GetTraceId())
		}
	}

	// 2. 更新账号
	{
		fmt.Println("\n--- 2. 更新账号 ---")
		request := &accountv2.UpdateAccountRequestV2{
			AccountId: "zhangsan",
		}

		configuration := &accountv2.AccountConfiguration{
			P2pChatBanned: true,
		}
		request.Configuration = configuration
		request.NeedKick = true
		request.KickNotifyExtension = "xxxx"

		result, err := services.GetAccountService().UpdateAccount(request)
		if err != nil {
			var sdkErr *base.YunxinSdkError
			if errors.As(err, &sdkErr) {
				fmt.Printf("✗ 更新失败（异常）, endpoint=%s, traceId=%s, err=%v\n",
					sdkErr.Endpoint, sdkErr.TraceId, sdkErr.Err)
			} else {
				logrus.Errorf("更新失败: %v", err)
			}
		} else if result.IsSuccess() {
			response := result.GetResponse()
			jsonData, _ := json.MarshalIndent(response, "", "  ")
			fmt.Printf("✓ 更新成功, response=%s, traceId=%s\n",
				string(jsonData), result.GetTraceId())
		} else {
			fmt.Printf("✗ 更新失败, code=%d, msg=%s, traceId=%s\n",
				result.GetCode(), result.GetMsg(), result.GetTraceId())
		}
	}

	// 3. 批量查询账号
	{
		fmt.Println("\n--- 3. 批量查询账号 ---")
		request := &accountv2.BatchQueryAccountsRequestV2{
			AccountList: []string{"zhangsan", "lisi"},
		}

		result, err := services.GetAccountService().BatchQueryAccounts(request)
		if err != nil {
			var sdkErr *base.YunxinSdkError
			if errors.As(err, &sdkErr) {
				fmt.Printf("✗ 查询失败（异常）, endpoint=%s, traceId=%s, err=%v\n",
					sdkErr.Endpoint, sdkErr.TraceId, sdkErr.Err)
			} else {
				logrus.Errorf("查询失败: %v", err)
			}
		} else if result.IsSuccess() {
			response := result.GetResponse()
			successJSON, _ := json.MarshalIndent(response.SuccessList, "", "  ")
			failedJSON, _ := json.MarshalIndent(response.FailedList, "", "  ")
			fmt.Printf("✓ 查询成功, success_list=%s\n", string(successJSON))
			fmt.Printf("  failed_list=%s, traceId=%s\n",
				string(failedJSON), result.GetTraceId())
		} else {
			fmt.Printf("✗ 查询失败, code=%d, msg=%s, traceId=%s\n",
				result.GetCode(), result.GetMsg(), result.GetTraceId())
		}
	}

	client.Shutdown()
}
