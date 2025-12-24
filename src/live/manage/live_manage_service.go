package manage

import (
	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
)

// LiveManageService 直播管理服务接口
type LiveManageService interface {
	// CreateChannel 创建直播频道
	CreateChannel(req *LiveCreateChannelRequest) (*core.Result[*LiveCreateChannelResponse], error)
}
