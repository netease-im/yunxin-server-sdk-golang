package friend

import (
	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
)

// FriendV1Service 好友服务接口
type FriendV1Service interface {
	// Add 添加好友
	Add(req *AddFriendRequestV1) (*core.Result[*AddFriendResponseV1], error)

	// Update 更新好友相关信息
	Update(req *UpdateFriendRequestV1) (*core.Result[*UpdateFriendResponseV1], error)

	// Delete 删除好友关系
	Delete(req *DeleteFriendRequestV1) (*core.Result[*DeleteFriendResponseV1], error)

	// Get 获取好友列表
	Get(req *GetFriendListRequestV1) (*core.Result[*GetFriendListResponseV1], error)

	// GetByAccid 获取好友关系
	GetByAccid(req *GetFriendRelationshipRequestV1) (*core.Result[*GetFriendRelationshipResponseV1], error)
}

// FriendV1ServiceImpl 好友服务实现
type FriendV1ServiceImpl struct {
	httpClient core.YunxinApiHttpClient
}

// NewFriendV1Service 创建好友服务实例
func NewFriendV1Service(httpClient core.YunxinApiHttpClient) FriendV1Service {
	return &FriendV1ServiceImpl{
		httpClient: httpClient,
	}
}
