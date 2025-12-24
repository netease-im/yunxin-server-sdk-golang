package tool

import (
	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
)

// ToolV2Service 工具服务接口
type ToolV2Service interface {
	// Asr 语音转文字
	Asr(req *AsrRequestV2) (*core.Result[*AsrResponseV2], error)

	// TranslateText 文本翻译
	TranslateText(req *TranslateTextRequestV2) (*core.Result[*TranslateTextResponseV2], error)
}

// ToolV2ServiceImpl 工具服务实现
type ToolV2ServiceImpl struct {
	httpClient core.YunxinApiHttpClient
}

// NewToolV2Service 创建工具服务实例
func NewToolV2Service(httpClient core.YunxinApiHttpClient) ToolV2Service {
	return &ToolV2ServiceImpl{
		httpClient: httpClient,
	}
}
