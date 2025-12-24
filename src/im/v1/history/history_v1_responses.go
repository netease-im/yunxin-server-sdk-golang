package history

// History V1 Response Structures

// DeleteChatroomHistoryMessageResponseV1 DeleteChatroomHistoryMessage响应
type DeleteChatroomHistoryMessageResponseV1 struct {
}

// QueryBroadcastHistoryMessageByIdResponseV1 QueryBroadcastHistoryMessageById响应
type QueryBroadcastHistoryMessageByIdResponseV1 struct {
	Msg BroadcastMessage `json:"msg,omitempty"`
}

// BroadcastMessage BroadcastMessage
type BroadcastMessage struct {
	BroadcastId int64    `json:"broadcastId,omitempty"`
	Body        string   `json:"body,omitempty"`
	CreateTime  int64    `json:"createTime,omitempty"`
	ExpireTime  int64    `json:"expireTime,omitempty"`
	From        string   `json:"from,omitempty"`
	IsOffline   bool     `json:"isOffline,omitempty"`
	TargetOs    []string `json:"targetOs,omitempty"`
}

// QueryBroadcastHistoryMessageResponseV1 QueryBroadcastHistoryMessage响应
type QueryBroadcastHistoryMessageResponseV1 struct {
	Msgs []BroadcastMessage `json:"msgs,omitempty"`
}

// QueryChatroomHistoryMessageResponseV1 QueryChatroomHistoryMessage响应
type QueryChatroomHistoryMessageResponseV1 struct {
	Size int               `json:"size,omitempty"`
	Msgs []ChatroomMessage `json:"msgs,omitempty"`
}

// QuerySessionHistoryMessageResponseV1 QuerySessionHistoryMessage响应
type QuerySessionHistoryMessageResponseV1 struct {
	Size int       `json:"size,omitempty"`
	Msgs []Message `json:"msgs,omitempty"`
}

// QuerySessionListResponseV1 QuerySessionList响应
type QuerySessionListResponseV1 struct {
	HasMore  bool      `json:"hasMore,omitempty"`
	Sessions []Session `json:"sessions,omitempty"`
}

// QueryTeamHistoryMessageResponseV1 QueryTeamHistoryMessage响应
type QueryTeamHistoryMessageResponseV1 struct {
	Size int       `json:"size,omitempty"`
	Msgs []Message `json:"msgs,omitempty"`
}

// QueryUserEventsResponseV1 QueryUserEvents响应
type QueryUserEventsResponseV1 struct {
	Size   int         `json:"size,omitempty"`
	Events []UserEvent `json:"events,omitempty"`
}

// UserEvent UserEvent
type UserEvent struct {
	Accid      string `json:"accid,omitempty"`
	Timestamp  int64  `json:"timestamp,omitempty"`
	EventType  int    `json:"eventType,omitempty"`
	ClientIp   string `json:"clientIp,omitempty"`
	SdkVersion int    `json:"sdkVersion,omitempty"`
	ClientType string `json:"clientType,omitempty"`
	DeviceId   string `json:"deviceId,omitempty"`
	CustomTag  string `json:"customTag,omitempty"`
	Code       int    `json:"code,omitempty"`
}

// Session Session
type Session struct {
	SessionType int    `json:"sessionType,omitempty"`
	Accid       string `json:"accid,omitempty"`
	Tid         int64  `json:"tid,omitempty"`
	UpdateTime  int64  `json:"updateTime,omitempty"`
	Ext         string `json:"ext,omitempty"`
	LastMsgType string `json:"lastMsgType,omitempty"`
	LastMsg     string `json:"lastMsg,omitempty"`
}
