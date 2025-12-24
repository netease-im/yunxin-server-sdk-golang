package neroom

import (
	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/neroom/user"
)

// YunxinNeroomApiServices 云信Neroom API服务聚合器
type YunxinNeroomApiServices struct {
	neroomUserService user.NeroomUserService
}

// NewYunxinNeroomApiServices 创建云信Neroom API服务聚合器
func NewYunxinNeroomApiServices(yunxinApiHttpClient core.YunxinApiHttpClient) *YunxinNeroomApiServices {
	return &YunxinNeroomApiServices{
		neroomUserService: user.NewNeroomUserService(yunxinApiHttpClient),
	}
}

// GetNeroomUserService 获取Neroom用户服务
func (s *YunxinNeroomApiServices) GetNeroomUserService() user.NeroomUserService {
	return s.neroomUserService
}
