package signal

// URL路径常量
const (
	CreateSignalRoom       = "/signal/create.action"       // 创建信令房间
	CloseSignalRoom        = "/signal/close.action"        // 关闭信令房间
	InviteSignalRoom       = "/signal/invite.action"       // 邀请进入信令房间
	CancelSignalRoomInvite = "/signal/cancelInvite.action" // 取消信令房间邀请
	KickSignalRoom         = "/signal/kick.action"         // 踢出信令房间
	CtrlSignalRoom         = "/signal/ctrl.action"         // 控制信令房间
	DelaySignalRoom        = "/signal/delay.action"        // 延长信令房间
	GetSignalRoomInfo      = "/signal/getRoomInfo.action"  // 获取信令房间信息
)
