package message

// Message V1 Request Structures

// BroadcastMessageRequestV1 BroadcastMessage请求
type BroadcastMessageRequestV1 struct {
	Body      string   `json:"body,omitempty"`
	From      string   `json:"from,omitempty"`
	IsOffline bool     `json:"isOffline,omitempty"`
	Ttl       int      `json:"ttl,omitempty"`
	TargetOs  []string `json:"targetOs,omitempty"`
}

// DeleteBroadcastMessageByIdRequestV1 DeleteBroadcastMessageById请求
type DeleteBroadcastMessageByIdRequestV1 struct {
	BroadcastId int64 `json:"broadcastId,omitempty"`
}

// DeleteFileRequestV1 DeleteFile请求
type DeleteFileRequestV1 struct {
	StartTime   int64  `json:"startTime,omitempty"`
	EndTime     int64  `json:"endTime,omitempty"`
	ContentType string `json:"contentType,omitempty"`
	Tag         string `json:"tag,omitempty"`
}

// DeleteMessageOneWayRequestV1 DeleteMessageOneWay请求
type DeleteMessageOneWayRequestV1 struct {
	DeleteMsgid int64  `json:"deleteMsgid,omitempty"`
	Timetag     int64  `json:"timetag,omitempty"`
	Type        int    `json:"type,omitempty"`
	From        string `json:"from,omitempty"`
	To          string `json:"to,omitempty"`
	Msg         string `json:"msg,omitempty"`
}

// DeleteMessageRequestV1 DeleteMessage请求
type DeleteMessageRequestV1 struct {
	DeleteMsgid int64  `json:"deleteMsgid,omitempty"`
	Timetag     int64  `json:"timetag,omitempty"`
	Type        int    `json:"type,omitempty"`
	From        string `json:"from,omitempty"`
	To          string `json:"to,omitempty"`
}

// DeleteRoamSessionRequestV1 DeleteRoamSession请求
type DeleteRoamSessionRequestV1 struct {
	Type int    `json:"type,omitempty"`
	From string `json:"from,omitempty"`
	To   string `json:"to,omitempty"`
}

// MarkReadMessageRequestV1 MarkReadMessage请求
type MarkReadMessageRequestV1 struct {
	From        string `json:"from,omitempty"`
	To          string `json:"to,omitempty"`
	MsgidClient string `json:"msgidClient,omitempty"`
}

// MarkReadTeamMessageRequestV1 MarkReadTeamMessage请求
type MarkReadTeamMessageRequestV1 struct {
	Accid  string  `json:"accid,omitempty"`
	Tid    int64   `json:"tid,omitempty"`
	Msgids []int64 `json:"msgids,omitempty"`
}

// RecallMessageRequestV1 RecallMessage请求
type RecallMessageRequestV1 struct {
	DeleteMsgid int64  `json:"deleteMsgid,omitempty"`
	Timetag     int64  `json:"timetag,omitempty"`
	Type        int    `json:"type,omitempty"`
	From        string `json:"from,omitempty"`
	To          string `json:"to,omitempty"`
	Msg         string `json:"msg,omitempty"`
	IgnoreTime  string `json:"ignoreTime,omitempty"`
	Pushcontent string `json:"pushcontent,omitempty"`
	Payload     string `json:"payload,omitempty"`
	Env         string `json:"env,omitempty"`
	Attach      string `json:"attach,omitempty"`
}

// SendBatchMessageRequestV1 SendBatchMessage请求
type SendBatchMessageRequestV1 struct {
	FromAccid         string   `json:"fromAccid,omitempty"`
	ToAccids          []string `json:"toAccids,omitempty"`
	Type              int      `json:"type,omitempty"`
	Body              string   `json:"body,omitempty"`
	MsgDesc           string   `json:"msgDesc,omitempty"`
	Option            string   `json:"option,omitempty"`
	PushContent       string   `json:"pushContent,omitempty"`
	Payload           string   `json:"payload,omitempty"`
	Ext               string   `json:"ext,omitempty"`
	UseYidun          int      `json:"useYidun,omitempty"`
	Bid               string   `json:"bid,omitempty"`
	YidunAntiCheating string   `json:"yidunAntiCheating,omitempty"`
	YidunAntiSpamExt  string   `json:"yidunAntiSpamExt,omitempty"`
	ReturnMsgid       bool     `json:"returnMsgid,omitempty"`
	Env               string   `json:"env,omitempty"`
	MsgSenderNoSense  int      `json:"msgSenderNoSense,omitempty"`
}

// SendMessageRequestV1 SendMessage请求
type SendMessageRequestV1 struct {
	From               string   `json:"from,omitempty"`
	Ope                int      `json:"ope,omitempty"`
	To                 string   `json:"to,omitempty"`
	Type               int      `json:"type,omitempty"`
	Body               string   `json:"body,omitempty"`
	MsgDesc            string   `json:"msgDesc,omitempty"`
	Antispam           bool     `json:"antispam,omitempty"`
	AntispamCustom     string   `json:"antispamCustom,omitempty"`
	Option             string   `json:"option,omitempty"`
	PushContent        string   `json:"pushContent,omitempty"`
	Payload            string   `json:"payload,omitempty"`
	Ext                string   `json:"ext,omitempty"`
	ForcePushAll       bool     `json:"forcePushAll,omitempty"`
	ForcePushList      []string `json:"forcePushList,omitempty"`
	ForcePushContent   string   `json:"forcePushContent,omitempty"`
	UseYidun           int      `json:"useYidun,omitempty"`
	Bid                string   `json:"bid,omitempty"`
	YidunAntiCheating  string   `json:"yidunAntiCheating,omitempty"`
	YidunAntiSpamExt   string   `json:"yidunAntiSpamExt,omitempty"`
	MarkRead           int      `json:"markRead,omitempty"`
	Async              bool     `json:"async,omitempty"`
	CheckFriend        bool     `json:"checkFriend,omitempty"`
	SubType            int      `json:"subType,omitempty"`
	MsgSenderNoSense   int      `json:"msgSenderNoSense,omitempty"`
	MsgReceiverNoSense int      `json:"msgReceiverNoSense,omitempty"`
	Env                string   `json:"env,omitempty"`
	RobotAccount       string   `json:"robotAccount,omitempty"`
	RobotTopic         string   `json:"robotTopic,omitempty"`
	RobotFunction      string   `json:"robotFunction,omitempty"`
	RobotCustomContent string   `json:"robotCustomContent,omitempty"`
}

// UploadFileRequestV1 UploadFile请求
type UploadFileRequestV1 struct {
	Content   string `json:"content,omitempty"`
	Type      string `json:"type,omitempty"`
	Ishttps   bool   `json:"ishttps,omitempty"`
	ExpireSec int    `json:"expireSec,omitempty"`
	Tag       string `json:"tag,omitempty"`
}
