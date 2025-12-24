## 快速开始（v1接口，使用服务封装）

```go
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/base"
	v1 "github.com/netease-im/yunxin-server-sdk-golang/src/im/v1"
	"github.com/netease-im/yunxin-server-sdk-golang/src/im/v1/account"
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
	services := v1.NewYunxinV1ApiServices(client)

	// 创建请求
	request := &account.CreateAccountRequestV1{
		Accid: "zhangsan",
	}

	// 执行请求
	result, err := services.GetAccountService().CreateAccount(request)
	if err != nil {
		// 请求异常（如超时等）
		logrus.Infof("register error, error = %v\n", err)
		return
	}

	if result.IsSuccess() {
		response := result.GetResponse()
		// 注册成功
		fmt.Printf("register success, accid = %s, token = %s, traceId = %s\n", 
			response.Accid, response.Token, result.GetTraceId())
	} else {
		// 注册失败，如参数错误、重复注册等
		logrus.Infof("register fail, code = %d, msg = %s, traceId = %s\n", 
			result.GetCode(), result.GetMsg(), result.GetTraceId())
	}
}
```
