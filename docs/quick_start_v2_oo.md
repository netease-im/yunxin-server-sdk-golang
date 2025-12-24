## 快速开始（v2接口，使用服务封装）

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/base"
	v2 "github.com/netease-im/yunxin-server-sdk-golang/src/im/v2"
	"github.com/netease-im/yunxin-server-sdk-golang/src/im/v2/account"
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

	// 创建服务聚合器
	services := v2.NewYunxinV2ApiServices(client)

	// 创建账号
	{
		request := &account.CreateAccountRequestV2{
			AccountId: "zhangsan",
			Configuration: &account.AccountConfiguration{
				ChatroomChatBanned: true,
			},
			UserInformation: &account.UserInformationV2{
				Email: "xxx@126.com",
			},
		}

		result, err := services.GetAccountService().CreateAccount(request)
		if err != nil {
			// 超时等异常
			logrus.Infof("register error, error = %v\n", err)
			return
		}

		if result.IsSuccess() {
			response := result.GetResponse()
			// 注册成功
			responseJson, _ := json.Marshal(response)
			fmt.Printf("register success, response = %s, traceId = %s\n", string(responseJson), result.GetTraceId())
		} else {
			// 注册失败，如参数错误、重复注册等
			logrus.Infof("register fail, code = %d, msg = %s, traceId = %s\n",
				result.GetCode(), result.GetMsg(), result.GetTraceId())
		}
	}

	// 更新账号
	{
		request := &account.UpdateAccountRequestV2{
			AccountId: "zhangsan",
			Configuration: &account.AccountConfiguration{
				P2pChatBanned: true,
			},
			NeedKick:             true,
			KickNotifyExtension:  "xxxx",
		}

		result, err := services.GetAccountService().UpdateAccount(request)
		if err != nil {
			// 超时等异常
			logrus.Infof("update error, error = %v\n", err)
			return
		}

		if result.IsSuccess() {
			response := result.GetResponse()
			// 更新成功
			responseJson, _ := json.Marshal(response)
			fmt.Printf("update success, response = %s, traceId = %s\n", string(responseJson), result.GetTraceId())
		} else {
			// 更新失败，如参数错误等
			logrus.Infof("update fail, code = %d, msg = %s, traceId = %s\n",
				result.GetCode(), result.GetMsg(), result.GetTraceId())
		}
	}

	// 批量查询账号
	{
		request := &account.BatchQueryAccountsRequestV2{
			AccountList: []string{"zhangsan", "lisi"},
		}

		result, err := services.GetAccountService().BatchQueryAccounts(request)
		if err != nil {
			// 超时等异常
			logrus.Infof("query error, error = %v\n", err)
			return
		}

		if result.IsSuccess() {
			response := result.GetResponse()
			// 查询结果
			successJson, _ := json.Marshal(response.SuccessList)
			failedJson, _ := json.Marshal(response.FailedList)
			fmt.Printf("query success, success_list = %s, traceId = %s\n", string(successJson), result.GetTraceId())
			fmt.Printf("query success, failed_list = %s, traceId = %s\n", string(failedJson), result.GetTraceId())
		} else {
			// 查询失败，如参数错误等
			logrus.Infof("query fail, code = %d, msg = %s, traceId = %s\n",
				result.GetCode(), result.GetMsg(), result.GetTraceId())
		}
	}
}

```
