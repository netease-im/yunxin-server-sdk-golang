package upload

import (
	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
)

// VodUploadService VOD上传服务接口
type VodUploadService interface {
	// UploadInit 初始化上传
	UploadInit(req *VodUploadInitRequest) (*core.Result[*VodUploadInitResponse], error)
}
