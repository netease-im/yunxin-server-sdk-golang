package conversation_unread

import (
	"fmt"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/http"
	"github.com/netease-im/yunxin-server-sdk-golang/src/im/v2/util"
)

// OverViewsConversation 获取会话概览
func (c *ConversationUnreadV2ServiceImpl) OverViewsConversation(req *OverViewsConversationRequestV2) (*core.Result[*OverViewsConversationResponseV2], error) {
	// Validate required parameters
	if req.AccountId == "" {
		return nil, fmt.Errorf("account ID cannot be empty")
	}

	pathParams := map[string]string{
		"account_id": req.AccountId,
	}

	apiResponse, err := c.httpClient.ExecuteV2Api(http.GET, OverviewsConversation, pathParams, nil, "")
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*OverViewsConversationResponseV2](apiResponse)
}

// ClearConversationUnread 清除会话未读数
func (c *ConversationUnreadV2ServiceImpl) ClearConversationUnread(req *ClearConversationUnreadRequestV2) (*core.Result[*ClearConversationUnreadResponseV2], error) {
	// Validate required parameters
	if req.ConversationId == "" {
		return nil, fmt.Errorf("conversation ID cannot be empty")
	}

	pathParams := map[string]string{
		"conversation_id": req.ConversationId,
	}

	apiResponse, err := c.httpClient.ExecuteV2Api(http.POST, ClearConversationUnread, pathParams, nil, "")
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*ClearConversationUnreadResponseV2](apiResponse)
}
