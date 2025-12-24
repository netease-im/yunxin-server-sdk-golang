package ai

// ChatAssistRequestV2 AI聊天助手请求
type ChatAssistRequestV2 struct {
	SenderId            string        `json:"sender_id"`                       // 发送者ID（必填）
	SenderTagList       []string      `json:"sender_tag_list,omitempty"`       // 发送者标签列表
	ReceiverId          string        `json:"receiver_id"`                     // 接收者ID（必填）
	ReceiverTagList     []string      `json:"receiver_tag_list,omitempty"`     // 接收者标签列表
	ReceiverLastMessage string        `json:"receiver_last_message,omitempty"` // 接收者最后一条消息
	StyleList           []string      `json:"style_list"`                      // 风格列表（必填）
	History             []HistoryItem `json:"history,omitempty"`               // 历史消息列表
}

// HistoryItem 历史消息项
type HistoryItem struct {
	SenderId string `json:"sender_id"` // 发送者ID
	Text     string `json:"text"`      // 消息文本
}
