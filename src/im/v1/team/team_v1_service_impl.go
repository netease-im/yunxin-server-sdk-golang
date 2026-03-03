package team

import (
	"encoding/json"
	"strconv"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/utils"
)

// CreateTeam CreateTeam
func (s *TeamV1ServiceImpl) CreateTeam(req *CreateTeamRequestV1) (*core.Result[*CreateTeamResponseV1], error) {
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*CreateTeamResponseV1)(nil)), nil
	}

	resp := &CreateTeamResponseV1{}
	// 按照 Java SDK 标准：直接从根对象解析tid
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
			if msgVal, ok := faccidMap["msg"]; ok {
				resp.Faccid.Msg = msgVal.(string)
			}
			if accidVal, ok := faccidMap["accid"]; ok {
				accidJson, _ := json.Marshal(accidVal)
				var accidList []string
				json.Unmarshal(accidJson, &accidList)
				resp.Faccid.AccidList = accidList
			}
		}
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// AddTeam AddTeam
func (s *TeamV1ServiceImpl) AddTeam(req *AddTeamRequestV1) (*core.Result[*AddTeamResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := s.httpClient.ExecuteV1Api(Add, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*AddTeamResponseV1)(nil)), nil
	}

	resp := &AddTeamResponseV1{}
	// 从根对象提取faccid对象，然后提取其中的msg和accid字段
	if faccidVal, ok := jsonObj["faccid"]; ok {
		faccidMap := faccidVal.(map[string]interface{})
		if msgVal, ok := faccidMap["msg"]; ok {
			resp.Msg = msgVal.(string)
		}
		if accidVal, ok := faccidMap["accid"]; ok {
			accidJson, _ := json.Marshal(accidVal)
			var accids []string
			json.Unmarshal(accidJson, &accids)
			resp.Accids = accids
		}
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// KickTeam KickTeam
func (s *TeamV1ServiceImpl) KickTeam(req *KickTeamRequestV1) (*core.Result[*KickTeamResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := s.httpClient.ExecuteV1Api(Kick, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*KickTeamResponseV1)(nil)), nil
	}

	resp := &KickTeamResponseV1{}
	// 从根对象提取faccid对象，然后提取其中的msg和accid字段
	if faccidVal, ok := jsonObj["faccid"]; ok {
		faccidMap := faccidVal.(map[string]interface{})
		if msgVal, ok := faccidMap["msg"]; ok {
			resp.Msg = msgVal.(string)
		}
		if accidVal, ok := faccidMap["accid"]; ok {
			accidJson, _ := json.Marshal(accidVal)
			var accids []string
			json.Unmarshal(accidJson, &accids)
			resp.Accid = accids
		}
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// LeaveTeam LeaveTeam
func (s *TeamV1ServiceImpl) LeaveTeam(req *LeaveTeamRequestV1) (*core.Result[*LeaveTeamResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := s.httpClient.ExecuteV1Api(Leave, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*LeaveTeamResponseV1)(nil)), nil
	}

	resp := &LeaveTeamResponseV1{}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// UpdateTeam UpdateTeam
func (s *TeamV1ServiceImpl) UpdateTeam(req *UpdateTeamRequestV1) (*core.Result[*UpdateTeamResponseV1], error) {
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*UpdateTeamResponseV1)(nil)), nil
	}

	resp := &UpdateTeamResponseV1{}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// QueryTeam QueryTeam
func (s *TeamV1ServiceImpl) QueryTeam(req *QueryTeamRequestV1) (*core.Result[*QueryTeamResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := s.httpClient.ExecuteV1Api(Query, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*QueryTeamResponseV1)(nil)), nil
	}

	resp := &QueryTeamResponseV1{}
	// 按照 Java SDK 标准：tinfos是JSON字符串，需要先解析成字符串再解析成数组
	if tinfosVal, ok := jsonObj["tinfos"]; ok {
		var tinfosStr string
		if tinfosStrVal, ok := tinfosVal.(string); ok {
			tinfosStr = tinfosStrVal
		} else {
			// 如果不是字符串，尝试marshal成字符串
			tinfosJson, _ := json.Marshal(tinfosVal)
			tinfosStr = string(tinfosJson)
		}
		if tinfosStr != "" {
			var tinfosArray []map[string]interface{}
			json.Unmarshal([]byte(tinfosStr), &tinfosArray)
			var tinfos []TeamInfoV1
			for _, itemMap := range tinfosArray {
				var teamInfo TeamInfoV1
				itemJson, _ := json.Marshal(itemMap)
				json.Unmarshal(itemJson, &teamInfo)
				// 特殊处理isNotifyCloseOnline和isNotifyClosePersistent字段
				if val, ok := itemMap["isNotifyCloseOnline"]; ok {
					if boolVal, ok := val.(bool); ok {
						teamInfo.IsNotifyCloseOnline = boolVal
					}
				}
				if val, ok := itemMap["isNotifyClosePersistent"]; ok {
					if boolVal, ok := val.(bool); ok {
						teamInfo.IsNotifyClosePersistent = boolVal
					}
				}
				tinfos = append(tinfos, teamInfo)
			}
			resp.Tinfos = tinfos
		}
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// QueryTeamInfoDetails QueryTeamInfoDetails
func (s *TeamV1ServiceImpl) QueryTeamInfoDetails(req *QueryTeamInfoDetailsRequestV1) (*core.Result[*QueryTeamInfoDetailsResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := s.httpClient.ExecuteV1Api(QueryDetail, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*QueryTeamInfoDetailsResponseV1)(nil)), nil
	}

	resp := &QueryTeamInfoDetailsResponseV1{}
	// 响应包含tinfo字段，解析到resp.Tinfo
	if tinfoVal, ok := jsonObj["tinfo"]; ok {
		tinfoJson, _ := json.Marshal(tinfoVal)
		var tinfo TeamInfo
		json.Unmarshal(tinfoJson, &tinfo)
		resp.Tinfo = tinfo
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// AddManagerTeam AddManagerTeam
func (s *TeamV1ServiceImpl) AddManagerTeam(req *AddManagerTeamRequestV1) (*core.Result[*AddManagerTeamResponseV1], error) {
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*AddManagerTeamResponseV1)(nil)), nil
	}

	resp := &AddManagerTeamResponseV1{}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// RemoveManagerTeam RemoveManagerTeam
func (s *TeamV1ServiceImpl) RemoveManagerTeam(req *RemoveManagerTeamRequestV1) (*core.Result[*RemoveManagerTeamResponseV1], error) {
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*RemoveManagerTeamResponseV1)(nil)), nil
	}

	resp := &RemoveManagerTeamResponseV1{}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// ChangeOwnerTeam ChangeOwnerTeam
func (s *TeamV1ServiceImpl) ChangeOwnerTeam(req *ChangeOwnerTeamRequestV1) (*core.Result[*ChangeOwnerTeamResponseV1], error) {
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*ChangeOwnerTeamResponseV1)(nil)), nil
	}

	resp := &ChangeOwnerTeamResponseV1{}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// DismissTeam DismissTeam
func (s *TeamV1ServiceImpl) DismissTeam(req *DismissTeamRequestV1) (*core.Result[*DismissTeamResponseV1], error) {
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*DismissTeamResponseV1)(nil)), nil
	}

	resp := &DismissTeamResponseV1{}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// MuteTeam MuteTeam
func (s *TeamV1ServiceImpl) MuteTeam(req *MuteTeamRequestV1) (*core.Result[*MuteTeamResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := s.httpClient.ExecuteV1Api(MuteTeam, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*MuteTeamResponseV1)(nil)), nil
	}

	resp := &MuteTeamResponseV1{}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// MuteTeamAllMember MuteTeamAllMember
func (s *TeamV1ServiceImpl) MuteTeamAllMember(req *MuteTeamAllMemberRequestV1) (*core.Result[*MuteTeamAllMemberResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := s.httpClient.ExecuteV1Api(MuteTeamAllMember, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*MuteTeamAllMemberResponseV1)(nil)), nil
	}

	resp := &MuteTeamAllMemberResponseV1{}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// MuteTeamTargetMember MuteTeamTargetMember
func (s *TeamV1ServiceImpl) MuteTeamTargetMember(req *MuteTeamTargetMemberRequestV1) (*core.Result[*MuteTeamTargetMemberResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := s.httpClient.ExecuteV1Api(MuteTeamTargetMember, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*MuteTeamTargetMemberResponseV1)(nil)), nil
	}

	resp := &MuteTeamTargetMemberResponseV1{}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// QueryMuteTeamMembers QueryMuteTeamMembers
func (s *TeamV1ServiceImpl) QueryMuteTeamMembers(req *QueryMuteTeamMembersRequestV1) (*core.Result[*QueryMuteTeamMembersResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := s.httpClient.ExecuteV1Api(QueryMuteTeamMembers, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*QueryMuteTeamMembersResponseV1)(nil)), nil
	}

	resp := &QueryMuteTeamMembersResponseV1{}
	// 按照 Java SDK 标准：mutes是JSON字符串，需要先解析成字符串再解析成数组
	if mutesVal, ok := jsonObj["mutes"]; ok {
		var mutesStr string
		if mutesStrVal, ok := mutesVal.(string); ok {
			mutesStr = mutesStrVal
		} else {
			mutesJson, _ := json.Marshal(mutesVal)
			mutesStr = string(mutesJson)
		}
		if mutesStr != "" {
			var mutes []MuteInfo
			json.Unmarshal([]byte(mutesStr), &mutes)
			resp.Mutes = mutes
		}
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// UpdateTeamNick UpdateTeamNick
func (s *TeamV1ServiceImpl) UpdateTeamNick(req *UpdateTeamNickRequestV1) (*core.Result[*UpdateTeamNickResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := s.httpClient.ExecuteV1Api(UpdateTeamNick, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*UpdateTeamNickResponseV1)(nil)), nil
	}

	resp := &UpdateTeamNickResponseV1{}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// JoinsTeam JoinsTeam
func (s *TeamV1ServiceImpl) JoinsTeam(req *JoinsTeamRequestV1) (*core.Result[*JoinsTeamResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := s.httpClient.ExecuteV1Api(JoinsTeam, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*JoinsTeamResponseV1)(nil)), nil
	}

	resp := &JoinsTeamResponseV1{}
	// 按照 Java SDK 标准：infos是JSON字符串，需要先解析成字符串再解析成数组
	if infosVal, ok := jsonObj["infos"]; ok {
		var infosStr string
		if infosStrVal, ok := infosVal.(string); ok {
			infosStr = infosStrVal
		} else {
			infosJson, _ := json.Marshal(infosVal)
			infosStr = string(infosJson)
		}
		if infosStr != "" {
			var infos []JoinsTinfo
			json.Unmarshal([]byte(infosStr), &infos)
			resp.Infos = infos
		}
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// QueryOnlineTeamMember QueryOnlineTeamMember
func (s *TeamV1ServiceImpl) QueryOnlineTeamMember(req *QueryOnlineTeamMemberRequestV1) (*core.Result[*QueryOnlineTeamMemberResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := s.httpClient.ExecuteV1Api(QueryOnlineTeamMember, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*QueryOnlineTeamMemberResponseV1)(nil)), nil
	}

	resp := &QueryOnlineTeamMemberResponseV1{}
	// 按照 Java SDK 标准：从data对象解析
	if dataVal, ok := jsonObj["data"]; ok {
		if dataMap, ok := dataVal.(map[string]interface{}); ok {
			if countVal, ok := dataMap["count"]; ok {
				if countFloat, ok := countVal.(float64); ok {
					resp.Count = int(countFloat)
				}
			}
			// 特殊处理status字段：它是一个对象，key是accid，value是数组
			if statusVal, ok := dataMap["status"]; ok {
				if statusMap, ok := statusVal.(map[string]interface{}); ok {
					var onlineStatusList []OnlineStatus
					for accid, statusArrayVal := range statusMap {
						onlineStatus := OnlineStatus{Accid: accid}
						if statusArray, ok := statusArrayVal.([]interface{}); ok {
							var statusList []Status
							for _, statusItem := range statusArray {
								if statusItemMap, ok := statusItem.(map[string]interface{}); ok {
									var status Status
									if clientTypeVal, ok := statusItemMap["clientType"]; ok {
										if clientTypeFloat, ok := clientTypeVal.(float64); ok {
											status.ClientType = int(clientTypeFloat)
										}
									}
									if loginTimeVal, ok := statusItemMap["loginTime"]; ok {
										if loginTimeFloat, ok := loginTimeVal.(float64); ok {
											status.LoginTime = int64(loginTimeFloat)
										}
									}
									statusList = append(statusList, status)
								}
							}
							onlineStatus.StatusList = statusList
						}
						onlineStatusList = append(onlineStatusList, onlineStatus)
					}
					resp.Status = onlineStatusList
				}
			}
		}
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// BatchQueryOnlineTeamMemberCount BatchQueryOnlineTeamMemberCount
func (s *TeamV1ServiceImpl) BatchQueryOnlineTeamMemberCount(req *BatchQueryOnlineTeamMemberCountRequestV1) (*core.Result[*BatchQueryOnlineTeamMemberCountResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := s.httpClient.ExecuteV1Api(BatchQueryOnlineTeamMemberCount, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*BatchQueryOnlineTeamMemberCountResponseV1)(nil)), nil
	}

	resp := &BatchQueryOnlineTeamMemberCountResponseV1{}
	// 按照 Java SDK 标准：data是JSON字符串，需要先解析成字符串再解析成数组
	if dataVal, ok := jsonObj["data"]; ok {
		var dataStr string
		if dataStrVal, ok := dataVal.(string); ok {
			dataStr = dataStrVal
		} else {
			dataJson, _ := json.Marshal(dataVal)
			dataStr = string(dataJson)
		}
		if dataStr != "" {
			var dataList []teamOnlineCount
			json.Unmarshal([]byte(dataStr), &dataList)
			resp.Data = dataList
		}
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// QueryTeamMsgMarkReadInfo QueryTeamMsgMarkReadInfo
func (s *TeamV1ServiceImpl) QueryTeamMsgMarkReadInfo(req *QueryTeamMsgMarkReadInfoRequestV1) (*core.Result[*QueryTeamMsgMarkReadInfoResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := s.httpClient.ExecuteV1Api(QueryTeamMsgMarkReadInfo, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*QueryTeamMsgMarkReadInfoResponseV1)(nil)), nil
	}

	resp := &QueryTeamMsgMarkReadInfoResponseV1{}
	if dataVal, ok := jsonObj["data"]; ok {
		dataJson, _ := json.Marshal(dataVal)
		json.Unmarshal(dataJson, resp)
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// QueryAllJoinedTeamMemberInfoByAccId QueryAllJoinedTeamMemberInfoByAccId
func (s *TeamV1ServiceImpl) QueryAllJoinedTeamMemberInfoByAccId(req *QueryAllJoinedTeamMemberInfoByAccIdRequestV1) (*core.Result[*QueryAllJoinedTeamMemberInfoByAccIdResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := s.httpClient.ExecuteV1Api(QueryAllJoinedTeamMemberInfoByAccId, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*QueryAllJoinedTeamMemberInfoByAccIdResponseV1)(nil)), nil
	}

	resp := &QueryAllJoinedTeamMemberInfoByAccIdResponseV1{}
	// 按照 Java SDK 标准：data是JSON字符串，需要先解析成字符串再解析成数组
	if dataVal, ok := jsonObj["data"]; ok {
		var dataStr string
		if dataStrVal, ok := dataVal.(string); ok {
			dataStr = dataStrVal
		} else {
			dataJson, _ := json.Marshal(dataVal)
			dataStr = string(dataJson)
		}
		if dataStr != "" {
			var dataList []TeamMemberInfo
			json.Unmarshal([]byte(dataStr), &dataList)
			resp.Data = dataList
		}
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}
