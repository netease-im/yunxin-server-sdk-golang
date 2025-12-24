package chatroom_message

// Chatroom Message V1 Request Structures

// BatchChatroomTargetMsgRequestV1 BatchChatroomTargetMsg请求
type BatchChatroomTargetMsgRequestV1 struct {
	RoomId            int64     `json:"roomid,omitempty"`
	MsgList           []Message `json:"msgList,omitempty"`
	FromAccid         string    `json:"fromAccid,omitempty"`
	ToAccids          []string  `json:"toAccids,omitempty"`
	Route             int       `json:"route,omitempty"`
	UseYidun          int       `json:"useYidun,omitempty"`
	YidunAntiCheating string    `json:"yidunAntiCheating,omitempty"`
	YidunAntiSpamExt  string    `json:"yidunAntiSpamExt,omitempty"`
	Bid               string    `json:"bid,omitempty"`
	Antispam          bool      `json:"antispam,omitempty"`
	AntispamCustom    string    `json:"antispamCustom,omitempty"`
	Env               string    `json:"env,omitempty"`
}

// Message Message
type Message struct {
	MsgId      string `json:"msgId,omitempty"`
	MsgType    int    `json:"msgType,omitempty"`
	Attach     string `json:"attach,omitempty"`
	Ext        string `json:"ext,omitempty"`
	SubType    int    `json:"subType,omitempty"`
	ResendFlag int    `json:"resendFlag,omitempty"`
}

// BatchSendChatroomMsgRequestV1 BatchSendChatroomMsg请求
type BatchSendChatroomMsgRequestV1 struct {
	RoomId                         int64     `json:"roomid,omitempty"`
	MsgList                        []Message `json:"msgList,omitempty"`
	FromAccid                      string    `json:"fromAccid,omitempty"`
	IgnoreMute                     bool      `json:"ignoreMute,omitempty"`
	SkipHistory                    int       `json:"skipHistory,omitempty"`
	Route                          int       `json:"route,omitempty"`
	AbandonRatio                   int       `json:"abandonRatio,omitempty"`
	HighPriority                   bool      `json:"highPriority,omitempty"`
	NeedHighPriorityMsgResend      bool      `json:"needHighPriorityMsgResend,omitempty"`
	UseYidun                       int       `json:"useYidun,omitempty"`
	YidunAntiCheating              string    `json:"yidunAntiCheating,omitempty"`
	YidunAntiSpamExt               string    `json:"yidunAntiSpamExt,omitempty"`
	Bid                            string    `json:"bid,omitempty"`
	Antispam                       bool      `json:"antispam,omitempty"`
	NotifyTargetTags               string    `json:"notifyTargetTags,omitempty"`
	AntispamCustom                 string    `json:"antispamCustom,omitempty"`
	Env                            string    `json:"env,omitempty"`
	ChatMsgPriority                int       `json:"chatMsgPriority,omitempty"`
	ForbiddenIfHighPriorityMsgFreq int       `json:"forbiddenIfHighPriorityMsgFreq,omitempty"`
}

// ChatroomBroadcastRequestV1 ChatroomBroadcast请求
type ChatroomBroadcastRequestV1 struct {
	MsgId             string `json:"msgId,omitempty"`
	FromAccid         string `json:"fromAccid,omitempty"`
	MsgType           int    `json:"msgType,omitempty"`
	Attach            string `json:"attach,omitempty"`
	SubType           int    `json:"subType,omitempty"`
	ResendFlag        int    `json:"resendFlag,omitempty"`
	Route             int    `json:"route,omitempty"`
	Ext               string `json:"ext,omitempty"`
	UseYidun          int    `json:"useYidun,omitempty"`
	YidunAntiCheating string `json:"yidunAntiCheating,omitempty"`
	Bid               string `json:"bid,omitempty"`
	Antispam          bool   `json:"antispam,omitempty"`
	AntispamCustom    string `json:"antispamCustom,omitempty"`
	NotifyTargetTags  string `json:"notifyTargetTags,omitempty"`
	Env               string `json:"env,omitempty"`
	HighPriority      bool   `json:"highPriority,omitempty"`
}

// ChatroomTargetMsgRequestV1 ChatroomTargetMsg请求
type ChatroomTargetMsgRequestV1 struct {
	RoomId            int64    `json:"roomid,omitempty"`
	MsgId             string   `json:"msgId,omitempty"`
	Attach            string   `json:"attach,omitempty"`
	FromAccid         string   `json:"fromAccid,omitempty"`
	ToAccids          []string `json:"toAccids,omitempty"`
	MsgType           int      `json:"msgType,omitempty"`
	SubType           int      `json:"subType,omitempty"`
	ResendFlag        int      `json:"resendFlag,omitempty"`
	Route             int      `json:"route,omitempty"`
	Ext               string   `json:"ext,omitempty"`
	UseYidun          int      `json:"useYidun,omitempty"`
	YidunAntiCheating string   `json:"yidunAntiCheating,omitempty"`
	YidunAntiSpamExt  string   `json:"yidunAntiSpamExt,omitempty"`
	Bid               string   `json:"bid,omitempty"`
	Antispam          bool     `json:"antispam,omitempty"`
	AntispamCustom    string   `json:"antispamCustom,omitempty"`
	Env               string   `json:"env,omitempty"`
}

// RecallChatroomMsgRequestV1 RecallChatroomMsg请求
type RecallChatroomMsgRequestV1 struct {
	RoomId      int64  `json:"roomid,omitempty"`
	MsgTimetag  int64  `json:"msgTimetag,omitempty"`
	MsgId       string `json:"msgId,omitempty"`
	FromAcc     string `json:"fromAcc,omitempty"`
	OperatorAcc string `json:"operatorAcc,omitempty"`
	NotifyExt   string `json:"notifyExt,omitempty"`
}

// SendChatroomMsgRequestV1 SendChatroomMsg请求
type SendChatroomMsgRequestV1 struct {
	RoomId                         int64   `json:"roomid,omitempty"`
	MsgId                          string  `json:"msgId,omitempty"`
	Attach                         string  `json:"attach,omitempty"`
	FromAccid                      string  `json:"fromAccid,omitempty"`
	IgnoreMute                     bool    `json:"ignoreMute,omitempty"`
	MsgType                        int     `json:"msgType,omitempty"`
	SubType                        int     `json:"subType,omitempty"`
	ResendFlag                     int     `json:"resendFlag,omitempty"`
	Ext                            string  `json:"ext,omitempty"`
	Route                          int     `json:"route,omitempty"`
	SkipHistory                    int     `json:"skipHistory,omitempty"`
	AbandonRatio                   int     `json:"abandonRatio,omitempty"`
	HighPriority                   bool    `json:"highPriority,omitempty"`
	NeedHighPriorityMsgResend      bool    `json:"needHighPriorityMsgResend,omitempty"`
	UseYidun                       int     `json:"useYidun,omitempty"`
	YidunAntiCheating              string  `json:"yidunAntiCheating,omitempty"`
	YidunAntiSpamExt               string  `json:"yidunAntiSpamExt,omitempty"`
	Bid                            string  `json:"bid,omitempty"`
	Antispam                       bool    `json:"antispam,omitempty"`
	NotifyTargetTags               string  `json:"notifyTargetTags,omitempty"`
	AntispamCustom                 string  `json:"antispamCustom,omitempty"`
	Env                            string  `json:"env,omitempty"`
	ChatMsgPriority                int     `json:"chatMsgPriority,omitempty"`
	ForbiddenIfHighPriorityMsgFreq int     `json:"forbiddenIfHighPriorityMsgFreq,omitempty"`
	LocX                           float64 `json:"locX,omitempty"`
	LocY                           float64 `json:"locY,omitempty"`
	LocZ                           float64 `json:"locZ,omitempty"`
}
