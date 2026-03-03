package message

// URL路径常量
const (
	SendMessage                = "/msg/sendMsg.action"                 // 发送单聊消息
	SendBatchMessage           = "/msg/sendBatchMsg.action"            // 批量发送单聊消息
	UploadFile                 = "/msg/upload.action"                  // 文件上传
	RecallMessage              = "/msg/recall.action"                  // 消息撤回
	BroadcastMessage           = "/msg/broadcastMsg.action"            // 发送广播消息
	DeleteBroadcastMessageById = "/history/delBroadcastMsgById.action" // 根据ID删除广播消息
	DeleteMessage              = "/msg/delMsg.action"                  // 删除消息
	DeleteMessageOneWay        = "/msg/delMsgOneWay.action"            // 单向删除消息
	DeleteFile                 = "/job/nos/del.action"                 // 删除文件
	DeleteRoamSession          = "/msg/delRoamSession.action"          // 删除漫游会话
	MarkReadMessage            = "/msg/markReadMsg.action"             // 标记消息已读
	MarkReadTeamMessage        = "/msg/markReadTeamMsg.action"         // 标记群消息已读
)
