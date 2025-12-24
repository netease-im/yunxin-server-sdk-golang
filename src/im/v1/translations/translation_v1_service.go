package translations

import (
	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
)

// TranslationV1Service translations服务接口
type TranslationV1Service interface {
	// TranslatorText 文本翻译
	TranslatorText(req *TextTranslationRequestV1) (*core.Result[*TextTranslationResponseV1], error)
}

// TranslationV1ServiceImpl translations服务实现
type TranslationV1ServiceImpl struct {
	httpClient core.YunxinApiHttpClient
}

// NewTranslationsV1Service 创建translations服务实例
func NewTranslationsV1Service(httpClient core.YunxinApiHttpClient) TranslationV1Service {
	return &TranslationV1ServiceImpl{
		httpClient: httpClient,
	}
}
