package team

// URL路径常量
const (
	Create                              = "/nimserver/team/create.action"                              // 创建高级群
	Add                                 = "/nimserver/team/add.action"                                 // 拉人入群
	Kick                                = "/nimserver/team/kick.action"                                // 踢人出群
	Leave                               = "/nimserver/team/leave.action"                               // 主动退群
	Update                              = "/nimserver/team/update.action"                              // 更新群资料
	Query                               = "/nimserver/team/query.action"                               // 查询群组信息
	QueryDetail                         = "/nimserver/team/queryDetail.action"                         // 查询群组详细信息
	AddManager                          = "/nimserver/team/addManager.action"                          // 添加管理员
	RemoveManager                       = "/nimserver/team/removeManager.action"                       // 移除管理员
	ChangeOwner                         = "/nimserver/team/changeOwner.action"                         // 移交群主
	Dismiss                             = "/nimserver/team/dismiss.action"                             // 解散群
	MuteTeam                            = "/nimserver/team/muteTlist.action"                           // 禁言群成员
	MuteTeamAllMember                   = "/nimserver/team/muteTlistAll.action"                        // 禁言全体成员
	MuteTeamTargetMember                = "/nimserver/team/muteTeamTargetMember.action"                // 禁言指定成员
	QueryMuteTeamMembers                = "/nimserver/team/listTeamMute.action"                        // 获取群组禁言列表
	UpdateTeamNick                      = "/nimserver/team/updateTeamNick.action"                      // 修改群昵称
	JoinsTeam                           = "/nimserver/team/join.action"                                // 批量加入群
	QueryOnlineTeamMember               = "/nimserver/team/queryOnlineMemberCount.action"              // 查询在线群成员
	BatchQueryOnlineTeamMemberCount     = "/nimserver/team/batchQueryOnlineMemberCount.action"         // 批量查询在线群成员数
	QueryTeamMsgMarkReadInfo            = "/nimserver/team/queryTeamMsgMarkReadInfo.action"            // 查询群消息已读信息
	QueryAllJoinedTeamMemberInfoByAccId = "/nimserver/team/queryAllJoinedTeamMemberInfoByAccId.action" // 查询用户加入的所有群成员信息
)
