

## metrics

* 默认开启，可以关闭
* 支持配置统计周期
* 有2个维度，一个是endpoint维度，一个是uri维度，统计了请求的成功/失败数量，耗时的平均、最大、分位数（p50/p75/p90/p99/p999）
* 支持以prometheus格式输出，也支持自定义输出
* 可以定时获取数据，也可以使用MetricsCallback主动接受数据推送
* 数据字段参考 `Stats` 结构体和 `PrometheusConverter` 相关代码

```go
package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/base"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/metrics"
)

func main() {
	// 初始化
	appkey := "xx"
	appsecret := "xx"
	timeoutMillis := 5000 * time.Millisecond

	// 定义metrics回调函数
	metricsCallback := func(stats *metrics.Stats) {
		fmt.Println("receive stats callback")
		jsonBytes, _ := json.MarshalIndent(stats, "", "  ")
		fmt.Println(string(jsonBytes))
	}

	// 创建客户端
	client := core.NewYunxinApiHttpClientBuilder(base.BizIM, appkey, appsecret).
		TimeoutMillis(timeoutMillis).
		MetricEnable(true). // 默认true
		MetricCollectIntervalSeconds(60). // 默认60s
		MetricsCallback(metricsCallback). // 设置回调
		Build()

	// 你可以60s调用一次GetStats方法获取数据，也可以用MetricsCallback获取数据

	// 获取统计数据，可以自定义格式输出到监控系统
	stats := client.GetStats()
	jsonBytes, _ := json.MarshalIndent(stats, "", "  ")
	fmt.Println(string(jsonBytes))

	// 转换为prometheus格式
	prometheusStr := metrics.ConvertToPrometheus(stats)
	fmt.Println(prometheusStr)
}
```

