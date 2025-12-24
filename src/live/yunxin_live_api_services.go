package live

import (
	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/live/manage"
)

// YunxinLiveApiServices 云信直播API服务
type YunxinLiveApiServices struct {
	liveManageService manage.LiveManageService
}

// NewYunxinLiveApiServices 创建云信直播API服务
func NewYunxinLiveApiServices(yunxinApiHttpClient core.YunxinApiHttpClient) *YunxinLiveApiServices {
	return &YunxinLiveApiServices{
		liveManageService: manage.NewLiveManageService(yunxinApiHttpClient),
	}
}

// GetLiveManageService 获取直播管理服务
func (s *YunxinLiveApiServices) GetLiveManageService() manage.LiveManageService {
	return s.liveManageService
}
