package metrics

// Stats 统计信息
type Stats struct {
	BizName           string          `json:"biz_name"`
	EndpointStatsList []EndpointStats `json:"endpoint_stats_list"`
	UriStatsList      []UriStats      `json:"uri_stats_list"`
}

// GetBizName 获取业务名称
func (s *Stats) GetBizName() string {
	return s.BizName
}

// SetBizName 设置业务名称
func (s *Stats) SetBizName(bizName string) {
	s.BizName = bizName
}

// GetEndpointStatsList 获取端点统计列表
func (s *Stats) GetEndpointStatsList() []EndpointStats {
	return s.EndpointStatsList
}

// SetEndpointStatsList 设置端点统计列表
func (s *Stats) SetEndpointStatsList(endpointStatsList []EndpointStats) {
	s.EndpointStatsList = endpointStatsList
}

// GetUriStatsList 获取URI统计列表
func (s *Stats) GetUriStatsList() []UriStats {
	return s.UriStatsList
}

// SetUriStatsList 设置URI统计列表
func (s *Stats) SetUriStatsList(uriStatsList []UriStats) {
	s.UriStatsList = uriStatsList
}
