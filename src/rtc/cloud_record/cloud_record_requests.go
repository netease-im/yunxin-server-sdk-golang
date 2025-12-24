package cloud_record

// CloudRecordCreateRequestV2 V2版本创建云端录制任务请求
// See https://doc.yunxin.163.com/nertc/server-apis/TEyOTk2OTY?platform=server
type CloudRecordCreateRequestV2 struct {
	Cid          int64         `json:"cid"`                    // 房间ID
	Uid          uint64        `json:"uid,omitempty"`          // 录制机器人UID
	RecordConfig *RecordConfig `json:"recordConfig,omitempty"` // 录制配置
	Detect       *Detect       `json:"detect,omitempty"`       // 内容安全审核配置
	LayoutConfig *LayoutConfig `json:"layoutConfig,omitempty"` // 录制布局配置
	Watermark    *Watermark    `json:"watermark,omitempty"`    // 水印配置
}

// CloudRecordCreateRequestV3 V3版本创建云端录制任务请求
// See https://doc.yunxin.163.com/nertc/server-apis/TEyOTk2OTY?platform=server
type CloudRecordCreateRequestV3 struct {
	Cname        string        `json:"cname"`                  // 房间名称
	Uid          uint64        `json:"uid,omitempty"`          // 录制机器人UID
	RecordConfig *RecordConfig `json:"recordConfig,omitempty"` // 录制配置
	Detect       *Detect       `json:"detect,omitempty"`       // 内容安全审核配置
	LayoutConfig *LayoutConfig `json:"layoutConfig,omitempty"` // 录制布局配置
	Watermark    *Watermark    `json:"watermark,omitempty"`    // 水印配置
}

// RecordConfig 录制配置
type RecordConfig struct {
	RecordType   int         `json:"recordType,omitempty"`   // 录制类型: 0-音频, 1-视频, 2-音视频, 100-模板文件集合
	AudioProfile int         `json:"audioProfile,omitempty"` // 音频质量: 0-标清, 1-高清, 2-超清
	VideoProfile int         `json:"videoProfile,omitempty"` // 视频质量: 0-标清, 1-高清, 2-超清, 3-2K, 4-4K
	FileFormat   int         `json:"fileFormat,omitempty"`   // 文件格式: 0-MP4, 1-FLV, 2-MP3
	ModeList     []ModeItem  `json:"modeList,omitempty"`     // 录制模式列表(仅recordType=100时使用)
	Subs         []SubStream `json:"subs,omitempty"`         // 订阅名单
}

// ModeItem 录制模式项(用于模板文件集合模式)
type ModeItem struct {
	Mode       int     `json:"mode"`                 // 流模式: 0-音视频, 1-音频, 2-视频
	LayoutType int     `json:"layoutType,omitempty"` // 布局类型
	Layout     *Layout `json:"layout,omitempty"`     // 布局配置
}

// SubStream 订阅流
type SubStream struct {
	Uid        uint64 `json:"uid"`                  // 用户ID
	StreamType int    `json:"streamType,omitempty"` // 流类型: 0-音视频, 1-辅流
}

// Detect 内容安全审核配置
type Detect struct {
	Audio       bool `json:"audio,omitempty"`       // 是否开启音频检测
	Video       bool `json:"video,omitempty"`       // 是否开启视频检测
	ScFrequency int  `json:"scFrequency,omitempty"` // 安全审核截图检测频率(秒)，取值范围1-60
	DetectType  int  `json:"detectType,omitempty"`  // 检测类型: 0-视频与音频同时检测, 1-仅检测视频
}

// LayoutConfig 录制布局配置
type LayoutConfig struct {
	LayoutType int     `json:"layoutType,omitempty"` // 布局类型: 0-画廊, 1-主讲人, 2-自定义, 3-悬浮, 4-垂直
	Layout     *Layout `json:"layout,omitempty"`     // 布局详情
}

// Layout 布局详情
type Layout struct {
	Canvas     *Canvas       `json:"canvas,omitempty"`     // 画布
	SubStreams []LayoutUser  `json:"subStreams,omitempty"` // 子流列表
	Users      []LayoutUser  `json:"users,omitempty"`      // 用户列表
	Images     []LayoutImage `json:"images,omitempty"`     // 图片列表
}

// Canvas 画布
type Canvas struct {
	Width      int    `json:"width,omitempty"`      // 宽度
	Height     int    `json:"height,omitempty"`     // 高度
	Color      string `json:"color,omitempty"`      // 背景色
	BackgroudX int    `json:"backgroudX,omitempty"` // 背景图片x坐标
	BackgroudY int    `json:"backgroudY,omitempty"` // 背景图片y坐标
	ImageUrl   string `json:"imageUrl,omitempty"`   // 背景图片URL
}

// LayoutUser 布局用户
type LayoutUser struct {
	Uid        uint64 `json:"uid,omitempty"`        // 用户ID
	StreamType int    `json:"streamType,omitempty"` // 流类型: 0-音视频, 1-辅流
	X          int    `json:"x,omitempty"`          // x坐标
	Y          int    `json:"y,omitempty"`          // y坐标
	Width      int    `json:"width,omitempty"`      // 宽度
	Height     int    `json:"height,omitempty"`     // 高度
	Zorder     int    `json:"zorder,omitempty"`     // 层级
	Quality    int    `json:"quality,omitempty"`    // 画质: 0-清晰度优先, 1-流畅度优先
}

// LayoutImage 布局图片
type LayoutImage struct {
	Url    string `json:"url,omitempty"`    // 图片URL
	X      int    `json:"x,omitempty"`      // x坐标
	Y      int    `json:"y,omitempty"`      // y坐标
	Width  int    `json:"width,omitempty"`  // 宽度
	Height int    `json:"height,omitempty"` // 高度
}

// Watermark 水印配置
type Watermark struct {
	ImgWms []ImageWatermark `json:"imgWms,omitempty"` // 图片水印列表
}

// ImageWatermark 图片水印
type ImageWatermark struct {
	Url    string `json:"url,omitempty"`    // 水印图片URL
	X      int    `json:"x,omitempty"`      // x坐标
	Y      int    `json:"y,omitempty"`      // y坐标
	Width  int    `json:"width,omitempty"`  // 宽度
	Height int    `json:"height,omitempty"` // 高度
}

// CloudRecordUpdateLayoutRequestV2 V2版本更新录制布局请求
// See https://doc.yunxin.163.com/nertc/server-apis/DUyOTg1NDI?platform=server
type CloudRecordUpdateLayoutRequestV2 struct {
	Cid          int64         `json:"cid"`                    // 房间ID
	TaskId       string        `json:"taskId"`                 // 任务ID
	LayoutConfig *LayoutConfig `json:"layoutConfig,omitempty"` // 录制布局配置
}

// CloudRecordUpdateLayoutRequestV3 V3版本更新录制布局请求
// See https://doc.yunxin.163.com/nertc/server-apis/DUyOTg1NDI?platform=server
type CloudRecordUpdateLayoutRequestV3 struct {
	Cname        string        `json:"cname"`                  // 房间名称
	TaskId       string        `json:"taskId"`                 // 任务ID
	LayoutConfig *LayoutConfig `json:"layoutConfig,omitempty"` // 录制布局配置
}

// CloudRecordQueryTaskRequestV2 V2版本查询云端录制任务请求
// See https://doc.yunxin.163.com/nertc/server-apis/TU4ODczNDk?platform=server
type CloudRecordQueryTaskRequestV2 struct {
	Cid    int64  `json:"cid"`              // 房间ID
	TaskId string `json:"taskId,omitempty"` // 任务ID，不指定则查询所有进行中的任务
}

// CloudRecordQueryTaskRequestV3 V3版本查询云端录制任务请求
// See https://doc.yunxin.163.com/nertc/server-apis/TU4ODczNDk?platform=server
type CloudRecordQueryTaskRequestV3 struct {
	Cname  string `json:"cname"`            // 房间名称
	TaskId string `json:"taskId,omitempty"` // 任务ID，不指定则查询所有进行中的任务
}

// CloudRecordUpdateSubscriptionRequestV2 V2版本更新订阅名单请求
// See https://doc.yunxin.163.com/nertc/server-apis/Tg5OTI0ODA?platform=server
type CloudRecordUpdateSubscriptionRequestV2 struct {
	Cid    int64       `json:"cid"`            // 房间ID
	TaskId string      `json:"taskId"`         // 任务ID
	Subs   []SubStream `json:"subs,omitempty"` // 订阅名单
	Mode   int         `json:"mode,omitempty"` // 订阅模式: 0-白名单, 1-黑名单
}

// CloudRecordUpdateSubscriptionRequestV3 V3版本更新订阅名单请求
// See https://doc.yunxin.163.com/nertc/server-apis/Tg5OTI0ODA?platform=server
type CloudRecordUpdateSubscriptionRequestV3 struct {
	Cname  string      `json:"cname"`          // 房间名称
	TaskId string      `json:"taskId"`         // 任务ID
	Subs   []SubStream `json:"subs,omitempty"` // 订阅名单
	Mode   int         `json:"mode,omitempty"` // 订阅模式: 0-白名单, 1-黑名单
}

// CloudRecordStopRequestV2 V2版本停止录制任务请求
// See https://doc.yunxin.163.com/nertc/server-apis/zIwNDM4NDg?platform=server
type CloudRecordStopRequestV2 struct {
	Cid    int64  `json:"cid"`    // 房间ID
	TaskId string `json:"taskId"` // 任务ID
}

// CloudRecordStopRequestV3 V3版本停止录制任务请求
// See https://doc.yunxin.163.com/nertc/server-apis/zIwNDM4NDg?platform=server
type CloudRecordStopRequestV3 struct {
	Cname  string `json:"cname"`  // 房间名称
	TaskId string `json:"taskId"` // 任务ID
}
