
## 快速开始（使用底层HTTP客户端）

```go
package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
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
	client := core.NewYunxinApiHttpClientBuilder(base.BizMEETING, appkey, appsecret).
		TimeoutMillis(timeoutMillis).
		Build()

	// 构建请求体
	userUuid := strings.ReplaceAll(uuid.New().String(), "-", "")
	requestData := map[string]interface{}{
		"userUuid": userUuid,
		"name":     "zhangsan",
	}
	requestBody, _ := json.Marshal(requestData)

	// 执行请求
	response, err := client.ExecuteJson(http.POST, "/scene/meeting/api/v2/add-user", nil, nil, string(requestBody))
	if err != nil {
		fmt.Printf("request error: %v\n", err)
		return
	}

	// 解析响应
	var result map[string]interface{}
	json.Unmarshal([]byte(response.GetData()), &result)

	code := int(result["code"].(float64))
	if code != 200 {
		fmt.Printf("%v\n", result)
	} else {
		data := result["data"]
		fmt.Printf("%v\n", data)
	}
}
```
