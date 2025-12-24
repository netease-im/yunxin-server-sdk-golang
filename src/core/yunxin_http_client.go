package core

import (
	"fmt"
	"io"
	nethttp "net/http"
	neturl "net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/base"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/endpoint"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/http"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/metrics"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/trace"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/utils"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/version"
	"github.com/sirupsen/logrus"
)

// YunxinHttpClient 云信HTTP客户端接口
type YunxinHttpClient interface {
	// 执行HTTP请求
	Execute(method http.HttpMethod, contextType http.ContextType, apiVersion trace.ApiVersion, uri string, path string, queryParams map[string]string, data string) (*http.HttpResponse, error)

	// 获取统计信息
	GetStats() *metrics.Stats

	// 关闭客户端
	Close() error
}

// YunxinHttpClientConfig 云信HTTP客户端配置
type YunxinHttpClientConfig struct {
	ConnectTimeoutMillis int `json:"connect_timeout_millis"` // 连接超时时间(毫秒)
	ReadTimeoutMillis    int `json:"read_timeout_millis"`    // 读取超时时间(毫秒)
	WriteTimeoutMillis   int `json:"write_timeout_millis"`   // 写入超时时间(毫秒)
	MaxIdleConnections   int `json:"max_idle_connections"`   // 最大空闲连接数
	KeepAliveDuration    int `json:"keep_alive_duration"`    // 保持连接时间(秒)
}

// yunxinHttpClientImpl 云信HTTP客户端实现
type yunxinHttpClientImpl struct {
	biz              base.BizName
	appkey           string
	appsecret        string
	endpointConfig   *endpoint.EndpointConfig
	httpClientConfig *YunxinHttpClientConfig
	metricsConfig    *metrics.MetricsConfig
	httpClient       *nethttp.Client
	metricsCollector *metrics.YunxinApiSdkMetricsCollector
	running          bool
	runningMutex     sync.RWMutex
}

// NewYunxinHttpClient 创建云信HTTP客户端
func NewYunxinHttpClient(biz base.BizName, appkey, appsecret string, endpointConfig *endpoint.EndpointConfig, httpClientConfig *YunxinHttpClientConfig, metricsConfig *metrics.MetricsConfig) YunxinHttpClient {
	// 创建HTTP客户端
	client := &nethttp.Client{
		Timeout: time.Duration(httpClientConfig.ReadTimeoutMillis) * time.Millisecond,
		Transport: &nethttp.Transport{
			MaxIdleConns:        httpClientConfig.MaxIdleConnections,
			MaxIdleConnsPerHost: httpClientConfig.MaxIdleConnections,
			IdleConnTimeout:     time.Duration(httpClientConfig.KeepAliveDuration) * time.Second,
			DisableCompression:  false,
		},
	}

	yunxinHttpClient := &yunxinHttpClientImpl{
		biz:              biz,
		appkey:           appkey,
		appsecret:        appsecret,
		endpointConfig:   endpointConfig,
		httpClientConfig: httpClientConfig,
		metricsConfig:    metricsConfig,
		httpClient:       client,
		running:          true,
	}
	if metricsConfig != nil && metricsConfig.Enable {
		yunxinHttpClient.metricsCollector = metrics.NewYunxinApiSdkMetricsCollector(biz.Name,
			metricsConfig.CollectIntervalSeconds, metricsConfig.MetricsCallback)
	}
	yunxinHttpClient.endpointConfig.EndpointSelector.Init(client)

	return yunxinHttpClient
}

// Execute 执行HTTP请求
func (c *yunxinHttpClientImpl) Execute(method http.HttpMethod, contextType http.ContextType, apiVersion trace.ApiVersion, uri string, path string, queryParams map[string]string, data string) (*http.HttpResponse, error) {
	// 检查客户端是否已关闭
	if !c.isRunning() {
		return nil, fmt.Errorf("yunxin http client has been shutdown")
	}

	// 生成trace-id
	traceId := trace.Gen()
	// 选择端点
	yunxinEndpoint, err := c.selectNextEndpoint(nil)
	if err != nil {
		return nil, err
	}

	// 创建执行上下文
	executeContext := base.NewExecuteContext(c.biz, yunxinEndpoint, method, contextType, apiVersion, uri, path, queryParams, data, traceId)

	// 构建URL
	url := c.buildURL(path, queryParams)

	// 获取重试策略
	retryPolicy := c.endpointConfig.RetryPolicy
	maxRetry := retryPolicy.MaxRetry()
	if maxRetry <= 0 {
		maxRetry = 0
	}
	if maxRetry > 128 {
		maxRetry = 128
	}

	// 获取HTTP客户端
	httpClient := c.httpClient

	var lastError error
	var excludeEndpoints map[string]bool
	// 重试循环
	for i := 0; i <= maxRetry; i++ {
		// 构建完整URL
		fullURL := yunxinEndpoint + url

		// 构建HTTP请求
		req, err := c.buildRequest(method, contextType, fullURL, data)
		if err != nil {
			return nil, err
		}

		// 添加请求头
		c.addHeaders(req, apiVersion, traceId)

		startTime := time.Now()

		// 记录日志
		c.logDebug("execute", executeContext, queryParams, data)

		// 发送HTTP请求
		resp, err := httpClient.Do(req)
		if err != nil {
			// 处理网络错误
			requestResult := c.classifyError(err)
			c.updateEndpointSelector(yunxinEndpoint, requestResult)
			c.collectMetrics(yunxinEndpoint, method, contextType, apiVersion, uri, err.Error(), startTime)

			lastError = base.NewYunxinSdkError(c.biz, yunxinEndpoint, traceId, err)

			// 判断是否重试
			retryAction := retryPolicy.OnError(executeContext, i, err)
			if !retryAction.IsRetry() {
				return nil, lastError
			}

			// 重试间隔
			interval := retryPolicy.RetryInterval(executeContext, i)
			if interval > 0 {
				time.Sleep(time.Duration(interval) * time.Millisecond)
			}

			// 是否切换端点
			if retryAction.IsNextEndpoint() {
				if excludeEndpoints == nil {
					excludeEndpoints = make(map[string]bool)
				}
				excludeEndpoints[yunxinEndpoint] = true
				yunxinEndpoint, _ = c.selectNextEndpoint(excludeEndpoints)
			}
			continue
		}

		defer resp.Body.Close()

		// 读取响应体
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		httpCode := resp.StatusCode
		bodyString := string(body)

		// 检查HTTP状态码
		if httpCode != 200 {
			requestResult := c.classifyHttpError(httpCode)
			c.updateEndpointSelector(yunxinEndpoint, requestResult)
			c.collectMetrics(yunxinEndpoint, method, contextType, apiVersion, uri, fmt.Sprintf("http_code_%d", httpCode), startTime)
			if !c.biz.HttpCodeAlways200 {
				c.updateEndpointSelector(yunxinEndpoint, c.classifyHttpError(httpCode))
				return http.NewHttpResponse(yunxinEndpoint, httpCode, bodyString, traceId), nil
			}

			httpCodeErr := base.NewHttpCodeError(c.biz, yunxinEndpoint, httpCode, bodyString)
			lastError = httpCodeErr

			// 判断是否重试
			retryAction := retryPolicy.OnError(executeContext, i, httpCodeErr)
			if !retryAction.IsRetry() {
				return nil, lastError
			}

			// 重试间隔
			interval := retryPolicy.RetryInterval(executeContext, i)
			if interval > 0 {
				time.Sleep(time.Duration(interval) * time.Millisecond)
			}

			// 是否切换端点
			if retryAction.IsNextEndpoint() {
				if excludeEndpoints == nil {
					excludeEndpoints = make(map[string]bool)
				}
				excludeEndpoints[yunxinEndpoint] = true
				yunxinEndpoint, _ = c.selectNextEndpoint(excludeEndpoints)
			}
			continue
		}

		// 成功响应
		c.updateEndpointSelector(yunxinEndpoint, endpoint.SUCCESS)
		c.collectMetrics(yunxinEndpoint, method, contextType, apiVersion, uri, "success", startTime)

		return http.NewHttpResponse(yunxinEndpoint, httpCode, bodyString, traceId), nil
	}

	// 重试次数用完，返回最后一个错误
	if lastError != nil {
		return nil, lastError
	}

	return nil, fmt.Errorf("request failed after %d retries", maxRetry)
}

func (c *yunxinHttpClientImpl) GetStats() *metrics.Stats {
	if !c.running {
		return nil
	}

	if c.metricsCollector == nil {
		return nil
	}

	return c.metricsCollector.GetStats()
}

func (c *yunxinHttpClientImpl) Close() error {
	// 关闭HTTP客户端连接池
	c.httpClient.CloseIdleConnections()
	return nil
}

// 辅助方法实现

// isRunning 检查客户端是否运行中
func (c *yunxinHttpClientImpl) isRunning() bool {
	c.runningMutex.RLock()
	defer c.runningMutex.RUnlock()
	return c.running
}

// setRunning 设置运行状态
func (c *yunxinHttpClientImpl) setRunning(running bool) {
	c.runningMutex.Lock()
	defer c.runningMutex.Unlock()
	c.running = running
}

// selectNextEndpoint 选择下一个端点,排除指定的端点集合
func (c *yunxinHttpClientImpl) selectNextEndpoint(excludeEndpoints map[string]bool) (string, error) {
	if c.endpointConfig == nil || c.endpointConfig.EndpointSelector == nil {
		return "", fmt.Errorf("endpoint selector not configured")
	}
	yunxinEndpoint := c.endpointConfig.EndpointSelector.SelectEndpoint(excludeEndpoints)
	if yunxinEndpoint == "" {
		return "", fmt.Errorf("no alternative endpoint available")
	}
	return yunxinEndpoint, nil
}

// buildURL 构建URL
func (c *yunxinHttpClientImpl) buildURL(path string, queryParams map[string]string) string {
	if queryParams == nil || len(queryParams) == 0 {
		return path
	}

	// 构建查询字符串
	var params []string
	for key, value := range queryParams {
		params = append(params, key+"="+neturl.QueryEscape(value))
	}

	if len(params) > 0 {
		return path + "?" + strings.Join(params, "&")
	}
	return path
}

// getHttpClient 获取HTTP客户端（支持自定义超时）
func (c *yunxinHttpClientImpl) getHttpClient(timeoutMillis *int64) *nethttp.Client {
	if timeoutMillis == nil {
		return c.httpClient
	}

	// 创建一个新的客户端实例，使用自定义超时
	transport := c.httpClient.Transport.(*nethttp.Transport).Clone()

	return &nethttp.Client{
		Transport: transport,
		Timeout:   time.Duration(*timeoutMillis) * time.Millisecond,
	}
}

// addHeaders 添加请求头
func (c *yunxinHttpClientImpl) addHeaders(req *nethttp.Request, apiVersion trace.ApiVersion, traceId string) {
	curTime := strconv.FormatInt(time.Now().Unix(), 10)
	nonce := strings.ReplaceAll(uuid.New().String(), "-", "")

	// 基础认证头
	req.Header.Set("AppKey", c.appkey)
	req.Header.Set("CurTime", curTime)
	req.Header.Set("Nonce", nonce)
	req.Header.Set("CheckSum", utils.GetCheckSum(c.appsecret, nonce, curTime))

	// 根据业务类型和API版本设置trace-id头
	if c.biz.Value == base.BizIM.Value {
		if apiVersion == trace.V1 {
			req.Header.Set("RequestId", traceId)
		} else if apiVersion == trace.V2 {
			req.Header.Set("X-custom-traceid", traceId)
		}
	} else {
		req.Header.Set("X-custom-traceid", traceId)
	}

	// User-Agent
	req.Header.Set("User-Agent", "yunxin-server-sdk/"+version.YunxinApiSdkVersion)
}

// buildRequest 重新定义构建HTTP请求（增加ContextType支持）
func (c *yunxinHttpClientImpl) buildRequest(method http.HttpMethod, contextType http.ContextType, url string, data string) (*nethttp.Request, error) {
	var body io.Reader

	// 处理请求体
	if method == http.POST || method == http.PUT || method == http.PATCH {
		if data != "" {
			body = strings.NewReader(data)
		}
	}

	req, err := nethttp.NewRequest(method.String(), url, body)
	if err != nil {
		return nil, err
	}

	// 根据ContextType设置Content-Type
	if method == http.POST || method == http.PUT || method == http.PATCH {
		if data != "" {
			req.Header.Set("Content-Type", contextType.GetValue())
		}
	}

	return req, nil
}

// logDebug 记录调试日志
func (c *yunxinHttpClientImpl) logDebug(action string, executeContext base.ExecuteContext, queryParams map[string]string, data string) {
	logrus.Debugf("%s, bizName=%s, endpoint=%s, method=%s, contextType=%s, apiVersion=%s, uri=%s, path=%s, traceId=%s, queryParams=%v, data=%s",
		action, executeContext.Biz.Name, executeContext.Endpoint, executeContext.HttpMethod.String(),
		executeContext.ContextType.String(), executeContext.ApiVersion.String(), executeContext.Uri,
		executeContext.Path, executeContext.TraceId, queryParams, data)
}

// classifyError 分类错误
func (c *yunxinHttpClientImpl) classifyError(err error) endpoint.RequestResult {
	if utils.IsTimeoutErr(err) {
		return endpoint.READ_WRITE_TIMEOUT
	}
	if utils.IsConnectError(err) {
		return endpoint.CONNECT_TIMEOUT
	}
	return endpoint.OTHER_ERRORS
}

// classifyHttpError 分类HTTP错误
func (c *yunxinHttpClientImpl) classifyHttpError(httpCode int) endpoint.RequestResult {
	return endpoint.GetRequestResultByHttpCode(httpCode)
}

// updateEndpointSelector 更新端点选择器状态
func (c *yunxinHttpClientImpl) updateEndpointSelector(endpoint string, result endpoint.RequestResult) {
	if c.endpointConfig != nil && c.endpointConfig.EndpointSelector != nil {
		c.endpointConfig.EndpointSelector.Update(endpoint, result)
	}
}

// collectMetrics 收集度量数据
func (c *yunxinHttpClientImpl) collectMetrics(endpoint string, method http.HttpMethod, contextType http.ContextType,
	apiVersion trace.ApiVersion, uri string, result string, startTime time.Time) {
	if c.metricsCollector != nil {
		duration := time.Since(startTime).Milliseconds()
		c.metricsCollector.Collect(endpoint, method, contextType, apiVersion, uri, result, duration)
	}
}
