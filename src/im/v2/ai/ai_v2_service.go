package ai

import (
	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
)

// AiV2Service AI服务接口
type AiV2Service interface {
	// ChatAssist AI聊天助手
	ChatAssist(req *ChatAssistRequestV2) (*core.Result[*ChatAssistResponseV2], error)
}

// AiV2ServiceImpl AI服务实现
type AiV2ServiceImpl struct {
	httpClient core.YunxinApiHttpClient
}

// NewAiV2Service 创建AI服务实例
func NewAiV2Service(httpClient core.YunxinApiHttpClient) AiV2Service {
	return &AiV2ServiceImpl{
		httpClient: httpClient,
	}
}
