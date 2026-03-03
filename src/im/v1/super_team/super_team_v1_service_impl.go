package super_team

import (
	"encoding/json"
	"strconv"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/utils"
)

// SuperTeamCreate SuperTeamCreate
func (s *SuperTeamV1ServiceImpl) SuperTeamCreate(req *SuperTeamCreateRequestV1) (*core.Result[*CreateSuperTeamResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := s.httpClient.ExecuteV1Api(Create, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*CreateSuperTeamResponseV1)(nil)), nil
	}

	resp := &CreateSuperTeamResponseV1{}
	// 按照 Java SDK 标准：从根对象解析tid和faccid
	if tidVal, ok := jsonObj["tid"]; ok && tidVal != nil {
		if tidFloat, ok := tidVal.(float64); ok {
			resp.Tid = int64(tidFloat)
		} else if tidStr, ok := tidVal.(string); ok {
			// tid可能是字符串类型
			if tidInt, err := strconv.ParseInt(tidStr, 10, 64); err == nil {
				resp.Tid = tidInt
			}
		}
	}
	// 按照 Java SDK 标准：从根对象的faccid对象解析
	if faccidVal, ok := jsonObj["faccid"]; ok && faccidVal != nil {
		if faccidMap, ok := faccidVal.(map[string]interface{}); ok {
			resp.Faccid = &SuperTeamFailAccountList{}
			if msgVal, ok := faccidMap["msg"]; ok {
				resp.Faccid.Msg = msgVal.(string)
			}
			if accidVal, ok := faccidMap["accid"]; ok {
				var accidStr string
				if accidStrVal, ok := accidVal.(string); ok {
					accidStr = accidStrVal
				} else {
					accidJson, _ := json.Marshal(accidVal)
					accidStr = string(accidJson)
				}
				if accidStr != "" {
					var accidList []string
					json.Unmarshal([]byte(accidStr), &accidList)
					resp.Faccid.AccidList = accidList
				}
			}
		}
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// SuperTeamUpdate SuperTeamUpdate
func (s *SuperTeamV1ServiceImpl) SuperTeamUpdate(req *SuperTeamUpdateRequestV1) (*core.Result[*SuperTeamUpdateResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := s.httpClient.ExecuteV1Api(Update, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*SuperTeamUpdateResponseV1)(nil)), nil
	}

	// 按照 Java SDK 标准：返回空响应
	resp := &SuperTeamUpdateResponseV1{}
	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// SuperTeamDismiss SuperTeamDismiss
func (s *SuperTeamV1ServiceImpl) SuperTeamDismiss(req *SuperTeamDismissRequestV1) (*core.Result[*SuperTeamDismissResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := s.httpClient.ExecuteV1Api(Dismiss, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*SuperTeamDismissResponseV1)(nil)), nil
	}

	// 按照 Java SDK 标准：返回空响应
	resp := &SuperTeamDismissResponseV1{}
	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// SuperTeamInvite SuperTeamInvite
func (s *SuperTeamV1ServiceImpl) SuperTeamInvite(req *SuperTeamInviteRequestV1) (*core.Result[*SuperTeamInviteResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := s.httpClient.ExecuteV1Api(Invite, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*SuperTeamInviteResponseV1)(nil)), nil
	}

	resp := &SuperTeamInviteResponseV1{}
	// 按照 Java SDK 标准：从根对象的faccid对象解析
	if faccidVal, ok := jsonObj["faccid"]; ok && faccidVal != nil {
		if faccidMap, ok := faccidVal.(map[string]interface{}); ok {
			resp.Faccid = &SuperTeamFailAccountList{}
			if msgVal, ok := faccidMap["msg"]; ok {
				resp.Faccid.Msg = msgVal.(string)
			}
			if accidVal, ok := faccidMap["accid"]; ok {
				var accidStr string
				if accidStrVal, ok := accidVal.(string); ok {
					accidStr = accidStrVal
				} else {
					accidJson, _ := json.Marshal(accidVal)
					accidStr = string(accidJson)
				}
				if accidStr != "" {
					var accidList []string
					json.Unmarshal([]byte(accidStr), &accidList)
					resp.Faccid.AccidList = accidList
				}
			}
		}
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// SuperTeamKickMember SuperTeamKickMember
func (s *SuperTeamV1ServiceImpl) SuperTeamKickMember(req *SuperTeamKickMemberRequestV1) (*core.Result[*SuperTeamKickMemberResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := s.httpClient.ExecuteV1Api(KickMember, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*SuperTeamKickMemberResponseV1)(nil)), nil
	}

	// 按照 Java SDK 标准：返回空响应
	resp := &SuperTeamKickMemberResponseV1{}
	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// SuperTeamMemberLeave SuperTeamMemberLeave
func (s *SuperTeamV1ServiceImpl) SuperTeamMemberLeave(req *SuperTeamMemberLeaveRequestV1) (*core.Result[*SuperTeamMemberLeaveResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := s.httpClient.ExecuteV1Api(MemberLeave, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*SuperTeamMemberLeaveResponseV1)(nil)), nil
	}

	// 按照 Java SDK 标准：返回空响应
	resp := &SuperTeamMemberLeaveResponseV1{}
	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// SuperTeamAddManager SuperTeamAddManager
func (s *SuperTeamV1ServiceImpl) SuperTeamAddManager(req *SuperTeamAddManagerRequestV1) (*core.Result[*SuperTeamAddManagerResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := s.httpClient.ExecuteV1Api(AddManager, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*SuperTeamAddManagerResponseV1)(nil)), nil
	}

	// 按照 Java SDK 标准：返回空响应
	resp := &SuperTeamAddManagerResponseV1{}
	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// SuperTeamRemoveManager SuperTeamRemoveManager
func (s *SuperTeamV1ServiceImpl) SuperTeamRemoveManager(req *SuperTeamRemoveManagerRequestV1) (*core.Result[*SuperTeamRemoveManagerResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := s.httpClient.ExecuteV1Api(RemoveManager, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*SuperTeamRemoveManagerResponseV1)(nil)), nil
	}

	// 按照 Java SDK 标准：返回空响应
	resp := &SuperTeamRemoveManagerResponseV1{}
	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// SuperTeamChangeOwner SuperTeamChangeOwner
func (s *SuperTeamV1ServiceImpl) SuperTeamChangeOwner(req *SuperTeamChangeOwnerRequestV1) (*core.Result[*SuperTeamChangeOwnerResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := s.httpClient.ExecuteV1Api(ChangeOwner, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*SuperTeamChangeOwnerResponseV1)(nil)), nil
	}

	// 按照 Java SDK 标准：返回空响应
	resp := &SuperTeamChangeOwnerResponseV1{}
	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// GetSuperTeam GetSuperTeam
func (s *SuperTeamV1ServiceImpl) GetSuperTeam(req *GetSuperTeamRequestV1) (*core.Result[*GetSuperTeamResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := s.httpClient.ExecuteV1Api(GetSuperTeam, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*GetSuperTeamResponseV1)(nil)), nil
	}

	resp := &GetSuperTeamResponseV1{}
	// 按照 Java SDK 标准：从根对象解析tinfos（JSON字符串）
	if tinfosVal, ok := jsonObj["tinfos"]; ok && tinfosVal != nil {
		var tinfosStr string
		if tinfosStrVal, ok := tinfosVal.(string); ok {
			tinfosStr = tinfosStrVal
		} else {
			tinfosJson, _ := json.Marshal(tinfosVal)
			tinfosStr = string(tinfosJson)
		}
		if tinfosStr != "" {
			var tinfos []SuperTeamInfo
			json.Unmarshal([]byte(tinfosStr), &tinfos)
			resp.Tinfos = tinfos
		}
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// GetSuperTeamMember GetSuperTeamMember
func (s *SuperTeamV1ServiceImpl) GetSuperTeamMember(req *GetSuperTeamMemberRequestV1) (*core.Result[*GetSuperTeamMemberResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := s.httpClient.ExecuteV1Api(GetMember, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*GetSuperTeamMemberResponseV1)(nil)), nil
	}

	resp := &GetSuperTeamMemberResponseV1{}
	// 按照 Java SDK 标准：从根对象解析tlists（JSON字符串）
	if tlistsVal, ok := jsonObj["tlists"]; ok && tlistsVal != nil {
		var tlistsStr string
		if tlistsStrVal, ok := tlistsVal.(string); ok {
			tlistsStr = tlistsStrVal
		} else {
			tlistsJson, _ := json.Marshal(tlistsVal)
			tlistsStr = string(tlistsJson)
		}
		if tlistsStr != "" {
			var tlists []SuperTeamMemberInfo
			json.Unmarshal([]byte(tlistsStr), &tlists)
			resp.Tlists = tlists
		}
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// GetJoinSuperTeam GetJoinSuperTeam
func (s *SuperTeamV1ServiceImpl) GetJoinSuperTeam(req *GetJoinSuperTeamRequestV1) (*core.Result[*GetJoinSuperTeamResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := s.httpClient.ExecuteV1Api(GetJoinSuperTeam, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*GetJoinSuperTeamResponseV1)(nil)), nil
	}

	resp := &GetJoinSuperTeamResponseV1{}
	// 按照 Java SDK 标准：从根对象解析tinfos（JSON字符串）
	if tinfosVal, ok := jsonObj["tinfos"]; ok && tinfosVal != nil {
		var tinfosStr string
		if tinfosStrVal, ok := tinfosVal.(string); ok {
			tinfosStr = tinfosStrVal
		} else {
			tinfosJson, _ := json.Marshal(tinfosVal)
			tinfosStr = string(tinfosJson)
		}
		if tinfosStr != "" {
			var tinfos []GetJoinSuperTeamInfo
			json.Unmarshal([]byte(tinfosStr), &tinfos)
			resp.Tinfos = tinfos
		}
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// SuperTeamUpdateMemberInfo SuperTeamUpdateMemberInfo
func (s *SuperTeamV1ServiceImpl) SuperTeamUpdateMemberInfo(req *SuperTeamUpdateMemberInfoRequestV1) (*core.Result[*SuperTeamUpdateMemberInfoResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := s.httpClient.ExecuteV1Api(UpdateMemberInfo, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*SuperTeamUpdateMemberInfoResponseV1)(nil)), nil
	}

	// 按照 Java SDK 标准：返回空响应
	resp := &SuperTeamUpdateMemberInfoResponseV1{}
	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// SuperTeamUpdateNick SuperTeamUpdateNick
func (s *SuperTeamV1ServiceImpl) SuperTeamUpdateNick(req *SuperTeamUpdateNickRequestV1) (*core.Result[*SuperTeamUpdateNickResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := s.httpClient.ExecuteV1Api(UpdateNick, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*SuperTeamUpdateNickResponseV1)(nil)), nil
	}

	// 按照 Java SDK 标准：返回空响应
	resp := &SuperTeamUpdateNickResponseV1{}
	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// SuperTeamMute SuperTeamMute
func (s *SuperTeamV1ServiceImpl) SuperTeamMute(req *SuperTeamMuteRequestV1) (*core.Result[*SuperTeamMuteResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := s.httpClient.ExecuteV1Api(Mute, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*SuperTeamMuteResponseV1)(nil)), nil
	}

	// 按照 Java SDK 标准：返回空响应
	resp := &SuperTeamMuteResponseV1{}
	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// SuperTeamMuteTlist SuperTeamMuteTlist
func (s *SuperTeamV1ServiceImpl) SuperTeamMuteTlist(req *SuperTeamMuteTlistRequestV1) (*core.Result[*SuperTeamMuteTlistResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := s.httpClient.ExecuteV1Api(MuteTlist, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*SuperTeamMuteTlistResponseV1)(nil)), nil
	}

	// 按照 Java SDK 标准：返回空响应
	resp := &SuperTeamMuteTlistResponseV1{}
	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// GetSuperTeamMuteMember GetSuperTeamMuteMember
func (s *SuperTeamV1ServiceImpl) GetSuperTeamMuteMember(req *GetSuperTeamMuteMemberRequestV1) (*core.Result[*GetSuperTeamMuteMemberResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := s.httpClient.ExecuteV1Api(GetMuteMember, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*GetSuperTeamMuteMemberResponseV1)(nil)), nil
	}

	resp := &GetSuperTeamMuteMemberResponseV1{}
	// 按照 Java SDK 标准：从根对象解析muteTlists（JSON字符串）
	if muteTlistsVal, ok := jsonObj["muteTlists"]; ok && muteTlistsVal != nil {
		var muteTlistsStr string
		if muteTlistsStrVal, ok := muteTlistsVal.(string); ok {
			muteTlistsStr = muteTlistsStrVal
		} else {
			muteTlistsJson, _ := json.Marshal(muteTlistsVal)
			muteTlistsStr = string(muteTlistsJson)
		}
		if muteTlistsStr != "" {
			var muteTlists []SuperTeamMuteMemberInfo
			json.Unmarshal([]byte(muteTlistsStr), &muteTlists)
			resp.Tlists = muteTlists
		}
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// SuperTeamChangeLevel SuperTeamChangeLevel
func (s *SuperTeamV1ServiceImpl) SuperTeamChangeLevel(req *SuperTeamChangeLevelRequestV1) (*core.Result[*SuperTeamChangeLevelResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := s.httpClient.ExecuteV1Api(ChangeLevel, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*SuperTeamChangeLevelResponseV1)(nil)), nil
	}

	// 按照 Java SDK 标准：返回空响应
	resp := &SuperTeamChangeLevelResponseV1{}
	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// SendSuperTeamMessage SendSuperTeamMessage
func (s *SuperTeamV1ServiceImpl) SendSuperTeamMessage(req *SendSuperTeamMessageRequestV1) (*core.Result[*SendSuperTeamMessageResponseV1], error) {
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*SendSuperTeamMessageResponseV1)(nil)), nil
	}

	resp := &SendSuperTeamMessageResponseV1{}
	// 按照 Java SDK 标准：从data字段解析（JSON字符串）
	if dataVal, ok := jsonObj["data"]; ok && dataVal != nil {
		var dataStr string
		if dataStrVal, ok := dataVal.(string); ok {
			dataStr = dataStrVal
		} else {
			dataJson, _ := json.Marshal(dataVal)
			dataStr = string(dataJson)
		}
		if dataStr != "" {
			json.Unmarshal([]byte(dataStr), resp)
		}
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// SendAttachSuperTeamMessage SendAttachSuperTeamMessage
func (s *SuperTeamV1ServiceImpl) SendAttachSuperTeamMessage(req *SendAttachSuperTeamMessageRequestV1) (*core.Result[*SendAttachSuperTeamMessageResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := s.httpClient.ExecuteV1Api(SendAttachMessage, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*SendAttachSuperTeamMessageResponseV1)(nil)), nil
	}

	// 按照 Java SDK 标准：返回空响应
	resp := &SendAttachSuperTeamMessageResponseV1{}
	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// RecallSuperTeamMessage RecallSuperTeamMessage
func (s *SuperTeamV1ServiceImpl) RecallSuperTeamMessage(req *RecallSuperTeamMessageRequestV1) (*core.Result[*RecallSuperTeamMessageResponseV1], error) {
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*RecallSuperTeamMessageResponseV1)(nil)), nil
	}

	// 按照 Java SDK 标准：返回空响应
	resp := &RecallSuperTeamMessageResponseV1{}
	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// GetSuperTeamMessage GetSuperTeamMessage
func (s *SuperTeamV1ServiceImpl) GetSuperTeamMessage(req *GetSuperTeamMessageRequestV1) (*core.Result[*GetSuperTeamMessageResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := s.httpClient.ExecuteV1Api(GetMessage, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*GetSuperTeamMessageResponseV1)(nil)), nil
	}

	resp := &GetSuperTeamMessageResponseV1{}
	// 按照 Java SDK 标准：从根对象解析size和msgs
	resp.Tid = req.Tid
	if sizeVal, ok := jsonObj["size"]; ok && sizeVal != nil {
		if sizeFloat, ok := sizeVal.(float64); ok {
			resp.Size = int(sizeFloat)
		}
	}
	if msgsVal, ok := jsonObj["msgs"]; ok && msgsVal != nil {
		var msgsStr string
		if msgsStrVal, ok := msgsVal.(string); ok {
			msgsStr = msgsStrVal
		} else {
			msgsJson, _ := json.Marshal(msgsVal)
			msgsStr = string(msgsJson)
		}
		if msgsStr != "" {
			var msgs []GetSuperTeamMessage
			json.Unmarshal([]byte(msgsStr), &msgs)
			resp.Msgs = msgs
		}
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// GetSuperTeamMessageByIds GetSuperTeamMessageByIds
func (s *SuperTeamV1ServiceImpl) GetSuperTeamMessageByIds(req *GetSuperTeamMessageByIdsRequestV1) (*core.Result[*GetSuperTeamMessageByIdsResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := s.httpClient.ExecuteV1Api(GetMessageByIds, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*GetSuperTeamMessageByIdsResponseV1)(nil)), nil
	}

	resp := &GetSuperTeamMessageByIdsResponseV1{}
	// 按照 Java SDK 标准：从根对象解析msgs（JSON字符串）
	if msgsVal, ok := jsonObj["msgs"]; ok && msgsVal != nil {
		var msgsStr string
		if msgsStrVal, ok := msgsVal.(string); ok {
			msgsStr = msgsStrVal
		} else {
			msgsJson, _ := json.Marshal(msgsVal)
			msgsStr = string(msgsJson)
		}
		if msgsStr != "" {
			var msgs []GetSuperTeamMessage
			json.Unmarshal([]byte(msgsStr), &msgs)
			resp.Msgs = msgs
		}
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}
