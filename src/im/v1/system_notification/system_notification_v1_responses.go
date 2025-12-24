package system_notification

// System Notification V1 Response Structures

// SendAttachMsgResponseV1 SendAttachMsg响应
type SendAttachMsgResponseV1 struct {
}

// SendBatchAttachMsgResponseV1 SendBatchAttachMsg响应
type SendBatchAttachMsgResponseV1 struct {
	Unregister []string `json:"unregister,omitempty"`
}
