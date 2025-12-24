package core

import (
	"fmt"
	"strings"
	"sync"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core/base"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/endpoint"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/http"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/metrics"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/trace"
)

// YunxinApiHttpClient 云信API HTTP客户端
type YunxinApiHttpClient interface {
	// ExecuteV1Api 执行V1 API请求 (POST方法，form_url_encoded)
	ExecuteV1Api(path string, queryParams map[string]string) (*YunxinApiResponse, error)

	// ExecuteV2Api 执行V2 API请求
	ExecuteV2Api(method http.HttpMethod, uri string, pathParams, queryParams map[string]string, data string) (*YunxinApiResponse, error)

	// Execute 执行API请求
	Execute(method http.HttpMethod, contextType http.ContextType, uri string, path string, queryParams map[string]string, data string) (*YunxinApiResponse, error)

	// ExecuteJson 执行JSON API请求 (content-type=application/json)
	ExecuteJson(method http.HttpMethod, uri string, pathParams, queryParams map[string]string, data string) (*YunxinApiResponse, error)

	// ExecuteForm 执行Form API请求 (content-type=application/x-www-form-urlencoded)
	ExecuteForm(method http.HttpMethod, uri string, pathParams, queryParams map[string]string, data string) (*YunxinApiResponse, error)

	// GetStats 获取统计信息
	GetStats() *metrics.Stats

	// Shutdown 关闭客户端
	Shutdown()
}

// yunxinApiHttpClientImpl 云信API HTTP客户端实现
type yunxinApiHttpClientImpl struct {
	biz          base.BizName
	httpClient   YunxinHttpClient
	duplicateKey string
}

var (
	clientMap sync.Map // 线程安全的map，key为string，value为*yunxinApiHttpClientImpl
)

// YunxinApiHttpClientBuilder 云信API HTTP客户端构建器
type YunxinApiHttpClientBuilder struct {
	biz              base.BizName
	appkey           string
	appsecret        string
	endpointConfig   *endpoint.EndpointConfig
	httpClientConfig *YunxinHttpClientConfig
	metricsConfig    *metrics.MetricsConfig
	region           endpoint.Region
}

// NewYunxinApiHttpClientBuilder 创建云信API HTTP客户端构建器
func NewYunxinApiHttpClientBuilder(biz base.BizName, appkey, appsecret string) *YunxinApiHttpClientBuilder {
	cacheKey := fmt.Sprintf("%s/%s/%d", appkey, appsecret, biz.Value)

	if biz.Value != base.BizCUSTOM.Value {
		if _, exists := clientMap.Load(cacheKey); exists {
			panic(fmt.Sprintf("YunxinApiHttpClient with appkey = [%s] and bizName = [%s] duplicate init", appkey, biz.Name))
		}
	}

	endpointConfig := &endpoint.EndpointConfig{
		RetryPolicy: biz.DefaultRetryPolicy,
	}

	return &YunxinApiHttpClientBuilder{
		biz:            biz,
		appkey:         appkey,
		appsecret:      appsecret,
		endpointConfig: endpointConfig,
		httpClientConfig: &YunxinHttpClientConfig{
			ConnectTimeoutMillis: base.HttpConfig.ConnectTimeoutMillis,
			ReadTimeoutMillis:    base.HttpConfig.ReadTimeoutMillis,
			WriteTimeoutMillis:   base.HttpConfig.WriteTimeoutMillis,
			MaxIdleConnections:   base.HttpConfig.MaxIdleConnections,
			KeepAliveDuration:    base.HttpConfig.KeepAliveSeconds,
		},
		metricsConfig: &metrics.MetricsConfig{
			Enable:                 base.MetricConfig.Enable,
			CollectIntervalSeconds: base.MetricConfig.CollectIntervalSeconds,
		},
	}
}

// RetryPolicy 设置重试策略
func (b *YunxinApiHttpClientBuilder) RetryPolicy(retryPolicy base.RetryPolicy) *YunxinApiHttpClientBuilder {
	if retryPolicy == nil {
		panic("retry policy null")
	}
	b.endpointConfig.RetryPolicy = retryPolicy
	return b
}

// Endpoint 设置固定端点
func (b *YunxinApiHttpClientBuilder) Endpoint(endpointUrl string) *YunxinApiHttpClientBuilder {
	if endpointUrl == "" {
		panic("endpoint null")
	}
	b.endpointConfig.EndpointSelector = endpoint.NewFixedEndpointSelector(endpointUrl)
	return b
}

// EndpointSelector 设置端点选择器
func (b *YunxinApiHttpClientBuilder) EndpointSelector(endpointSelector endpoint.EndpointSelector) *YunxinApiHttpClientBuilder {
	if endpointSelector == nil {
		panic("endpointSelector null")
	}
	b.endpointConfig.EndpointSelector = endpointSelector
	return b
}

// TimeoutMillis 设置超时时间
func (b *YunxinApiHttpClientBuilder) TimeoutMillis(timeoutMillis int) *YunxinApiHttpClientBuilder {
	if timeoutMillis <= 0 {
		panic("illegal timeoutMillis")
	}
	b.httpClientConfig.ConnectTimeoutMillis = timeoutMillis
	b.httpClientConfig.ReadTimeoutMillis = timeoutMillis
	b.httpClientConfig.WriteTimeoutMillis = timeoutMillis
	return b
}

// HttpClientConfig 设置HTTP客户端配置
func (b *YunxinApiHttpClientBuilder) HttpClientConfig(httpClientConfig *YunxinHttpClientConfig) *YunxinApiHttpClientBuilder {
	if httpClientConfig == nil {
		panic("httpClientConfig null")
	}
	b.httpClientConfig = httpClientConfig
	return b
}

// MetricEnable 设置是否启用指标
func (b *YunxinApiHttpClientBuilder) MetricEnable(enable bool) *YunxinApiHttpClientBuilder {
	b.metricsConfig.Enable = enable
	return b
}

// MetricCollectIntervalSeconds 设置指标收集间隔
func (b *YunxinApiHttpClientBuilder) MetricCollectIntervalSeconds(collectIntervalSeconds int) *YunxinApiHttpClientBuilder {
	if collectIntervalSeconds <= 0 {
		panic("illegal collectIntervalSeconds")
	}
	b.metricsConfig.CollectIntervalSeconds = collectIntervalSeconds
	return b
}

// MetricsCallback 设置指标回调
func (b *YunxinApiHttpClientBuilder) MetricsCallback(metricsCallback metrics.MetricsCallback) *YunxinApiHttpClientBuilder {
	b.metricsConfig.MetricsCallback = metricsCallback
	return b
}

// Region 设置区域
func (b *YunxinApiHttpClientBuilder) Region(region endpoint.Region) *YunxinApiHttpClientBuilder {
	b.region = region
	return b
}

// Build 构建云信API HTTP客户端
func (b *YunxinApiHttpClientBuilder) Build() YunxinApiHttpClient {
	cacheKey := fmt.Sprintf("%s/%s/%d", b.appkey, b.appsecret, b.biz.Value)

	if b.biz.Value != base.BizCUSTOM.Value {
		if _, exists := clientMap.Load(cacheKey); exists {
			panic(fmt.Sprintf("YunxinApiHttpClient with appkey = [%s] and bizName = [%s] duplicate init", b.appkey, b.biz.Name))
		}
	} else {
		if b.endpointConfig.EndpointSelector == nil {
			panic("bizName with CUSTOM must specify endpoints")
		}
	}

	if b.endpointConfig.EndpointSelector == nil {
		// 创建动态端点选择器
		fetcher := endpoint.NewDynamicEndpointFetcher(b.biz, b.appkey, b.region)
		b.endpointConfig.EndpointSelector = endpoint.NewDynamicEndpointSelectorWithBizName(b.biz, fetcher)
	}

	// 转换EndpointConfig为EndpointConfigHttp
	endpointConfigHttp := &endpoint.EndpointConfig{
		RetryPolicy:      b.endpointConfig.RetryPolicy,
		EndpointSelector: b.endpointConfig.EndpointSelector.(endpoint.EndpointSelector),
	}

	httpClient := NewYunxinHttpClient(b.biz, b.appkey, b.appsecret, endpointConfigHttp, b.httpClientConfig, b.metricsConfig)

	client := &yunxinApiHttpClientImpl{
		biz:          b.biz,
		httpClient:   httpClient,
		duplicateKey: cacheKey,
	}

	if b.biz != base.BizCUSTOM {
		clientMap.Store(cacheKey, client)
	}

	return client
}

// ExecuteV1Api 执行V1 API请求
func (c *yunxinApiHttpClientImpl) ExecuteV1Api(path string, queryParams map[string]string) (*YunxinApiResponse, error) {
	if c.biz != base.BizIM {
		return nil, fmt.Errorf("only support bizName = IM")
	}

	builder := http.NewParamBuilder()
	for key, value := range queryParams {
		builder.AddParam(key, value)
	}

	response, err := c.httpClient.Execute(http.POST, http.FormUrlEncoded, trace.V1, path, path, nil, builder.Build())
	if err != nil {
		return nil, err
	}

	return NewYunxinApiResponse(response.GetEndpoint(), response.GetHttpCode(), response.GetData(), response.GetTraceId()), nil
}

// ExecuteV2Api 执行V2 API请求
func (c *yunxinApiHttpClientImpl) ExecuteV2Api(method http.HttpMethod, uri string, pathParams, queryParams map[string]string, data string) (*YunxinApiResponse, error) {
	if c.biz != base.BizIM {
		return nil, fmt.Errorf("only support bizName = IM")
	}

	path := buildPath(uri, pathParams)
	response, err := c.httpClient.Execute(method, http.JSON, trace.V2, uri, path, queryParams, data)
	if err != nil {
		return nil, err
	}

	return NewYunxinApiResponse(response.GetEndpoint(), response.GetHttpCode(), response.GetData(), response.GetTraceId()), nil
}

// Execute 执行API请求
func (c *yunxinApiHttpClientImpl) Execute(method http.HttpMethod, contextType http.ContextType, uri string, path string, queryParams map[string]string, data string) (*YunxinApiResponse, error) {
	response, err := c.httpClient.Execute(method, contextType, trace.V1, uri, path, queryParams, data)
	if err != nil {
		return nil, err
	}

	return NewYunxinApiResponse(response.GetEndpoint(), response.GetHttpCode(), response.GetData(), response.GetTraceId()), nil
}

// ExecuteJson 执行JSON API请求
func (c *yunxinApiHttpClientImpl) ExecuteJson(method http.HttpMethod, uri string, pathParams, queryParams map[string]string, data string) (*YunxinApiResponse, error) {
	path := buildPath(uri, pathParams)
	response, err := c.httpClient.Execute(method, http.JSON, trace.V1, uri, path, queryParams, data)
	if err != nil {
		return nil, err
	}

	return NewYunxinApiResponse(response.GetEndpoint(), response.GetHttpCode(), response.GetData(), response.GetTraceId()), nil
}

// ExecuteForm 执行Form API请求
func (c *yunxinApiHttpClientImpl) ExecuteForm(method http.HttpMethod, uri string, pathParams, queryParams map[string]string, data string) (*YunxinApiResponse, error) {
	path := buildPath(uri, pathParams)
	response, err := c.httpClient.Execute(method, http.FormUrlEncoded, trace.V1, path, path, queryParams, data)
	if err != nil {
		return nil, err
	}

	return NewYunxinApiResponse(response.GetEndpoint(), response.GetHttpCode(), response.GetData(), response.GetTraceId()), nil
}

// 替换路径中的参数
func buildPath(uri string, pathParams map[string]string) string {
	if len(pathParams) == 0 {
		return uri
	}

	args := make([]string, 0, len(pathParams)*2)
	for k, v := range pathParams {
		args = append(args, "{"+k+"}", v)
	}

	return strings.NewReplacer(args...).Replace(uri)
}

// GetStats 获取统计信息
func (c *yunxinApiHttpClientImpl) GetStats() *metrics.Stats {
	return c.httpClient.GetStats()
}

// Shutdown 关闭客户端
func (c *yunxinApiHttpClientImpl) Shutdown() {
	c.httpClient.Close()

	clientMap.Delete(c.duplicateKey)
}
