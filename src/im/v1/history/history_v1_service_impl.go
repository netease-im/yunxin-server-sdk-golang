package history

import (
	"encoding/json"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/utils"
)

// QuerySessionHistoryMessage QuerySessionHistoryMessage
func (s *HistoryV1ServiceImpl) QuerySessionHistoryMessage(req *QuerySessionHistoryMessageRequestV1) (*core.Result[*QuerySessionHistoryMessageResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := s.httpClient.ExecuteV1Api(QuerySessionMsg, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*QuerySessionHistoryMessageResponseV1)(nil)), nil
	}

	resp := &QuerySessionHistoryMessageResponseV1{}
	if sizeVal, ok := jsonObj["size"]; ok {
		resp.Size = int(sizeVal.(float64))
	}
	if msgsVal, ok := jsonObj["msgs"]; ok && msgsVal != nil {
		msgsJson, _ := json.Marshal(msgsVal)
		json.Unmarshal(msgsJson, &resp.Msgs)
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// QueryTeamHistoryMessage QueryTeamHistoryMessage
func (s *HistoryV1ServiceImpl) QueryTeamHistoryMessage(req *QueryTeamHistoryMessageRequestV1) (*core.Result[*QueryTeamHistoryMessageResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := s.httpClient.ExecuteV1Api(QueryTeamMsg, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*QueryTeamHistoryMessageResponseV1)(nil)), nil
	}

	resp := &QueryTeamHistoryMessageResponseV1{}
	if sizeVal, ok := jsonObj["size"]; ok {
		resp.Size = int(sizeVal.(float64))
	}
	if msgsVal, ok := jsonObj["msgs"]; ok && msgsVal != nil {
		msgsJson, _ := json.Marshal(msgsVal)
		json.Unmarshal(msgsJson, &resp.Msgs)
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// QueryChatroomHistoryMessage QueryChatroomHistoryMessage
func (s *HistoryV1ServiceImpl) QueryChatroomHistoryMessage(req *QueryChatroomHistoryMessageRequestV1) (*core.Result[*QueryChatroomHistoryMessageResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := s.httpClient.ExecuteV1Api(QueryChatroomMsg, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*QueryChatroomHistoryMessageResponseV1)(nil)), nil
	}

	resp := &QueryChatroomHistoryMessageResponseV1{}
	if sizeVal, ok := jsonObj["size"]; ok {
		resp.Size = int(sizeVal.(float64))
	}
	if msgsVal, ok := jsonObj["msgs"]; ok && msgsVal != nil {
		msgsJson, _ := json.Marshal(msgsVal)
		json.Unmarshal(msgsJson, &resp.Msgs)
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// DeleteChatroomHistoryMessage DeleteChatroomHistoryMessage
func (s *HistoryV1ServiceImpl) DeleteChatroomHistoryMessage(req *DeleteChatroomHistoryMessageRequestV1) (*core.Result[*DeleteChatroomHistoryMessageResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := s.httpClient.ExecuteV1Api(DeleteChatroomHistoryMessage, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*DeleteChatroomHistoryMessageResponseV1)(nil)), nil
	}

	resp := &DeleteChatroomHistoryMessageResponseV1{}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// QuerySessionList QuerySessionList
func (s *HistoryV1ServiceImpl) QuerySessionList(req *QuerySessionListRequestV1) (*core.Result[*QuerySessionListResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := s.httpClient.ExecuteV1Api(QuerySessionList, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*QuerySessionListResponseV1)(nil)), nil
	}

	resp := &QuerySessionListResponseV1{}
	if sessionsVal, ok := jsonObj["sessions"]; ok {
		sessionsJson, _ := json.Marshal(sessionsVal)
		json.Unmarshal(sessionsJson, &resp.Sessions)
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// QueryBroadcastHistoryMessageById QueryBroadcastHistoryMessageById
func (s *HistoryV1ServiceImpl) QueryBroadcastHistoryMessageById(req *QueryBroadcastHistoryMessageByIdRequestV1) (*core.Result[*QueryBroadcastHistoryMessageByIdResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := s.httpClient.ExecuteV1Api(QueryBroadcastMsgById, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*QueryBroadcastHistoryMessageByIdResponseV1)(nil)), nil
	}

	resp := &QueryBroadcastHistoryMessageByIdResponseV1{}
	if dataVal, ok := jsonObj["data"]; ok {
		dataJson, _ := json.Marshal(dataVal)
		json.Unmarshal(dataJson, resp)
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// QueryBroadcastHistoryMessage QueryBroadcastHistoryMessage
func (s *HistoryV1ServiceImpl) QueryBroadcastHistoryMessage(req *QueryBroadcastHistoryMessageRequestV1) (*core.Result[*QueryBroadcastHistoryMessageResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := s.httpClient.ExecuteV1Api(QueryBroadcastMsg, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*QueryBroadcastHistoryMessageResponseV1)(nil)), nil
	}

	resp := &QueryBroadcastHistoryMessageResponseV1{}
	if dataVal, ok := jsonObj["data"]; ok {
		dataJson, _ := json.Marshal(dataVal)
		json.Unmarshal(dataJson, resp)
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// QueryUserEvents QueryUserEvents
func (s *HistoryV1ServiceImpl) QueryUserEvents(req *QueryUserEventsRequestV1) (*core.Result[*QueryUserEventsResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := s.httpClient.ExecuteV1Api(QueryUserEvents, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*QueryUserEventsResponseV1)(nil)), nil
	}

	resp := &QueryUserEventsResponseV1{}
	if eventsVal, ok := jsonObj["events"]; ok {
		eventsJson, _ := json.Marshal(eventsVal)
		json.Unmarshal(eventsJson, &resp.Events)
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}
