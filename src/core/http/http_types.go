package http

// HttpMethod HTTP方法枚举
type HttpMethod int

const (
	POST HttpMethod = iota
	GET
	PATCH
	PUT
	DELETE
)

// String 返回HTTP方法的字符串表示
func (h HttpMethod) String() string {
	switch h {
	case POST:
		return "POST"
	case GET:
		return "GET"
	case PATCH:
		return "PATCH"
	case PUT:
		return "PUT"
	case DELETE:
		return "DELETE"
	default:
		return "UNKNOWN"
	}
}

// ContextType 内容类型枚举
type ContextType int

const (
	FormUrlEncoded ContextType = iota // application/x-www-form-urlencoded;charset=UTF-8
	JSON                              // application/json;charset=UTF-8
)

// 为了兼容保留旧的常量名
var (
	FORM_URL_ENCODED = FormUrlEncoded
)

// GetValue 获取内容类型的值
func (c ContextType) GetValue() string {
	switch c {
	case FormUrlEncoded:
		return "application/x-www-form-urlencoded;charset=UTF-8"
	case JSON:
		return "application/json;charset=UTF-8"
	default:
		return ""
	}
}

// String 返回内容类型的字符串表示
func (c ContextType) String() string {
	switch c {
	case FormUrlEncoded:
		return "FORM_URL_ENCODED"
	case JSON:
		return "JSON"
	default:
		return "UNKNOWN"
	}
}

// HttpResponse HTTP响应
type HttpResponse struct {
	Endpoint string `json:"endpoint"`  // 端点
	HttpCode int    `json:"http_code"` // HTTP状态码
	Data     string `json:"data"`      // 响应数据
	TraceId  string `json:"trace_id"`  // 跟踪ID
}

// NewHttpResponse 创建新的HTTP响应
func NewHttpResponse(endpoint string, httpCode int, data string, traceId string) *HttpResponse {
	return &HttpResponse{
		Endpoint: endpoint,
		HttpCode: httpCode,
		Data:     data,
		TraceId:  traceId,
	}
}

// GetEndpoint 获取端点
func (h *HttpResponse) GetEndpoint() string {
	return h.Endpoint
}

// GetHttpCode 获取HTTP状态码
func (h *HttpResponse) GetHttpCode() int {
	return h.HttpCode
}

// GetData 获取响应数据
func (h *HttpResponse) GetData() string {
	return h.Data
}

// GetTraceId 获取跟踪ID
func (h *HttpResponse) GetTraceId() string {
	return h.TraceId
}
