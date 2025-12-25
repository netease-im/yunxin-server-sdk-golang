package room

// RtcRoomMember RTC房间成员信息
// See https://doc.yunxin.163.com/interactive-streaming/outdated_articles/zY5OTk5NTM?platform=client
type RtcRoomMember struct {
	Uid       uint64 `json:"uid"`       // 用户ID
	StartTime int64  `json:"starttime"` // 开始时间
	UserRole  int    `json:"userRole"`  // 用户角色
}

// RtcCreateRoomResponse 创建房间响应
type RtcCreateRoomResponse struct {
	Cid int64 `json:"cid"` // 房间ID
}

// RtcGetRoomResponse 获取房间信息响应
type RtcGetRoomResponse struct {
	Cid         int64  `json:"cid"`         // 房间ID
	Cname       string `json:"cname"`       // 房间名称
	Uid         uint64 `json:"uid"`         // 用户ID
	Total       int    `json:"total"`       // 总数
	Stats       int    `json:"stats"`       // 状态
	CreateTime  int64  `json:"createtime"`  // 创建时间
	DestroyTime int64  `json:"destroytime"` // 销毁时间
}

// RtcListRoomMembersResponse 列出房间成员响应
type RtcListRoomMembersResponse struct {
	Cname     string          `json:"cname"`     // 房间名称
	Cid       int64           `json:"cid"`       // 房间ID
	Total     int             `json:"total"`     // 总数
	Members   []RtcRoomMember `json:"members"`   // 成员列表
	Code      int             `json:"code"`      // 状态码
	RequestId string          `json:"requestId"` // 请求ID
	Errmsg    string          `json:"errmsg"`    // 错误信息
}

// RtcAddMemberToKicklistResponse 添加成员到踢出列表响应
type RtcAddMemberToKicklistResponse struct {
	Code      int    `json:"code"`      // 状态码
	RequestId string `json:"requestId"` // 请求ID
	Errmsg    string `json:"errmsg"`    // 错误信息
}

// RtcDeleteRoomResponse 删除房间响应
type RtcDeleteRoomResponse struct {
	Code      int    `json:"code"`      // 状态码
	RequestId string `json:"requestId"` // 请求ID
	Errmsg    string `json:"errmsg"`    // 错误信息
}
