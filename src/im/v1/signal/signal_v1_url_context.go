package signal

// URL路径常量，用于所有信令相关的API请求
const (
	CreateSignalRoom       = "/signal/createRoom.action"   // 创建信令房间
	DelaySignalRoom        = "/signal/delayRoom.action"    // 延长信令房间有效期
	CloseSignalRoom        = "/signal/closeRoom.action"    // 关闭信令房间
	GetSignalRoomInfo      = "/signal/getRoomInfo.action"  // 查询信令房间信息
	CtrlSignalRoom         = "/signal/ctrlRoom.action"     // 发送控制指令
	InviteSignalRoom       = "/signal/invite.action"       // 邀请加入信令房间
	CancelSignalRoomInvite = "/signal/cancelInvite.action" // 取消邀请加入信令房间
	KickSignalRoom         = "/signal/kick.action"         // 将成员踢出信令房间
)
