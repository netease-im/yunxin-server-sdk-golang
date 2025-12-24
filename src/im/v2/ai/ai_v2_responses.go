package ai

// ChatAssistResponseV2 AI聊天助手响应
type ChatAssistResponseV2 struct {
	Items []AssistItem `json:"items,omitempty"` // 助手回复项列表
}

// AssistItem 助手回复项
type AssistItem struct {
	Style     string `json:"style,omitempty"`      // 风格
	StyleName string `json:"style_name,omitempty"` // 风格名称
	Answer    string `json:"answer,omitempty"`     // 回复内容
}
