package chatroom

// Chatroom V1 Response Structures

// AddRobotResponseV1 AddRobot响应
type AddRobotResponseV1 struct {
	FailAccids    []string `json:"failAccids,omitempty"`
	SuccessAccids []string `json:"successAccids,omitempty"`
	OldAccids     []string `json:"oldAccids,omitempty"`
}

// CleanRobotResponseV1 CleanRobot响应
type CleanRobotResponseV1 struct {
	Size int `json:"size,omitempty"`
}

// CreateChatroomResponseV1 CreateChatroom响应
type CreateChatroomResponseV1 struct {
	Roomid       int64     `json:"roomid,omitempty"`
	Valid        bool      `json:"valid,omitempty"`
	Announcement string    `json:"announcement,omitempty"`
	Ext          string    `json:"ext,omitempty"`
	Creator      string    `json:"creator,omitempty"`
	Name         string    `json:"name,omitempty"`
	Broadcasturl string    `json:"broadcasturl,omitempty"`
	Queuelevel   int       `json:"queuelevel,omitempty"`
	Muted        bool      `json:"muted,omitempty"`
	DelayInfo    DelayInfo `json:"delayInfo,omitempty"`
}

// DelayInfo DelayInfo
type DelayInfo struct {
	DelayCloseEnable bool  `json:"delayCloseEnable,omitempty"`
	DelayClosePolicy int   `json:"delayClosePolicy,omitempty"`
	DelaySeconds     int64 `json:"delaySeconds,omitempty"`
	Status           int   `json:"status,omitempty"`
	StartTime        int64 `json:"startTime,omitempty"`
}

// KickMemberResponseV1 KickMember响应
type KickMemberResponseV1 struct {
}

// MuteRoomResponseV1 MuteRoom响应
type MuteRoomResponseV1 struct {
	Success bool `json:"success,omitempty"`
}

// QueryChatroomAddressResponseV1 QueryChatroomAddress响应
type QueryChatroomAddressResponseV1 struct {
	ChatroomAddress []string `json:"addr,omitempty"`
}

// QueryChatroomInfoResponseV1 QueryChatroomInfo响应
type QueryChatroomInfoResponseV1 struct {
	RoomId          int64  `json:"roomid,omitempty"`
	Valid           bool   `json:"valid,omitempty"`
	Muted           bool   `json:"muted,omitempty"`
	Announcement    string `json:"announcement,omitempty"`
	Name            string `json:"name,omitempty"`
	LiveUrl         string `json:"broadcasturl,omitempty"`
	OnlineUserCount int    `json:"onlineusercount,omitempty"`
	Extension       string `json:"ext,omitempty"`
	Creator         string `json:"creator,omitempty"`
	QueueLevel      int    `json:"queuelevel,omitempty"`
	IoNotify        bool   `json:"ionotify,omitempty"`
}

// QueryChatroomInfosResponseV1 QueryChatroomInfos响应
type QueryChatroomInfosResponseV1 struct {
	NoExistRooms []int64    `json:"noExistRooms,omitempty"`
	RoomInfos    []RoomInfo `json:"succRooms,omitempty"`
	FailRooms    []int64    `json:"failRooms,omitempty"`
}

// RoomInfo RoomInfo
type RoomInfo struct {
	RoomId          int64  `json:"roomid,omitempty"`
	Valid           bool   `json:"valid,omitempty"`
	Muted           bool   `json:"muted,omitempty"`
	Announcement    string `json:"announcement,omitempty"`
	Name            string `json:"name,omitempty"`
	LiveUrl         string `json:"broadcasturl,omitempty"`
	OnlineUserCount int    `json:"onlineusercount,omitempty"`
	Extension       string `json:"ext,omitempty"`
	Creator         string `json:"creator,omitempty"`
	QueueLevel      int    `json:"queuelevel,omitempty"`
	IoNotify        bool   `json:"ionotify,omitempty"`
}

// QueryMembersByPageResponseV1 QueryMembersByPage响应
type QueryMembersByPageResponseV1 struct {
	Data []QueryMembersByPageBean `json:"data,omitempty"`
}

// QueryMembersByPage QueryMembersByPage
type QueryMembersByPageBean struct {
	RoomId        int64  `json:"roomid,omitempty"`
	Accid         string `json:"accid,omitempty"`
	Nick          string `json:"nick,omitempty"`
	Avator        string `json:"avator,omitempty"`
	Ext           string `json:"ext,omitempty"`
	Type          string `json:"type,omitempty"`
	Level         int    `json:"level,omitempty"`
	OnlineStat    bool   `json:"onlineStat,omitempty"`
	EnterTime     int64  `json:"enterTime,omitempty"`
	Blacklisted   bool   `json:"blacklisted,omitempty"`
	Muted         bool   `json:"muted,omitempty"`
	TempMuted     bool   `json:"tempMuted,omitempty"`
	TempMuteTtl   int64  `json:"tempMuteTtl,omitempty"`
	IsRobot       bool   `json:"isRobot,omitempty"`
	RobotExpireAt int    `json:"robotExpirAt,omitempty"`
}

// QueryMembersByRolesResponseV1 QueryMembersByRoles响应
type QueryMembersByRolesResponseV1 struct {
	Data []QueryMembersByRolesBean `json:"data,omitempty"`
}

// QueryMembersByRoles QueryMembersByRoles
type QueryMembersByRolesBean struct {
	RoomId        int64  `json:"roomid,omitempty"`
	Accid         string `json:"accid,omitempty"`
	Nick          string `json:"nick,omitempty"`
	Avator        string `json:"avator,omitempty"`
	Ext           string `json:"ext,omitempty"`
	Type          string `json:"type,omitempty"`
	Level         int    `json:"level,omitempty"`
	OnlineStat    bool   `json:"onlineStat,omitempty"`
	EnterTime     int64  `json:"enterTime,omitempty"`
	Blacklisted   bool   `json:"blacklisted,omitempty"`
	Muted         bool   `json:"muted,omitempty"`
	TempMuted     bool   `json:"tempMuted,omitempty"`
	TempMuteTtl   int64  `json:"tempMuteTtl,omitempty"`
	IsRobot       bool   `json:"isRobot,omitempty"`
	RobotExpireAt int    `json:"robotExpirAt,omitempty"`
}

// QueryMembersResponseV1 QueryMembers响应
type QueryMembersResponseV1 struct {
	Data []QueryMembersBean `json:"data,omitempty"`
}

// QueryMembers QueryMembers
type QueryMembersBean struct {
	RoomId        int64  `json:"roomid,omitempty"`
	Accid         string `json:"accid,omitempty"`
	Nick          string `json:"nick,omitempty"`
	Avator        string `json:"avator,omitempty"`
	Ext           string `json:"ext,omitempty"`
	Type          string `json:"type,omitempty"`
	Level         int    `json:"level,omitempty"`
	OnlineStat    bool   `json:"onlineStat,omitempty"`
	EnterTime     int64  `json:"enterTime,omitempty"`
	Blacklisted   bool   `json:"blacklisted,omitempty"`
	Muted         bool   `json:"muted,omitempty"`
	TempMuted     bool   `json:"tempMuted,omitempty"`
	TempMuteTtl   int64  `json:"tempMuteTtl,omitempty"`
	IsRobot       bool   `json:"isRobot,omitempty"`
	RobotExpireAt int    `json:"robotExpirAt,omitempty"`
}

// QueryTagHistoryMsgResponseV1 QueryTagHistoryMsg响应
type QueryTagHistoryMsgResponseV1 struct {
	Msgs []TagMsgBean `json:"msgs,omitempty"`
}

// TagMsgBean TagMsgBean
type TagMsgBean struct {
	Msgid          string `json:"msgid,omitempty"`
	From           string `json:"from,omitempty"`
	Type           int    `json:"type,omitempty"`
	Sendtime       int64  `json:"sendtime,omitempty"`
	Body           string `json:"body,omitempty"`
	Fromclienttype int    `json:"fromclienttype,omitempty"`
}

// QueryUserRoomIdsResponseV1 QueryUserRoomIds响应
type QueryUserRoomIdsResponseV1 struct {
	RoomIds []int64 `json:"roomIds,omitempty"`
}

// QueueBatchOfferResponseV1 QueueBatchOffer响应
type QueueBatchOfferResponseV1 struct {
	HighPriority bool     `json:"highPriority,omitempty"`
	FailedKeys   []string `json:"failedKeys,omitempty"`
}

// QueueBatchUpdateResponseV1 QueueBatchUpdate响应
type QueueBatchUpdateResponseV1 struct {
	HighPriority      bool     `json:"highPriority,omitempty"`
	NoExistElementKey []string `json:"noExistElementKey,omitempty"`
}

// QueueDropResponseV1 QueueDrop响应
type QueueDropResponseV1 struct {
	HighPriority bool `json:"highPriority,omitempty"`
}

// QueueGetResponseV1 QueueGet响应
type QueueGetResponseV1 struct {
	List []QueueElement `json:"list,omitempty"`
}

// QueueInitResponseV1 QueueInit响应
type QueueInitResponseV1 struct {
}

// QueueListResponseV1 QueueList响应
type QueueListResponseV1 struct {
	List []QueueElement `json:"list,omitempty"`
}

// QueueOfferResponseV1 QueueOffer响应
type QueueOfferResponseV1 struct {
	HighPriority bool `json:"highPriority,omitempty"`
}

// QueuePollResponseV1 QueuePoll响应
type QueuePollResponseV1 struct {
	HighPriority bool   `json:"highPriority,omitempty"`
	Key          string `json:"key,omitempty"`
	Value        string `json:"value,omitempty"`
}

// RemoveRobotResponseV1 RemoveRobot响应
type RemoveRobotResponseV1 struct {
	FailAccids    []string `json:"failAccids,omitempty"`
	SuccessAccids []string `json:"successAccids,omitempty"`
}

// SetMemberRoleResponseV1 SetMemberRole响应
type SetMemberRoleResponseV1 struct {
	RoomId int64  `json:"roomid,omitempty"`
	Level  int    `json:"level,omitempty"`
	Accid  string `json:"accid,omitempty"`
	Type   string `json:"type,omitempty"`
}

// TagMembersCountResponseV1 TagMembersCount响应
type TagMembersCountResponseV1 struct {
	Tag             string `json:"tag,omitempty"`
	OnlineUserCount int    `json:"onlineUserCount,omitempty"`
}

// TagMembersQueryResponseV1 TagMembersQuery响应
type TagMembersQueryResponseV1 struct {
	Data []TagMemberQueryBean `json:"data,omitempty"`
}

// TagMemberQueryBean TagMemberQueryBean
type TagMemberQueryBean struct {
	RoomId           int64  `json:"roomid,omitempty"`
	Accid            string `json:"accid,omitempty"`
	Nick             string `json:"nick,omitempty"`
	Avator           string `json:"avator,omitempty"`
	Ext              string `json:"ext,omitempty"`
	Type             string `json:"type,omitempty"`
	Level            int    `json:"level,omitempty"`
	OnlineStat       bool   `json:"onlineStat,omitempty"`
	EnterTime        int64  `json:"enterTime,omitempty"`
	Blacklisted      bool   `json:"blacklisted,omitempty"`
	Muted            bool   `json:"muted,omitempty"`
	TempMuted        bool   `json:"tempMuted,omitempty"`
	TempMuteTtl      int    `json:"tempMuteTtl,omitempty"`
	IsRobot          bool   `json:"isRobot,omitempty"`
	RobotExpirAt     int    `json:"robotExpirAt,omitempty"`
	Tags             string `json:"tags,omitempty"`
	NotifyTargetTags string `json:"notifyTargetTags,omitempty"`
}

// TagTemporaryMuteResponseV1 TagTemporaryMute响应
type TagTemporaryMuteResponseV1 struct {
	MuteDuration int `json:"muteDuration,omitempty"`
}

// TemporaryMuteResponseV1 TemporaryMute响应
type TemporaryMuteResponseV1 struct {
	MuteDuration int64 `json:"muteDuration,omitempty"`
}

// ToggleCloseChatroomStatResponseV1 ToggleCloseChatroomStat响应
type ToggleCloseChatroomStatResponseV1 struct {
	RoomId       int64     `json:"roomid,omitempty"`
	Valid        bool      `json:"valid,omitempty"`
	Announcement string    `json:"announcement,omitempty"`
	Name         string    `json:"name,omitempty"`
	LiveUrl      string    `json:"broadcasturl,omitempty"`
	Extension    string    `json:"ext,omitempty"`
	QueueLevel   int       `json:"queuelevel,omitempty"`
	Muted        bool      `json:"muted,omitempty"`
	Creator      string    `json:"creator,omitempty"`
	DelayInfo    DelayInfo `json:"delayInfo,omitempty"`
}

// UpdateChatRoomRoleTagResponseV1 UpdateChatRoomRoleTag响应
type UpdateChatRoomRoleTagResponseV1 struct {
}

// UpdateChatroomDelayClosePolicyResponseV1 UpdateChatroomDelayClosePolicy响应
type UpdateChatroomDelayClosePolicyResponseV1 struct {
	RoomId       int64     `json:"roomid,omitempty"`
	Valid        bool      `json:"valid,omitempty"`
	Announcement string    `json:"announcement,omitempty"`
	Name         string    `json:"name,omitempty"`
	LiveUrl      string    `json:"broadcasturl,omitempty"`
	Extension    string    `json:"ext,omitempty"`
	QueueLevel   int       `json:"queuelevel,omitempty"`
	Muted        bool      `json:"muted,omitempty"`
	Creator      string    `json:"creator,omitempty"`
	DelayInfo    DelayInfo `json:"delayInfo,omitempty"`
}

// UpdateChatroomInOutNotificationResponseV1 UpdateChatroomInOutNotification响应
type UpdateChatroomInOutNotificationResponseV1 struct {
}

// UpdateChatroomResponseV1 UpdateChatroom响应
type UpdateChatroomResponseV1 struct {
	Roomid       int64  `json:"roomid,omitempty"`
	Valid        bool   `json:"valid,omitempty"`
	Announcement string `json:"announcement,omitempty"`
	Ext          string `json:"ext,omitempty"`
	Creator      string `json:"creator,omitempty"`
	Name         string `json:"name,omitempty"`
	Broadcasturl string `json:"broadcasturl,omitempty"`
}

// UpdateMyRoomRoleResponseV1 UpdateMyRoomRole响应
type UpdateMyRoomRoleResponseV1 struct {
}
