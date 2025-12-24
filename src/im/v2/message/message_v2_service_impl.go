package message

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/http"
	"github.com/netease-im/yunxin-server-sdk-golang/src/im/v2/util"
)

// SendMessage 发送消息到会话(单聊或群聊)
func (m *MessageV2ServiceImpl) SendMessage(req *SendMessageRequestV2) (*core.Result[*SendMessageResponseV2], error) {
	pathParams := map[string]string{
		"conversation_id": req.ConversationId,
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := m.httpClient.ExecuteV2Api(http.POST, SendMessage, pathParams, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*SendMessageResponseV2](apiResponse)
}

// StreamMessage 发送流式消息
func (m *MessageV2ServiceImpl) StreamMessage(req *StreamMessageRequestV2) (*core.Result[*StreamMessageResponseV2], error) {
	pathParams := map[string]string{
		"conversation_id": req.ConversationId,
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := m.httpClient.ExecuteV2Api(http.POST, StreamMessage, pathParams, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*StreamMessageResponseV2](apiResponse)
}

// BatchSendP2PMessage 批量发送点对点消息
func (m *MessageV2ServiceImpl) BatchSendP2PMessage(req *BatchSendP2PMessageRequestV2) (*core.Result[*BatchSendP2PMessageResponseV2], error) {
	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := m.httpClient.ExecuteV2Api(http.POST, BatchSendP2PMessage, nil, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*BatchSendP2PMessageResponseV2](apiResponse)
}

// ModifyMessage 修改消息
func (m *MessageV2ServiceImpl) ModifyMessage(req *ModifyMessageRequestV2) (*core.Result[*ModifyMessageResponseV2], error) {
	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := m.httpClient.ExecuteV2Api(http.POST, ModifyMessage, nil, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*ModifyMessageResponseV2](apiResponse)
}

// WithdrawMessage 撤回消息
func (m *MessageV2ServiceImpl) WithdrawMessage(req *WithdrawMessageRequestV2) (*core.Result[*WithdrawMessageResponseV2], error) {
	pathParams := map[string]string{
		"conversation_id":   req.ConversationId,
		"message_server_id": req.MessageServerId,
	}

	queryParams := map[string]string{
		"type": strconv.Itoa(req.Type),
	}

	if req.ServerExtension != "" {
		queryParams["server_extension"] = req.ServerExtension
	}
	if req.Notify != "" {
		queryParams["notify"] = req.Notify
	}

	apiResponse, err := m.httpClient.ExecuteV2Api(http.DELETE, WithdrawMessage, pathParams, queryParams, "")
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*WithdrawMessageResponseV2](apiResponse)
}

// DeleteConversationMessages 删除会话所有消息
func (m *MessageV2ServiceImpl) DeleteConversationMessages(req *DeleteConversationMessagesRequestV2) (*core.Result[*DeleteConversationMessagesResponseV2], error) {
	pathParams := map[string]string{
		"conversation_id": req.ConversationId,
	}

	queryParams := map[string]string{
		"delete_type": strconv.Itoa(req.DeleteType),
	}

	apiResponse, err := m.httpClient.ExecuteV2Api(http.DELETE, DeleteConversationMessages, pathParams, queryParams, "")
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*DeleteConversationMessagesResponseV2](apiResponse)
}

// SendP2PReadReceipt 发送点对点消息已读回执
func (m *MessageV2ServiceImpl) SendP2PReadReceipt(req *SendP2PReadReceiptRequestV2) (*core.Result[*SendP2PReadReceiptResponseV2], error) {
	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := m.httpClient.ExecuteV2Api(http.POST, SendP2PReadReceipt, nil, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*SendP2PReadReceiptResponseV2](apiResponse)
}

// SendTeamReadReceipt 发送群消息已读回执
func (m *MessageV2ServiceImpl) SendTeamReadReceipt(req *SendTeamReadReceiptRequestV2) (*core.Result[*SendTeamReadReceiptResponseV2], error) {
	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := m.httpClient.ExecuteV2Api(http.POST, SendTeamReadReceipt, nil, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*SendTeamReadReceiptResponseV2](apiResponse)
}

// QueryTeamReadReceipt 查询群消息已读回执详情
func (m *MessageV2ServiceImpl) QueryTeamReadReceipt(req *QueryTeamReadReceiptRequestV2) (*core.Result[*QueryTeamReadReceiptResponseV2], error) {
	queryParams := map[string]string{
		"send_account_id":   req.SendAccountId,
		"team_id":           req.TeamId,
		"message_server_id": req.MessageServerId,
		"include_account":   fmt.Sprintf("%t", req.IncludeAccount),
	}

	apiResponse, err := m.httpClient.ExecuteV2Api(http.GET, QueryTeamReadReceipt, nil, queryParams, "")
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*QueryTeamReadReceiptResponseV2](apiResponse)
}

// QueryMessage 查询单条消息详情
func (m *MessageV2ServiceImpl) QueryMessage(req *QueryMessageRequestV2) (*core.Result[*QueryMessageResponseV2], error) {
	pathParams := map[string]string{
		"conversation_id":   req.ConversationId,
		"message_server_id": req.MessageServerId,
	}

	apiResponse, err := m.httpClient.ExecuteV2Api(http.GET, QueryMessage, pathParams, nil, "")
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*QueryMessageResponseV2](apiResponse)
}

// SearchMessages 搜索历史消息
func (m *MessageV2ServiceImpl) SearchMessages(req *SearchMessagesRequestV2) (*core.Result[*SearchMessagesResponseV2], error) {
	queryParams := map[string]string{}

	if req.SendAccountId != "" {
		queryParams["send_account_id"] = req.SendAccountId
	}
	if req.ConversationId != "" {
		queryParams["conversation_id"] = req.ConversationId
	}
	if req.MsgType > 0 {
		queryParams["msg_type"] = strconv.Itoa(req.MsgType)
	}
	if req.Keyword != "" {
		queryParams["keyword"] = req.Keyword
	}
	if req.BeginTime > 0 {
		queryParams["begin_time"] = strconv.FormatInt(req.BeginTime, 10)
	}
	if req.EndTime > 0 {
		queryParams["end_time"] = strconv.FormatInt(req.EndTime, 10)
	}
	if req.Limit > 0 {
		queryParams["limit"] = strconv.Itoa(req.Limit)
	}
	queryParams["reverse"] = fmt.Sprintf("%t", req.Reverse)

	apiResponse, err := m.httpClient.ExecuteV2Api(http.GET, SearchMessages, nil, queryParams, "")
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*SearchMessagesResponseV2](apiResponse)
}

// QueryMessagesByPage 分页查询会话消息
func (m *MessageV2ServiceImpl) QueryMessagesByPage(req *QueryMessagesByPageRequestV2) (*core.Result[*QueryMessagesByPageResponseV2], error) {
	pathParams := map[string]string{
		"conversation_id": req.ConversationId,
	}

	queryParams := map[string]string{
		"begin_time": strconv.FormatInt(req.BeginTime, 10),
		"end_time":   strconv.FormatInt(req.EndTime, 10),
		"limit":      strconv.Itoa(req.Limit),
		"reverse":    fmt.Sprintf("%t", req.Reverse),
	}

	if req.MsgTypes != "" {
		queryParams["msg_types"] = req.MsgTypes
	}

	apiResponse, err := m.httpClient.ExecuteV2Api(http.GET, QueryConversationMessages, pathParams, queryParams, "")
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*QueryMessagesByPageResponseV2](apiResponse)
}

// BatchQueryMessages 批量查询消息
func (m *MessageV2ServiceImpl) BatchQueryMessages(req *BatchQueryMessagesByIdRequestV2) (*core.Result[*BatchQueryMessagesByIdResponseV2], error) {
	pathParams := map[string]string{
		"conversation_id": req.ConversationId,
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := m.httpClient.ExecuteV2Api(http.POST, BatchQueryMessages, pathParams, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*BatchQueryMessagesByIdResponseV2](apiResponse)
}

// QueryThreadMessages 查询Thread消息
func (m *MessageV2ServiceImpl) QueryThreadMessages(req *QueryThreadMessagesRequestV2) (*core.Result[*QueryThreadMessagesResponseV2], error) {
	queryParams := map[string]string{
		"thread_root":     req.ThreadRoot,
		"conversation_id": req.ConversationId,
		"begin_time":      strconv.FormatInt(req.BeginTime, 10),
		"end_time":        strconv.FormatInt(req.EndTime, 10),
		"limit":           strconv.Itoa(req.Limit),
		"reverse":         fmt.Sprintf("%t", req.Reverse),
		"exclude_root":    fmt.Sprintf("%t", req.ExcludeRoot),
	}

	apiResponse, err := m.httpClient.ExecuteV2Api(http.GET, QueryThreadMessages, nil, queryParams, "")
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*QueryThreadMessagesResponseV2](apiResponse)
}

// AddQuickComment 添加快捷评论
func (m *MessageV2ServiceImpl) AddQuickComment(req *AddQuickCommentRequestV2) (*core.Result[*AddQuickCommentResponseV2], error) {
	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := m.httpClient.ExecuteV2Api(http.PATCH, AddQuickComment, nil, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*AddQuickCommentResponseV2](apiResponse)
}

// DeleteQuickComment 删除快捷评论
func (m *MessageV2ServiceImpl) DeleteQuickComment(req *DeleteQuickCommentRequestV2) (*core.Result[*DeleteQuickCommentResponseV2], error) {
	queryParams := map[string]string{
		"operator_account_id": req.OperatorAccountId,
		"conversation_id":     req.ConversationId,
		"message_server_id":   req.MessageServerId,
		"comment_idx":         strconv.Itoa(req.CommentIdx),
	}

	apiResponse, err := m.httpClient.ExecuteV2Api(http.DELETE, DeleteQuickComment, nil, queryParams, "")
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*DeleteQuickCommentResponseV2](apiResponse)
}

// BatchQueryQuickComments 批量查询快捷评论
func (m *MessageV2ServiceImpl) BatchQueryQuickComments(req *BatchQueryQuickCommentsRequestV2) (*core.Result[*BatchQueryQuickCommentsResponseV2], error) {
	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := m.httpClient.ExecuteV2Api(http.POST, BatchQueryQuickComments, nil, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*BatchQueryQuickCommentsResponseV2](apiResponse)
}
