## 快速开始（v2接口，使用底层HTTP客户端）

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/base"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/http"
)

func main() {
	// 初始化
	appkey := "xx"
	appsecret := "xx"
	timeoutMillis := 5000 * time.Millisecond

	// 创建客户端
	client := core.NewYunxinApiHttpClientBuilder(base.BizIM, appkey, appsecret).
		TimeoutMillis(timeoutMillis).
		Build()

	// 创建账号
	{
		uri := "/im/v2/accounts"   // 仅仅用于统计
		path := "/im/v2/accounts"

		// 构建请求体
		requestData := map[string]interface{}{
			"account_id": "zhangsan",
		}
		requestBody, _ := json.Marshal(requestData)

		// 执行请求
		response, err := client.ExecuteV2Api(http.POST, uri, nil, nil, string(requestBody))
		if err != nil {
			// 请求失败
			logrus.Infof("register error, error = %v\n", err)
			return
		}

		// 获取结果
		responseData := response.GetData()

		// 解析结果
		var result map[string]interface{}
		json.Unmarshal([]byte(responseData), &result)

		code := int(result["code"].(float64))
		if code != 200 {
			// 注册失败
			logrus.Infof("register fail, response = %s, traceId = %s\n", responseData, response.GetTraceId())
		} else {
			// 注册成功
			data := result["data"].(map[string]interface{})
			accid := data["account_id"].(string)
			token := data["token"].(string)
			fmt.Printf("register success, accid = %s, token = %s, traceId = %s\n", accid, token, response.GetTraceId())
		}
	}

	// 更新账号
	{
		uri := "/im/v2/accounts/{account_id}" // 仅仅用于统计
		path := "/im/v2/accounts/zhangsan"

		// 构建请求体
		requestData := map[string]interface{}{
			"token": "abc",
		}
		requestBody, _ := json.Marshal(requestData)

		// 执行请求
		response, err := client.ExecuteV2Api(http.PATCH, uri, nil, nil, string(requestBody))
		if err != nil {
			// 请求失败
			logrus.Infof("update error, error = %v\n", err)
			return
		}

		// 获取结果
		responseData := response.GetData()

		// 解析结果
		var result map[string]interface{}
		json.Unmarshal([]byte(responseData), &result)

		code := int(result["code"].(float64))
		if code != 200 {
			// 更新失败
			logrus.Infof("update fail, response = %s, traceId = %s\n", responseData, response.GetTraceId())
		} else {
			// 更新成功
			data := result["data"].(map[string]interface{})
			accid := data["account_id"].(string)
			token := data["token"].(string)
			fmt.Printf("update success, accid = %s, token = %s, traceId = %s\n", accid, token, response.GetTraceId())
		}
	}

	// 批量查询账号信息
	{
		uri := "/im/v2/accounts"  // 仅仅用于统计
		path := "/im/v2/accounts"

		// 查询参数
		queryParams := map[string]string{
			"account_ids": "account1,account2,account3",
		}

		// 执行请求
		response, err := client.ExecuteV2Api(http.GET, uri, nil, queryParams, "")
		if err != nil {
			// 请求失败
			logrus.Infof("query error, error = %v\n", err)
			return
		}

		// 获取结果
		responseData := response.GetData()

		// 解析结果
		var result map[string]interface{}
		json.Unmarshal([]byte(responseData), &result)

		code := int(result["code"].(float64))
		if code != 200 {
			// 查询失败
			logrus.Infof("query fail, response = %s, traceId = %s\n", responseData, response.GetTraceId())
		} else {
			// 查询成功
			data := result["data"].(map[string]interface{})
			successList := data["success_list"]
			failedList := data["failed_list"]
			fmt.Printf("success_list = %v\n", successList)
			fmt.Printf("failed_list = %v\n", failedList)
		}
	}
}
```
