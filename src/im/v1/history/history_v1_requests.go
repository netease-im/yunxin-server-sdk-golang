package history

// History V1 Request Structures

// ChatroomMessage ChatroomMessage
type ChatroomMessage struct {
	From           string `json:"from,omitempty"`
	Msgid          string `json:"msgid,omitempty"`
	Sendtime       int64  `json:"sendtime,omitempty"`
	Type           int    `json:"type,omitempty"`
	Fromclienttype int    `json:"fromclienttype,omitempty"`
	Body           string `json:"body,omitempty"`
}

// DeleteChatroomHistoryMessageRequestV1 DeleteChatroomHistoryMessage请求
type DeleteChatroomHistoryMessageRequestV1 struct {
	Roomid     int64  `json:"roomid,omitempty"`
	FromAcc    string `json:"fromAcc,omitempty"`
	MsgTimetag int64  `json:"msgTimetag,omitempty"`
}

// Message Message
type Message struct {
	From           string `json:"from,omitempty"`
	Msgid          int64  `json:"msgid,omitempty"`
	Msgidclient    string `json:"msgidclient,omitempty"`
	Sendtime       int64  `json:"sendtime,omitempty"`
	Type           int    `json:"type,omitempty"`
	Fromclienttype int    `json:"fromclienttype,omitempty"`
	Body           string `json:"body,omitempty"`
}

// QueryBroadcastHistoryMessageByIdRequestV1 QueryBroadcastHistoryMessageById请求
type QueryBroadcastHistoryMessageByIdRequestV1 struct {
	BroadcastId int64 `json:"broadcastId,omitempty"`
}

// QueryBroadcastHistoryMessageRequestV1 QueryBroadcastHistoryMessage请求
type QueryBroadcastHistoryMessageRequestV1 struct {
	BroadcastId int64 `json:"broadcastId,omitempty"`
	Limit       int   `json:"limit,omitempty"`
	Type        int64 `json:"type,omitempty"`
}

// QueryChatroomHistoryMessageRequestV1 QueryChatroomHistoryMessage请求
type QueryChatroomHistoryMessageRequestV1 struct {
	Roomid  int64  `json:"roomid,omitempty"`
	Accid   string `json:"accid,omitempty"`
	Timetag int64  `json:"timetag,omitempty"`
	Limit   int    `json:"limit,omitempty"`
	Reverse int    `json:"reverse,omitempty"`
	Type    string `json:"type,omitempty"`
}

// QuerySessionHistoryMessageRequestV1 QuerySessionHistoryMessage请求
type QuerySessionHistoryMessageRequestV1 struct {
	From              string `json:"from,omitempty"`
	To                string `json:"to,omitempty"`
	Begintime         int64  `json:"begintime,omitempty"`
	Endtime           int64  `json:"endtime,omitempty"`
	Limit             int    `json:"limit,omitempty"`
	Reverse           int    `json:"reverse,omitempty"`
	Type              string `json:"type,omitempty"`
	IncludeNoSenseMsg bool   `json:"includeNoSenseMsg,omitempty"`
	ExcludeMsgid      string `json:"excludeMsgid,omitempty"`
}

// QuerySessionListRequestV1 QuerySessionList请求
type QuerySessionListRequestV1 struct {
	Accid        string `json:"accid,omitempty"`
	MinTimestamp int64  `json:"minTimestamp,omitempty"`
	MaxTimestamp int64  `json:"maxTimestamp,omitempty"`
	Limit        int    `json:"limit,omitempty"`
	NeedLastMsg  int    `json:"needLastMsg,omitempty"`
}

// QueryTeamHistoryMessageRequestV1 QueryTeamHistoryMessage请求
type QueryTeamHistoryMessageRequestV1 struct {
	Tid               int64  `json:"tid,omitempty"`
	Accid             string `json:"accid,omitempty"`
	Begintime         int64  `json:"begintime,omitempty"`
	Endtime           int64  `json:"endtime,omitempty"`
	Limit             int    `json:"limit,omitempty"`
	Reverse           int    `json:"reverse,omitempty"`
	Type              string `json:"type,omitempty"`
	CheckTeamValid    bool   `json:"checkTeamValid,omitempty"`
	IncludeNoSenseMsg bool   `json:"includeNoSenseMsg,omitempty"`
	ExcludeMsgid      string `json:"excludeMsgid,omitempty"`
}

// QueryUserEventsRequestV1 QueryUserEvents请求
type QueryUserEventsRequestV1 struct {
	Accid     string `json:"accid,omitempty"`
	Begintime int64  `json:"begintime,omitempty"`
	Endtime   int64  `json:"endtime,omitempty"`
	Limit     int    `json:"limit,omitempty"`
	Reverse   int    `json:"reverse,omitempty"`
}
