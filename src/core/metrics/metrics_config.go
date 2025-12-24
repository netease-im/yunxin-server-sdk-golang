package metrics

// MetricsConfig 指标配置
type MetricsConfig struct {
	Enable                 bool            `json:"enable"`                   // 是否启用指标收集
	CollectIntervalSeconds int             `json:"collect_interval_seconds"` // 收集间隔(秒)
	MetricsCallback        MetricsCallback `json:"-"`                        // 指标回调函数
}
