package tool

// TranslatorConfig 翻译配置
type TranslatorConfig struct {
	Vendor string `json:"vendor,omitempty"` // 翻译服务提供商 (可选)
}

// AsrRequestV2 语音转文字请求
type AsrRequestV2 struct {
	OperatorAccountId string `json:"operator_account_id"`   // 操作者账号ID (必填)
	Format            string `json:"format"`                // 音频格式，如: wav, opus, mp3, aac (必填)
	Url               string `json:"url"`                   // 音频URL (必填)
	SampleRate        int    `json:"sample_rate,omitempty"` // 采样率 (可选)
	Duration          int64  `json:"duration,omitempty"`    // 音频时长，毫秒 (可选)
}

// TranslateTextRequestV2 文本翻译请求
type TranslateTextRequestV2 struct {
	OperatorAccountId string            `json:"operator_account_id"`         // 操作者账号ID (必填)
	SourceText        string            `json:"source_text"`                 // 待翻译的文本 (必填)
	SourceLanguage    string            `json:"source_language,omitempty"`   // 源语言 (可选，默认auto)
	TargetLanguage    string            `json:"target_language"`             // 目标语言 (必填)
	TranslatorConfig  *TranslatorConfig `json:"translator_config,omitempty"` // 翻译配置 (可选)
}
