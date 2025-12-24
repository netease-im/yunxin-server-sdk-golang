package endpoint

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core/base"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/version"
)

// DynamicEndpointFetcher 动态端点获取器
type DynamicEndpointFetcher struct {
	biz                   base.BizName
	appkey                string
	lbsList               []string
	reloadIntervalSeconds int
	detectPath            string

	httpClient    *http.Client
	md5           string
	endpoints     *Endpoints
	nextFetchTime int64

	mu       sync.RWMutex
	stopCh   chan struct{}
	stopOnce sync.Once
}

// NewDynamicEndpointFetcher 创建指定业务的动态端点获取器
func NewDynamicEndpointFetcher(biz base.BizName, appkey string, region Region) *DynamicEndpointFetcher {
	fetcher := &DynamicEndpointFetcher{
		biz:                   biz,
		appkey:                appkey,
		detectPath:            biz.DetectPath,
		reloadIntervalSeconds: base.Endpoint.ScheduleFetchIntervalSeconds,
		stopCh:                make(chan struct{}),
	}

	// 根据地区设置LBS列表
	switch region {
	case CN:
		fetcher.lbsList = []string{base.Endpoint.LBS.CNLBS}
	case SG:
		fetcher.lbsList = []string{base.Endpoint.LBS.SGLBS}
	default:
		fetcher.lbsList = []string{
			base.Endpoint.LBS.DefaultLBS,
			base.Endpoint.LBS.CNLBS,
			base.Endpoint.LBS.SGLBS,
		}
	}

	return fetcher
}

// NewDynamicEndpointFetcherWithCustom 创建自定义配置的动态端点获取器
func NewDynamicEndpointFetcherWithCustom(biz base.BizName, appkey string, lbsList []string, reloadIntervalSeconds int) *DynamicEndpointFetcher {
	return &DynamicEndpointFetcher{
		biz:                   biz,
		appkey:                appkey,
		detectPath:            biz.DetectPath,
		lbsList:               lbsList,
		reloadIntervalSeconds: reloadIntervalSeconds,
		stopCh:                make(chan struct{}),
	}
}

// Shutdown 关闭获取器 - 实现EndpointFetcher接口
func (f *DynamicEndpointFetcher) Shutdown() {
	f.Stop()
}

// Init 初始化端点获取器
func (f *DynamicEndpointFetcher) Init(httpClient *http.Client) {
	f.httpClient = httpClient

	// 初始化时多试几次
	for i := 0; i < 3; i++ {
		for _, lbs := range f.lbsList {
			if err := f.reload(lbs); err == nil {
				break
			}
		}
		if f.endpoints != nil {
			break
		}
	}

	if f.endpoints == nil {
		panic("init endpoints error")
	}

	// 启动定时刷新
	go f.startScheduleReload()
}

func (f *DynamicEndpointFetcher) Get() *Endpoints {
	return f.endpoints
}

// Stop 停止定时刷新
func (f *DynamicEndpointFetcher) Stop() {
	f.stopOnce.Do(func() {
		close(f.stopCh)
	})
}

// startScheduleReload 启动定时刷新
func (f *DynamicEndpointFetcher) startScheduleReload() {
	ticker := time.NewTicker(time.Duration(f.reloadIntervalSeconds) * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			for _, lbs := range f.lbsList {
				if err := f.reload(lbs); err == nil {
					break
				}
			}
		case <-f.stopCh:
			return
		}
	}
}

// reload 重新加载端点
func (f *DynamicEndpointFetcher) reload(lbs string) error {
	if time.Now().UnixMilli() < f.nextFetchTime {
		return nil
	}

	// 构建请求URL
	params := fmt.Sprintf("k=%s&sv=%s&biz=%d&p=go", f.appkey, version.YunxinApiSdkVersion, f.biz.Value)
	if f.md5 != "" {
		params += "&md5=" + f.md5
	}

	url := lbs + "?" + params
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	resp, err := f.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("http.code=%d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return err
	}

	code, ok := result["code"].(float64)
	if !ok {
		return fmt.Errorf("illegal response: %s", string(body))
	}

	if code == 304 { // 没有发生变更
		return nil
	}

	if code != 200 {
		return fmt.Errorf("fetch endpoints error, response: %s", string(body))
	}

	// 解析endpoints
	data, ok := result["data"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("illegal data format")
	}

	defaultEndpoint := ""
	if val, ok := data["default.endpoint"].(string); ok {
		if f.check(val) {
			defaultEndpoint = val
		}
	}

	var backupEndpoints []string
	if backupArray, ok := data["backup.endpoints"].([]interface{}); ok {
		for _, endpoint := range backupArray {
			if endpointStr, ok := endpoint.(string); ok {
				if f.check(endpointStr) {
					if defaultEndpoint == "" {
						defaultEndpoint = endpointStr
					} else {
						backupEndpoints = append(backupEndpoints, endpointStr)
					}
				}
			}
		}
	}

	if defaultEndpoint == "" {
		return fmt.Errorf("no valid endpoints")
	}

	// 设置下次获取时间
	ttl := 30
	if ttlVal, ok := data["ttl"].(float64); ok {
		if ttlInt := int(ttlVal); ttlInt > 0 && ttlInt <= 86400 {
			ttl = ttlInt
		}
	}

	f.mu.Lock()
	f.nextFetchTime = time.Now().UnixMilli() + int64(ttl*1000)
	if md5Val, ok := data["md5"].(string); ok {
		f.md5 = md5Val
	}
	f.endpoints = &Endpoints{
		DefaultEndpoint: defaultEndpoint,
		BackupEndpoints: backupEndpoints,
	}
	f.mu.Unlock()

	return nil
}

// check 检查端点是否可用
func (f *DynamicEndpointFetcher) check(endpoint string) bool {
	if f.httpClient == nil {
		return true
	}

	url := endpoint
	if f.detectPath != "" {
		url = endpoint + f.detectPath
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return false
	}

	resp, err := f.httpClient.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode == 200
}
