package task

// TaskCreateResponse 创建云端播放任务响应
type TaskCreateResponse struct {
	Cid       int64    `json:"cid"`       // 房间ID
	Code      int      `json:"code"`      // 状态码
	RequestId string   `json:"requestId"` // 请求ID
	Errmsg    string   `json:"errmsg"`    // 错误信息
	Result    TaskInfo `json:"result"`    // 任务信息
}

// TaskUpdateResponse 更新云端播放任务响应
type TaskUpdateResponse struct {
	Code      int      `json:"code"`      // 状态码
	RequestId string   `json:"requestId"` // 请求ID
	Errmsg    string   `json:"errmsg"`    // 错误信息
	Result    TaskInfo `json:"result"`    // 任务信息
}

// TaskQueryResponse 查询云端播放任务响应
type TaskQueryResponse struct {
	Result    []*TaskInfo `json:"result"`    // 任务列表
	Code      int         `json:"code"`      // 状态码
	RequestId string      `json:"requestId"` // 请求ID
	Errmsg    string      `json:"errmsg"`    // 错误信息
}

// TaskInfo 任务信息
type TaskInfo struct {
	Cid    int64     `json:"cid,omitempty"`    // 房间ID
	TaskId string    `json:"taskId,omitempty"` // 任务ID
	Data   *TaskData `json:"data,omitempty"`   // 云端播放任务信息
}

// TaskPauseResponse 暂停云端播放任务响应
type TaskPauseResponse struct {
	Code      int      `json:"code"`      // 状态码
	RequestId string   `json:"requestId"` // 请求ID
	Errmsg    string   `json:"errmsg"`    // 错误信息
	Result    TaskInfo `json:"result"`    // 任务信息
}

// TaskResumeResponse 恢复云端播放任务响应
type TaskResumeResponse struct {
	Code      int      `json:"code"`      // 状态码
	RequestId string   `json:"requestId"` // 请求ID
	Errmsg    string   `json:"errmsg"`    // 错误信息
	Result    TaskInfo `json:"result"`    // 任务信息
}

// TaskDestroyResponse 销毁云端播放任务响应
type TaskDestroyResponse struct {
	Code      int    `json:"code"`      // 状态码
	RequestId string `json:"requestId"` // 请求ID
	Errmsg    string `json:"errmsg"`    // 错误信息
}
