package super_team

// URL路径常量
const (
	Create            = "/nimserver/superteam/create.action"           // 创建超大群
	Update            = "/nimserver/superteam/update.action"           // 更新超大群信息
	Dismiss           = "/nimserver/superteam/dismiss.action"          // 解散超大群
	Invite            = "/nimserver/superteam/invite.action"           // 邀请加入超大群
	KickMember        = "/nimserver/superteam/kick.action"             // 踢出超大群成员
	MemberLeave       = "/nimserver/superteam/leave.action"            // 超大群成员离开
	AddManager        = "/nimserver/superteam/addManager.action"       // 添加管理员
	RemoveManager     = "/nimserver/superteam/removeManager.action"    // 移除管理员
	ChangeOwner       = "/nimserver/superteam/changeOwner.action"      // 移交群主
	GetSuperTeam      = "/nimserver/superteam/query.action"            // 查询超大群信息
	GetMember         = "/nimserver/superteam/queryDetail.action"      // 查询超大群成员
	GetJoinSuperTeam  = "/nimserver/superteam/queryUserJoin.action"    // 查询用户加入的超大群
	UpdateMemberInfo  = "/nimserver/superteam/updateMemberInfo.action" // 更新成员信息
	UpdateNick        = "/nimserver/superteam/updateNick.action"       // 更新群昵称
	Mute              = "/nimserver/superteam/muteTlist.action"        // 禁言群成员
	MuteTlist         = "/nimserver/superteam/muteTlistAll.action"     // 禁言全体成员
	GetMuteMember     = "/nimserver/superteam/listTeamMute.action"     // 获取禁言成员列表
	ChangeLevel       = "/nimserver/superteam/changeLevel.action"      // 更新群成员等级
	SendMessage       = "/nimserver/superteam/sendMsg.action"          // 发送超大群消息
	SendAttachMessage = "/nimserver/superteam/sendAttachMsg.action"    // 发送超大群自定义系统通知
	RecallMessage     = "/nimserver/superteam/recallMsg.action"        // 撤回超大群消息
	GetMessage        = "/nimserver/superteam/queryMsg.action"         // 查询超大群消息
	GetMessageByIds   = "/nimserver/superteam/queryMsgByIds.action"    // 根据ID查询超大群消息
)
