package chatroom_message

// Chatroom Message V1 Response Structures

// BatchChatroomTargetMsgResponseV1 BatchChatroomTargetMsg响应
type BatchChatroomTargetMsgResponseV1 struct {
	FailList    []FailedMessage               `json:"failList,omitempty"`
	SuccessList []ChatroomTargetMsgResponseV1 `json:"successList,omitempty"`
}

// FailedMessage FailedMessage
type FailedMessage struct {
	ClientMsgId string `json:"clientMsgId,omitempty"`
	Reason      string `json:"reason,omitempty"`
}

// BatchSendChatroomMsgResponseV1 BatchSendChatroomMsg响应
type BatchSendChatroomMsgResponseV1 struct {
	FailList    []FailedMessage     `json:"failList,omitempty"`
	SuccessList []SuccessfulMessage `json:"successList,omitempty"`
}

// SuccessfulMessage SuccessfulMessage
type SuccessfulMessage struct {
	ClientMsgId      string `json:"clientMsgId,omitempty"`
	FromNick         string `json:"fromNick,omitempty"`
	FromAccount      string `json:"fromAccount,omitempty"`
	Attach           string `json:"attach,omitempty"`
	Time             string `json:"time,omitempty"`
	Type             int    `json:"type,omitempty"`
	HighPriorityFlag string `json:"highPriorityFlag,omitempty"`
	RoomId           int64  `json:"roomId,omitempty"`
	Ext              string `json:"ext,omitempty"`
	FromClientType   string `json:"fromClientType,omitempty"`
}

// ChatroomTargetMsgResponseV1 ChatroomTargetMsg响应
type ChatroomTargetMsgResponseV1 struct {
	Time           int64  `json:"time,omitempty"`
	FromAvator     string `json:"fromAvator,omitempty"`
	MsgidClient    string `json:"msgid_client,omitempty"`
	FromClientType string `json:"fromClientType,omitempty"`
	RoomId         int64  `json:"roomId,omitempty"`
	FromAccount    string `json:"fromAccount,omitempty"`
	FromNick       string `json:"fromNick,omitempty"`
	Attach         string `json:"attach,omitempty"`
	Type           int    `json:"type,omitempty"`
	Ext            string `json:"ext,omitempty"`
}

// RecallChatroomMsgResponseV1 RecallChatroomMsg响应
type RecallChatroomMsgResponseV1 struct {
}

// SendChatroomMsgResponseV1 SendChatroomMsg响应
type SendChatroomMsgResponseV1 struct {
	Time             int64  `json:"time,omitempty"`
	FromAvator       string `json:"fromAvator,omitempty"`
	Msgid_client     string `json:"msgid_client,omitempty"`
	FromClientType   string `json:"fromClientType,omitempty"`
	RoomId           int64  `json:"roomId,omitempty"`
	FromAccount      string `json:"fromAccount,omitempty"`
	FromNick         string `json:"fromNick,omitempty"`
	Attach           string `json:"attach,omitempty"`
	Type             int    `json:"type,omitempty"`
	Ext              string `json:"ext,omitempty"`
	HighPriorityFlag int    `json:"highPriorityFlag,omitempty"`
	MsgAbandonFlag   int    `json:"msgAbandonFlag,omitempty"`
}
