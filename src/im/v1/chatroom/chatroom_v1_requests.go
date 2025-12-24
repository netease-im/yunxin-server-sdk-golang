package chatroom

// Chatroom V1 Request Structures

// AddRobotRequestV1 AddRobot请求
type AddRobotRequestV1 struct {
	Roomid    int64    `json:"roomid,omitempty"`
	Accids    []string `json:"accids,omitempty"`
	RoleExt   string   `json:"roleExt,omitempty"`
	NotifyExt string   `json:"notifyExt,omitempty"`
}

// CleanRobotRequestV1 CleanRobot请求
type CleanRobotRequestV1 struct {
	Roomid int64 `json:"roomid,omitempty"`
	Notify bool  `json:"notify,omitempty"`
}

// CreateChatroomRequestV1 CreateChatroom请求
type CreateChatroomRequestV1 struct {
	Creator           string `json:"creator,omitempty"`
	Name              string `json:"name,omitempty"`
	Announcement      string `json:"announcement,omitempty"`
	Broadcasturl      string `json:"broadcasturl,omitempty"`
	Ext               string `json:"ext,omitempty"`
	Bid               string `json:"bid,omitempty"`
	DelayClosePolicy  int    `json:"delayClosePolicy,omitempty"`
	DelaySeconds      int64  `json:"delaySeconds,omitempty"`
	QueueLevel        int    `json:"queuelevel,omitempty"`
	InOutNotification int    `json:"inOutNotification,omitempty"`
}

// KickMemberRequestV1 KickMember请求
type KickMemberRequestV1 struct {
	Roomid      int64  `json:"roomid,omitempty"`
	TargetAccid string `json:"targetAccid,omitempty"`
	Accid       string `json:"accid,omitempty"`
	NotifyExt   string `json:"notifyExt,omitempty"`
}

// MuteRoomRequestV1 MuteRoom请求
type MuteRoomRequestV1 struct {
	Roomid     int64  `json:"roomid,omitempty"`
	Operator   string `json:"operator,omitempty"`
	Mute       bool   `json:"mute,omitempty"`
	NeedNotify bool   `json:"needNotify,omitempty"`
	NotifyExt  string `json:"notifyExt,omitempty"`
}

// QueryChatroomAddressRequestV1 QueryChatroomAddress请求
type QueryChatroomAddressRequestV1 struct {
	Roomid     int64  `json:"roomid,omitempty"`
	Accid      string `json:"accid,omitempty"`
	ClientType int    `json:"clienttype,omitempty"`
	ClientIp   string `json:"clientip,omitempty"`
}

// QueryChatroomInfoRequestV1 QueryChatroomInfo请求
type QueryChatroomInfoRequestV1 struct {
	Roomid              int64 `json:"roomid,omitempty"`
	NeedOnlineUserCount bool  `json:"needOnlineUserCount,omitempty"`
}

// QueryChatroomInfosRequestV1 QueryChatroomInfos请求
type QueryChatroomInfosRequestV1 struct {
	Roomids             []int64 `json:"roomids,omitempty"`
	NeedOnlineUserCount bool    `json:"needOnlineUserCount,omitempty"`
}

// QueryMembersByPageRequestV1 QueryMembersByPage请求
type QueryMembersByPageRequestV1 struct {
	Roomid  int64 `json:"roomid,omitempty"`
	Type    int   `json:"type,omitempty"`
	EndTime int64 `json:"endtime,omitempty"`
	Limit   int64 `json:"limit,omitempty"`
}

// QueryMembersByRolesRequestV1 QueryMembersByRoles请求
type QueryMembersByRolesRequestV1 struct {
	Roomid int64  `json:"roomid,omitempty"`
	Roles  string `json:"roles,omitempty"`
}

// QueryMembersRequestV1 QueryMembers请求
type QueryMembersRequestV1 struct {
	Roomid int64    `json:"roomid,omitempty"`
	Accids []string `json:"accids,omitempty"`
}

// QueryTagHistoryMsgRequestV1 QueryTagHistoryMsg请求
type QueryTagHistoryMsgRequestV1 struct {
	RoomId   int64    `json:"roomId,omitempty"`
	Tags     []string `json:"tags,omitempty"`
	Types    string   `json:"types,omitempty"`
	FromTime int64    `json:"fromTime,omitempty"`
	ToTime   int64    `json:"toTime,omitempty"`
	Limit    int      `json:"limit,omitempty"`
	Reverse  int      `json:"reverse,omitempty"`
}

// QueryUserRoomIdsRequestV1 QueryUserRoomIds请求
type QueryUserRoomIdsRequestV1 struct {
	Creator string `json:"creator,omitempty"`
}

// QueueBatchOfferRequestV1 QueueBatchOffer请求
type QueueBatchOfferRequestV1 struct {
	Roomid             int64          `json:"roomid,omitempty"`
	Operator           string         `json:"operator,omitempty"`
	Transient          bool           `json:"transient,omitempty"`
	Elements           []QueueElement `json:"elements,omitempty"`
	NeedNotify         bool           `json:"needNotify,omitempty"`
	NotifyExt          string         `json:"notifyExt,omitempty"`
	HighPriority       int            `json:"highPriority,omitempty"`
	HighPriorityPolicy int            `json:"highPriorityPolicy,omitempty"`
}

// QueueElement QueueElement
type QueueElement struct {
	Key       string `json:"key,omitempty"`
	Value     string `json:"value,omitempty"`
	Accid     string `json:"accid,omitempty"`
	Transient bool   `json:"transient,omitempty"`
}

// QueueBatchUpdateRequestV1 QueueBatchUpdate请求
type QueueBatchUpdateRequestV1 struct {
	Roomid             int64             `json:"roomid,omitempty"`
	Operator           string            `json:"operator,omitempty"`
	Elements           map[string]string `json:"elements,omitempty"`
	NeedNotify         bool              `json:"needNotify,omitempty"`
	NotifyExt          string            `json:"notifyExt,omitempty"`
	HighPriority       int               `json:"highPriority,omitempty"`
	HighPriorityPolicy int               `json:"highPriorityPolicy,omitempty"`
}

// QueueDropRequestV1 QueueDrop请求
type QueueDropRequestV1 struct {
	Roomid             int64 `json:"roomid,omitempty"`
	HighPriority       int   `json:"highPriority,omitempty"`
	HighPriorityPolicy int   `json:"highPriorityPolicy,omitempty"`
}

// QueueGetRequestV1 QueueGet请求
type QueueGetRequestV1 struct {
	Roomid int64    `json:"roomid,omitempty"`
	Keys   []string `json:"keys,omitempty"`
}

// QueueInitRequestV1 QueueInit请求
type QueueInitRequestV1 struct {
	Roomid    int64 `json:"roomid,omitempty"`
	SizeLimit int   `json:"sizeLimit,omitempty"`
}

// QueueListRequestV1 QueueList请求
type QueueListRequestV1 struct {
	Roomid int64 `json:"roomid,omitempty"`
}

// QueueOfferRequestV1 QueueOffer请求
type QueueOfferRequestV1 struct {
	Roomid             int64  `json:"roomid,omitempty"`
	Key                string `json:"key,omitempty"`
	Value              string `json:"value,omitempty"`
	Operator           string `json:"operator,omitempty"`
	Transient          bool   `json:"transient,omitempty"`
	HighPriority       int    `json:"highPriority,omitempty"`
	HighPriorityPolicy int    `json:"highPriorityPolicy,omitempty"`
}

// QueuePollRequestV1 QueuePoll请求
type QueuePollRequestV1 struct {
	Roomid             int64  `json:"roomid,omitempty"`
	Key                string `json:"key,omitempty"`
	HighPriority       int    `json:"highPriority,omitempty"`
	HighPriorityPolicy int    `json:"highPriorityPolicy,omitempty"`
}

// RemoveRobotRequestV1 RemoveRobot请求
type RemoveRobotRequestV1 struct {
	Roomid int64    `json:"roomid,omitempty"`
	Accids []string `json:"accids,omitempty"`
}

// SetMemberRoleRequestV1 SetMemberRole请求
type SetMemberRoleRequestV1 struct {
	Roomid    int64  `json:"roomid,omitempty"`
	Operator  string `json:"operator,omitempty"`
	Target    string `json:"target,omitempty"`
	Opt       int    `json:"opt,omitempty"`
	OptValue  bool   `json:"optvalue,omitempty"`
	NotifyExt string `json:"notifyExt,omitempty"`
}

// TagMembersCountRequestV1 TagMembersCount请求
type TagMembersCountRequestV1 struct {
	Roomid int64  `json:"roomid,omitempty"`
	Tag    string `json:"tag,omitempty"`
}

// TagMembersQueryRequestV1 TagMembersQuery请求
type TagMembersQueryRequestV1 struct {
	Roomid  int64  `json:"roomid,omitempty"`
	Tag     string `json:"tag,omitempty"`
	EndTime int64  `json:"endTime,omitempty"`
	Limit   int    `json:"limit,omitempty"`
}

// TagTemporaryMuteRequestV1 TagTemporaryMute请求
type TagTemporaryMuteRequestV1 struct {
	Roomid           int64  `json:"roomid,omitempty"`
	Operator         string `json:"operator,omitempty"`
	TargetTag        string `json:"targetTag,omitempty"`
	NeedNotify       bool   `json:"needNotify,omitempty"`
	NotifyExt        string `json:"notifyExt,omitempty"`
	MuteDuration     int    `json:"muteDuration,omitempty"`
	NotifyTargetTags string `json:"notifyTargetTags,omitempty"`
}

// TemporaryMuteRequestV1 TemporaryMute请求
type TemporaryMuteRequestV1 struct {
	Roomid       int64  `json:"roomid,omitempty"`
	Operator     string `json:"operator,omitempty"`
	Target       string `json:"target,omitempty"`
	MuteDuration int64  `json:"muteDuration,omitempty"`
	NeedNotify   bool   `json:"needNotify,omitempty"`
	NotifyExt    string `json:"notifyExt,omitempty"`
}

// ToggleCloseChatroomStatRequestV1 ToggleCloseChatroomStat请求
type ToggleCloseChatroomStatRequestV1 struct {
	Roomid           int64  `json:"roomid,omitempty"`
	Operator         string `json:"operator,omitempty"`
	Valid            bool   `json:"valid,omitempty"`
	DelayClosePolicy int    `json:"delayClosePolicy,omitempty"`
	DelaySeconds     int64  `json:"delaySeconds,omitempty"`
}

// UpdateChatRoomRoleTagRequestV1 UpdateChatRoomRoleTag请求
type UpdateChatRoomRoleTagRequestV1 struct {
	RoomId           int64    `json:"roomId,omitempty"`
	Tags             []string `json:"tags,omitempty"`
	Accid            string   `json:"accid,omitempty"`
	NotifyTargetTags string   `json:"notifyTargetTags,omitempty"`
}

// UpdateChatroomDelayClosePolicyRequestV1 UpdateChatroomDelayClosePolicy请求
type UpdateChatroomDelayClosePolicyRequestV1 struct {
	Roomid           int64 `json:"roomid,omitempty"`
	DelayClosePolicy int   `json:"delayClosePolicy,omitempty"`
	DelaySeconds     int64 `json:"delaySeconds,omitempty"`
}

// UpdateChatroomInOutNotificationRequestV1 UpdateChatroomInOutNotification请求
type UpdateChatroomInOutNotificationRequestV1 struct {
	Roomid int64 `json:"roomid,omitempty"`
	Close  bool  `json:"close,omitempty"`
}

// UpdateChatroomRequestV1 UpdateChatroom请求
type UpdateChatroomRequestV1 struct {
	Roomid       int64  `json:"roomid,omitempty"`
	Name         string `json:"name,omitempty"`
	Announcement string `json:"announcement,omitempty"`
	Broadcasturl string `json:"broadcasturl,omitempty"`
	Ext          string `json:"ext,omitempty"`
	NeedNotify   bool   `json:"needNotify,omitempty"`
	NotifyExt    string `json:"notifyExt,omitempty"`
	Queuelevel   int    `json:"queuelevel,omitempty"`
	Bid          string `json:"bid,omitempty"`
}

// UpdateMyRoomRoleRequestV1 UpdateMyRoomRole请求
type UpdateMyRoomRoleRequestV1 struct {
	Roomid     int64  `json:"roomid,omitempty"`
	Accid      string `json:"accid,omitempty"`
	Save       bool   `json:"save,omitempty"`
	NeedNotify bool   `json:"needNotify,omitempty"`
	NotifyExt  string `json:"notifyExt,omitempty"`
	Nick       string `json:"nick,omitempty"`
	Avator     string `json:"avator,omitempty"`
	Ext        string `json:"ext,omitempty"`
	Bid        string `json:"bid,omitempty"`
}
