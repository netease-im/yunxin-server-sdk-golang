
## 重试机制

* 重试机制由三个部分组成，分别是最大重试次数（最多允许128次）、重试策略、重试间隔，由 `RetryPolicy` 接口定义
* 你可以自定义 `RetryPolicy` 实现，默认为 `DefaultRetryPolicy`，此时重试逻辑如下：
* 默认最大重试为：`1` 次，可以自定义最大重试次数，如果设置为 `0`，则表示不重试
* IM、LIVE、VOD、MEETING、NEROOM，以上BizName，默认重试策略为：`http.code=502` 或者 `连接超时` 时重试，否则不重试
* RTC、SMS、CUSTOM，以上BizName，默认重试策略为：`连接超时` 时重试，否则不重试
* 默认重试间隔为：第一次重试没有间隔，之后重试间隔依次为 `5, 10, 50, 100, 500, 1000, 5000` (单位:毫秒)，超过则固定为5000ms

### RetryPolicy接口定义

```go
package base

import "time"

// RetryPolicy 重试策略接口
type RetryPolicy interface {
	// MaxRetry 最大重试次数，0 代表不重试
	// 最大允许设置 128，如果超过 128 实际生效最多 128
	MaxRetry() int

	// RetryInterval 重试间隔
	// executeContext: 执行上下文
	// retry: 第几次重试，从0开始
	// 返回: 重试间隔时间
	RetryInterval(executeContext *ExecuteContext, retry int) time.Duration

	// OnError 是否要重试
	// executeContext: 执行上下文
	// retry: 第几次重试，从0开始
	// err: 错误信息
	// 返回: 重试操作
	OnError(executeContext *ExecuteContext, retry int, err error) RetryAction
}

// 默认重试间隔数组（单位：毫秒）
var IntervalArray = []int64{0, 5, 10, 50, 100, 500, 1000, 5000}
```

### 初始化时设置自定义RetryPolicy

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

	// RetryPolicy可以自定义
	maxRetry := 1
	retryOn502 := true
	retryPolicy := base.NewDefaultRetryPolicy(maxRetry, retryOn502)

	// 创建客户端
	client := core.NewYunxinApiHttpClientBuilder(base.BizIM, appkey, appsecret).
		TimeoutMillis(timeoutMillis).
		RetryPolicy(retryPolicy).
		Build()

	// 使用client进行API调用
	_ = client
}
```
