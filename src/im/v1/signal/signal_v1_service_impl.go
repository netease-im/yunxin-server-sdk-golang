package signal

import (
	"encoding/json"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/utils"
)

// CreateSignalRoom CreateSignalRoom
func (s *SignalV1ServiceImpl) CreateSignalRoom(req *CreateSignalRoomRequestV1) (*core.Result[*CreateSignalRoomResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := s.httpClient.ExecuteV1Api(CreateSignalRoom, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*CreateSignalRoomResponseV1)(nil)), nil
	}

	resp := &CreateSignalRoomResponseV1{}
	if dataVal, ok := jsonObj["data"]; ok {
		dataJson, _ := json.Marshal(dataVal)
		json.Unmarshal(dataJson, resp)
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// CloseSignalRoom CloseSignalRoom
func (s *SignalV1ServiceImpl) CloseSignalRoom(req *CloseSignalRoomRequestV1) (*core.Result[*CloseSignalRoomResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := s.httpClient.ExecuteV1Api(CloseSignalRoom, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*CloseSignalRoomResponseV1)(nil)), nil
	}

	resp := &CloseSignalRoomResponseV1{}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// InviteSignalRoom InviteSignalRoom
func (s *SignalV1ServiceImpl) InviteSignalRoom(req *InviteSignalRoomRequestV1) (*core.Result[*InviteSignalRoomResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := s.httpClient.ExecuteV1Api(InviteSignalRoom, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*InviteSignalRoomResponseV1)(nil)), nil
	}

	resp := &InviteSignalRoomResponseV1{}
	if dataVal, ok := jsonObj["data"]; ok {
		dataJson, _ := json.Marshal(dataVal)
		json.Unmarshal(dataJson, resp)
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// CancelSignalRoomInvite CancelSignalRoomInvite
func (s *SignalV1ServiceImpl) CancelSignalRoomInvite(req *CancelSignalRoomInviteRequestV1) (*core.Result[*CancelSignalRoomInviteResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := s.httpClient.ExecuteV1Api(CancelSignalRoomInvite, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*CancelSignalRoomInviteResponseV1)(nil)), nil
	}

	resp := &CancelSignalRoomInviteResponseV1{}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// KickSignalRoom KickSignalRoom
func (s *SignalV1ServiceImpl) KickSignalRoom(req *KickSignalRoomRequestV1) (*core.Result[*KickSignalRoomResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := s.httpClient.ExecuteV1Api(KickSignalRoom, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*KickSignalRoomResponseV1)(nil)), nil
	}

	resp := &KickSignalRoomResponseV1{}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// CtrlSignalRoom CtrlSignalRoom
func (s *SignalV1ServiceImpl) CtrlSignalRoom(req *CtrlSignalRoomRequestV1) (*core.Result[*CtrlSignalRoomResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := s.httpClient.ExecuteV1Api(CtrlSignalRoom, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*CtrlSignalRoomResponseV1)(nil)), nil
	}

	resp := &CtrlSignalRoomResponseV1{}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// DelaySignalRoom DelaySignalRoom
func (s *SignalV1ServiceImpl) DelaySignalRoom(req *DelaySignalRoomRequestV1) (*core.Result[*DelaySignalRoomResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := s.httpClient.ExecuteV1Api(DelaySignalRoom, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*DelaySignalRoomResponseV1)(nil)), nil
	}

	resp := &DelaySignalRoomResponseV1{}
	if dataVal, ok := jsonObj["data"]; ok {
		dataJson, _ := json.Marshal(dataVal)
		json.Unmarshal(dataJson, resp)
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// GetSignalRoomInfo GetSignalRoomInfo
func (s *SignalV1ServiceImpl) GetSignalRoomInfo(req *GetSignalRoomInfoRequestV1) (*core.Result[*GetSignalRoomInfoResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := s.httpClient.ExecuteV1Api(GetSignalRoomInfo, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*GetSignalRoomInfoResponseV1)(nil)), nil
	}

	resp := &GetSignalRoomInfoResponseV1{}
	if dataVal, ok := jsonObj["data"]; ok {
		dataJson, _ := json.Marshal(dataVal)
		json.Unmarshal(dataJson, resp)
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}
