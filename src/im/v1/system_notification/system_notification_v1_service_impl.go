package system_notification

import (
	"encoding/json"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/utils"
)

// SendAttachMsg SendAttachMsg
func (s *SystemNotificationV1ServiceImpl) SendAttachMsg(req *SendAttachMsgRequestV1) (*core.Result[*SendAttachMsgResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := s.httpClient.ExecuteV1Api(SendAttachMsg, paramMap)
	if err != nil {
		return nil, err
	}

	var jsonObj map[string]interface{}
	if err := json.Unmarshal([]byte(apiResponse.GetData()), &jsonObj); err != nil {
		return nil, err
	}

	code := int(jsonObj["code"].(float64))
	if code != 200 {
		desc := ""
		if descVal, ok := jsonObj["desc"]; ok {
			desc = descVal.(string)
		}
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*SendAttachMsgResponseV1)(nil)), nil
	}

	resp := &SendAttachMsgResponseV1{}
	// 按照 Java SDK 标准：从根对象解析desc字段
	if descVal, ok := jsonObj["desc"]; ok && descVal != nil {
		if descStr, ok := descVal.(string); ok {
			resp.Desc = descStr
		}
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// SendBatchAttachMsg SendBatchAttachMsg
func (s *SystemNotificationV1ServiceImpl) SendBatchAttachMsg(req *SendBatchAttachMsgRequestV1) (*core.Result[*SendBatchAttachMsgResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := s.httpClient.ExecuteV1Api(SendBatchAttachMsg, paramMap)
	if err != nil {
		return nil, err
	}

	var jsonObj map[string]interface{}
	if err := json.Unmarshal([]byte(apiResponse.GetData()), &jsonObj); err != nil {
		return nil, err
	}

	code := int(jsonObj["code"].(float64))
	if code != 200 {
		desc := ""
		if descVal, ok := jsonObj["desc"]; ok {
			desc = descVal.(string)
		}
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*SendBatchAttachMsgResponseV1)(nil)), nil
	}

	resp := &SendBatchAttachMsgResponseV1{}
	// 按照 Java SDK 标准：从根对象解析unregister字段（JSON字符串）
	if unregisterVal, ok := jsonObj["unregister"]; ok && unregisterVal != nil {
		var unregisterStr string
		if unregisterStrVal, ok := unregisterVal.(string); ok {
			unregisterStr = unregisterStrVal
		} else {
			unregisterJson, _ := json.Marshal(unregisterVal)
			unregisterStr = string(unregisterJson)
		}
		if unregisterStr != "" {
			var unregisterList []string
			json.Unmarshal([]byte(unregisterStr), &unregisterList)
			resp.Unregister = unregisterList
		}
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}
