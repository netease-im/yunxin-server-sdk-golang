package base

// Endpoint 端点相关常量
var Endpoint = struct {
	ScheduleDetectIntervalSeconds int
	ScheduleResultIntervalSeconds int
	ScheduleFetchIntervalSeconds  int
	SlideWindowBuckets            int
	SlideWindowTime               int64
	LBS                           LBSEndpoints
	IM                            IMEndpoints
	RTC                           RTCEndpoints
	SMS                           SMSEndpoints
	LIVE                          LIVEEndpoints
	VOD                           VODEndpoints
	MEETING                       MEETINGEndpoints
	NEROOM                        NERoomEndpoints
}{
	ScheduleDetectIntervalSeconds: 5,
	ScheduleResultIntervalSeconds: 10,
	ScheduleFetchIntervalSeconds:  60,
	SlideWindowBuckets:            12,
	SlideWindowTime:               5000,
}

var HttpConfig = struct {
	ConnectTimeoutMillis int
	ReadTimeoutMillis    int
	WriteTimeoutMillis   int
	MaxRequests          int
	MaxRequestsPerHost   int
	KeepAliveSeconds     int
	MaxIdleConnections   int
}{
	ConnectTimeoutMillis: 5000,
	ReadTimeoutMillis:    5000,
	WriteTimeoutMillis:   5000,
	MaxRequests:          4096,
	MaxRequestsPerHost:   2048,
	KeepAliveSeconds:     3,
	MaxIdleConnections:   512,
}

var MetricConfig = struct {
	Enable                 bool
	CollectIntervalSeconds int
}{
	Enable:                 true,
	CollectIntervalSeconds: 60,
}

// LBSEndpoints LBS端点配置
type LBSEndpoints struct {
	DefaultLBS string
	CNLBS      string
	SGLBS      string
}

// IMEndpoints IM端点配置
type IMEndpoints struct {
	CN CNIMEndpoints
	SG SGIMEndpoints
}

type CNIMEndpoints struct {
	Default string
	Backup  string
	Backup1 string
	Backup2 string
	Backup3 string
}

type SGIMEndpoints struct {
	Default string
	Backup  string
	Backup1 string
	Backup2 string
	Backup3 string
}

// RTCEndpoints RTC端点配置
type RTCEndpoints struct {
	CN CNRTCEndpoints
	SG SGRTCEndpoints
}

type CNRTCEndpoints struct {
	Default string
	Backup1 string
	Backup2 string
}

type SGRTCEndpoints struct {
	Default string
	Backup1 string
}

// SMSEndpoints SMS端点配置
type SMSEndpoints struct {
	CN CNSMSEndpoints
}

type CNSMSEndpoints struct {
	Default string
	Backup1 string
}

// LIVEEndpoints LIVE端点配置
type LIVEEndpoints struct {
	CN CNLIVEEndpoints
	SG SGLIVEEndpoints
}

type CNLIVEEndpoints struct {
	Default string
	Backup1 string
}

type SGLIVEEndpoints struct {
	Default string
	Backup1 string
}

// VODEndpoints VOD端点配置
type VODEndpoints struct {
	CN CNVODEndpoints
	SG SGVODEndpoints
}

type CNVODEndpoints struct {
	Default string
	Backup1 string
}

type SGVODEndpoints struct {
	Default string
	Backup1 string
}

// MEETINGEndpoints MEETING端点配置
type MEETINGEndpoints struct {
	CN CNMEETINGEndpoints
}

type CNMEETINGEndpoints struct {
	Default string
	Backup1 string
}

// NERoomEndpoints NEROOM端点配置
type NERoomEndpoints struct {
	CN CNNERoomEndpoints
	SG SGNERoomEndpoints
}

type CNNERoomEndpoints struct {
	Default string
	Backup1 string
}

type SGNERoomEndpoints struct {
	Default string
	Backup1 string
}

func init() {
	// 初始化LBS端点
	Endpoint.LBS = LBSEndpoints{
		DefaultLBS: "https://srv-sdk-lbs.yunxinapi.com/srv-sdk/allocate",
		CNLBS:      "https://srv-sdk-lbs-cn.yunxinapi.com/srv-sdk/allocate",
		SGLBS:      "https://srv-sdk-lbs-sg.yunxinapi.com/srv-sdk/allocate",
	}

	// 初始化IM端点
	Endpoint.IM = IMEndpoints{
		CN: CNIMEndpoints{
			Default: "https://api-cn.yunxinapi.com/nimserver",
			Backup:  "https://api-cn-bak.yunxinapi.com/nimserver",
			Backup1: "https://api-cn-01.yunxinapi.com/nimserver",
			Backup2: "https://api-cn-02.yunxinapi.com/nimserver",
			Backup3: "https://api-cn-03.yunxinapi.com/nimserver",
		},
		SG: SGIMEndpoints{
			Default: "https://api-sg.yunxinapi.com/nimserver",
			Backup:  "https://api-sg-bak.yunxinapi.com/nimserver",
			Backup1: "https://api-sg-01.yunxinapi.com/nimserver",
			Backup2: "https://api-sg-02.yunxinapi.com/nimserver",
			Backup3: "https://api-sg-03.yunxinapi.com/nimserver",
		},
	}

	// 初始化RTC端点
	Endpoint.RTC = RTCEndpoints{
		CN: CNRTCEndpoints{
			Default: "https://rtc.yunxinapi.com",
			Backup1: "https://rtc-pri.yunxinapi.com",
			Backup2: "https://rtc-bak.yunxinapi.com",
		},
		SG: SGRTCEndpoints{
			Default: "https://rtc-ap.yunxinapi.com",
			Backup1: "https://rtc-ap-bak.yunxinapi.com",
		},
	}

	// 初始化SMS端点
	Endpoint.SMS = SMSEndpoints{
		CN: CNSMSEndpoints{
			Default: "https://sms.yunxinapi.com/sms",
			Backup1: "https://sms-backup.yunxinapi.com/sms",
		},
	}

	// 初始化LIVE端点
	Endpoint.LIVE = LIVEEndpoints{
		CN: CNLIVEEndpoints{
			Default: "https://vcloud.yunxinapi.com",
			Backup1: "https://vcloud.163.com",
		},
		SG: SGLIVEEndpoints{
			Default: "https://vcloud-sea.yunxinapi.com",
			Backup1: "https://api-sea.yunxinvcloud.com",
		},
	}

	// 初始化VOD端点
	Endpoint.VOD = VODEndpoints{
		CN: CNVODEndpoints{
			Default: "https://vcloud.yunxinapi.com",
			Backup1: "https://vcloud.163.com",
		},
		SG: SGVODEndpoints{
			Default: "https://vcloud-sea.yunxinapi.com",
			Backup1: "https://api-sea.yunxinvcloud.com",
		},
	}

	// 初始化MEETING端点
	Endpoint.MEETING = MEETINGEndpoints{
		CN: CNMEETINGEndpoints{
			Default: "https://roomkit-alt1.yunxinapi.com",
			Backup1: "https://roomkit-alt2.yunxinapi.com",
		},
	}

	// 初始化NEROOM端点
	Endpoint.NEROOM = NERoomEndpoints{
		CN: CNNERoomEndpoints{
			Default: "https://roomkit-alt1.yunxinapi.com",
			Backup1: "https://roomkit-alt2.yunxinapi.com",
		},
		SG: SGNERoomEndpoints{
			Default: "https://roomkit-sg.yunxinapi.com",
			Backup1: "https://roomkit-sg-alt1.yunxinapi.com",
		},
	}
}
