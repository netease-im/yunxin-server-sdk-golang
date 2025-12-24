package tool

// AsrResponseV2 语音转文字响应
type AsrResponseV2 struct {
	Text string `json:"text"` // 识别出的文本内容
}

// TranslateTextResponseV2 文本翻译响应
type TranslateTextResponseV2 struct {
	TargetText string `json:"target_text"` // 翻译后的文本
}
