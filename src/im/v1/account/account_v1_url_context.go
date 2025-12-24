package account

// URL路径常量，用于所有账户相关的API请求
const (
	CreateAccount     = "/user/create.action"           // 创建账户
	UpdateToken       = "/user/update.action"           // 更新Token
	RefreshToken      = "/user/refreshToken.action"     // 刷新Token
	BlockAccount      = "/user/block.action"            // 封禁账户
	UnblockAccount    = "/user/unblock.action"          // 解除账户封禁
	MuteAccount       = "/user/mute.action"             // 账户禁言
	MuteModule        = "/user/muteModule.action"       // 模块禁言
	SetDonnop         = "/user/setDonnop.action"        // 设置免打扰
	QueryOnlineStatus = "/user/userOnlineStatus.action" // 查询账户在线状态
	QueryUserInfos    = "/user/getUinfos.action"        // 获取用户信息
	UpdateUinfo       = "/user/updateUinfo.action"      // 更新用户信息
)
