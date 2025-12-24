package history

import (
	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
)

// HistoryV1Service history服务接口
type HistoryV1Service interface {
	// QuerySessionHistoryMessage QuerySessionHistoryMessage
	QuerySessionHistoryMessage(req *QuerySessionHistoryMessageRequestV1) (*core.Result[*QuerySessionHistoryMessageResponseV1], error)

	// QueryTeamHistoryMessage QueryTeamHistoryMessage
	QueryTeamHistoryMessage(req *QueryTeamHistoryMessageRequestV1) (*core.Result[*QueryTeamHistoryMessageResponseV1], error)

	// QueryChatroomHistoryMessage QueryChatroomHistoryMessage
	QueryChatroomHistoryMessage(req *QueryChatroomHistoryMessageRequestV1) (*core.Result[*QueryChatroomHistoryMessageResponseV1], error)

	// DeleteChatroomHistoryMessage DeleteChatroomHistoryMessage
	DeleteChatroomHistoryMessage(req *DeleteChatroomHistoryMessageRequestV1) (*core.Result[*DeleteChatroomHistoryMessageResponseV1], error)

	// QuerySessionList QuerySessionList
	QuerySessionList(req *QuerySessionListRequestV1) (*core.Result[*QuerySessionListResponseV1], error)

	// QueryBroadcastHistoryMessageById QueryBroadcastHistoryMessageById
	QueryBroadcastHistoryMessageById(req *QueryBroadcastHistoryMessageByIdRequestV1) (*core.Result[*QueryBroadcastHistoryMessageByIdResponseV1], error)

	// QueryBroadcastHistoryMessage QueryBroadcastHistoryMessage
	QueryBroadcastHistoryMessage(req *QueryBroadcastHistoryMessageRequestV1) (*core.Result[*QueryBroadcastHistoryMessageResponseV1], error)

	// QueryUserEvents QueryUserEvents
	QueryUserEvents(req *QueryUserEventsRequestV1) (*core.Result[*QueryUserEventsResponseV1], error)
}

// HistoryV1ServiceImpl history服务实现
type HistoryV1ServiceImpl struct {
	httpClient core.YunxinApiHttpClient
}

// NewHistoryV1Service 创建history服务实例
func NewHistoryV1Service(httpClient core.YunxinApiHttpClient) HistoryV1Service {
	return &HistoryV1ServiceImpl{
		httpClient: httpClient,
	}
}
