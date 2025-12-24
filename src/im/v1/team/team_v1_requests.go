package team

// Team V1 Request Structures

// AddManagerTeamRequestV1 AddManagerTeam请求
type AddManagerTeamRequestV1 struct {
	Tid     int64    `json:"tid,omitempty"`
	Owner   string   `json:"owner,omitempty"`
	Attach  string   `json:"attach,omitempty"`
	Members []string `json:"members,omitempty"`
}

// AddTeamRequestV1 AddTeam请求
type AddTeamRequestV1 struct {
	Tid     int64    `json:"tid,omitempty"`
	Owner   string   `json:"owner,omitempty"`
	Attach  string   `json:"attach,omitempty"`
	Members []string `json:"members,omitempty"`
	Magree  int      `json:"magree,omitempty"`
	Msg     string   `json:"msg,omitempty"`
}

// BatchQueryOnlineTeamMemberCountRequestV1 BatchQueryOnlineTeamMemberCount请求
type BatchQueryOnlineTeamMemberCountRequestV1 struct {
	Tids []int64 `json:"tids,omitempty"`
}

// ChangeOwnerTeamRequestV1 ChangeOwnerTeam请求
type ChangeOwnerTeamRequestV1 struct {
	Tid      int64  `json:"tid,omitempty"`
	Owner    string `json:"owner,omitempty"`
	NewOwner string `json:"newowner,omitempty"`
	Leave    int    `json:"leave,omitempty"`
	Attach   string `json:"attach,omitempty"`
}

// CreateTeamRequestV1 CreateTeam请求
type CreateTeamRequestV1 struct {
	Attach                  string   `json:"attach,omitempty"`
	Tname                   string   `json:"tname,omitempty"`
	Owner                   string   `json:"owner,omitempty"`
	Announcement            string   `json:"announcement,omitempty"`
	Intro                   string   `json:"intro,omitempty"`
	Custom                  string   `json:"custom,omitempty"`
	Icon                    string   `json:"icon,omitempty"`
	Bid                     string   `json:"bid,omitempty"`
	Joinmode                int      `json:"joinmode,omitempty"`
	Beinvitemode            int      `json:"beinvitemode,omitempty"`
	Invitemode              int      `json:"invitemode,omitempty"`
	Uptinfomode             int      `json:"uptinfomode,omitempty"`
	Upcustommode            int      `json:"upcustommode,omitempty"`
	IsNotifyCloseOnline     int      `json:"isNotifyCloseOnline,omitempty"`
	IsNotifyClosePersistent int      `json:"isNotifyClosePersistent,omitempty"`
	Msg                     string   `json:"msg,omitempty"`
	Members                 []string `json:"members,omitempty"`
	TeamMemberLimit         int      `json:"teamMemberLimit,omitempty"`
	Magree                  int      `json:"magree,omitempty"`
}

// DismissTeamRequestV1 DismissTeam请求
type DismissTeamRequestV1 struct {
	Tid    int64  `json:"tid,omitempty"`
	Attach string `json:"attach,omitempty"`
	Owner  string `json:"owner,omitempty"`
}

// JoinsTeamRequestV1 JoinsTeam请求
type JoinsTeamRequestV1 struct {
	Accid string `json:"accid,omitempty"`
}

// KickTeamRequestV1 KickTeam请求
type KickTeamRequestV1 struct {
	Tid     int64    `json:"tid,omitempty"`
	Owner   string   `json:"owner,omitempty"`
	Attach  string   `json:"attach,omitempty"`
	Members []string `json:"members,omitempty"`
}

// LeaveTeamRequestV1 LeaveTeam请求
type LeaveTeamRequestV1 struct {
	Tid    int64  `json:"tid,omitempty"`
	Attach string `json:"attach,omitempty"`
	Accid  string `json:"accid,omitempty"`
}

// MuteTeamAllMemberRequestV1 MuteTeamAllMember请求
type MuteTeamAllMemberRequestV1 struct {
	Tid      int64  `json:"tid,omitempty"`
	Owner    string `json:"owner,omitempty"`
	Mute     bool   `json:"mute,omitempty"`
	MuteType int    `json:"muteType,omitempty"`
	Attach   string `json:"attach,omitempty"`
}

// MuteTeamRequestV1 MuteTeam请求
type MuteTeamRequestV1 struct {
	Tid   int64  `json:"tid,omitempty"`
	Ope   int    `json:"ope,omitempty"`
	Accid string `json:"accid,omitempty"`
}

// MuteTeamTargetMemberRequestV1 MuteTeamTargetMember请求
type MuteTeamTargetMemberRequestV1 struct {
	Tid    int64  `json:"tid,omitempty"`
	Accid  string `json:"accid,omitempty"`
	Owner  string `json:"owner,omitempty"`
	Mute   int    `json:"mute,omitempty"`
	Attach string `json:"attach,omitempty"`
}

// QueryAllJoinedTeamMemberInfoByAccIdRequestV1 QueryAllJoinedTeamMemberInfoByAccId请求
type QueryAllJoinedTeamMemberInfoByAccIdRequestV1 struct {
	Accid string `json:"accid,omitempty"`
}

// QueryMuteTeamMembersRequestV1 QueryMuteTeamMembers请求
type QueryMuteTeamMembersRequestV1 struct {
	Tid   int64  `json:"tid,omitempty"`
	Owner string `json:"owner,omitempty"`
}

// QueryOnlineTeamMemberRequestV1 QueryOnlineTeamMember请求
type QueryOnlineTeamMemberRequestV1 struct {
	Tid int64 `json:"tid,omitempty"`
}

// QueryTeamInfoDetailsRequestV1 QueryTeamInfoDetails请求
type QueryTeamInfoDetailsRequestV1 struct {
	Tid int64 `json:"tid,omitempty"`
}

// QueryTeamMsgMarkReadInfoRequestV1 QueryTeamMsgMarkReadInfo请求
type QueryTeamMsgMarkReadInfoRequestV1 struct {
	Tid       int64  `json:"tid,omitempty"`
	MsgId     int64  `json:"msgid,omitempty"`
	FromAccid string `json:"fromAccid,omitempty"`
	Snapshot  bool   `json:"snapshot,omitempty"`
}

// QueryTeamRequestV1 QueryTeam请求
type QueryTeamRequestV1 struct {
	Tids []int64 `json:"tids,omitempty"`
	Ope  int     `json:"ope,omitempty"`
}

// RemoveManagerTeamRequestV1 RemoveManagerTeam请求
type RemoveManagerTeamRequestV1 struct {
	Tid     int64    `json:"tid,omitempty"`
	Owner   string   `json:"owner,omitempty"`
	Attach  string   `json:"attach,omitempty"`
	Members []string `json:"members,omitempty"`
}

// UpdateTeamNickRequestV1 UpdateTeamNick请求
type UpdateTeamNickRequestV1 struct {
	Tid    int64  `json:"tid,omitempty"`
	Owner  string `json:"owner,omitempty"`
	Accid  string `json:"accid,omitempty"`
	Nick   string `json:"nick,omitempty"`
	Custom string `json:"custom,omitempty"`
}

// UpdateTeamRequestV1 UpdateTeam请求
type UpdateTeamRequestV1 struct {
	Tid             int64  `json:"tid,omitempty"`
	Tname           string `json:"tname,omitempty"`
	Owner           string `json:"owner,omitempty"`
	Announcement    string `json:"announcement,omitempty"`
	Intro           string `json:"intro,omitempty"`
	Custom          string `json:"custom,omitempty"`
	Icon            string `json:"icon,omitempty"`
	Bid             string `json:"bid,omitempty"`
	Joinmode        int    `json:"joinmode,omitempty"`
	Beinvitemode    int    `json:"beinvitemode,omitempty"`
	Invitemode      int    `json:"invitemode,omitempty"`
	Uptinfomode     int    `json:"uptinfomode,omitempty"`
	Upcustommode    int    `json:"upcustommode,omitempty"`
	Attach          string `json:"attach,omitempty"`
	TeamMemberLimit int    `json:"teamMemberLimit,omitempty"`
}
