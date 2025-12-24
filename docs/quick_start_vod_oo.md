
## 快速开始（使用服务封装）

```go
package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/base"
	"github.com/netease-im/yunxin-server-sdk-golang/src/vod"
	"github.com/netease-im/yunxin-server-sdk-golang/src/vod/upload"
)

func main() {
	// 初始化
	appkey := "xx"
	appsecret := "xx"
	timeoutMillis := 5000 * time.Millisecond

	// 创建客户端
	client := core.NewYunxinApiHttpClientBuilder(base.BizVOD, appkey, appsecret).
		TimeoutMillis(timeoutMillis).
		Build()

	// 创建服务聚合器
	vodApiServices := vod.NewYunxinVodApiServices(client)
	vodUploadService := vodApiServices.GetVodUploadService()

	// 创建请求
	request := &upload.VodUploadInitRequest{
		OriginFileName: "xxxx",
	}

	result, _ := vodUploadService.UploadInit(request)
	if result.IsSuccess() {
		responseJson, _ := json.Marshal(result.GetResponse())
		fmt.Println(string(responseJson))
	} else {
		resultJson, _ := json.Marshal(result)
		fmt.Println(string(resultJson))
	}
}
```
