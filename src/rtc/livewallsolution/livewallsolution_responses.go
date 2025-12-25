package livewallsolution

// CreateTaskResponse 创建安全通审核任务响应
type CreateTaskResponse struct {
	TaskId string `json:"taskId,omitempty"` // 审核任务id,用于后续审核任务的详情查询
	Status bool   `json:"status,omitempty"` // 任务创建状态
	DataId string `json:"dataId,omitempty"` // 网易云信为您自动生成的数据唯一标识
}

// StopTaskResponse 停止安全通审核任务响应
type StopTaskResponse struct {
	Result []StopTaskResult `json:"result,omitempty"` // 接口返回结果列表
}

// StopTaskResult 停止任务结果
type StopTaskResult struct {
	TaskId      string `json:"taskId,omitempty"`      // 实时音视频安全通审核ID
	Result      int    `json:"result,omitempty"`      // 请求结果,0:成功,1:失败,2:数据不存在
	ChannelName string `json:"channelName,omitempty"` // 停止实时音视频安全通审核任务的房间名称
}

// QueryImageResponse 查询审核视频截图响应
type QueryImageResponse struct {
	ChannelName string        `json:"channelName,omitempty"` // 房间名称
	Result      []ImageResult `json:"result,omitempty"`      // 视频截图信息列表
}

// ImageResult 视频截图信息
type ImageResult struct {
	Id           int64     `json:"id,omitempty"`           // 截图ID
	SegmentId    string    `json:"segmentId,omitempty"`    // 截图证据ID
	TaskId       string    `json:"taskId,omitempty"`       // 任务ID
	Uid          int64     `json:"uid,omitempty"`          // UID
	BeginTime    int64     `json:"beginTime,omitempty"`    // 截图开始时间
	EndTime      int64     `json:"endTime,omitempty"`      // 截图结束时间
	Type         int       `json:"type,omitempty"`         // 截图类型
	Url          string    `json:"url,omitempty"`          // 截图URL
	Callback     string    `json:"callback,omitempty"`     // 回调参数
	CensorSource int       `json:"censorSource,omitempty"` // 审核来源
	Segments     []Segment `json:"segments,omitempty"`     // 审核分类
	Records      []Record  `json:"records,omitempty"`      // 录制文件信息
}

// Segment 审核分类信息
type Segment struct {
	Label     int      `json:"label,omitempty"`     // 标签
	Level     int      `json:"level,omitempty"`     // 风险等级
	SubLabels []string `json:"subLabels,omitempty"` // 子标签列表
}

// Record 录制文件信息
type Record struct {
	StartTime int64  `json:"startTime,omitempty"` // 录制开始时间
	EndTime   int64  `json:"endTime,omitempty"`   // 录制结束时间
	Url       string `json:"url,omitempty"`       // 录制文件URL
}

// QueryAudioTaskResponse 查询审核音频断句响应
type QueryAudioTaskResponse struct {
	ChannelName string        `json:"channelName,omitempty"` // 房间名称
	Result      []AudioResult `json:"result,omitempty"`      // 音频断句信息列表
}

// AudioResult 音频断句信息
type AudioResult struct {
	Id           int64     `json:"id,omitempty"`           // 断句ID
	SegmentId    string    `json:"segmentId,omitempty"`    // 断句证据ID
	TaskId       string    `json:"taskId,omitempty"`       // 任务ID
	Uid          int64     `json:"uid,omitempty"`          // UID
	Action       int       `json:"action,omitempty"`       // 断句操作,0:未知,1:正常,2:违规
	AsrStatus    int       `json:"asrStatus,omitempty"`    // 语音识别状态,0:成功,1:失败,2:音频质量不足,3:无人声
	Content      string    `json:"content,omitempty"`      // 识别出的文本内容
	Callback     string    `json:"callback,omitempty"`     // 回调参数
	CensorSource int       `json:"censorSource,omitempty"` // 审核来源
	StartTime    int64     `json:"startTime,omitempty"`    // 断句开始时间
	EndTime      int64     `json:"endTime,omitempty"`      // 断句结束时间
	Segments     []Segment `json:"segments,omitempty"`     // 审核分类
	Records      []Record  `json:"records,omitempty"`      // 录制文件信息
}
