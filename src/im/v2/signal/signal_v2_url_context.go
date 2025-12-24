package signal

// URL路径常量类，用于所有信令相关的API请求 V2
const (
	CreateRoom   = "/im/v2/signalling_room"                       // 创建信令房间
	DelayRoom    = "/im/v2/signalling_room"                       // 延长信令房间有效期 (同一个URL，不同HTTP方法)
	CloseRoom    = "/im/v2/signalling_room"                       // 关闭信令房间 (同一个URL，不同HTTP方法)
	QueryRoom    = "/im/v2/signalling_room"                       // 查询信令房间 (同一个URL，不同HTTP方法)
	SendControl  = "/im/v2/signalling_room/actions/control"       // 信令房间发送控制指令
	Invite       = "/im/v2/signalling_room/actions/invite"        // 邀请加入信令房间
	CancelInvite = "/im/v2/signalling_room/actions/cancel_invite" // 取消邀请加入信令房间
	KickMember   = "/im/v2/signalling_room/actions/kick"          // 将成员踢出信令房间
)
