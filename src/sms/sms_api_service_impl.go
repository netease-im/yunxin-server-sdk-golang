package sms

import (
	"encoding/json"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/http"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/utils"
)

// SendCode 发送短信验证码
func (s *SmsApiServiceImpl) SendCode(request *SmsSendCodeRequest) (*core.Result[*SmsSendCodeResponse], error) {
	// 将请求对象转换为参数Map
	paramMap := utils.Convert(request)

	// 使用ParamBuilder构建Form数据
	builder := http.NewParamBuilder()
	for key, value := range paramMap {
		builder.AddParam(key, value)
	}
	formData := builder.Build()

	// 执行Form请求
	apiResponse, err := s.httpClient.ExecuteForm(http.POST, SendCode, nil, nil, formData)
	if err != nil {
		return nil, err
	}

	// 解析响应JSON
	var jsonObj map[string]interface{}
	if err := json.Unmarshal([]byte(apiResponse.GetData()), &jsonObj); err != nil {
		return nil, err
	}

	// 获取code
	code := 0
	if codeFloat, ok := jsonObj["code"].(float64); ok {
		code = int(codeFloat)
	}

	// 如果code不是200，返回错误结果
	if code != 200 {
		msg, _ := jsonObj["msg"].(string)
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), msg, (*SmsSendCodeResponse)(nil)), nil
	}

	// 解析sendid (在msg字段中)
	var sendid int64
	if msgFloat, ok := jsonObj["msg"].(float64); ok {
		sendid = int64(msgFloat)
	}

	// 解析authCode (在obj字段中)
	authCode, _ := jsonObj["obj"].(string)

	// 构建响应对象
	response := &SmsSendCodeResponse{
		Sendid:   sendid,
		AuthCode: authCode,
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", response), nil
}
