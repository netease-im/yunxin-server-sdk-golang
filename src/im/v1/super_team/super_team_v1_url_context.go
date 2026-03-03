package super_team

// URL路径常量，参考 Java SDK
const (
	Create            = "/superteam/create.action"               // 创建超大群
	Update            = "/superteam/updateTinfo.action"          // 更新超大群信息
	Dismiss           = "/superteam/dismiss.action"              // 解散超大群
	Invite            = "/superteam/invite.action"               // 邀请加入超大群
	KickMember        = "/superteam/kick.action"                 // 踢出超大群成员
	MemberLeave       = "/superteam/leave.action"                // 超大群成员离开
	AddManager        = "/superteam/addManager.action"           // 添加管理员
	RemoveManager     = "/superteam/removeManager.action"        // 移除管理员
	ChangeOwner       = "/superteam/changeOwner.action"          // 移交群主
	GetSuperTeam      = "/superteam/getTinfos.action"            // 查询超大群信息
	GetMember         = "/superteam/getTlists.action"            // 查询超大群成员
	GetJoinSuperTeam  = "/superteam/joinTeams.action"            // 查询用户加入的超大群
	UpdateMemberInfo  = "/superteam/updateTlist.action"          // 更新成员信息
	UpdateNick        = "/superteam/updateTeamNick.action"       // 更新群昵称
	Mute              = "/superteam/mute.action"                 // 禁言超大群
	MuteTlist         = "/superteam/muteTlist.action"            // 禁言群成员
	GetMuteMember     = "/superteam/getMuteTlists.action"        // 获取禁言成员列表
	ChangeLevel       = "/superteam/changeLevel.action"          // 更新群成员等级
	SendMessage       = "/superteam/sendMsg.action"              // 发送超大群消息
	SendAttachMessage = "/superteam/sendAttachMsg.action"        // 发送超大群自定义系统通知
	RecallMessage     = "/superteam/recallMsg.action"            // 撤回超大群消息
	GetMessage        = "/superteam/queryHistoryMsg.action"      // 查询超大群历史消息
	GetMessageByIds   = "/superteam/queryHistoryMsgByIds.action" // 根据ID查询超大群历史消息
)
