

## 多域名切换机制

* 内置了多个api域名，并且会定时从云信服务器更新最新api域名列表
* sdk会定时探测各个api域名
* sdk会根据各个api域名的请求结果和探测结果来动态调度（选择延迟最低、成功率最高的api域名）
* 默认切换机制参考：`DynamicEndpointSelector`，你也可以自定义

```go
package main

import (
	"time"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/base"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/endpoint"
)

func main() {
	// 初始化
	appkey := "xx"
	appsecret := "xx"
	timeoutMillis := 5000 * time.Millisecond

	// EndpointSelector可以自定义
	endpointFetcher := endpoint.NewDynamicEndpointFetcher(base.BizIM, appkey)
	endpointSelector := endpoint.NewDynamicEndpointSelector(base.BizIM, endpointFetcher)

	// 创建客户端
	client := core.NewYunxinApiHttpClientBuilder(base.BizIM, appkey, appsecret).
		TimeoutMillis(timeoutMillis).
		EndpointSelector(endpointSelector).
		Build()

	// 使用client进行API调用
	_ = client
}
```
