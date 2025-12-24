package account

import (
	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
)

// AccountV2Service Account V2服务接口
type AccountV2Service interface {
	// CreateAccount 创建账号
	CreateAccount(req *CreateAccountRequestV2) (*core.Result[*CreateAccountResponseV2], error)

	// UpdateAccount 更新账号
	UpdateAccount(req *UpdateAccountRequestV2) (*core.Result[*UpdateAccountResponseV2], error)

	// BatchQueryAccounts 批量查询账号
	BatchQueryAccounts(req *BatchQueryAccountsRequestV2) (*core.Result[*BatchQueryAccountsResponseV2], error)

	// DisableAccount 禁用/启用账号
	DisableAccount(req *DisableAccountRequestV2) (*core.Result[*DisableAccountResponseV2], error)

	// SetPushConfig 设置推送配置
	SetPushConfig(req *SetPushConfigRequestV2) (*core.Result[*SetPushConfigResponseV2], error)

	// GetAccountDetails 获取账号详情
	GetAccountDetails(req *GetAccountDetailsRequestV2) (*core.Result[*GetAccountDetailsResponseV2], error)

	// KickAccount 踢账号下线
	KickAccount(req *KickAccountRequestV2) (*core.Result[*KickAccountResponseV2], error)

	// RefreshToken 刷新Token
	RefreshToken(req *RefreshTokenRequestV2) (*core.Result[*RefreshTokenResponseV2], error)
}

// AccountV2ServiceImpl Account V2服务实现
type AccountV2ServiceImpl struct {
	httpClient core.YunxinApiHttpClient
}

// NewAccountV2Service 创建Account V2服务实例
func NewAccountV2Service(httpClient core.YunxinApiHttpClient) AccountV2Service {
	return &AccountV2ServiceImpl{
		httpClient: httpClient,
	}
}
