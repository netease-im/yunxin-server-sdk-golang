package account

import (
	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
)

// AccountV1Service 账户服务接口
type AccountV1Service interface {
	// CreateAccount 创建账户
	CreateAccount(req *CreateAccountRequestV1) (*core.Result[*CreateAccountResponseV1], error)

	// UpdateToken 更新Token
	UpdateToken(req *UpdateTokenRequestV1) (*core.Result[*UpdateTokenResponseV1], error)

	// RefreshToken 刷新Token
	RefreshToken(req *RefreshTokenRequestV1) (*core.Result[*RefreshTokenResponseV1], error)

	// Block 封禁账户
	Block(req *BlockAccountRequestV1) (*core.Result[*BlockAccountResponseV1], error)

	// Unblock 解除账户封禁
	Unblock(req *UnBlockAccountRequestV1) (*core.Result[*UnBlockAccountResponseV1], error)

	// Mute 账户禁言
	Mute(req *MuteAccountRequestV1) (*core.Result[*MuteAccountResponseV1], error)

	// MuteModule 模块禁言
	MuteModule(req *MuteModuleRequestV1) (*core.Result[*MuteModuleResponseV1], error)

	// SetDonnop 设置免打扰
	SetDonnop(req *SetDonnopRequestV1) (*core.Result[*SetDonnopResponseV1], error)

	// QueryOnlineStatus 查询账户在线状态
	QueryOnlineStatus(req *QueryAccountOnlineStatusRequestV1) (*core.Result[*QueryAccountOnlineStatusResponseV1], error)

	// QueryUserInfos 获取用户信息
	QueryUserInfos(req *QueryUserInfosRequestV1) (*core.Result[*QueryUserInfosResponseV1], error)

	// UpdateUinfo 更新用户信息
	UpdateUinfo(req *UpdateUinfoRequestV1) (*core.Result[*UpdateUinfoResponseV1], error)
}

// AccountV1ServiceImpl 账户服务实现
type AccountV1ServiceImpl struct {
	httpClient core.YunxinApiHttpClient
}

// NewAccountV1Service 创建账户服务实例
func NewAccountV1Service(httpClient core.YunxinApiHttpClient) AccountV1Service {
	return &AccountV1ServiceImpl{
		httpClient: httpClient,
	}
}
