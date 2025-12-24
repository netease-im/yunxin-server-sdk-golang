package translations

// Translations V1 Response Structures

// TextTranslationResponseV1 文本翻译响应
type TextTranslationResponseV1 struct {
	Translation string `json:"translation,omitempty"` // 翻译结果
	Language    string `json:"language,omitempty"`    // 源语言和目标语言，由"2"分割，"2"前面是源语言，后面是目标语言
	Timestamp   int64  `json:"timestamp,omitempty"`   // 翻译完成的时间
}
