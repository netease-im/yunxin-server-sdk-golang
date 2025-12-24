package chatroom

// URL路径常量，用于所有聊天室相关的API请求
const (
	Create                  = "/chatroom/create.action"                   // 创建聊天室
	Update                  = "/chatroom/update.action"                   // 更新聊天室信息
	RequestAddr             = "/chatroom/requestAddr.action"              // 查询聊天室地址
	Get                     = "/chatroom/get.action"                      // 获取聊天室信息
	GetBatch                = "/chatroom/getBatch.action"                 // 批量获取聊天室信息
	ToggleCloseStat         = "/chatroom/toggleCloseStat.action"          // 切换聊天室关闭状态
	UpdateDelayClosePolicy  = "/chatroom/updateDelayClosePolicy.action"   // 更新聊天室延迟关闭策略
	UpdateInOutNotification = "/chatroom/updateInOutNotification.action"  // 更新聊天室进出通知
	KickMember              = "/chatroom/kickMember.action"               // 踢出聊天室成员
	SetMemberRole           = "/chatroom/setMemberRole.action"            // 设置聊天室成员角色
	UpdateMyRoomRole        = "/chatroom/updateMyRoomRole.action"         // 更新我的聊天室角色
	MembersByPage           = "/chatroom/membersByPage.action"            // 分页获取聊天室成员
	QueryMembersByRole      = "/chatroom/queryMembersByRole.action"       // 按角色获取聊天室成员
	QueryMembers            = "/chatroom/queryMembers.action"             // 查询聊天室成员
	AddRobot                = "/chatroom/addRobot.action"                 // 添加机器人
	RemoveRobot             = "/chatroom/removeRobot.action"              // 移除机器人
	CleanRobot              = "/chatroom/cleanRobot.action"               // 清空机器人
	TemporaryMute           = "/chatroom/temporaryMute.action"            // 设置临时禁言状态
	MuteRoom                = "/chatroom/muteRoom.action"                 // 将聊天室整体禁言
	TagMute                 = "/chatroom/tagTemporaryMute.action"         // 标签临时禁言
	TagCount                = "/chatroom/tagMembersCount.action"          // 标签成员计数
	TagQuery                = "/chatroom/tagMembersQuery.action"          // 标签成员查询
	QueryTagMsg             = "/history/queryTagHistoryMsg.action"        // 查询标签历史消息
	UpdateChatRoomRoleTag   = "/chatroom/updateChatRoomRoleTag.action"    // 更新聊天室角色标签
	QueryUserRoomIds        = "/chatroom/queryUserRoomIds.action"         // 查询开放状态的聊天室
	QueueInit               = "/chatroom/queueInit.action"                // 初始化队列
	QueueDrop               = "/chatroom/queueDrop.action"                // 删除清理队列
	QueueOffer              = "/chatroom/queueOffer.action"               // 新加元素到队列
	QueueBatchOffer         = "/chatroom/queueBatchOffer.action"          // 批量添加队列元素
	QueueBatchUpdate        = "/chatroom/queueBatchUpdateElements.action" // 批量更新队列元素
	QueueList               = "/chatroom/queueList.action"                // 排序列出队列中所有元素
	QueuePoll               = "/chatroom/queuePoll.action"                // 从队列中取出元素
	QueueMultiGet           = "/chatroom/queueMultiGet.action"            // 获取指定的队列元素
)
