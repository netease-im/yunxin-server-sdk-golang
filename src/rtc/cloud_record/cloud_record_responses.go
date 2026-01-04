package cloud_record

// CloudRecordCreateResponse 创建云端录制任务响应
type CloudRecordCreateResponse struct {
	RequestId string `json:"requestId"` // 请求ID
	Errmsg    string `json:"errmsg"`    // 错误信息
	Record    struct {
		TaskId string `json:"taskId"` // 任务ID
		Cid    int64  `json:"cid"`    // 房间ID
	} `json:"record"`
}

// CloudRecordUpdateLayoutResponse 更新录制布局响应
type CloudRecordUpdateLayoutResponse struct {
	Code      int    `json:"code"`      // 状态码
	RequestId string `json:"requestId"` // 请求ID
	Errmsg    string `json:"errmsg"`    // 错误信息
	Record    struct {
		TaskId string `json:"taskId"` // 任务ID
		Cid    int64  `json:"cid"`    // 房间ID
	} `json:"record"`
}

// CloudRecordQueryTaskResponse 查询云端录制任务响应
type CloudRecordQueryTaskResponse struct {
	Record struct {
		Cid    int64        `json:"cid"`    // 房间ID
		Record []RecordInfo `json:"record"` // 录制信息
	} `json:"record"` // 录制信息
	Code      int    `json:"code"`      // 状态码
	RequestId string `json:"requestId"` // 请求ID
	Errmsg    string `json:"errmsg"`    // 错误信息
}

// RecordInfo 录制信息
type RecordInfo struct {
	State           int              `json:"state"`                  // 录制状态: 0-未开始, 1-录制中, 2-已结束, 3-异常
	TaskId          string           `json:"taskId"`                 // 任务ID
	StreamSubscribe *StreamSubscribe `json:"streamSubscribe"`        // 订阅信息
	RecordConfig    *RecordConfig    `json:"recordConfig,omitempty"` // 录制配置
	LayoutConfig    *LayoutConfig    `json:"layoutConfig,omitempty"` // 布局配置
	Detect          *Detect          `json:"detect,omitempty"`       // 内容安全审核配置
	Watermark       *Watermark       `json:"watermark,omitempty"`    // 水印配置
}

// CloudRecordUpdateSubscriptionResponse 更新订阅名单响应
type CloudRecordUpdateSubscriptionResponse struct {
	Code      int    `json:"code"`      // 状态码
	RequestId string `json:"requestId"` // 请求ID
	Errmsg    string `json:"errmsg"`    // 错误信息
}

// CloudRecordStopResponse 停止录制任务响应
type CloudRecordStopResponse struct {
	Code      int    `json:"code"`      // 状态码
	RequestId string `json:"requestId"` // 请求ID
	Errmsg    string `json:"errmsg"`    // 错误信息
}
