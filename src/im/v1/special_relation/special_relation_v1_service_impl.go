package special_relation

import (
	"encoding/json"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/utils"
)

// SetSpecialRelation SetSpecialRelation
func (s *SpecialRelationV1ServiceImpl) SetSpecialRelation(req *SetSpecialRelationRequestV1) (*core.Result[*SetSpecialRelationResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := s.httpClient.ExecuteV1Api(SetSpecialRelation, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*SetSpecialRelationResponseV1)(nil)), nil
	}

	resp := &SetSpecialRelationResponseV1{}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// ListSpecialRelation ListSpecialRelation
func (s *SpecialRelationV1ServiceImpl) ListSpecialRelation(req *ListSpecialRelationRequestV1) (*core.Result[*ListSpecialRelationResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := s.httpClient.ExecuteV1Api(ListSpecialRelation, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*ListSpecialRelationResponseV1)(nil)), nil
	}

	resp := &ListSpecialRelationResponseV1{}
	// 按照 Java SDK 标准：从根对象解析mutelist和blacklist（JSON字符串）
	if mutelistVal, ok := jsonObj["mutelist"]; ok && mutelistVal != nil {
		var mutelistStr string
		if mutelistStrVal, ok := mutelistVal.(string); ok {
			mutelistStr = mutelistStrVal
		} else {
			mutelistJson, _ := json.Marshal(mutelistVal)
			mutelistStr = string(mutelistJson)
		}
		if mutelistStr != "" {
			var mutelist []string
			json.Unmarshal([]byte(mutelistStr), &mutelist)
			resp.Mutelist = mutelist
		}
	}
	if blacklistVal, ok := jsonObj["blacklist"]; ok && blacklistVal != nil {
		var blacklistStr string
		if blacklistStrVal, ok := blacklistVal.(string); ok {
			blacklistStr = blacklistStrVal
		} else {
			blacklistJson, _ := json.Marshal(blacklistVal)
			blacklistStr = string(blacklistJson)
		}
		if blacklistStr != "" {
			var blacklist []string
			json.Unmarshal([]byte(blacklistStr), &blacklist)
			resp.Blacklist = blacklist
		}
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}
