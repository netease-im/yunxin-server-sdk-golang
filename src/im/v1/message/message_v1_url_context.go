package message

// URL路径常量
const (
	SendMessage                = "/nimserver/msg/sendMsg.action"                // 发送单聊消息
	SendBatchMessage           = "/nimserver/msg/sendBatchMsg.action"           // 批量发送单聊消息
	UploadFile                 = "/nimserver/msg/upload.action"                 // 文件上传
	RecallMessage              = "/nimserver/msg/recall.action"                 // 消息撤回
	BroadcastMessage           = "/nimserver/msg/broadcastMsg.action"           // 发送广播消息
	DeleteBroadcastMessageById = "/nimserver/msg/deleteBroadcastMsgById.action" // 根据ID删除广播消息
	DeleteMessage              = "/nimserver/msg/delete.action"                 // 删除消息
	DeleteMessageOneWay        = "/nimserver/msg/deleteOneWay.action"           // 单向删除消息
	DeleteFile                 = "/nimserver/msg/deleteFile.action"             // 删除文件
	DeleteRoamSession          = "/nimserver/msg/deleteRoamSession.action"      // 删除漫游会话
	MarkReadMessage            = "/nimserver/msg/markRead.action"               // 标记消息已读
	MarkReadTeamMessage        = "/nimserver/msg/markTeamRead.action"           // 标记群消息已读
)
