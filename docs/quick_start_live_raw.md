
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

	// 创建客户端
	client := core.NewYunxinApiHttpClientBuilder(base.BizLIVE, appkey, appsecret).
		TimeoutMillis(timeoutMillis).
		Build()

	// 构建请求体
	requestData := map[string]interface{}{
		"name": "xxx",
		"type": 0,
	}
	requestBody, _ := json.Marshal(requestData)

	// 执行请求
	apiResponse, err := client.ExecuteJson(http.POST, "/app/channel/create", nil, nil, string(requestBody))
	if err != nil {
		fmt.Printf("request error: %v\n", err)
		return
	}

	// 解析响应
	var result map[string]interface{}
	json.Unmarshal([]byte(apiResponse.GetData()), &result)

	code := int(result["code"].(float64))
	if code != 200 {
		responseJson, _ := json.Marshal(apiResponse)
		fmt.Println(string(responseJson))
	} else {
		ret := result["ret"]
		fmt.Printf("%v\n", ret)
	}
}
```
