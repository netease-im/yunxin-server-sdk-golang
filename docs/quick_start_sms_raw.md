
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
	client := core.NewYunxinApiHttpClientBuilder(base.BizSMS, appkey, appsecret).
		TimeoutMillis(timeoutMillis).
		Build()

	// 构建表单参数
	paramBuilder := http.NewParamBuilder()
	paramBuilder.AddParam("mobile", "13012340000")

	// 执行请求
	apiResponse, err := client.ExecuteForm(http.POST, "/sendcode.action", nil, nil, paramBuilder.Build())
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
		sendid := int64(result["msg"].(float64))
		authCode := result["obj"].(string)
		fmt.Printf("sendid = %d, authcode = %s\n", sendid, authCode)
	}
}
```
