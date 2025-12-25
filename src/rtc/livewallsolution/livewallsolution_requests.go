package livewallsolution

// CreateTaskRequest 创建安全通审核任务请求
type CreateTaskRequest struct {
	MonitorUid     int64          `json:"monitorUid"`               // 安全通虚拟用户加入房间时使用的uid,必选
	ChannelName    string         `json:"channelName"`              // 需要进行内容审核的音视频房间名称,必选
	DetectType     int            `json:"detectType,omitempty"`     // 安全通机器过检的类型,0:视频与音频同时检测,1:仅检测视频,2:仅检测音频,可选,默认0
	ScFrequency    int            `json:"scFrequency,omitempty"`    // 截图检测频率,取值范围1-600秒,可选,默认5秒
	SecretId       string         `json:"secretId"`                 // 安全通自定义产品的 secretId
	CallbackUrl    string         `json:"callbackUrl,omitempty"`    // 接收审核结果的回调地址,可选
	AutoMaskConfig AutoMaskConfig `json:"autoMaskConfig,omitempty"` // 自动打码配置,可选
	StopInSecond   int            `json:"stopInSecond,omitempty"`   // 设定一个时间值，该秒数后由网易云信自动停止安全通检测任务。取值范围为 [1, 259200]
}

// AutoMaskConfig 自动打码配置
type AutoMaskConfig struct {
	EnableMask bool    `json:"enableMask,omitempty"` // 是否开启自动打码,可选
	MaskType   int     `json:"maskType,omitempty"`   // 打码类型,0:黑屏,1:模糊,可选,默认0
	Duration   int     `json:"duration,omitempty"`   // 打码持续时间,单位秒,可选,默认5秒
	UnmaskUids []int64 `json:"unmaskUids,omitempty"` // 不打码的用户UID列表,可选
	MaskArea   int     `json:"maskArea,omitempty"`   // 打码区域,0:全部区域打码,1:精准区域打码,可选,默认0
}

// StopTaskRequest 停止安全通审核任务请求
type StopTaskRequest struct {
	RealTimeInfoList []RealTimeInfo `json:"realTimeInfoList"` // 直播实时信息列表,最多100条,必选
}

// RealTimeInfo 直播实时信息
type RealTimeInfo struct {
	ChannelName string `json:"channelName"` // 需要停止实时音视频安全通审核任务的房间名称,必选
	Status      int    `json:"status"`      // 检测状态,请设置为100表示停止检测,必选
}

// QueryImageRequest 查询审核视频截图请求
type QueryImageRequest struct {
	TaskId   string `json:"taskId,omitempty"`   // 音视频内容安全审核任务ID,是其唯一标识,可选
	Levels   string `json:"levels,omitempty"`   // String形式的数组,0:正常;1:嫌疑;2:确定,可选
	PageSize int    `json:"pageSize,omitempty"` // 每页查询数据条数,取值范围0~1000,可选
	PageNum  int    `json:"pageNum,omitempty"`  // 查询页码,可选
}

// QueryAudioTaskRequest 查询审核音频断句请求
type QueryAudioTaskRequest struct {
	TaskId      string `json:"taskId,omitempty"`      // 音视频内容审核任务ID,是其唯一标识,可选
	ChannelName string `json:"channelName,omitempty"` // 进行内容审核的音视频房间名称,可选
	StartTime   int64  `json:"startTime"`             // 断句开始时间,Unix时间戳,精确到毫秒,必选
	EndTime     int64  `json:"endTime"`               // 断句结束时间,必选
	Record      int    `json:"record,omitempty"`      // 是否返回录制文件地址,0:(默认)不返回,1:返回,可选
}
