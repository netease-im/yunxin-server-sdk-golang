package block

import (
	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
)

// BlockV2Service Block V2服务接口
type BlockV2Service interface {
	// AddBlockContact 添加黑名单联系人
	AddBlockContact(req *AddBlockContactRequestV2) (*core.Result[*AddBlockContactResponseV2], error)

	// RemoveBlockContact 移除黑名单联系人
	RemoveBlockContact(req *RemoveBlockContactRequestV2) (*core.Result[*RemoveBlockContactResponseV2], error)

	// ListBlockContacts 查询黑名单列表
	ListBlockContacts(req *ListBlockContactsRequestV2) (*core.Result[*ListBlockContactsResponseV2], error)
}

// BlockV2ServiceImpl Block V2服务实现
type BlockV2ServiceImpl struct {
	httpClient core.YunxinApiHttpClient
}

// NewBlockV2Service 创建Block V2服务实例
func NewBlockV2Service(httpClient core.YunxinApiHttpClient) BlockV2Service {
	return &BlockV2ServiceImpl{
		httpClient: httpClient,
	}
}
