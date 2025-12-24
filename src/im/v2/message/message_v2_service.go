package message

import (
	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
)

// MessageV2Service Interface for Message V2 operations
type MessageV2Service interface {
	// SendMessage 发送消息到会话(单聊或群聊)
	SendMessage(req *SendMessageRequestV2) (*core.Result[*SendMessageResponseV2], error)

	// StreamMessage 发送流式消息
	StreamMessage(req *StreamMessageRequestV2) (*core.Result[*StreamMessageResponseV2], error)

	// BatchSendP2PMessage 批量发送点对点消息
	BatchSendP2PMessage(req *BatchSendP2PMessageRequestV2) (*core.Result[*BatchSendP2PMessageResponseV2], error)

	// ModifyMessage 修改消息
	ModifyMessage(req *ModifyMessageRequestV2) (*core.Result[*ModifyMessageResponseV2], error)

	// WithdrawMessage 撤回消息
	WithdrawMessage(req *WithdrawMessageRequestV2) (*core.Result[*WithdrawMessageResponseV2], error)

	// DeleteConversationMessages 删除会话所有消息
	DeleteConversationMessages(req *DeleteConversationMessagesRequestV2) (*core.Result[*DeleteConversationMessagesResponseV2], error)

	// SendP2PReadReceipt 发送点对点消息已读回执
	SendP2PReadReceipt(req *SendP2PReadReceiptRequestV2) (*core.Result[*SendP2PReadReceiptResponseV2], error)

	// SendTeamReadReceipt 发送群消息已读回执
	SendTeamReadReceipt(req *SendTeamReadReceiptRequestV2) (*core.Result[*SendTeamReadReceiptResponseV2], error)

	// QueryTeamReadReceipt 查询群消息已读回执详情
	QueryTeamReadReceipt(req *QueryTeamReadReceiptRequestV2) (*core.Result[*QueryTeamReadReceiptResponseV2], error)

	// QueryMessage 查询单条消息详情
	QueryMessage(req *QueryMessageRequestV2) (*core.Result[*QueryMessageResponseV2], error)

	// SearchMessages 搜索历史消息
	SearchMessages(req *SearchMessagesRequestV2) (*core.Result[*SearchMessagesResponseV2], error)

	// QueryMessagesByPage 分页查询会话消息
	QueryMessagesByPage(req *QueryMessagesByPageRequestV2) (*core.Result[*QueryMessagesByPageResponseV2], error)

	// BatchQueryMessages 批量查询消息
	BatchQueryMessages(req *BatchQueryMessagesByIdRequestV2) (*core.Result[*BatchQueryMessagesByIdResponseV2], error)

	// QueryThreadMessages 查询Thread消息
	QueryThreadMessages(req *QueryThreadMessagesRequestV2) (*core.Result[*QueryThreadMessagesResponseV2], error)

	// AddQuickComment 添加快捷评论
	AddQuickComment(req *AddQuickCommentRequestV2) (*core.Result[*AddQuickCommentResponseV2], error)

	// DeleteQuickComment 删除快捷评论
	DeleteQuickComment(req *DeleteQuickCommentRequestV2) (*core.Result[*DeleteQuickCommentResponseV2], error)

	// BatchQueryQuickComments 批量查询快捷评论
	BatchQueryQuickComments(req *BatchQueryQuickCommentsRequestV2) (*core.Result[*BatchQueryQuickCommentsResponseV2], error)
}

// MessageV2ServiceImpl Message V2服务实现
type MessageV2ServiceImpl struct {
	httpClient core.YunxinApiHttpClient
}

// NewMessageV2Service 创建Message V2服务实例
func NewMessageV2Service(httpClient core.YunxinApiHttpClient) MessageV2Service {
	return &MessageV2ServiceImpl{
		httpClient: httpClient,
	}
}
