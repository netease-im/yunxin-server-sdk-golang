package mute

import (
	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
)

// MuteV2Service Mute V2服务接口
type MuteV2Service interface {
	// AddMuteContact 添加免打扰联系人
	AddMuteContact(req *AddMuteContactRequestV2) (*core.Result[*AddMuteContactResponseV2], error)

	// RemoveMuteContact 移除免打扰联系人
	RemoveMuteContact(req *RemoveMuteContactRequestV2) (*core.Result[*RemoveMuteContactResponseV2], error)

	// ListMuteContacts 列出免打扰联系人
	ListMuteContacts(req *ListMuteContactsRequestV2) (*core.Result[*ListMuteContactsResponseV2], error)
}

// MuteV2ServiceImpl Mute V2服务实现
type MuteV2ServiceImpl struct {
	httpClient core.YunxinApiHttpClient
}

// NewMuteV2Service 创建Mute V2服务实例
func NewMuteV2Service(httpClient core.YunxinApiHttpClient) MuteV2Service {
	return &MuteV2ServiceImpl{
		httpClient: httpClient,
	}
}
