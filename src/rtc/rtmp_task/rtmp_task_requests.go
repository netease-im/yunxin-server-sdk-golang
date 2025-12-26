package rtmp_task

// RtmpTaskCreateRequestV2 V2版本创建旁路推流任务请求
// See https://doc.yunxin.163.com/nertc/server-apis/DM0MTg2NTM?platform=server
type RtmpTaskCreateRequestV2 struct {
	Cid       int64       `json:"cid"`                 // 房间ID
	TaskId    string      `json:"taskId"`              // 自定义的推流任务ID,请保证此ID唯一
	StreamUrl string      `json:"streamUrl"`           // 推流地址,例如 rtmp://test.url
	Layout    RtmpLayout  `json:"layout"`              // 互动直播中的布局相关参数
	Record    bool        `json:"record,omitempty"`    // 旁路推流是否需要进行音视频录制
	Version   int         `json:"version"`             // 推流任务版本,设置为1
	HostUid   uint64      `json:"hostUid,omitempty"`   // 主播的UID
	Config    *RtmpConfig `json:"config,omitempty"`    // 音视频流配置
	ExtraInfo string      `json:"extraInfo,omitempty"` // 自定义的媒体补充增强信息
}

// RtmpTaskCreateRequestV3 V3版本创建旁路推流任务请求
// See https://doc.yunxin.163.com/nertc/server-apis/DM0MTg2NTM?platform=server
type RtmpTaskCreateRequestV3 struct {
	Cname     string      `json:"cname"`               // 房间名称
	TaskId    string      `json:"taskId"`              // 自定义的推流任务ID,请保证此ID唯一
	StreamUrl string      `json:"streamUrl"`           // 推流地址,例如 rtmp://test.url
	Layout    RtmpLayout  `json:"layout"`              // 互动直播中的布局相关参数
	Record    bool        `json:"record,omitempty"`    // 旁路推流是否需要进行音视频录制
	Version   int         `json:"version"`             // 推流任务版本,设置为1
	HostUid   uint64      `json:"hostUid,omitempty"`   // 主播的UID
	Config    *RtmpConfig `json:"config,omitempty"`    // 音视频流配置
	ExtraInfo string      `json:"extraInfo,omitempty"` // 自定义的媒体补充增强信息
}

// RtmpLayout 互动直播中的布局相关参数
type RtmpLayout struct {
	Canvas     Canvas      `json:"canvas"`               // 背景画布配置
	Users      []User      `json:"users,omitempty"`      // 用户画面配置列表
	SubStreams []SubStream `json:"subStreams,omitempty"` // 辅流画面配置列表
	Images     []Image     `json:"images,omitempty"`     // 占位图片配置列表
}

// Canvas 背景画布配置
type Canvas struct {
	Width  int `json:"width"`           // 画布宽度,单位为像素
	Height int `json:"height"`          // 画布高度,单位为像素
	Color  int `json:"color,omitempty"` // 画布背景色,格式为RGB,默认为0(黑色)
}

// User 用户画面配置
type User struct {
	Uid       uint64 `json:"uid"`                 // 用户ID
	X         int    `json:"x"`                   // 用户画面左上角的横坐标,单位为像素
	Y         int    `json:"y"`                   // 用户画面左上角的纵坐标,单位为像素
	Width     int    `json:"width"`               // 用户画面宽度,单位为像素
	Height    int    `json:"height"`              // 用户画面高度,单位为像素
	Adaption  int    `json:"adaption,omitempty"`  // 视窗填充策略: 0-视频尺寸等比缩放,1-视频尺寸非等比缩放
	PushAudio bool   `json:"pushAudio,omitempty"` // 是否推送该用户的音频流,默认为true
	PushVideo bool   `json:"pushVideo,omitempty"` // 是否推送该用户的视频流,默认为true
	ZOrder    int    `json:"zOrder,omitempty"`    // 图层编号,取值范围为[0, 100]
}

// SubStream 辅流画面配置
type SubStream struct {
	Uid       uint64 `json:"uid"`                 // 用户ID
	X         int    `json:"x"`                   // 辅流画面左上角的横坐标,单位为像素
	Y         int    `json:"y"`                   // 辅流画面左上角的纵坐标,单位为像素
	Width     int    `json:"width"`               // 辅流画面宽度,单位为像素
	Height    int    `json:"height"`              // 辅流画面高度,单位为像素
	Adaption  int    `json:"adaption,omitempty"`  // 视窗填充策略: 0-视频尺寸等比缩放,1-视频尺寸非等比缩放
	PushVideo bool   `json:"pushVideo,omitempty"` // 是否推送该辅流,默认为true
	ZOrder    int    `json:"zOrder,omitempty"`    // 图层编号,取值范围为[0, 100]
}

// Image 占位图片配置
type Image struct {
	Url      string `json:"url"`                // 图片URL
	X        int    `json:"x"`                  // 图片左上角的横坐标,单位为像素
	Y        int    `json:"y"`                  // 图片左上角的纵坐标,单位为像素
	Width    int    `json:"width"`              // 图片宽度,单位为像素
	Height   int    `json:"height"`             // 图片高度,单位为像素
	Adaption int    `json:"adaption,omitempty"` // 视窗填充策略: 0-视频尺寸等比缩放,1-视频尺寸非等比缩放
	ZOrder   int    `json:"zOrder,omitempty"`   // 图层编号,取值范围为[0, 100]
}

// RtmpConfig 音视频流配置
type RtmpConfig struct {
	AudioBitrate int `json:"audioBitrate,omitempty"` // 音频推流码率,单位为kbps,取值范围为[10, 192]
	SampleRate   int `json:"sampleRate,omitempty"`   // 音频推流采样率,单位为Hz,默认为48000
	Channels     int `json:"channels,omitempty"`     // 音频推流声道数,默认为2(双声道)
	CodecProfile int `json:"codecProfile,omitempty"` // 音频编码规格: 0-LC-AAC,1-HE-AAC,2-HE-AACv2
	VideoBitrate int `json:"videoBitrate,omitempty"` // 视频推流码率,单位为kbps,取值范围为[10, 5000]
	VideoFps     int `json:"videoFps,omitempty"`     // 视频推流帧率,单位为fps,默认为15
	VideoWidth   int `json:"videoWidth,omitempty"`   // 视频推流宽度,单位为像素
	VideoHeight  int `json:"videoHeight,omitempty"`  // 视频推流高度,单位为像素
	Bitrate      int `json:"bitrate,omitempty"`      // 视频推流码率,单位为kbps(已废弃,请使用videoBitrate)
	Fps          int `json:"fps,omitempty"`          // 视频推流帧率,单位为fps(已废弃,请使用videoFps)
}

// RtmpTaskUpdateRequestV2 V2版本更新旁路推流任务请求
// See https://doc.yunxin.163.com/nertc/server-apis/DM0MTg2NTM?platform=server
type RtmpTaskUpdateRequestV2 struct {
	Cid       int64       `json:"cid"`                 // 房间ID
	TaskId    string      `json:"taskId"`              // 推流任务ID
	StreamUrl string      `json:"streamUrl,omitempty"` // 推流地址
	Layout    *RtmpLayout `json:"layout,omitempty"`    // 布局相关参数
	Record    bool        `json:"record,omitempty"`    // 是否进行音视频录制
	Version   int         `json:"version,omitempty"`   // 推流任务版本
	HostUid   uint64      `json:"hostUid,omitempty"`   // 主播的UID
	Config    *RtmpConfig `json:"config,omitempty"`    // 音视频流配置
	ExtraInfo string      `json:"extraInfo,omitempty"` // 自定义的媒体补充增强信息
}

// RtmpTaskUpdateRequestV3 V3版本更新旁路推流任务请求
// See https://doc.yunxin.163.com/nertc/server-apis/DM0MTg2NTM?platform=server
type RtmpTaskUpdateRequestV3 struct {
	Cname     string      `json:"cname"`               // 房间名称
	TaskId    string      `json:"taskId"`              // 推流任务ID
	StreamUrl string      `json:"streamUrl,omitempty"` // 推流地址
	Layout    *RtmpLayout `json:"layout,omitempty"`    // 布局相关参数
	Record    bool        `json:"record,omitempty"`    // 是否进行音视频录制
	Version   int         `json:"version,omitempty"`   // 推流任务版本
	HostUid   uint64      `json:"hostUid,omitempty"`   // 主播的UID
	Config    *RtmpConfig `json:"config,omitempty"`    // 音视频流配置
	ExtraInfo string      `json:"extraInfo,omitempty"` // 自定义的媒体补充增强信息
}

// RtmpTaskQueryRequestV2 V2版本查询指定旁路推流任务请求
// See https://doc.yunxin.163.com/nertc/server-apis/jE5OTE5NDA?platform=server
type RtmpTaskQueryRequestV2 struct {
	Cid    int64  `json:"cid"`    // 房间ID
	TaskId string `json:"taskId"` // 推流任务ID
}

// RtmpTaskQueryRequestV3 V3版本查询指定旁路推流任务请求
// See https://doc.yunxin.163.com/nertc/server-apis/jE5OTE5NDA?platform=server
type RtmpTaskQueryRequestV3 struct {
	Cname  string `json:"cname"`  // 房间名称
	TaskId string `json:"taskId"` // 推流任务ID
}

// RtmpTaskQueryAllRequestV2 V2版本查询所有旁路推流任务请求
// See https://doc.yunxin.163.com/nertc/server-apis/TQ0MjA3NjA?platform=server
type RtmpTaskQueryAllRequestV2 struct {
	Cid int64 `json:"cid"` // 房间ID
}

// RtmpTaskQueryAllRequestV3 V3版本查询所有旁路推流任务请求
// See https://doc.yunxin.163.com/nertc/server-apis/TQ0MjA3NjA?platform=server
type RtmpTaskQueryAllRequestV3 struct {
	Cname string `json:"cname"` // 房间名称
}

// RtmpTaskDeleteRequestV2 V2版本停止旁路推流任务请求
// See https://doc.yunxin.163.com/nertc/server-apis/zY3NDA3MTc?platform=server
type RtmpTaskDeleteRequestV2 struct {
	Cid    int64  `json:"cid"`    // 房间ID
	TaskId string `json:"taskId"` // 推流任务ID
}

// RtmpTaskDeleteRequestV3 V3版本停止旁路推流任务请求
// See https://doc.yunxin.163.com/nertc/server-apis/zY3NDA3MTc?platform=server
type RtmpTaskDeleteRequestV3 struct {
	Cname  string `json:"cname"`  // 房间名称
	TaskId string `json:"taskId"` // 推流任务ID
}
