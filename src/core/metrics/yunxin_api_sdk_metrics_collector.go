package metrics

import (
	"context"
	"sync"
	"time"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core/http"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/trace"
)

// YunxinApiSdkMetricsCollector 云信API SDK指标收集器
type YunxinApiSdkMetricsCollector struct {
	bizName         string
	map1            sync.Map // Key1 -> *Statistics
	map2            sync.Map // Key2 -> *Statistics
	metricsCallback MetricsCallback
	stats           *Stats

	stopCh   chan struct{}
	stopOnce sync.Once

	mu sync.RWMutex
}

// Key1 URI级别的统计键
type Key1 struct {
	Endpoint    string           `json:"endpoint"`
	Method      http.HttpMethod  `json:"method"`
	ContextType http.ContextType `json:"context_type"`
	ApiVersion  trace.ApiVersion `json:"api_version"`
	Uri         string           `json:"uri"`
	Result      string           `json:"result"`
}

// Key2 端点级别的统计键
type Key2 struct {
	Endpoint string `json:"endpoint"`
	Result   string `json:"result"`
}

// NewYunxinApiSdkMetricsCollector 创建指标收集器
func NewYunxinApiSdkMetricsCollector(bizName string, collectIntervalSeconds int, metricsCallback MetricsCallback) *YunxinApiSdkMetricsCollector {
	if collectIntervalSeconds <= 0 {
		panic("illegal collectIntervalSeconds")
	}

	collector := &YunxinApiSdkMetricsCollector{
		bizName:         bizName,
		metricsCallback: metricsCallback,
		stats:           &Stats{},
		stopCh:          make(chan struct{}),
	}

	// 启动定时计算
	go collector.startCalculation(collectIntervalSeconds)

	return collector
}

// Collect 收集指标数据
func (c *YunxinApiSdkMetricsCollector) Collect(endpoint string, method http.HttpMethod, contextType http.ContextType, apiVersion trace.ApiVersion, uri string, result string, spendMs int64) {
	// URI级别统计
	key1 := Key1{
		Endpoint:    endpoint,
		Method:      method,
		ContextType: contextType,
		ApiVersion:  apiVersion,
		Uri:         uri,
		Result:      result,
	}

	if stats1, ok := c.map1.Load(key1); ok {
		stats1.(*Statistics).Update(spendMs)
	} else {
		newStats := NewStatistics()
		if actual, loaded := c.map1.LoadOrStore(key1, newStats); loaded {
			actual.(*Statistics).Update(spendMs)
		} else {
			newStats.Update(spendMs)
		}
	}

	// 端点级别统计
	key2 := Key2{
		Endpoint: endpoint,
		Result:   result,
	}

	if stats2, ok := c.map2.Load(key2); ok {
		stats2.(*Statistics).Update(spendMs)
	} else {
		newStats := NewStatistics()
		if actual, loaded := c.map2.LoadOrStore(key2, newStats); loaded {
			actual.(*Statistics).Update(spendMs)
		} else {
			newStats.Update(spendMs)
		}
	}
}

// Shutdown 关闭收集器
func (c *YunxinApiSdkMetricsCollector) Shutdown() {
	c.stopOnce.Do(func() {
		close(c.stopCh)
	})
}

// GetStats 获取统计信息
func (c *YunxinApiSdkMetricsCollector) GetStats() *Stats {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.stats
}

// startCalculation 启动定时计算
func (c *YunxinApiSdkMetricsCollector) startCalculation(intervalSeconds int) {
	ticker := time.NewTicker(time.Duration(intervalSeconds) * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			c.calculate()
		case <-c.stopCh:
			return
		}
	}
}

// calculate 执行指标计算
func (c *YunxinApiSdkMetricsCollector) calculate() {
	// 获取并重置map
	oldMap1 := c.swapMap1()
	oldMap2 := c.swapMap2()

	var endpointStatsList []EndpointStats
	var uriStatsList []UriStats

	// 处理URI级别统计
	oldMap1.Range(func(key, value interface{}) bool {
		k1 := key.(Key1)
		stats := value.(*Statistics)
		data := stats.GetStatsDataAndReset()

		if data.Count == 0 {
			return true
		}

		uriStats := UriStats{
			Endpoint:    k1.Endpoint,
			Method:      k1.Method,
			ContextType: k1.ContextType,
			ApiVersion:  k1.ApiVersion,
			Uri:         k1.Uri,
			Result:      k1.Result,
			Count:       data.Count,
			Avg:         data.Avg,
			Max:         data.Max,
			P50:         data.P50,
			P75:         data.P75,
			P90:         data.P90,
			P95:         data.P95,
			P99:         data.P99,
			P999:        data.P999,
		}
		uriStatsList = append(uriStatsList, uriStats)
		return true
	})

	// 处理端点级别统计
	oldMap2.Range(func(key, value interface{}) bool {
		k2 := key.(Key2)
		stats := value.(*Statistics)
		data := stats.GetStatsDataAndReset()

		if data.Count == 0 {
			return true
		}

		endpointStats := EndpointStats{
			Endpoint: k2.Endpoint,
			Result:   k2.Result,
			Count:    data.Count,
			Avg:      data.Avg,
			Max:      data.Max,
			P50:      data.P50,
			P75:      data.P75,
			P90:      data.P90,
			P95:      data.P95,
			P99:      data.P99,
			P999:     data.P999,
		}
		endpointStatsList = append(endpointStatsList, endpointStats)
		return true
	})

	// 更新统计信息
	newStats := &Stats{
		BizName:           c.bizName,
		UriStatsList:      uriStatsList,
		EndpointStatsList: endpointStatsList,
	}

	c.mu.Lock()
	c.stats = newStats
	c.mu.Unlock()

	// 执行回调
	if c.metricsCallback != nil {
		go c.executeCallback(newStats)
	}
}

// swapMap1 原子交换map1
func (c *YunxinApiSdkMetricsCollector) swapMap1() *sync.Map {
	oldMap := &sync.Map{}
	c.map1.Range(func(key, value interface{}) bool {
		oldMap.Store(key, value)
		c.map1.Delete(key)
		return true
	})
	return oldMap
}

// swapMap2 原子交换map2
func (c *YunxinApiSdkMetricsCollector) swapMap2() *sync.Map {
	oldMap := &sync.Map{}
	c.map2.Range(func(key, value interface{}) bool {
		oldMap.Store(key, value)
		c.map2.Delete(key)
		return true
	})
	return oldMap
}

// executeCallback 执行回调
func (c *YunxinApiSdkMetricsCollector) executeCallback(stats *Stats) {
	// 使用context控制超时
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	done := make(chan struct{})
	go func() {
		defer close(done)
		c.metricsCallback.OnMetrics(stats)
	}()

	select {
	case <-done:
		// 回调完成
	case <-ctx.Done():
		// 超时处理
	}
}
