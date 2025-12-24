
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
	appkey := "xxx"
	appsecret := "xxx"
	timeoutMillis := 5000 * time.Millisecond

	// 创建客户端
	client := core.NewYunxinApiHttpClientBuilder(base.BizRTC, appkey, appsecret).
		TimeoutMillis(timeoutMillis).
		Build()

	// 根据cname查询房间信息
	queryParams := map[string]string{
		"cname": "xxxxx",
	}

	response, err := client.ExecuteJson(http.GET, "/v3/api/rooms", nil, queryParams, "")
	if err != nil {
		fmt.Printf("request error: %v\n", err)
		return
	}

	httpCode := response.GetHttpCode()
	if httpCode != 200 {
		// HTTP状态码非200
		responseJson, _ := json.Marshal(response)
		fmt.Println(string(responseJson))
	} else {
		// 解析业务code
		var result map[string]interface{}
		json.Unmarshal([]byte(response.GetData()), &result)

		code := int(result["code"].(float64))
		if code != 200 {
			// 业务code非200
			fmt.Println(response.GetData())
		} else {
			// 成功
			fmt.Printf("%v\n", result)
		}
	}
}
```

* rtc相关api，http.code和业务code均可能非200，需要同时处理
