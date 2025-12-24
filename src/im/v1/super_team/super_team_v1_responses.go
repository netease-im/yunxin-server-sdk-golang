package super_team

// Super Team V1 Response Structures

// CreateSuperTeamResponseV1 CreateSuperTeam响应
type CreateSuperTeamResponseV1 struct {
	Tid int64 `json:"tid,omitempty"`
}

// GetJoinSuperTeamInfo GetJoinSuperTeamInfo
type GetJoinSuperTeamInfo struct {
	Tid          int64  `json:"tid,omitempty"`
	Tname        string `json:"tname,omitempty"`
	Creator      string `json:"creator,omitempty"`
	Tlevel       int    `json:"tlevel,omitempty"`
	Size         int    `json:"size,omitempty"`
	Intro        string `json:"intro,omitempty"`
	Announcement string `json:"announcement,omitempty"`
	ClientCustom string `json:"clientCustom,omitempty"`
	ServerCustom string `json:"serverCustom,omitempty"`
	Icon         string `json:"icon,omitempty"`
	MuteType     int    `json:"muteType,omitempty"`
	Joinmode     int    `json:"joinmode,omitempty"`
	Beinvitemode int    `json:"beinvitemode,omitempty"`
	Invitemode   int    `json:"invitemode,omitempty"`
	Uptinfomode  int    `json:"uptinfomode,omitempty"`
	Upcustommode int    `json:"upcustommode,omitempty"`
	Createtime   int64  `json:"createtime,omitempty"`
	Updatetime   int64  `json:"updatetime,omitempty"`
}

// GetJoinSuperTeamResponseV1 GetJoinSuperTeam响应
type GetJoinSuperTeamResponseV1 struct {
	Tinfos []GetJoinSuperTeamInfo `json:"tinfos,omitempty"`
}

// GetSuperTeamMemberResponseV1 GetSuperTeamMember响应
type GetSuperTeamMemberResponseV1 struct {
	Tlists []SuperTeamMemberInfo `json:"tlists,omitempty"`
}

// GetSuperTeamMessage GetSuperTeamMessage
type GetSuperTeamMessage struct {
	Msgid    int64  `json:"msgid,omitempty"`
	Sendtime int64  `json:"sendtime,omitempty"`
	From     string `json:"from,omitempty"`
	Type     int    `json:"type,omitempty"`
	Ext      string `json:"ext,omitempty"`
	Body     string `json:"body,omitempty"`
}

// GetSuperTeamMessageByIdsResponseV1 GetSuperTeamMessageByIds响应
type GetSuperTeamMessageByIdsResponseV1 struct {
	Msgs []GetSuperTeamMessage `json:"msgs,omitempty"`
}

// GetSuperTeamMessageResponseV1 GetSuperTeamMessage响应
type GetSuperTeamMessageResponseV1 struct {
	Size int                   `json:"size,omitempty"`
	Tid  int64                 `json:"tid,omitempty"`
	Msgs []GetSuperTeamMessage `json:"msgs,omitempty"`
}

// GetSuperTeamMuteMemberResponseV1 GetSuperTeamMuteMember响应
type GetSuperTeamMuteMemberResponseV1 struct {
	Tlists []SuperTeamMuteMemberInfo `json:"tlists,omitempty"`
}

// GetSuperTeamResponseV1 GetSuperTeam响应
type GetSuperTeamResponseV1 struct {
	Tinfos []SuperTeamInfo `json:"tinfos,omitempty"`
}

// RecallSuperTeamMessageResponseV1 RecallSuperTeamMessage响应
type RecallSuperTeamMessageResponseV1 struct {
}

// SendAttachSuperTeamMessageResponseV1 SendAttachSuperTeamMessage响应
type SendAttachSuperTeamMessageResponseV1 struct {
}

// SendSuperTeamMessageResponseV1 SendSuperTeamMessage响应
type SendSuperTeamMessageResponseV1 struct {
	Msgid    int64 `json:"msgid,omitempty"`
	Timetag  int64 `json:"timetag,omitempty"`
	Antispam bool  `json:"antispam,omitempty"`
}

// SuperTeamAddManagerResponseV1 SuperTeamAddManager响应
type SuperTeamAddManagerResponseV1 struct {
}

// SuperTeamChangeLevelResponseV1 SuperTeamChangeLevel响应
type SuperTeamChangeLevelResponseV1 struct {
}

// SuperTeamChangeOwnerResponseV1 SuperTeamChangeOwner响应
type SuperTeamChangeOwnerResponseV1 struct {
}

// SuperTeamDismissResponseV1 SuperTeamDismiss响应
type SuperTeamDismissResponseV1 struct {
}

// SuperTeamFailAccountList SuperTeamFailAccountList
type SuperTeamFailAccountList struct {
	AccidList []string `json:"accidList,omitempty"`
	Msg       string   `json:"msg,omitempty"`
}

// SuperTeamInfo SuperTeamInfo
type SuperTeamInfo struct {
	Tid          int64  `json:"tid,omitempty"`
	Tname        string `json:"tname,omitempty"`
	Creator      string `json:"creator,omitempty"`
	Tlevel       int    `json:"tlevel,omitempty"`
	Size         int    `json:"size,omitempty"`
	Intro        string `json:"intro,omitempty"`
	Announcement string `json:"announcement,omitempty"`
	ClientCustom string `json:"clientCustom,omitempty"`
	ServerCustom string `json:"serverCustom,omitempty"`
	Icon         string `json:"icon,omitempty"`
	MuteType     int    `json:"muteType,omitempty"`
	Joinmode     int    `json:"joinmode,omitempty"`
	Beinvitemode int    `json:"beinvitemode,omitempty"`
	Invitemode   int    `json:"invitemode,omitempty"`
	Uptinfomode  int    `json:"uptinfomode,omitempty"`
	Upcustommode int    `json:"upcustommode,omitempty"`
	Createtime   int64  `json:"createtime,omitempty"`
	Updatetime   int64  `json:"updatetime,omitempty"`
}

// SuperTeamInviteResponseV1 SuperTeamInvite响应
type SuperTeamInviteResponseV1 struct {
}

// SuperTeamKickMemberResponseV1 SuperTeamKickMember响应
type SuperTeamKickMemberResponseV1 struct {
}

// SuperTeamMemberInfo SuperTeamMemberInfo
type SuperTeamMemberInfo struct {
	Tid        int64  `json:"tid,omitempty"`
	Uid        int64  `json:"uid,omitempty"`
	Accid      string `json:"accid,omitempty"`
	Type       int    `json:"type,omitempty"`
	Nick       string `json:"nick,omitempty"`
	Jointime   string `json:"jointime,omitempty"`
	Mute       int    `json:"mute,omitempty"`
	Invitoruid int64  `json:"invitoruid,omitempty"`
	Invitor    string `json:"invitor,omitempty"`
	Validflag  int    `json:"validflag,omitempty"`
	Custom     string `json:"custom,omitempty"`
	State      int    `json:"state,omitempty"`
	Bits       int64  `json:"bits,omitempty"`
	Createtime int64  `json:"createtime,omitempty"`
	Updatetime int64  `json:"updatetime,omitempty"`
}

// SuperTeamMemberLeaveResponseV1 SuperTeamMemberLeave响应
type SuperTeamMemberLeaveResponseV1 struct {
}

// SuperTeamMuteMemberInfo SuperTeamMuteMemberInfo
type SuperTeamMuteMemberInfo struct {
	Nick       string `json:"nick,omitempty"`
	Createtime int64  `json:"createtime,omitempty"`
	Jointime   int64  `json:"jointime,omitempty"`
	Updatetime int64  `json:"updatetime,omitempty"`
	Custom     string `json:"custom,omitempty"`
	Accid      string `json:"accid,omitempty"`
	Tid        int64  `json:"tid,omitempty"`
	Type       int    `json:"type,omitempty"`
	Mute       int    `json:"mute,omitempty"`
	Invitor    string `json:"invitor,omitempty"`
}

// SuperTeamMuteResponseV1 SuperTeamMute响应
type SuperTeamMuteResponseV1 struct {
}

// SuperTeamMuteTlistResponseV1 SuperTeamMuteTlist响应
type SuperTeamMuteTlistResponseV1 struct {
}

// SuperTeamRemoveManagerResponseV1 SuperTeamRemoveManager响应
type SuperTeamRemoveManagerResponseV1 struct {
}

// SuperTeamUpdateMemberInfoResponseV1 SuperTeamUpdateMemberInfo响应
type SuperTeamUpdateMemberInfoResponseV1 struct {
}

// SuperTeamUpdateNickResponseV1 SuperTeamUpdateNick响应
type SuperTeamUpdateNickResponseV1 struct {
}

// SuperTeamUpdateResponseV1 SuperTeamUpdate响应
type SuperTeamUpdateResponseV1 struct {
}
