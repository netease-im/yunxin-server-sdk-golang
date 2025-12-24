package vod

import (
	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/vod/upload"
)

// YunxinVodApiServices 云信VOD API服务
type YunxinVodApiServices struct {
	vodUploadService upload.VodUploadService
}

// NewYunxinVodApiServices 创建云信VOD API服务
func NewYunxinVodApiServices(yunxinApiHttpClient core.YunxinApiHttpClient) *YunxinVodApiServices {
	return &YunxinVodApiServices{
		vodUploadService: upload.NewVodUploadService(yunxinApiHttpClient),
	}
}

// GetVodUploadService 获取VOD上传服务
func (s *YunxinVodApiServices) GetVodUploadService() upload.VodUploadService {
	return s.vodUploadService
}
