package message

import (
	"encoding/json"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/utils"
)

// SendMessage SendMessage
func (s *MessageV1ServiceImpl) SendMessage(req *SendMessageRequestV1) (*core.Result[*SendMessageResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := s.httpClient.ExecuteV1Api(SendMessage, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*SendMessageResponseV1)(nil)), nil
	}

	resp := &SendMessageResponseV1{}
	if dataVal, ok := jsonObj["data"]; ok {
		dataJson, _ := json.Marshal(dataVal)
		json.Unmarshal(dataJson, resp)
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// SendBatchMessage SendBatchMessage
func (s *MessageV1ServiceImpl) SendBatchMessage(req *SendBatchMessageRequestV1) (*core.Result[*SendBatchMessageResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := s.httpClient.ExecuteV1Api(SendBatchMessage, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*SendBatchMessageResponseV1)(nil)), nil
	}

	resp := &SendBatchMessageResponseV1{}
	// 按照 Java SDK 标准：直接从根对象解析字段
	if unregisterVal, ok := jsonObj["unregister"]; ok {
		unregisterJson, _ := json.Marshal(unregisterVal)
		var unregisterList []string
		json.Unmarshal(unregisterJson, &unregisterList)
		resp.Unregister = unregisterList
	}
	if timetagVal, ok := jsonObj["timetag"]; ok {
		resp.Timetag = int64(timetagVal.(float64))
	}
	if msgidsVal, ok := jsonObj["msgids"]; ok {
		msgidsJson, _ := json.Marshal(msgidsVal)
		var msgidsMap map[string]int64
		json.Unmarshal(msgidsJson, &msgidsMap)
		resp.Msgids = msgidsMap
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// UploadFile UploadFile
func (s *MessageV1ServiceImpl) UploadFile(req *UploadFileRequestV1) (*core.Result[*UploadFileResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := s.httpClient.ExecuteV1Api(UploadFile, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*UploadFileResponseV1)(nil)), nil
	}

	resp := &UploadFileResponseV1{}
	// 按照 Java SDK 标准：url 字段直接在根对象中
	if urlVal, ok := jsonObj["url"]; ok {
		resp.Url = urlVal.(string)
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// RecallMessage RecallMessage
func (s *MessageV1ServiceImpl) RecallMessage(req *RecallMessageRequestV1) (*core.Result[*RecallMessageResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := s.httpClient.ExecuteV1Api(RecallMessage, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*RecallMessageResponseV1)(nil)), nil
	}

	resp := &RecallMessageResponseV1{}
	if dataVal, ok := jsonObj["data"]; ok {
		dataJson, _ := json.Marshal(dataVal)
		json.Unmarshal(dataJson, resp)
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// BroadcastMessage BroadcastMessage
func (s *MessageV1ServiceImpl) BroadcastMessage(req *BroadcastMessageRequestV1) (*core.Result[*BroadcastMessageResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := s.httpClient.ExecuteV1Api(BroadcastMessage, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*BroadcastMessageResponseV1)(nil)), nil
	}

	resp := &BroadcastMessageResponseV1{}
	// 按照 Java SDK 标准：响应数据在 msg 字段中
	if msgVal, ok := jsonObj["msg"]; ok {
		if msgMap, ok := msgVal.(map[string]interface{}); ok {
			if broadcastIdVal, ok := msgMap["broadcastId"]; ok {
				resp.BroadcastId = int64(broadcastIdVal.(float64))
			}
			if fromVal, ok := msgMap["from"]; ok {
				resp.From = fromVal.(string)
			}
			if bodyVal, ok := msgMap["body"]; ok {
				resp.Body = bodyVal.(string)
			}
			if targetOsVal, ok := msgMap["targetOs"]; ok {
				targetOsJson, _ := json.Marshal(targetOsVal)
				var targetOsList []string
				json.Unmarshal(targetOsJson, &targetOsList)
				resp.TargetOs = targetOsList
			}
			if isOfflineVal, ok := msgMap["isOffline"]; ok {
				resp.IsOffline = isOfflineVal.(bool)
			}
			if createTimeVal, ok := msgMap["createTime"]; ok {
				resp.CreateTime = int64(createTimeVal.(float64))
			}
			if expireTimeVal, ok := msgMap["expireTime"]; ok {
				resp.ExpireTime = int64(expireTimeVal.(float64))
			}
		}
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// DeleteBroadcastMessageById DeleteBroadcastMessageById
func (s *MessageV1ServiceImpl) DeleteBroadcastMessageById(req *DeleteBroadcastMessageByIdRequestV1) (*core.Result[*DeleteBroadcastMessageByIdResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := s.httpClient.ExecuteV1Api(DeleteBroadcastMessageById, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*DeleteBroadcastMessageByIdResponseV1)(nil)), nil
	}

	resp := &DeleteBroadcastMessageByIdResponseV1{}
	if dataVal, ok := jsonObj["data"]; ok {
		dataJson, _ := json.Marshal(dataVal)
		json.Unmarshal(dataJson, resp)
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// DeleteMessage DeleteMessage
func (s *MessageV1ServiceImpl) DeleteMessage(req *DeleteMessageRequestV1) (*core.Result[*DeleteMessageResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := s.httpClient.ExecuteV1Api(DeleteMessage, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*DeleteMessageResponseV1)(nil)), nil
	}

	resp := &DeleteMessageResponseV1{}
	// 按照 Java SDK 标准：从 data 字段解析
	if dataVal, ok := jsonObj["data"]; ok {
		dataJson, _ := json.Marshal(dataVal)
		json.Unmarshal(dataJson, resp)
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// DeleteMessageOneWay DeleteMessageOneWay
func (s *MessageV1ServiceImpl) DeleteMessageOneWay(req *DeleteMessageOneWayRequestV1) (*core.Result[*DeleteMessageOneWayResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := s.httpClient.ExecuteV1Api(DeleteMessageOneWay, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*DeleteMessageOneWayResponseV1)(nil)), nil
	}

	resp := &DeleteMessageOneWayResponseV1{}
	// 按照 Java SDK 标准：从 data 字段解析
	if dataVal, ok := jsonObj["data"]; ok {
		dataJson, _ := json.Marshal(dataVal)
		json.Unmarshal(dataJson, resp)
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// DeleteFile DeleteFile
func (s *MessageV1ServiceImpl) DeleteFile(req *DeleteFileRequestV1) (*core.Result[*DeleteFileResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := s.httpClient.ExecuteV1Api(DeleteFile, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*DeleteFileResponseV1)(nil)), nil
	}

	resp := &DeleteFileResponseV1{}
	if dataVal, ok := jsonObj["data"]; ok {
		dataJson, _ := json.Marshal(dataVal)
		json.Unmarshal(dataJson, resp)
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// DeleteRoamSession DeleteRoamSession
func (s *MessageV1ServiceImpl) DeleteRoamSession(req *DeleteRoamSessionRequestV1) (*core.Result[*DeleteRoamSessionResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := s.httpClient.ExecuteV1Api(DeleteRoamSession, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*DeleteRoamSessionResponseV1)(nil)), nil
	}

	resp := &DeleteRoamSessionResponseV1{}
	// 按照 Java SDK 标准：从 data 字段解析
	if dataVal, ok := jsonObj["data"]; ok {
		dataJson, _ := json.Marshal(dataVal)
		json.Unmarshal(dataJson, resp)
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// MarkReadMessage MarkReadMessage
func (s *MessageV1ServiceImpl) MarkReadMessage(req *MarkReadMessageRequestV1) (*core.Result[*MarkReadMessageResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := s.httpClient.ExecuteV1Api(MarkReadMessage, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*MarkReadMessageResponseV1)(nil)), nil
	}

	resp := &MarkReadMessageResponseV1{}
	if dataVal, ok := jsonObj["data"]; ok {
		dataJson, _ := json.Marshal(dataVal)
		json.Unmarshal(dataJson, resp)
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// MarkReadTeamMessage MarkReadTeamMessage
func (s *MessageV1ServiceImpl) MarkReadTeamMessage(req *MarkReadTeamMessageRequestV1) (*core.Result[*MarkReadTeamMessageResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := s.httpClient.ExecuteV1Api(MarkReadTeamMessage, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*MarkReadTeamMessageResponseV1)(nil)), nil
	}

	resp := &MarkReadTeamMessageResponseV1{}
	// 按照 Java SDK 标准：从 data.errMsgIds 解析
	if dataVal, ok := jsonObj["data"]; ok {
		if dataMap, ok := dataVal.(map[string]interface{}); ok {
			if errMsgIdsVal, ok := dataMap["errMsgIds"]; ok {
				errMsgIdsJson, _ := json.Marshal(errMsgIdsVal)
				var errMsgIdsList []int64
				json.Unmarshal(errMsgIdsJson, &errMsgIdsList)
				resp.ErrMsgIds = errMsgIdsList
			}
		}
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}
