package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/base"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/endpoint"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/http"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/metrics"
	"github.com/sirupsen/logrus"
)

func main() {
	// 初始化配置
	appkey := "b882128a2a30430fd3269b2a272565df"
	appsecret := "4eb5f9177a7b"
	timeoutMillis := 5000

	fmt.Println("=== 云信 SDK core功能测试 ===")

	// 测试1: 自定义重试策略
	testCustomRetryPolicy(appkey, appsecret, timeoutMillis)

	// 测试2: 自定义端点选择器
	testCustomEndpointSelector(appkey, appsecret, timeoutMillis)

	// 测试3: 启用指标监控
	testMetricsCallback(appkey, appsecret, timeoutMillis)

	// 测试4: 自定义固定端点
	testCustomEndpoint(appkey, appsecret, timeoutMillis)

	// 测试5: 设置区域
	testRegionSetting(appkey, appsecret, timeoutMillis)

	// 测试6: 使用本地端点获取器
	testLocalEndpointFetcher(appkey, appsecret, timeoutMillis)

	fmt.Println("\n=== 所有测试完成 ===")
}

// testCustomRetryPolicy 测试自定义重试策略
func testCustomRetryPolicy(appkey, appsecret string, timeoutMillis int) {
	fmt.Println("\n========================================")
	fmt.Println("测试1: 自定义重试策略")
	fmt.Println("========================================")

	// RetryPolicy 可以自定义
	retryPolicy := base.NewDefaultRetryPolicy(1, true)

	client := core.NewYunxinApiHttpClientBuilder(base.BizIM, appkey, appsecret).
		TimeoutMillis(timeoutMillis).
		RetryPolicy(retryPolicy).
		Build()

	fmt.Printf("✓ 客户端创建成功，配置: maxRetry=%d, retryOn502=%v\n", 1, true)

	client.Shutdown()
}

// testCustomEndpointSelector 测试自定义端点选择器
func testCustomEndpointSelector(appkey, appsecret string, timeoutMillis int) {
	fmt.Println("\n========================================")
	fmt.Println("测试2: 自定义端点选择器")
	fmt.Println("========================================")

	// EndpointSelector 可以自定义
	endpointFetcher := endpoint.NewDynamicEndpointFetcher(base.BizIM, appkey, endpoint.CN)
	endpointSelector := endpoint.NewDynamicEndpointSelector(endpointFetcher)

	client := core.NewYunxinApiHttpClientBuilder(base.BizIM, appkey, appsecret).
		TimeoutMillis(timeoutMillis).
		EndpointSelector(endpointSelector).
		Build()

	fmt.Println("✓ 客户端创建成功，使用动态端点选择器")

	client.Shutdown()
}

// customMetricsCallback 自定义指标回调实现
type customMetricsCallback struct{}

func (c *customMetricsCallback) OnMetrics(stats *metrics.Stats) {
	fmt.Println("✓ 收到指标回调")
	jsonData, err := json.MarshalIndent(stats, "", "  ")
	if err != nil {
		logrus.Errorf("JSON序列化失败: %v", err)
	} else {
		fmt.Printf("指标数据: %s\n", string(jsonData))
	}
}

// testMetricsCallback 测试启用指标监控
func testMetricsCallback(appkey, appsecret string, timeoutMillis int) {
	fmt.Println("\n========================================")
	fmt.Println("测试3: 启用指标监控")
	fmt.Println("========================================")

	metricsCallback := &customMetricsCallback{}

	client := core.NewYunxinApiHttpClientBuilder(base.BizIM, appkey, appsecret).
		TimeoutMillis(timeoutMillis).
		MetricEnable(true).
		MetricCollectIntervalSeconds(10). // 默认 60s
		MetricsCallback(metricsCallback). // 设置回调
		Build()

	uri := "/im/v2/accounts"
	queryString := map[string]string{
		"account_ids": "zhagnsan,lisi",
	}
	_, _ = client.ExecuteV2Api(http.GET, uri, nil, queryString, "")

	// 可以通过 GetStats 方法获取数据
	stats := client.GetStats()
	jsonData, _ := json.MarshalIndent(stats, "", "  ")
	fmt.Printf("✓ 手动获取指标数据: %s\n", string(jsonData))

	// 转换为 Prometheus 格式
	prometheusFormat := metrics.Convert(stats)
	fmt.Printf("✓ Prometheus 格式:\n%s\n", prometheusFormat)

	// 等待回调被调用
	time.Sleep(time.Duration(20) * time.Second)

	client.Shutdown()
}

// testCustomEndpoint 测试自定义固定端点
func testCustomEndpoint(appkey, appsecret string, timeoutMillis int) {
	fmt.Println("\n========================================")
	fmt.Println("测试4: 自定义固定端点")
	fmt.Println("========================================")

	customEndpoint := "https://xxxx.com"

	client := core.NewYunxinApiHttpClientBuilder(base.BizIM, appkey, appsecret).
		TimeoutMillis(timeoutMillis).
		Endpoint(customEndpoint).
		Build()

	fmt.Printf("✓ 客户端创建成功，使用自定义端点: %s\n", customEndpoint)

	client.Shutdown()
}

// testRegionSetting 测试设置区域
func testRegionSetting(appkey, appsecret string, timeoutMillis int) {
	fmt.Println("\n========================================")
	fmt.Println("测试5: 设置区域")
	fmt.Println("========================================")

	client := core.NewYunxinApiHttpClientBuilder(base.BizIM, appkey, appsecret).
		TimeoutMillis(timeoutMillis).
		Region(endpoint.SG). // 限制调度服务域名的地区，默认可以不填
		Build()

	fmt.Println("✓ 客户端创建成功，区域设置为: SG（新加坡）")

	client.Shutdown()
}

// testLocalEndpointFetcher 测试使用本地端点获取器
func testLocalEndpointFetcher(appkey, appsecret string, timeoutMillis int) {
	fmt.Println("\n========================================")
	fmt.Println("测试6: 使用本地端点获取器")
	fmt.Println("========================================")

	endpointMain := "https://xxxx-01.com"
	endpointBackup := "https://xxxx-02.com"

	endpointFetcher := endpoint.NewLocalEndpointFetcher(endpointMain, endpointBackup)
	endpointSelector := endpoint.NewDynamicEndpointSelector(endpointFetcher)

	client := core.NewYunxinApiHttpClientBuilder(base.BizIM, appkey, appsecret).
		TimeoutMillis(timeoutMillis).
		EndpointSelector(endpointSelector).
		Build()

	fmt.Printf("✓ 客户端创建成功，使用本地端点: 主=%s, 备=%s\n",
		endpointMain, endpointBackup)

	client.Shutdown()
}
