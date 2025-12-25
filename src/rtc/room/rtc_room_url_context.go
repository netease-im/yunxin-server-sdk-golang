package room

// URL路径常量，用于所有RTC房间相关的API请求
const (
	CreateRoom                        = "/v2/api/room"                                    // 创建房间
	GetRoomByCid                      = "/v2/api/rooms/{cid}"                             // 根据CID获取房间信息
	GetRoomByCname                    = "/v3/api/rooms"                                   // 根据CNAME获取房间信息
	ListMembersV2                     = "/v2/api/rooms/{cid}/members"                     // V2版本列出房间成员
	ListMembersV3                     = "/v3/api/rooms/members"                           // V3版本列出房间成员
	AddMemberToKicklistV2             = "/v2/api/kicklist/{cid}/members/{uid}"            // V2版本添加成员到踢出列表
	AddMemberToKicklistV2WithDuration = "/v2/api/kicklist/{cid}/members/{uid}/{duration}" // V2版本添加成员到踢出列表(带时长)
	AddMemberToKicklistV3             = "/v3/api/kicklist/members"                        // V3版本添加成员到踢出列表
	DeleteRoomV2                      = "/v2/api/rooms/{cid}"                             // V2版本删除房间
	DeleteRoomV3                      = "/v3/api/rooms"                                   // V3版本删除房间
)
