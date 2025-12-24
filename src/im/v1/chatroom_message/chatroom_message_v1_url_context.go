package chatroom_message

// URL路径常量，用于所有聊天室消息相关的API请求
const (
	SendMsg               = "/chatroom/sendMsg.action"               // 发送聊天室消息
	BatchSendMsg          = "/chatroom/batchSendMsg.action"          // 批量发送聊天室消息
	Recall                = "/chatroom/recall.action"                // 撤回聊天室消息
	SendMsgToSomeone      = "/chatroom/sendMsgToSomeone.action"      // 发送聊天室定向消息
	BatchSendMsgToSomeone = "/chatroom/batchSendMsgToSomeone.action" // 批量发送聊天室定向消息
	Broadcast             = "/chatroom/broadcast.action"             // 发送聊天室全服广播消息
)
