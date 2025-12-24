package history

// URL路径常量
const (
	QuerySessionMsg              = "/history/querySessionMsg.action"       // 查询单聊云端历史记录
	QueryTeamMsg                 = "/history/queryTeamMsg.action"          // 查询群聊云端历史记录
	QueryChatroomMsg             = "/history/queryChatroomMsg.action"      // 查询聊天室云端历史记录
	DeleteChatroomHistoryMessage = "/chatroom/deleteHistoryMessage.action" // 删除聊天室云端历史记录
	QuerySessionList             = "/history/querySessionList.action"      // 查询会话列表
	QueryBroadcastMsgById        = "/history/queryBroadcastMsgById.action" // 根据ID查询广播消息
	QueryBroadcastMsg            = "/history/queryBroadcastMsg.action"     // 查询广播消息
	QueryUserEvents              = "/history/queryUserEvents.action"       // 查询用户事件
)
