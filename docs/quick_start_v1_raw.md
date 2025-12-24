
## 快速开始（v1接口，使用底层HTTP客户端）

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/base"
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

	// 请求地址
	path := "/user/create.action"

	// 请求参数
	paramMap := map[string]string{
		"accid": "zhangsan",
	}

	// 执行请求
	response, err := client.ExecuteV1Api(path, paramMap)
	if err != nil {
		// 请求失败
		logrus.Infof("register error, error = %v\n", err)
		return
	}

	// 获取结果
	data := response.GetData()

	// 解析结果
	var result map[string]interface{}
	if err := json.Unmarshal([]byte(data), &result); err != nil {
		logrus.Infof("parse json error, error = %v\n", err)
		return
	}

	code := int(result["code"].(float64))
	if code != 200 {
		// 注册失败
		logrus.Infof("register fail, response = %s, traceId = %s\n", data, response.GetTraceId())
	} else {
		// 注册成功
		info := result["info"].(map[string]interface{})
		accid := info["accid"].(string)
		token := info["token"].(string)
		fmt.Printf("register success, accid = %s, token = %s, traceId = %s\n", accid, token, response.GetTraceId())
	}
}
```
