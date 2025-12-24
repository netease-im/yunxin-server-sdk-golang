package task

// TaskCreateResponse 创建云端播放任务响应
type TaskCreateResponse struct {
	Cid       int64  `json:"cid"`       // 房间ID
	TaskId    string `json:"taskId"`    // 任务唯一标识符
	Code      int    `json:"code"`      // 状态码
	RequestId string `json:"requestId"` // 请求ID
	Errmsg    string `json:"errmsg"`    // 错误信息
}

// TaskUpdateResponse 更新云端播放任务响应
type TaskUpdateResponse struct {
	TaskId    string `json:"taskId"`    // 任务唯一标识符
	Code      int    `json:"code"`      // 状态码
	RequestId string `json:"requestId"` // 请求ID
	Errmsg    string `json:"errmsg"`    // 错误信息
}

// TaskQueryResponse 查询云端播放任务响应
type TaskQueryResponse struct {
	Tasks     []*TaskInfo `json:"tasks"`     // 任务列表
	Code      int         `json:"code"`      // 状态码
	RequestId string      `json:"requestId"` // 请求ID
	Errmsg    string      `json:"errmsg"`    // 错误信息
}

// TaskInfo 任务信息
type TaskInfo struct {
	TaskId     string    `json:"taskId"`     // 任务ID
	Status     int       `json:"status"`     // 任务状态: 1-运行中, 2-已结束, 3-已销毁
	Data       *TaskData `json:"data"`       // 云端播放任务信息
	Cid        int64     `json:"cid"`        // 房间ID
	Cname      string    `json:"cname"`      // 房间名称
	CreateTime int64     `json:"createTime"` // 创建时间戳(毫秒)
	UpdateTime int64     `json:"updateTime"` // 更新时间戳(毫秒)
}

// TaskPauseResponse 暂停云端播放任务响应
type TaskPauseResponse struct {
	Code      int    `json:"code"`      // 状态码
	RequestId string `json:"requestId"` // 请求ID
	Errmsg    string `json:"errmsg"`    // 错误信息
}

// TaskResumeResponse 恢复云端播放任务响应
type TaskResumeResponse struct {
	Code      int    `json:"code"`      // 状态码
	RequestId string `json:"requestId"` // 请求ID
	Errmsg    string `json:"errmsg"`    // 错误信息
}

// TaskDestroyResponse 销毁云端播放任务响应
type TaskDestroyResponse struct {
	Code      int    `json:"code"`      // 状态码
	RequestId string `json:"requestId"` // 请求ID
	Errmsg    string `json:"errmsg"`    // 错误信息
}
