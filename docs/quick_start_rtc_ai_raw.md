
## 快速开始（使用底层HTTP客户端）

```go
package main

import (
	"encoding/json"
	"fmt"
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

	// 创建客户端（使用自定义endpoint）
	client := core.NewYunxinApiHttpClientBuilder(base.BizCUSTOM, appkey, appsecret).
		TimeoutMillis(timeoutMillis).
		Endpoint("https://rtc-ai.yunxinapi.com").
		Build()

	// 构建请求体
	requestData := map[string]interface{}{
		"cname":    "xx",
		"taskType": 7,
		// 其他参数...
	}
	requestBody, _ := json.Marshal(requestData)

	// 执行请求
	response, err := client.ExecuteJson(http.POST, "/v1/api/task/create", nil, nil, string(requestBody))
	if err != nil {
		fmt.Printf("request error: %v\n", err)
		return
	}

	// 输出响应
	responseJson, _ := json.Marshal(response)
	fmt.Println(string(responseJson))
}
```
