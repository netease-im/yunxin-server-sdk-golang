package core

// YunxinApiResponse 云信API响应
type YunxinApiResponse struct {
	Endpoint string `json:"endpoint"`  // 端点
	HttpCode int    `json:"http_code"` // HTTP状态码
	Data     string `json:"data"`      // 响应数据
	TraceId  string `json:"trace_id"`  // 跟踪ID
}

// NewYunxinApiResponse 创建新的云信API响应
func NewYunxinApiResponse(endpoint string, httpCode int, data string, traceId string) *YunxinApiResponse {
	return &YunxinApiResponse{
		Endpoint: endpoint,
		HttpCode: httpCode,
		Data:     data,
		TraceId:  traceId,
	}
}

// GetHttpCode 获取HTTP状态码
func (r *YunxinApiResponse) GetHttpCode() int {
	return r.HttpCode
}

// GetEndpoint 获取端点
func (r *YunxinApiResponse) GetEndpoint() string {
	return r.Endpoint
}

// GetData 获取响应数据
func (r *YunxinApiResponse) GetData() string {
	return r.Data
}

// GetTraceId 获取跟踪ID
func (r *YunxinApiResponse) GetTraceId() string {
	return r.TraceId
}
