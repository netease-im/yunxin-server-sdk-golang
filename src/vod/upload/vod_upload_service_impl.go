package upload

import (
	"encoding/json"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/http"
	"github.com/netease-im/yunxin-server-sdk-golang/src/im/v2/util"
)

// VodUploadServiceImpl VOD上传服务实现
type VodUploadServiceImpl struct {
	httpClient core.YunxinApiHttpClient
}

// NewVodUploadService 创建VOD上传服务
func NewVodUploadService(httpClient core.YunxinApiHttpClient) VodUploadService {
	return &VodUploadServiceImpl{
		httpClient: httpClient,
	}
}

// UploadInit 初始化上传
func (s *VodUploadServiceImpl) UploadInit(req *VodUploadInitRequest) (*core.Result[*VodUploadInitResponse], error) {
	// 序列化请求体
	requestData, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	// 使用HTTP客户端发送POST请求到VOD上传初始化API
	apiResponse, err := s.httpClient.ExecuteJson(http.POST, UploadInit, nil, nil, string(requestData))
	if err != nil {
		return nil, err
	}

	// 转换响应为Result类型
	return util.ConvertWithClass[*VodUploadInitResponse](apiResponse)
}
