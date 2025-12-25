package livewallsolution

// CreateTaskResponse 创建安全通审核任务响应
type CreateTaskResponse struct {
	Result CreateTaskResult `json:"result"` // 接口响应结果,通常包含业务相关的响应结果
}

// CreateTaskResult result字段说明
type CreateTaskResult struct {
	TaskId    string               `json:"taskId"`              // 本次请求数据标识,可以根据该标识查询数据最新结果
	Status    bool                 `json:"status"`              // 请求结果,true:提交成功,false:提交失败
	DataId    string               `json:"dataId"`              // 网易云信为您自动生成的数据唯一标识
	Evidences []CreateTaskEvidence `json:"evidences,omitempty"` // 机审证据信息
}

// CreateTaskEvidence evidences结构
type CreateTaskEvidence struct {
	BeginTime int64  `json:"beginTime"` // 视频直播当前时间点,单位为毫秒
	EndTime   int64  `json:"endTime"`   // 视频直播当前时间点,单位为毫秒
	Type      int    `json:"type"`      // 违规数据类型,1:图片,2:视频
	Url       string `json:"url"`       // 证据信息
	SpeakerId string `json:"speakerId"` // 针对接入SDK监听客户,返回说话人ID
	FrontPics string `json:"frontPics"` // 违规前截图信息
}

// StopTaskResponse 停止安全通审核任务响应
type StopTaskResponse struct {
	Result []StopTaskResult `json:"result"` // 接口返回结果,通常包含业务相关的返回结果
}

// StopTaskResult 停止任务结果
type StopTaskResult struct {
	TaskId      string `json:"taskId"`      // 实时音视频安全通审核ID,是其唯一标识
	Result      int    `json:"result"`      // 请求结果,0:成功,1:失败,2:数据不存在
	ChannelName string `json:"channelName"` // 停止实时音视频安全通审核任务的房间名称
}

// QueryImageResponse 查询审核视频截图响应
type QueryImageResponse struct {
	Result      QueryImageResult `json:"result"`      // 接口返回结果,通常包含业务相关的返回结果
	ChannelName string           `json:"channelName"` // 进行内容审核的音视频房间名称
}

// QueryImageResult result字段说明
type QueryImageResult struct {
	Status int            `json:"status"` // 数据状态,0:成功,20:taskId过期,30:taskId不存在
	Images QueryImageData `json:"images"` // 截图分页数据
}

// QueryImageData images结构
type QueryImageData struct {
	Count int              `json:"count"` // 查询截图张数总量
	Rows  []ImageResultRow `json:"rows"`  // 分页数据
}

// ImageResultRow rows结构
type ImageResultRow struct {
	Url            string `json:"url"`            // 截图地址
	Label          int    `json:"label"`          // 分类信息,100:色情,200:广告,400:违禁,500:涉政,600:谩骂,700:灌水
	LabelLevel     int    `json:"labelLevel"`     // 分类级别,1:不确定,2:确定
	BeginTime      int64  `json:"beginTime"`      // 截图开始时间
	EndTime        int64  `json:"endTime"`        // 截图结束时间
	Uid            int64  `json:"uid"`            // 此违规信息对应的用户ID
	CallbackStatus int    `json:"callbackStatus"` // 回调状态,0:不需要回调,1:等待回调,2:已回调,3:回调失败
}

// QueryAudioTaskResponse 查询审核音频断句响应
type QueryAudioTaskResponse struct {
	Result      []AudioResult `json:"result"`      // 接口返回结果,通常包含业务相关的返回结果
	ChannelName string        `json:"channelName"` // 进行内容审核的音视频房间名称
}

// AudioResult result字段说明
type AudioResult struct {
	SegmentId    string         `json:"segmentId"`    // 直播断播唯一id
	Id           int64          `json:"id"`           // 唯一序号,固定顺序为0,无需关注
	TaskId       string         `json:"taskId"`       // 音视频安全通审核任务id,是其唯一标识
	Action       int            `json:"action"`       // 检测结果,0:通过,1:嫌疑,2:不通过
	AsrStatus    int            `json:"asrStatus"`    // 音频翻译结果,2:检测中,3:检测完成,4:检测失败
	AsrResult    int            `json:"asrResult"`    // 音频翻译检测关闭原因,1:文件格式问题,2:文件下载失败,3:解析失败,4:音频文件不存在
	Callback     string         `json:"callback"`     // 调用方反全通审核接口时,传递的callback字段数据
	CensorSource int            `json:"censorSource"` // 审核来源,0:独立人审,1:客户人审,2:智能机审
	Uid          int64          `json:"uid"`          // 此违规信息对应的用户ID
	StartTime    int64          `json:"startTime"`    // 开始时间(毫秒级)
	EndTime      int64          `json:"endTime"`      // 结束时间(毫秒级)
	Segments     []AudioSegment `json:"segments"`     // 证据信息,涉过的数据为空
	Records      []AudioRecord  `json:"records"`      // 录制信息
}

// AudioSegment segments结构
type AudioSegment struct {
	Label     int      `json:"label"`     // 分类信息,100:色情,200:广告,300:禁毒,400:违禁,500:涉政,600:谩骂
	Level     int      `json:"level"`     // 分类级别,0:正常,1:不确定,2:确定
	Evidence  string   `json:"evidence"`  // 命中的检测违禁词关键词的部分
	SubLabels []string `json:"subLabels"` // 分类信息
}

// AudioRecord records结构
type AudioRecord struct {
	StartTime int64  `json:"startTime"` // 开始时间(毫秒级)
	EndTime   int64  `json:"endTime"`   // 结束时间(毫秒级)
	Url       string `json:"url"`       // 录制链接
}
