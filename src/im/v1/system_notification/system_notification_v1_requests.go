package system_notification

// System Notification V1 Request Structures

// SendAttachMsgRequestV1 SendAttachMsg请求
type SendAttachMsgRequestV1 struct {
	From                  string   `json:"from,omitempty"`
	MsgType               int      `json:"msgtype,omitempty"`
	To                    string   `json:"to,omitempty"`
	Attach                string   `json:"attach,omitempty"`
	PushContent           string   `json:"pushcontent,omitempty"`
	Payload               string   `json:"payload,omitempty"`
	Sound                 string   `json:"sound,omitempty"`
	Save                  int      `json:"save,omitempty"`
	Option                string   `json:"option,omitempty"`
	IsForcePush           bool     `json:"isForcePush,omitempty"`
	ForcePushContent      string   `json:"forcePushContent,omitempty"`
	ForcePushAll          bool     `json:"forcePushAll,omitempty"`
	ForcePushList         []string `json:"forcePushList,omitempty"`
	Env                   string   `json:"env,omitempty"`
	CheckAccidAsyncEnable bool     `json:"checkAccidAsyncEnable,omitempty"`
}

// SendBatchAttachMsgRequestV1 SendBatchAttachMsg请求
type SendBatchAttachMsgRequestV1 struct {
	FromAccid             string   `json:"fromAccid,omitempty"`
	ToAccids              []string `json:"toAccids,omitempty"`
	Attach                string   `json:"attach,omitempty"`
	PushContent           string   `json:"pushcontent,omitempty"`
	Payload               string   `json:"payload,omitempty"`
	Sound                 string   `json:"sound,omitempty"`
	Save                  int      `json:"save,omitempty"`
	Option                string   `json:"option,omitempty"`
	IsForcePush           bool     `json:"isForcePush,omitempty"`
	ForcePushContent      string   `json:"forcePushContent,omitempty"`
	Env                   string   `json:"env,omitempty"`
	CheckAccidAsyncEnable bool     `json:"checkAccidAsyncEnable,omitempty"`
}
