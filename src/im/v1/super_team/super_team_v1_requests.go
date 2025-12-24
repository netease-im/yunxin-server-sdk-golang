package super_team

// Super Team V1 Request Structures

// GetJoinSuperTeamRequestV1 GetJoinSuperTeam请求
type GetJoinSuperTeamRequestV1 struct {
	Accid string `json:"accid,omitempty"`
}

// GetSuperTeamMemberRequestV1 GetSuperTeamMember请求
type GetSuperTeamMemberRequestV1 struct {
	Tid     int64 `json:"tid,omitempty"`
	Timetag int64 `json:"timetag,omitempty"`
	Limit   int   `json:"limit,omitempty"`
	Reverse int   `json:"reverse,omitempty"`
}

// GetSuperTeamMessageByIdsRequestV1 GetSuperTeamMessageByIds请求
type GetSuperTeamMessageByIdsRequestV1 struct {
	Msgs []Msg `json:"msgs,omitempty"`
}

// Msg Msg
type Msg struct {
	Msgid int64  `json:"msgid,omitempty"`
	From  string `json:"from,omitempty"`
	To    int64  `json:"to,omitempty"`
	Time  int64  `json:"time,omitempty"`
}

// GetSuperTeamMessageRequestV1 GetSuperTeamMessage请求
type GetSuperTeamMessageRequestV1 struct {
	Tid          int64  `json:"tid,omitempty"`
	Accid        string `json:"accid,omitempty"`
	Begintime    int64  `json:"begintime,omitempty"`
	Endtime      int64  `json:"endtime,omitempty"`
	Limit        int    `json:"limit,omitempty"`
	Reverse      int    `json:"reverse,omitempty"`
	Type         string `json:"type,omitempty"`
	ExcludeMsgid string `json:"excludeMsgid,omitempty"`
}

// GetSuperTeamMuteMemberRequestV1 GetSuperTeamMuteMember请求
type GetSuperTeamMuteMemberRequestV1 struct {
	Tid     int64 `json:"tid,omitempty"`
	Timetag int64 `json:"timetag,omitempty"`
	Limit   int   `json:"limit,omitempty"`
	Reverse int   `json:"reverse,omitempty"`
}

// GetSuperTeamRequestV1 GetSuperTeam请求
type GetSuperTeamRequestV1 struct {
	Tids []int64 `json:"tids,omitempty"`
}

// RecallSuperTeamMessageRequestV1 RecallSuperTeamMessage请求
type RecallSuperTeamMessageRequestV1 struct {
	DeleteMsgid int64  `json:"deleteMsgid,omitempty"`
	Timetag     int64  `json:"timetag,omitempty"`
	From        string `json:"from,omitempty"`
	To          int64  `json:"to,omitempty"`
	Msg         string `json:"msg,omitempty"`
	IgnoreTime  int    `json:"ignoreTime,omitempty"`
	PushContent string `json:"pushContent,omitempty"`
	PushPayload string `json:"pushPayload,omitempty"`
}

// SendAttachSuperTeamMessageRequestV1 SendAttachSuperTeamMessage请求
type SendAttachSuperTeamMessageRequestV1 struct {
	From             string   `json:"from,omitempty"`
	To               int64    `json:"to,omitempty"`
	Attach           string   `json:"attach,omitempty"`
	PushContent      string   `json:"pushContent,omitempty"`
	PushPayload      string   `json:"pushPayload,omitempty"`
	Sound            string   `json:"sound,omitempty"`
	Save             int      `json:"save,omitempty"`
	Option           string   `json:"option,omitempty"`
	IsForcePush      bool     `json:"isForcePush,omitempty"`
	ForcePushContent string   `json:"forcePushContent,omitempty"`
	ForcePushAll     bool     `json:"forcePushAll,omitempty"`
	ForcePushList    []string `json:"forcePushList,omitempty"`
}

// SendSuperTeamMessageRequestV1 SendSuperTeamMessage请求
type SendSuperTeamMessageRequestV1 struct {
	Tid               int64    `json:"tid,omitempty"`
	FromAccid         string   `json:"fromAccid,omitempty"`
	Type              int      `json:"type,omitempty"`
	Body              string   `json:"body,omitempty"`
	MsgDesc           string   `json:"msgDesc,omitempty"`
	Antispam          bool     `json:"antispam,omitempty"`
	AntispamCustom    string   `json:"antispamCustom,omitempty"`
	UseYidun          int      `json:"useYidun,omitempty"`
	YidunAntiCheating string   `json:"yidunAntiCheating,omitempty"`
	YidunAntiSpamExt  string   `json:"yidunAntiSpamExt,omitempty"`
	Bid               string   `json:"bid,omitempty"`
	Option            string   `json:"option,omitempty"`
	Ext               string   `json:"ext,omitempty"`
	PushContent       string   `json:"pushContent,omitempty"`
	PushPayload       string   `json:"pushPayload,omitempty"`
	IsForcePush       bool     `json:"isForcePush,omitempty"`
	ForcePushContent  string   `json:"forcePushContent,omitempty"`
	ForcePushAll      bool     `json:"forcePushAll,omitempty"`
	ForcePushList     []string `json:"forcePushList,omitempty"`
	SubType           int      `json:"subType,omitempty"`
	Env               string   `json:"env,omitempty"`
	IsCheckMute       bool     `json:"isCheckMute,omitempty"`
}

// SuperTeamAddManagerRequestV1 SuperTeamAddManager请求
type SuperTeamAddManagerRequestV1 struct {
	Tid           int64    `json:"tid,omitempty"`
	Owner         string   `json:"owner,omitempty"`
	ManagerAccids []string `json:"managerAccids,omitempty"`
}

// SuperTeamChangeLevelRequestV1 SuperTeamChangeLevel请求
type SuperTeamChangeLevelRequestV1 struct {
	Tid    int64  `json:"tid,omitempty"`
	Owner  string `json:"owner,omitempty"`
	Tlevel int    `json:"tlevel,omitempty"`
}

// SuperTeamChangeOwnerRequestV1 SuperTeamChangeOwner请求
type SuperTeamChangeOwnerRequestV1 struct {
	Tid   int64  `json:"tid,omitempty"`
	Owner string `json:"owner,omitempty"`
	Accid string `json:"accid,omitempty"`
	Leave int    `json:"leave,omitempty"`
}

// SuperTeamCreateRequestV1 SuperTeamCreate请求
type SuperTeamCreateRequestV1 struct {
	Owner        string   `json:"owner,omitempty"`
	InviteAccids []string `json:"inviteAccids,omitempty"`
	Tname        string   `json:"tname,omitempty"`
	Intro        string   `json:"intro,omitempty"`
	Announcement string   `json:"announcement,omitempty"`
	ServerCustom string   `json:"serverCustom,omitempty"`
	Icon         string   `json:"icon,omitempty"`
	Msg          string   `json:"msg,omitempty"`
	Joinmode     int      `json:"joinmode,omitempty"`
	Beinvitemode int      `json:"beinvitemode,omitempty"`
	Invitemode   int      `json:"invitemode,omitempty"`
	Uptinfomode  int      `json:"uptinfomode,omitempty"`
	Upcustommode int      `json:"upcustommode,omitempty"`
	Tlevel       int      `json:"tlevel,omitempty"`
	Bid          string   `json:"bid,omitempty"`
}

// SuperTeamDismissRequestV1 SuperTeamDismiss请求
type SuperTeamDismissRequestV1 struct {
	Tid   int64  `json:"tid,omitempty"`
	Owner string `json:"owner,omitempty"`
}

// SuperTeamInviteRequestV1 SuperTeamInvite请求
type SuperTeamInviteRequestV1 struct {
	Tid          int64    `json:"tid,omitempty"`
	Owner        string   `json:"owner,omitempty"`
	InviteAccids []string `json:"inviteAccids,omitempty"`
	Msg          string   `json:"msg,omitempty"`
}

// SuperTeamKickMemberRequestV1 SuperTeamKickMember请求
type SuperTeamKickMemberRequestV1 struct {
	Tid        int64    `json:"tid,omitempty"`
	Owner      string   `json:"owner,omitempty"`
	KickAccids []string `json:"kickAccids,omitempty"`
}

// SuperTeamMemberLeaveRequestV1 SuperTeamMemberLeave请求
type SuperTeamMemberLeaveRequestV1 struct {
	Tid   int64  `json:"tid,omitempty"`
	Accid string `json:"accid,omitempty"`
}

// SuperTeamMuteRequestV1 SuperTeamMute请求
type SuperTeamMuteRequestV1 struct {
	Tid      int64  `json:"tid,omitempty"`
	Owner    string `json:"owner,omitempty"`
	MuteType int    `json:"muteType,omitempty"`
}

// SuperTeamMuteTlistRequestV1 SuperTeamMuteTlist请求
type SuperTeamMuteTlistRequestV1 struct {
	Tid        int64    `json:"tid,omitempty"`
	Owner      string   `json:"owner,omitempty"`
	MuteAccids []string `json:"muteAccids,omitempty"`
	Mute       int      `json:"mute,omitempty"`
}

// SuperTeamRemoveManagerRequestV1 SuperTeamRemoveManager请求
type SuperTeamRemoveManagerRequestV1 struct {
	Tid           int64    `json:"tid,omitempty"`
	Owner         string   `json:"owner,omitempty"`
	ManagerAccids []string `json:"managerAccids,omitempty"`
}

// SuperTeamUpdateMemberInfoRequestV1 SuperTeamUpdateMemberInfo请求
type SuperTeamUpdateMemberInfoRequestV1 struct {
	Tid        int64  `json:"tid,omitempty"`
	Accid      string `json:"accid,omitempty"`
	Nick       string `json:"nick,omitempty"`
	SilentType int    `json:"silentType,omitempty"`
	Custom     string `json:"custom,omitempty"`
}

// SuperTeamUpdateNickRequestV1 SuperTeamUpdateNick请求
type SuperTeamUpdateNickRequestV1 struct {
	Tid   int64  `json:"tid,omitempty"`
	Owner string `json:"owner,omitempty"`
	Accid string `json:"accid,omitempty"`
	Nick  string `json:"nick,omitempty"`
}

// SuperTeamUpdateRequestV1 SuperTeamUpdate请求
type SuperTeamUpdateRequestV1 struct {
	Tid          int64  `json:"tid,omitempty"`
	Owner        string `json:"owner,omitempty"`
	Tname        string `json:"tname,omitempty"`
	Intro        string `json:"intro,omitempty"`
	Announcement string `json:"announcement,omitempty"`
	ServerCustom string `json:"serverCustom,omitempty"`
	Icon         string `json:"icon,omitempty"`
	Joinmode     int    `json:"joinmode,omitempty"`
	Beinvitemode int    `json:"beinvitemode,omitempty"`
	Invitemode   int    `json:"invitemode,omitempty"`
	Uptinfomode  int    `json:"uptinfomode,omitempty"`
	Upcustommode int    `json:"upcustommode,omitempty"`
	Bid          string `json:"bid,omitempty"`
}
