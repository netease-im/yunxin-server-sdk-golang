package rtmp_task

// RtmpTaskCreateResponse 创建旁路推流任务响应
type RtmpTaskCreateResponse struct {
	Code      int    `json:"code"`      // 状态码
	RequestId string `json:"requestId"` // 请求ID
	Errmsg    string `json:"errmsg"`    // 错误信息
}

// RtmpTaskUpdateResponse 更新旁路推流任务响应
type RtmpTaskUpdateResponse struct {
	Code      int    `json:"code"`      // 状态码
	RequestId string `json:"requestId"` // 请求ID
	Errmsg    string `json:"errmsg"`    // 错误信息
}

// RtmpTaskQueryResponse 查询指定旁路推流任务响应
type RtmpTaskQueryResponse struct {
	Code      int            `json:"code"`               // 状态码
	RequestId string         `json:"requestId"`          // 请求ID
	Errmsg    string         `json:"errmsg"`             // 错误信息
	Status    string         `json:"status"`             // 任务状态
	RtmpTask  RtmpTaskDetail `json:"rtmpTask,omitempty"` // 推流任务详细信息
}

// RtmpTaskQueryAllResponse 查询所有旁路推流任务响应
type RtmpTaskQueryAllResponse struct {
	Code      int              `json:"code"`                // 状态码
	RequestId string           `json:"requestId"`           // 请求ID
	Errmsg    string           `json:"errmsg"`              // 错误信息
	RtmpTasks []RtmpTaskDetail `json:"rtmpTasks,omitempty"` // 推流任务列表
}

// RtmpTaskDeleteResponse 停止旁路推流任务响应
type RtmpTaskDeleteResponse struct {
	Code      int    `json:"code"`      // 状态码
	RequestId string `json:"requestId"` // 请求ID
	Errmsg    string `json:"errmsg"`    // 错误信息
}

// RtmpTaskDetail 推流任务详细信息
type RtmpTaskDetail struct {
	TaskId    string     `json:"taskId,omitempty"`    // 推流任务ID
	StreamUrl string     `json:"streamUrl,omitempty"` // 推流地址
	Layout    RtmpLayout `json:"layout,omitempty"`    // 互动直播中的布局相关参数
	Record    bool       `json:"record,omitempty"`    // 旁路推流是否需要进行音视频录制
	Version   int        `json:"version,omitempty"`   // 推流任务版本
	HostUid   uint64     `json:"hostUid,omitempty"`   // 主播的UID
	Config    RtmpConfig `json:"config,omitempty"`    // 音视频流配置
	ExtraInfo string     `json:"extraInfo,omitempty"` // 自定义的媒体补充增强信息
}
