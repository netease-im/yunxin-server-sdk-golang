package message

// Message V1 Response Structures

// BroadcastMessageResponseV1 BroadcastMessage响应
type BroadcastMessageResponseV1 struct {
	BroadcastId int64    `json:"broadcastId,omitempty"`
	From        string   `json:"from,omitempty"`
	Body        string   `json:"body,omitempty"`
	TargetOs    []string `json:"targetOs,omitempty"`
	IsOffline   bool     `json:"isOffline,omitempty"`
	CreateTime  int64    `json:"createTime,omitempty"`
	ExpireTime  int64    `json:"expireTime,omitempty"`
}

// DeleteBroadcastMessageByIdResponseV1 DeleteBroadcastMessageById响应
type DeleteBroadcastMessageByIdResponseV1 struct {
}

// DeleteFileResponseV1 DeleteFile响应
type DeleteFileResponseV1 struct {
	Taskid string `json:"taskid,omitempty"`
}

// DeleteMessageOneWayResponseV1 DeleteMessageOneWay响应
type DeleteMessageOneWayResponseV1 struct {
}

// DeleteMessageResponseV1 DeleteMessage响应
type DeleteMessageResponseV1 struct {
}

// DeleteRoamSessionResponseV1 DeleteRoamSession响应
type DeleteRoamSessionResponseV1 struct {
}

// MarkReadMessageResponseV1 MarkReadMessage响应
type MarkReadMessageResponseV1 struct {
	Timestamp int64 `json:"timestamp,omitempty"`
}

// MarkReadTeamMessageResponseV1 MarkReadTeamMessage响应
type MarkReadTeamMessageResponseV1 struct {
	ErrMsgIds []int64 `json:"errMsgIds,omitempty"`
}

// RecallMessageResponseV1 RecallMessage响应
type RecallMessageResponseV1 struct {
}

// SendBatchMessageResponseV1 SendBatchMessage响应
type SendBatchMessageResponseV1 struct {
	Unregister []string         `json:"unregister,omitempty"`
	Timetag    int64            `json:"timetag,omitempty"`
	Msgids     map[string]int64 `json:"msgids,omitempty"`
}

// SendMessageResponseV1 SendMessage响应
type SendMessageResponseV1 struct {
	Msgid    int64 `json:"msgid,omitempty"`
	Antispam bool  `json:"antispam,omitempty"`
	Timetag  int64 `json:"timetag,omitempty"`
}

// UploadFileResponseV1 UploadFile响应
type UploadFileResponseV1 struct {
	Url string `json:"url,omitempty"`
}
