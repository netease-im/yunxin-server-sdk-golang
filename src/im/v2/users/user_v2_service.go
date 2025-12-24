package users

import (
	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
)

// UserV2Service 用户服务接口
type UserV2Service interface {
	// UpdateUser 更新用户信息
	UpdateUser(req *UpdateUserRequestV2) (*core.Result[*UpdateUserResponseV2], error)

	// GetUser 获取用户信息
	GetUser(req *GetUserRequestV2) (*core.Result[*GetUserResponseV2], error)

	// BatchGetUsers 批量获取用户信息
	BatchGetUsers(req *BatchGetUsersRequestV2) (*core.Result[*BatchGetUsersResponseV2], error)

	// GetUsersOnlineStatus 获取用户在线状态
	GetUsersOnlineStatus(req *GetUserOnlineStatusRequestV2) (*core.Result[*GetUserOnlineStatusResponseV2], error)
}

// UserV2ServiceImpl 用户服务实现
type UserV2ServiceImpl struct {
	httpClient core.YunxinApiHttpClient
}

// NewUserV2Service 创建用户服务实例
func NewUserV2Service(httpClient core.YunxinApiHttpClient) UserV2Service {
	return &UserV2ServiceImpl{
		httpClient: httpClient,
	}
}
