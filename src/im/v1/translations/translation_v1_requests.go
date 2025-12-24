package translations

// Translations V1 Request Structures

// TextTranslationRequestV1 文本翻译请求
type TextTranslationRequestV1 struct {
	Accid string `json:"accid,omitempty"` // 云信帐号
	Text  string `json:"text,omitempty"`  // 需要翻译的文本
	To    string `json:"to,omitempty"`    // 目标语言
	From  string `json:"from,omitempty"`  // 源语言，不传会自动识别源语言
}
