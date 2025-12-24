package ai

import (
	"encoding/json"
	"errors"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/http"
	"github.com/netease-im/yunxin-server-sdk-golang/src/im/v2/util"
)

// ChatAssist AI聊天助手
func (a *AiV2ServiceImpl) ChatAssist(req *ChatAssistRequestV2) (*core.Result[*ChatAssistResponseV2], error) {
	// 参数校验
	if req.SenderId == "" {
		return nil, errors.New("sender_id is required")
	}
	if req.ReceiverId == "" {
		return nil, errors.New("receiver_id is required")
	}
	if len(req.StyleList) == 0 {
		return nil, errors.New("style_list is required")
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := a.httpClient.ExecuteV2Api(http.POST, ChatAssistURL, nil, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*ChatAssistResponseV2](apiResponse)
}
