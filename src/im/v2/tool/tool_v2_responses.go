package tool

// AsrResponseV2 语音转文字响应
type AsrResponseV2 struct {
	Text string `json:"text"` // 识别出的文本内容
}

// TranslateTextResponseV2 文本翻译响应
type TranslateTextResponseV2 struct {
	TranslatedText string `json:"translated_text"` // 翻译后的文本
	SourceLanguage string `json:"source_language"` // 源语言
	TargetLanguage string `json:"target_language"` // 目标语言
}
