package metrics

// MetricsCallback 指标回调接口
type MetricsCallback interface {
	// OnMetrics 当产生指标时的回调
	OnMetrics(stats *Stats)
}

// StatsData 统计数据
type StatsData struct {
	Count int64   `json:"count"` // 计数
	Avg   float64 `json:"avg"`   // 平均值
	Max   float64 `json:"max"`   // 最大值
	P50   float64 `json:"p50"`   // 50%分位数
	P75   float64 `json:"p75"`   // 75%分位数
	P90   float64 `json:"p90"`   // 90%分位数
	P95   float64 `json:"p95"`   // 95%分位数
	P99   float64 `json:"p99"`   // 99%分位数
	P999  float64 `json:"p999"`  // 99.9%分位数
}
