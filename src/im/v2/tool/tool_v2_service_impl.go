package tool

import (
	"encoding/json"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/http"
	"github.com/netease-im/yunxin-server-sdk-golang/src/im/v2/util"
)

// Asr 语音转文字
func (s *ToolV2ServiceImpl) Asr(req *AsrRequestV2) (*core.Result[*AsrResponseV2], error) {
	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := s.httpClient.ExecuteV2Api(http.POST, AsrUrl, nil, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*AsrResponseV2](apiResponse)
}

// TranslateText 文本翻译
func (s *ToolV2ServiceImpl) TranslateText(req *TranslateTextRequestV2) (*core.Result[*TranslateTextResponseV2], error) {
	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := s.httpClient.ExecuteV2Api(http.POST, TranslateText, nil, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*TranslateTextResponseV2](apiResponse)
}
