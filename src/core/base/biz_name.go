package base

// BizName 业务
type BizName struct {
	Value              int
	Name               string
	HttpCodeAlways200  bool
	DefaultRetryPolicy RetryPolicy
	DetectPath         string
}

var (
	BizIM      = BizName{Value: 1, Name: "IM", HttpCodeAlways200: true, DefaultRetryPolicy: &retryOn502, DetectPath: "/health/liveness.action"}
	BizRTC     = BizName{Value: 2, Name: "RTC", HttpCodeAlways200: false, DefaultRetryPolicy: &notRetryOn502, DetectPath: "/index.html"}
	BizSMS     = BizName{Value: 3, Name: "SMS", HttpCodeAlways200: true, DefaultRetryPolicy: &notRetryOn502, DetectPath: "/health/liveness.action"}
	BizLIVE    = BizName{Value: 4, Name: "LIVE", HttpCodeAlways200: true, DefaultRetryPolicy: &retryOn502, DetectPath: "/health/status"}
	BizVOD     = BizName{Value: 5, Name: "VOD", HttpCodeAlways200: true, DefaultRetryPolicy: &retryOn502, DetectPath: "/health/status"}
	BizMEETING = BizName{Value: 6, Name: "MEETING", HttpCodeAlways200: true, DefaultRetryPolicy: &retryOn502, DetectPath: "/status.html"}
	BizNEROOM  = BizName{Value: 7, Name: "NEROOM", HttpCodeAlways200: true, DefaultRetryPolicy: &retryOn502, DetectPath: "/status.html"}
	BizCUSTOM  = BizName{Value: 1000, Name: "CUSTOM", HttpCodeAlways200: false, DefaultRetryPolicy: &notRetryOn502, DetectPath: ""}
)
