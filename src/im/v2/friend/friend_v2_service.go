package friend

import (
	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
)

// FriendV2Service Friend V2服务接口
type FriendV2Service interface {
	// AddFriend 添加好友
	AddFriend(req *AddFriendRequestV2) (*core.Result[*AddFriendResponseV2], error)

	// DeleteFriend 删除好友
	DeleteFriend(req *DeleteFriendRequestV2) (*core.Result[*DeleteFriendResponseV2], error)

	// UpdateFriend 更新好友
	UpdateFriend(req *UpdateFriendRequestV2) (*core.Result[*UpdateFriendResponseV2], error)

	// GetFriend 获取好友
	GetFriend(req *GetFriendRequestV2) (*core.Result[*GetFriendResponseV2], error)

	// ListFriends 列出好友
	ListFriends(req *ListFriendsRequestV2) (*core.Result[*ListFriendsResponseV2], error)

	// HandleFriendAddition 处理好友添加
	HandleFriendAddition(req *HandleFriendAdditionRequestV2) (*core.Result[*HandleFriendAdditionResponseV2], error)
}

// FriendV2ServiceImpl Friend V2服务实现
type FriendV2ServiceImpl struct {
	httpClient core.YunxinApiHttpClient
}

// NewFriendV2Service 创建Friend V2服务实例
func NewFriendV2Service(httpClient core.YunxinApiHttpClient) FriendV2Service {
	return &FriendV2ServiceImpl{
		httpClient: httpClient,
	}
}
