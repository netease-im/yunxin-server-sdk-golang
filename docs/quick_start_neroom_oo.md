
## 快速开始（使用服务封装）

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
	"github.com/netease-im/yunxin-server-sdk-golang/src/neroom"
	"github.com/netease-im/yunxin-server-sdk-golang/src/neroom/user"
)

func main() {
	// 初始化
	appkey := "xx"
	appsecret := "xx"
	timeoutMillis := 5000 * time.Millisecond

	// 创建客户端
	client := core.NewYunxinApiHttpClientBuilder(base.BizNEROOM, appkey, appsecret).
		TimeoutMillis(timeoutMillis).
		Build()

	// 创建服务聚合器
	neroomApiServices := neroom.NewYunxinNeroomApiServices(client)
	neroomUserService := neroomApiServices.GetNeroomUserService()

	// 创建请求
	userUuid := strings.ReplaceAll(uuid.New().String(), "-", "")
	request := &user.CreateNeroomAccountRequest{
		UserUuid: userUuid,
	}

	result, _ := neroomUserService.CreateAccount(request)
	if result.IsSuccess() {
		responseJson, _ := json.Marshal(result.GetResponse())
		fmt.Println(string(responseJson))
	} else {
		resultJson, _ := json.Marshal(result)
		fmt.Println(string(resultJson))
	}
}
```
