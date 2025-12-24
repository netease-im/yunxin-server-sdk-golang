package endpoint

import (
	"context"
	"net/http"
	"sort"
	"sync"
	"time"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core/base"
	"github.com/sirupsen/logrus"
)

// DynamicEndpointSelector 动态端点选择器
type DynamicEndpointSelector struct {
	fetcher                   EndpointFetcher
	scheduleDetectIntervalSec int
	scheduleResultIntervalSec int
	slideWindowBuckets        int
	slideWindowTime           int64

	defaultEndpoint  string
	backupEndpoints  []string
	detectPath       string
	orderedEndpoints []string

	ratioMap     map[string]*RatioSlideWindows
	latencyMap   map[string]*LatencySlideWindows
	ratioMutex   sync.RWMutex
	latencyMutex sync.RWMutex

	httpClient *http.Client

	scheduleResultFuture *time.Ticker
	scheduleDetectFuture *time.Ticker
	cancelFunc           context.CancelFunc
}

// NewDynamicEndpointSelector 创建一个新的动态端点选择器实例
func NewDynamicEndpointSelector(fetcher EndpointFetcher) *DynamicEndpointSelector {
	return &DynamicEndpointSelector{
		fetcher:                   fetcher,
		detectPath:                base.BizIM.DetectPath,
		scheduleDetectIntervalSec: base.Endpoint.ScheduleDetectIntervalSeconds,
		scheduleResultIntervalSec: base.Endpoint.ScheduleResultIntervalSeconds,
		slideWindowBuckets:        base.Endpoint.SlideWindowBuckets,
		slideWindowTime:           base.Endpoint.SlideWindowTime,
		ratioMap:                  make(map[string]*RatioSlideWindows),
		latencyMap:                make(map[string]*LatencySlideWindows),
		httpClient:                &http.Client{},
	}
}

// NewDynamicEndpointSelectorWithBizName 创建一个带有业务名称的动态端点选择器实例
func NewDynamicEndpointSelectorWithBizName(biz base.BizName, fetcher EndpointFetcher) *DynamicEndpointSelector {
	return &DynamicEndpointSelector{
		fetcher:                   fetcher,
		detectPath:                biz.DetectPath,
		scheduleDetectIntervalSec: base.Endpoint.ScheduleDetectIntervalSeconds,
		scheduleResultIntervalSec: base.Endpoint.ScheduleResultIntervalSeconds,
		slideWindowBuckets:        base.Endpoint.SlideWindowBuckets,
		slideWindowTime:           base.Endpoint.SlideWindowTime,
		ratioMap:                  make(map[string]*RatioSlideWindows),
		latencyMap:                make(map[string]*LatencySlideWindows),
		httpClient:                &http.Client{},
	}
}

// NewDynamicEndpointSelectorWithParams 创建一个带完整参数的动态端点选择器实例
func NewDynamicEndpointSelectorWithParams(
	fetcher EndpointFetcher,
	detectPath string,
	scheduleDetectIntervalSec int,
	scheduleResultIntervalSec int,
	slideWindowBuckets int,
	slideWindowTime int64) *DynamicEndpointSelector {

	return &DynamicEndpointSelector{
		fetcher:                   fetcher,
		detectPath:                detectPath,
		scheduleDetectIntervalSec: scheduleDetectIntervalSec,
		scheduleResultIntervalSec: scheduleResultIntervalSec,
		slideWindowBuckets:        slideWindowBuckets,
		slideWindowTime:           slideWindowTime,
		ratioMap:                  make(map[string]*RatioSlideWindows),
		latencyMap:                make(map[string]*LatencySlideWindows),
		httpClient:                &http.Client{},
	}
}

// NewDynamicEndpointSelectorWithBizAndParams 创建一个带业务名和完整参数的动态端点选择器实例
func NewDynamicEndpointSelectorWithBizAndParams(
	biz base.BizName,
	fetcher EndpointFetcher,
	scheduleDetectIntervalSec int,
	scheduleResultIntervalSec int,
	slideWindowBuckets int,
	slideWindowTime int64) *DynamicEndpointSelector {

	return &DynamicEndpointSelector{
		fetcher:                   fetcher,
		detectPath:                biz.DetectPath,
		scheduleDetectIntervalSec: scheduleDetectIntervalSec,
		scheduleResultIntervalSec: scheduleResultIntervalSec,
		slideWindowBuckets:        slideWindowBuckets,
		slideWindowTime:           slideWindowTime,
		ratioMap:                  make(map[string]*RatioSlideWindows),
		latencyMap:                make(map[string]*LatencySlideWindows),
		httpClient:                &http.Client{},
	}
}

// Init 初始化端点选择器
func (d *DynamicEndpointSelector) Init(httpClient *http.Client) {
	d.httpClient = httpClient
	d.fetcher.Init(httpClient)

	endpoints := d.fetcher.Get()
	if endpoints == nil {
		panic("endpoints fetch null")
	}

	d.defaultEndpoint = endpoints.DefaultEndpoint
	d.backupEndpoints = endpoints.BackupEndpoints

	if d.defaultEndpoint == "" {
		panic("default endpoint is null")
	}

	d.orderedEndpoints = append(d.orderedEndpoints, d.defaultEndpoint)
	if d.backupEndpoints != nil {
		d.orderedEndpoints = append(d.orderedEndpoints, d.backupEndpoints...)
	}

	ctx, cancel := context.WithCancel(context.Background())
	d.cancelFunc = cancel

	// 启动定时任务
	if d.detectPath != "" {
		go d.startScheduleDetect(ctx)
	}
	go d.startScheduleResult(ctx)
}

// startScheduleDetect 启动探测定时任务
func (d *DynamicEndpointSelector) startScheduleDetect(ctx context.Context) {
	ticker := time.NewTicker(time.Duration(d.scheduleDetectIntervalSec) * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			d.scheduleDetect()
		}
	}
}

// startScheduleResult 启动结果定时任务
func (d *DynamicEndpointSelector) startScheduleResult(ctx context.Context) {
	ticker := time.NewTicker(time.Duration(d.scheduleResultIntervalSec) * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			d.scheduleResult()
		}
	}
}

// Shutdown 关闭端点选择器
func (d *DynamicEndpointSelector) Shutdown() {
	if d.cancelFunc != nil {
		d.cancelFunc()
	}
}

// Update 更新端点请求结果
func (d *DynamicEndpointSelector) Update(endpoint string, result RequestResult) {
	windows := d.getRatioSlideWindows(endpoint)
	if result == SUCCESS {
		windows.IncrementSuccess()
	} else {
		windows.IncrementFail()
	}
}

// SelectEndpoint 选择端点,excludeEndpoints 为需要排除的端点集合
func (d *DynamicEndpointSelector) SelectEndpoint(excludeEndpoints map[string]bool) string {
	if d.backupEndpoints == nil || len(d.backupEndpoints) == 0 {
		return d.defaultEndpoint
	}

	defer func() {
		if r := recover(); r != nil {
			// 错误处理，返回默认端点
		}
	}()

	// 选择第一个不在排除列表中的端点
	for _, endpoint := range d.orderedEndpoints {
		if _, excluded := excludeEndpoints[endpoint]; !excluded {
			return endpoint
		}
	}

	// 如果所有端点都被排除,返回默认端点
	return d.defaultEndpoint
}

// scheduleResult 定时计算结果
func (d *DynamicEndpointSelector) scheduleResult() {
	defer func() {
		if r := recover(); r != nil {
			logrus.Errorf("schedule result error: %v", r)
		}
	}()

	d.fetchEndpoints()

	ratioSlideWindows := d.getRatioSlideWindows(d.defaultEndpoint)
	latencySlideWindows := d.getLatencySlideWindows(d.defaultEndpoint)

	list := make([]*Entry, 0)
	list = append(list, &Entry{
		Endpoint:        d.defaultEndpoint,
		Ratio:           ratioSlideWindows.GetSuccessRatio(),
		Latency:         latencySlideWindows.GetLatency(),
		DefaultEndpoint: true,
	})

	for _, backupEndpoint := range d.backupEndpoints {
		backUpRatioSlideWindows := d.getRatioSlideWindows(backupEndpoint)
		backUpLatencySlideWindows := d.getLatencySlideWindows(backupEndpoint)
		list = append(list, &Entry{
			Endpoint:        backupEndpoint,
			Ratio:           backUpRatioSlideWindows.GetSuccessRatio(),
			Latency:         backUpLatencySlideWindows.GetLatency(),
			DefaultEndpoint: false,
		})
	}

	sort.Sort(ByPriority(list))

	result := make([]string, 0, len(list))
	for _, entry := range list {
		result = append(result, entry.Endpoint)

		logrus.Debugf("endpoint = %s, default-endpoint = %t, ratio = %f, latency = %f",
			entry.Endpoint, entry.DefaultEndpoint, entry.Ratio, entry.Latency)
	}

	d.orderedEndpoints = result
}

// scheduleDetect 定时探测
func (d *DynamicEndpointSelector) scheduleDetect() {
	defer func() {
		if r := recover(); r != nil {
			logrus.Errorf("schedule detect error: %v", r)
		}
	}()

	d.fetchEndpoints()
	d.check(d.defaultEndpoint)

	for _, backupEndpoint := range d.backupEndpoints {
		d.check(backupEndpoint)
	}
}

// fetchEndpoints 获取端点列表
func (d *DynamicEndpointSelector) fetchEndpoints() {
	defer func() {
		if r := recover(); r != nil {
			logrus.Errorf("fetch error: %v", r)
		}
	}()

	endpoints := d.fetcher.Get()
	if endpoints == nil {
		logrus.Error("fetch endpoints error, endpoints is null")
		return
	}

	defaultEndpoint := endpoints.DefaultEndpoint
	if defaultEndpoint == "" {
		logrus.Error("fetch endpoints error, default endpoint is null")
		return
	}

	d.defaultEndpoint = defaultEndpoint
	d.backupEndpoints = endpoints.BackupEndpoints
}

// check 检查端点状态
func (d *DynamicEndpointSelector) check(endpoint string) {
	if d.httpClient == nil {
		return
	}

	url := endpoint + d.detectPath
	startTime := time.Now().UnixMilli()
	success := false

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logrus.Debugf("detect error, endpoint = %s, path = %s, err = %v", endpoint, d.detectPath, err)
		success = false
	} else {
		resp, err := d.httpClient.Do(req)
		if err != nil {
			logrus.Debugf("detect error, endpoint = %s, path = %s, err = %v", endpoint, d.detectPath, err)
			success = false
		} else {
			defer resp.Body.Close()

			if resp.StatusCode == 200 {
				success = true
			}

			logrus.Debugf("detect, endpoint = %s, path = %s, code = %d", endpoint, d.detectPath, resp.StatusCode)
		}
	}

	latencySlideWindows := d.getLatencySlideWindows(endpoint)
	latencySlideWindows.Increment(time.Now().UnixMilli() - startTime)

	ratioSlideWindows := d.getRatioSlideWindows(endpoint)
	if success {
		ratioSlideWindows.IncrementSuccess()
	} else {
		ratioSlideWindows.IncrementFail()
	}
}

// getRatioSlideWindows 获取成功率滑动窗口
func (d *DynamicEndpointSelector) getRatioSlideWindows(endpoint string) *RatioSlideWindows {
	d.ratioMutex.RLock()
	windows, exists := d.ratioMap[endpoint]
	d.ratioMutex.RUnlock()

	if !exists {
		d.ratioMutex.Lock()
		// 双重检查
		windows, exists = d.ratioMap[endpoint]
		if !exists {
			windows = NewRatioSlideWindows(d.slideWindowBuckets, d.slideWindowTime)
			d.ratioMap[endpoint] = windows
		}
		d.ratioMutex.Unlock()
	}

	return windows
}

// getLatencySlideWindows 获取延迟滑动窗口
func (d *DynamicEndpointSelector) getLatencySlideWindows(endpoint string) *LatencySlideWindows {
	d.latencyMutex.RLock()
	windows, exists := d.latencyMap[endpoint]
	d.latencyMutex.RUnlock()

	if !exists {
		d.latencyMutex.Lock()
		// 双重检查
		windows, exists = d.latencyMap[endpoint]
		if !exists {
			windows = NewLatencySlideWindows(d.slideWindowBuckets, d.slideWindowTime)
			d.latencyMap[endpoint] = windows
		}
		d.latencyMutex.Unlock()
	}

	return windows
}

// Entry 用于排序的实体
type Entry struct {
	Endpoint        string
	Ratio           float64
	Latency         float64
	DefaultEndpoint bool
}

// ByPriority 实现排序接口
type ByPriority []*Entry

func (a ByPriority) Len() int { return len(a) }

func (a ByPriority) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (a ByPriority) Less(i, j int) bool {
	if a[i].Ratio >= 1.0 && a[j].Ratio >= 1.0 {
		if a[i].Latency == a[j].Latency {
			if a[i].DefaultEndpoint {
				return true
			}
			return false
		}
		return a[i].Latency < a[j].Latency
	} else if a[i].Ratio < 1.0 && a[i].Ratio >= 0.999 && a[j].Ratio < 1.0 && a[j].Ratio >= 0.999 {
		return a[i].Latency < a[j].Latency
	} else {
		if a[j].Ratio != a[i].Ratio {
			return a[j].Ratio < a[i].Ratio
		}
		return a[i].Latency < a[j].Latency
	}
}
