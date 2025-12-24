
## 快速开始（使用服务封装）

```go
package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/base"
	"github.com/netease-im/yunxin-server-sdk-golang/src/sms"
)

func main() {
	// 初始化
	appkey := "xx"
	appsecret := "xx"
	timeoutMillis := 5000 * time.Millisecond

	// 创建客户端
	client := core.NewYunxinApiHttpClientBuilder(base.BizSMS, appkey, appsecret).
		TimeoutMillis(timeoutMillis).
		Build()

	// 创建服务聚合器
	services := sms.NewYunxinSmsApiServices(client)
	smsApiService := services.GetSmsApiService()

	// 发送验证码短信
	request := &sms.SmsSendCodeRequest{
		Mobile: "13012340000",
	}

	result, _ := smsApiService.SendCode(request)
	if result.GetCode() != 200 {
		fmt.Println(result.GetCode())
	} else {
		response := result.GetResponse()
		responseJson, _ := json.Marshal(response)
		fmt.Println(string(responseJson))
	}
}
```
