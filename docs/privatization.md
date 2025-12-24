

## 私有化

* 私有化场景下，可以手动指定endpoint，此时多域名切换机制将失效

```go
package main

import (
	"time"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/base"
)

func main() {
	// 初始化
	appkey := "xx"
	appsecret := "xx"
	timeoutMillis := 5000 * time.Millisecond

	endpoint := "https://xxxx.com"

	// 创建客户端
	client := core.NewYunxinApiHttpClientBuilder(base.BizIM, appkey, appsecret).
		TimeoutMillis(timeoutMillis).
		Endpoint(endpoint).
		Build()

	// 使用client进行API调用
	_ = client
}
```

* 如果私有化场景下，也有多个域名，此时你可以这样配置

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

	endpoint := "https://xxxx-01.com"
	endpointBackUp := "https://xxxx-02.com"

	// 使用本地endpoint获取器
	endpointFetcher := endpoint.NewLocalEndpointFetcher(endpoint, endpointBackUp)
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
