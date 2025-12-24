package cloud_record

// CloudRecordCreateResponse 创建云端录制任务响应
type CloudRecordCreateResponse struct {
	TaskId    string `json:"taskId"`    // 任务ID
	Code      int    `json:"code"`      // 状态码
	RequestId string `json:"requestId"` // 请求ID
	Errmsg    string `json:"errmsg"`    // 错误信息
}

// CloudRecordUpdateLayoutResponse 更新录制布局响应
type CloudRecordUpdateLayoutResponse struct {
	Code      int    `json:"code"`      // 状态码
	RequestId string `json:"requestId"` // 请求ID
	Errmsg    string `json:"errmsg"`    // 错误信息
}

// CloudRecordQueryTaskResponse 查询云端录制任务响应
type CloudRecordQueryTaskResponse struct {
	Record    *RecordInfo `json:"record"`    // 录制信息
	Code      int         `json:"code"`      // 状态码
	RequestId string      `json:"requestId"` // 请求ID
	Errmsg    string      `json:"errmsg"`    // 错误信息
}

// RecordInfo 录制信息
type RecordInfo struct {
	Status       int           `json:"status"`                 // 录制状态: 0-未开始, 1-录制中, 2-已结束, 3-异常
	TaskId       string        `json:"taskId"`                 // 任务ID
	RecordId     string        `json:"recordId"`               // 录制ID
	RecordConfig *RecordConfig `json:"recordConfig,omitempty"` // 录制配置
	LayoutConfig *LayoutConfig `json:"layoutConfig,omitempty"` // 布局配置
	Detect       *Detect       `json:"detect,omitempty"`       // 内容安全审核配置
	Watermark    *Watermark    `json:"watermark,omitempty"`    // 水印配置
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
