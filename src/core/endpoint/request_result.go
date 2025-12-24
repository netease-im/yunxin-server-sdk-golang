package endpoint

// RequestResult 请求结果枚举
type RequestResult int

const (
	SUCCESS            RequestResult = iota // 成功
	CONNECT_TIMEOUT                         // 连接超时
	READ_WRITE_TIMEOUT                      // 请求读写超时
	HTTP_CODE_502                           // http错误码502
	HTTP_CODE_500                           // http错误码500
	HTTP_CODE_400                           // http错误码400
	HTTP_CODE_NOT_200                       // http非200错误码
	OTHER_ERRORS                            // 其他错误
)

// String 返回枚举的字符串表示
func (r RequestResult) String() string {
	switch r {
	case SUCCESS:
		return "SUCCESS"
	case CONNECT_TIMEOUT:
		return "CONNECT_TIMEOUT"
	case READ_WRITE_TIMEOUT:
		return "READ_WRITE_TIMEOUT"
	case HTTP_CODE_502:
		return "HTTP_CODE_502"
	case HTTP_CODE_500:
		return "HTTP_CODE_500"
	case HTTP_CODE_400:
		return "HTTP_CODE_400"
	case HTTP_CODE_NOT_200:
		return "HTTP_CODE_NOT_200"
	case OTHER_ERRORS:
		return "OTHER_ERRORS"
	default:
		return "UNKNOWN"
	}
}

func GetRequestResultByHttpCode(httpCode int) RequestResult {
	switch httpCode {
	case 200:
		return SUCCESS
	case 502:
		return HTTP_CODE_502
	case 500:
		return HTTP_CODE_500
	case 400:
		return HTTP_CODE_400
	default:
		return HTTP_CODE_NOT_200
	}
}
