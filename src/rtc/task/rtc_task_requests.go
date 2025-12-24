package task

// TaskCreateRequest 创建云端播放任务请求
// See https://doc.yunxin.163.com/nertc/server-apis/DgwMTQ2NzA?platform=server
type TaskCreateRequest struct {
	Cname     string    `json:"cname"`               // 房间名称
	RequestId string    `json:"requestId,omitempty"` // 随机字符串,用于问题排查
	TaskType  int       `json:"taskType"`            // 服务器任务类型,固定为2
	Data      *TaskData `json:"data"`                // 云端播放任务信息
}

// TaskData 云端播放任务数据
type TaskData struct {
	StreamUrl   string `json:"streamUrl"`             // 点播或者直播流地址(支持 RTMP、HLS、HTTP-FLV、HTTPS-FLV 协议的直播流,支持FLV、MP4、MPEG-TS、Matroska(MKV)、MP3、wav格式的点播文件)
	Token       string `json:"token,omitempty"`       // 用于鉴权的安全认证签名(Token)
	Uid         uint64 `json:"uid"`                   // 云端播放器在房间内的用户ID
	IdleTimeout int    `json:"idleTimeout,omitempty"` // 云端播放器处于空闲状态的最大时长,单位为秒(s),超时后自动销毁
	PlayTs      int64  `json:"playTs,omitempty"`      // 云端播放器开始播放在线媒体流时的Unix时间戳,单位为秒(s)
	MediaType   int    `json:"mediaType"`             // 播放流的媒体类型: 0-音频, 1-视频, 2-音视频
}

// TaskUpdateRequestV2 V2版本更新云端播放任务请求
// See https://doc.yunxin.163.com/nertc/server-apis/TM1ODMyMTA?platform=server
type TaskUpdateRequestV2 struct {
	Cid       int64     `json:"cid"`                 // 房间ID
	RequestId string    `json:"requestId,omitempty"` // 随机字符串,用于问题排查
	TaskType  int       `json:"taskType"`            // 服务器任务类型,固定为2
	TaskId    string    `json:"taskId"`              // 任务唯一标识符
	Data      *TaskData `json:"data"`                // 云端播放任务信息
}

// TaskUpdateRequestV3 V3版本更新云端播放任务请求
// See https://doc.yunxin.163.com/nertc/server-apis/TM1ODMyMTA?platform=server
type TaskUpdateRequestV3 struct {
	Cname     string    `json:"cname"`               // 房间名称
	RequestId string    `json:"requestId,omitempty"` // 随机字符串,用于问题排查
	TaskType  int       `json:"taskType"`            // 服务器任务类型,固定为2
	TaskId    string    `json:"taskId"`              // 任务唯一标识符
	Data      *TaskData `json:"data"`                // 云端播放任务信息
}

// TaskQueryRequestV2 V2版本查询云端播放任务请求
// See https://doc.yunxin.163.com/nertc/server-apis/TM5NTMwNzY?platform=server
type TaskQueryRequestV2 struct {
	Cid      int64  `json:"cid"`              // 房间ID
	TaskType int    `json:"taskType"`         // 服务器任务类型,固定为2
	TaskId   string `json:"taskId"`           // 任务唯一标识符
	Status   int    `json:"status,omitempty"` // 任务状态: 0-查询全部, 1-查询运行中的任务
}

// TaskQueryRequestV3 V3版本查询云端播放任务请求
// See https://doc.yunxin.163.com/nertc/server-apis/TM5NTMwNzY?platform=server
type TaskQueryRequestV3 struct {
	Cname    string `json:"cname"`            // 房间名称
	TaskType int    `json:"taskType"`         // 服务器任务类型,固定为2
	TaskId   string `json:"taskId"`           // 任务唯一标识符
	Status   int    `json:"status,omitempty"` // 任务状态: 0-查询全部, 1-查询运行中的任务
}

// TaskPauseRequestV2 V2版本暂停云端播放任务请求
// See https://doc.yunxin.163.com/nertc/server-apis/jE4MzIwNzQ?platform=server
type TaskPauseRequestV2 struct {
	Cid      int64  `json:"cid"`      // 房间ID
	TaskType int    `json:"taskType"` // 服务器任务类型,固定为2
	TaskId   string `json:"taskId"`   // 任务唯一标识符
}

// TaskPauseRequestV3 V3版本暂停云端播放任务请求
// See https://doc.yunxin.163.com/nertc/server-apis/jE4MzIwNzQ?platform=server
type TaskPauseRequestV3 struct {
	Cname    string `json:"cname"`    // 房间名称
	TaskType int    `json:"taskType"` // 服务器任务类型,固定为2
	TaskId   string `json:"taskId"`   // 任务唯一标识符
}

// TaskResumeRequestV2 V2版本恢复云端播放任务请求
// See https://doc.yunxin.163.com/nertc/server-apis/DQ2ODI0NjQ?platform=server
type TaskResumeRequestV2 struct {
	Cid      int64  `json:"cid"`      // 房间ID
	TaskType int    `json:"taskType"` // 服务器任务类型,固定为2
	TaskId   string `json:"taskId"`   // 任务唯一标识符
}

// TaskResumeRequestV3 V3版本恢复云端播放任务请求
// See https://doc.yunxin.163.com/nertc/server-apis/DQ2ODI0NjQ?platform=server
type TaskResumeRequestV3 struct {
	Cname    string `json:"cname"`    // 房间名称
	TaskType int    `json:"taskType"` // 服务器任务类型,固定为2
	TaskId   string `json:"taskId"`   // 任务唯一标识符
}

// TaskDestroyRequestV2 V2版本销毁云端播放任务请求
// See https://doc.yunxin.163.com/nertc/server-apis/zAzNjE5NTM?platform=server
type TaskDestroyRequestV2 struct {
	Cid      int64  `json:"cid"`      // 房间ID
	TaskType int    `json:"taskType"` // 服务器任务类型,固定为2
	TaskId   string `json:"taskId"`   // 任务唯一标识符
}

// TaskDestroyRequestV3 V3版本销毁云端播放任务请求
// See https://doc.yunxin.163.com/nertc/server-apis/zAzNjE5NTM?platform=server
type TaskDestroyRequestV3 struct {
	Cname    string `json:"cname"`    // 房间名称
	TaskType int    `json:"taskType"` // 服务器任务类型,固定为2
	TaskId   string `json:"taskId"`   // 任务唯一标识符
}
