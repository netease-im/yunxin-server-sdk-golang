package team

// Team V1 Response Structures

// AddManagerTeamResponseV1 AddManagerTeam响应
type AddManagerTeamResponseV1 struct {
}

// AddTeamResponseV1 AddTeam响应
type AddTeamResponseV1 struct {
	Accids []string `json:"accids,omitempty"`
	Msg    string   `json:"msg,omitempty"`
}

// BatchQueryOnlineTeamMemberCountResponseV1 BatchQueryOnlineTeamMemberCount响应
type BatchQueryOnlineTeamMemberCountResponseV1 struct {
	Data []teamOnlineCount `json:"data,omitempty"`
}

// teamOnlineCount teamOnlineCount
type teamOnlineCount struct {
	Tid              int64 `json:"tid,omitempty"`
	OnlineUserCount  int   `json:"onlineUserCount,omitempty"`
	OfflineUserCount int   `json:"offlineUserCount,omitempty"`
}

// ChangeOwnerTeamResponseV1 ChangeOwnerTeam响应
type ChangeOwnerTeamResponseV1 struct {
}

// CreateTeamResponseV1 CreateTeam响应
type CreateTeamResponseV1 struct {
	Tid int64 `json:"tid,omitempty"`
}

// TeamFailAccids TeamFailAccids
type TeamFailAccids struct {
	AccidList []string `json:"accidList,omitempty"`
	Msg       string   `json:"msg,omitempty"`
}

// DismissTeamResponseV1 DismissTeam响应
type DismissTeamResponseV1 struct {
}

// JoinsTeamResponseV1 JoinsTeam响应
type JoinsTeamResponseV1 struct {
	Infos []JoinsTinfo `json:"infos,omitempty"`
}

// JoinsTinfo JoinsTinfo
type JoinsTinfo struct {
	Owner    string `json:"owner,omitempty"`
	Tname    string `json:"tname,omitempty"`
	MaxUsers int    `json:"maxusers,omitempty"`
	Tid      int64  `json:"tid,omitempty"`
	Size     int    `json:"size,omitempty"`
	Custom   string `json:"custom,omitempty"`
}

// KickTeamResponseV1 KickTeam响应
type KickTeamResponseV1 struct {
	Accid []string `json:"accid,omitempty"`
	Msg   string   `json:"msg,omitempty"`
}

// LeaveTeamResponseV1 LeaveTeam响应
type LeaveTeamResponseV1 struct {
}

// MuteTeamAllMemberResponseV1 MuteTeamAllMember响应
type MuteTeamAllMemberResponseV1 struct {
}

// MuteTeamResponseV1 MuteTeam响应
type MuteTeamResponseV1 struct {
}

// MuteTeamTargetMemberResponseV1 MuteTeamTargetMember响应
type MuteTeamTargetMemberResponseV1 struct {
}

// QueryAllJoinedTeamMemberInfoByAccIdResponseV1 QueryAllJoinedTeamMemberInfoByAccId响应
type QueryAllJoinedTeamMemberInfoByAccIdResponseV1 struct {
	Data []TeamMemberInfo `json:"data,omitempty"`
}

// TeamMemberInfo TeamMemberInfo
type TeamMemberInfo struct {
	Tid               int64  `json:"tid,omitempty"`
	AccId             string `json:"accid,omitempty"`
	Nick              string `json:"nick,omitempty"`
	Mute              bool   `json:"mute,omitempty"`
	ManagerPushEnable bool   `json:"managerPushEnable,omitempty"`
	PushEnable        bool   `json:"pushEnable,omitempty"`
	CreateTime        int64  `json:"createtime,omitempty"`
	UpdateTime        int64  `json:"updatetime,omitempty"`
}

// QueryMuteTeamMembersResponseV1 QueryMuteTeamMembers响应
type QueryMuteTeamMembersResponseV1 struct {
	Mutes []MuteInfo `json:"mutes,omitempty"`
}

// MuteInfo MuteInfo
type MuteInfo struct {
	Nick  string `json:"nick,omitempty"`
	Accid string `json:"accid,omitempty"`
	Tid   int64  `json:"tid,omitempty"`
	Type  int    `json:"type,omitempty"`
}

// QueryOnlineTeamMemberResponseV1 QueryOnlineTeamMember响应
type QueryOnlineTeamMemberResponseV1 struct {
	Count  int            `json:"count,omitempty"`
	Status []OnlineStatus `json:"status,omitempty"`
}

// OnlineStatus OnlineStatus
type OnlineStatus struct {
	Accid      string   `json:"accid,omitempty"`
	StatusList []Status `json:"statusList,omitempty"`
}

// Status Status
type Status struct {
	LoginTime  int64 `json:"loginTime,omitempty"`
	ClientType int   `json:"clientType,omitempty"`
}

// QueryTeamInfoDetailsResponseV1 QueryTeamInfoDetails响应
type QueryTeamInfoDetailsResponseV1 struct {
	Tinfo TeamInfo `json:"tinfo,omitempty"`
}

// TeamInfo TeamInfo
type TeamInfo struct {
	Tname                   string           `json:"tname,omitempty"`
	Icon                    string           `json:"icon,omitempty"`
	Owner                   TeamMemberInfo   `json:"owner,omitempty"`
	MaxUsers                int              `json:"maxusers,omitempty"`
	Tid                     int64            `json:"tid,omitempty"`
	Size                    int              `json:"size,omitempty"`
	Announcement            string           `json:"announcement,omitempty"`
	Intro                   string           `json:"intro,omitempty"`
	JoinMode                int              `json:"joinmode,omitempty"`
	BeInviteMode            int              `json:"beinvitemode,omitempty"`
	InviteMode              int              `json:"invitemode,omitempty"`
	UpTinfoMode             int              `json:"uptinfomode,omitempty"`
	UpCustomMode            int              `json:"upcustommode,omitempty"`
	MuteType                int              `json:"muteType,omitempty"`
	IsNotifyCloseOnline     bool             `json:"isNotifyCloseOnline,omitempty"`
	IsNotifyClosePersistent bool             `json:"isNotifyClosePersistent,omitempty"`
	Custom                  string           `json:"custom,omitempty"`
	ClientCustom            string           `json:"clientCustom,omitempty"`
	Mute                    bool             `json:"mute,omitempty"`
	Admins                  []TeamMemberInfo `json:"admins,omitempty"`
	Members                 []TeamMemberInfo `json:"members,omitempty"`
	CreateTime              int64            `json:"createtime,omitempty"`
	UpdateTime              int64            `json:"updatetime,omitempty"`
}

// QueryTeamMsgMarkReadInfoResponseV1 QueryTeamMsgMarkReadInfo响应
type QueryTeamMsgMarkReadInfoResponseV1 struct {
	Data TeamMsgMarkReadInfo `json:"data,omitempty"`
}

// TeamMsgMarkReadInfo TeamMsgMarkReadInfo
type TeamMsgMarkReadInfo struct {
	ReadSize     int      `json:"readSize,omitempty"`
	UnreadSize   int      `json:"unreadSize,omitempty"`
	ReadAccids   []string `json:"readAccids,omitempty"`
	UnreadAccids []string `json:"unreadAccids,omitempty"`
}

// QueryTeamResponseV1 QueryTeam响应
type QueryTeamResponseV1 struct {
	Tinfos []TeamInfoV1 `json:"tinfos,omitempty"`
}

// TeamInfoV1 TeamInfoV1
type TeamInfoV1 struct {
	Tname                   string   `json:"tname,omitempty"`
	Icon                    string   `json:"icon,omitempty"`
	Owner                   string   `json:"owner,omitempty"`
	Maxusers                int      `json:"maxusers,omitempty"`
	Tid                     int64    `json:"tid,omitempty"`
	Size                    int      `json:"size,omitempty"`
	Announcement            string   `json:"announcement,omitempty"`
	Intro                   string   `json:"intro,omitempty"`
	Joinmode                int      `json:"joinmode,omitempty"`
	Beinvitemode            int      `json:"beinvitemode,omitempty"`
	Uptinfomode             int      `json:"uptinfomode,omitempty"`
	Upcustommode            int      `json:"upcustommode,omitempty"`
	Invitemode              int      `json:"invitemode,omitempty"`
	MuteType                int      `json:"muteType,omitempty"`
	IsNotifyCloseOnline     bool     `json:"isNotifyCloseOnline,omitempty"`
	IsNotifyClosePersistent bool     `json:"isNotifyClosePersistent,omitempty"`
	Custom                  string   `json:"custom,omitempty"`
	ClientCustom            string   `json:"clientCustom,omitempty"`
	Mute                    bool     `json:"mute,omitempty"`
	Admins                  []string `json:"admins,omitempty"`
	Members                 []string `json:"members,omitempty"`
	Createtime              int64    `json:"createtime,omitempty"`
	Updatetime              int64    `json:"updatetime,omitempty"`
}

// RemoveManagerTeamResponseV1 RemoveManagerTeam响应
type RemoveManagerTeamResponseV1 struct {
}

// UpdateTeamNickResponseV1 UpdateTeamNick响应
type UpdateTeamNickResponseV1 struct {
}

// UpdateTeamResponseV1 UpdateTeam响应
type UpdateTeamResponseV1 struct {
}
