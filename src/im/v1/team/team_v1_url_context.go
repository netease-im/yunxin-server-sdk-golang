package team

// URL路径常量
const (
	Create                              = "/team/create.action"              // 创建高级群
	Add                                 = "/team/add.action"                 // 拉人入群
	Kick                                = "/team/kick.action"                // 踢人出群
	Leave                               = "/team/leave.action"               // 主动退群
	Update                              = "/team/update.action"              // 更新群资料
	Query                               = "/team/query.action"               // 查询群组信息
	QueryDetail                         = "/team/queryDetail.action"         // 查询群组详细信息
	AddManager                          = "/team/addManager.action"          // 添加管理员
	RemoveManager                       = "/team/removeManager.action"       // 移除管理员
	ChangeOwner                         = "/team/changeOwner.action"         // 移交群主
	Dismiss                             = "/team/remove.action"              // 解散群
	MuteTeam                            = "/team/muteTeam.action"            // 设置群消息提醒开关
	MuteTeamAllMember                   = "/team/muteTlistAll.action"        // 禁言全体成员
	MuteTeamTargetMember                = "/team/muteTlist.action"           // 禁言指定成员
	QueryMuteTeamMembers                = "/team/listTeamMute.action"        // 获取群组禁言列表
	UpdateTeamNick                      = "/team/updateTeamNick.action"      // 修改群昵称
	JoinsTeam                           = "/team/joinTeams.action"           // 查询某用户所加入的群信息
	QueryOnlineTeamMember               = "/team/listOnlineUsers.action"     // 获取群组在线成员列表
	BatchQueryOnlineTeamMemberCount     = "/team/listOnlineUserCount.action" // 批量获取群组在线人数
	QueryTeamMsgMarkReadInfo            = "/team/getMarkReadInfo.action"     // 获取群组已读消息的已读详情信息
	QueryAllJoinedTeamMemberInfoByAccId = "/team/listMemberInfo.action"      // 查询群组成员信息
)
