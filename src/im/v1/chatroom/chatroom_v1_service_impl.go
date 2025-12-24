package chatroom

import (
	"encoding/json"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/utils"
)

// CreateChatroom CreateChatroom
func (c *ChatroomV1ServiceImpl) CreateChatroom(req *CreateChatroomRequestV1) (*core.Result[*CreateChatroomResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := c.httpClient.ExecuteV1Api(Create, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*CreateChatroomResponseV1)(nil)), nil
	}

	resp := &CreateChatroomResponseV1{}
	if dataVal, ok := jsonObj["chatroom"]; ok {
		dataJson, _ := json.Marshal(dataVal)
		json.Unmarshal(dataJson, resp)
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// UpdateChatroom UpdateChatroom
func (c *ChatroomV1ServiceImpl) UpdateChatroom(req *UpdateChatroomRequestV1) (*core.Result[*UpdateChatroomResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := c.httpClient.ExecuteV1Api(Update, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*UpdateChatroomResponseV1)(nil)), nil
	}

	resp := &UpdateChatroomResponseV1{}
	if dataVal, ok := jsonObj["chatroom"]; ok {
		dataJson, _ := json.Marshal(dataVal)
		json.Unmarshal(dataJson, resp)
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// QueryChatroomAddress QueryChatroomAddress
func (c *ChatroomV1ServiceImpl) QueryChatroomAddress(req *QueryChatroomAddressRequestV1) (*core.Result[*QueryChatroomAddressResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := c.httpClient.ExecuteV1Api(RequestAddr, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*QueryChatroomAddressResponseV1)(nil)), nil
	}

	resp := &QueryChatroomAddressResponseV1{}
	if addr, ok := jsonObj["addr"].(string); ok && addr != "" {
		json.Unmarshal([]byte(addr), &resp.ChatroomAddress)
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// QueryChatroomInfo QueryChatroomInfo
func (c *ChatroomV1ServiceImpl) QueryChatroomInfo(req *QueryChatroomInfoRequestV1) (*core.Result[*QueryChatroomInfoResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := c.httpClient.ExecuteV1Api(Get, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*QueryChatroomInfoResponseV1)(nil)), nil
	}

	resp := &QueryChatroomInfoResponseV1{}
	if dataVal, ok := jsonObj["chatroom"]; ok {
		dataJson, _ := json.Marshal(dataVal)
		json.Unmarshal(dataJson, resp)
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// QueryChatroomInfos QueryChatroomInfos
func (c *ChatroomV1ServiceImpl) QueryChatroomInfos(req *QueryChatroomInfosRequestV1) (*core.Result[*QueryChatroomInfosResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := c.httpClient.ExecuteV1Api(GetBatch, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*QueryChatroomInfosResponseV1)(nil)), nil
	}

	resp := &QueryChatroomInfosResponseV1{}
	if succRooms, ok := jsonObj["succRooms"].(string); ok && succRooms != "" {
		json.Unmarshal([]byte(succRooms), &resp.RoomInfos)
	}
	if noExistRooms, ok := jsonObj["noExistRooms"].(string); ok && noExistRooms != "" {
		json.Unmarshal([]byte(noExistRooms), &resp.NoExistRooms)
	}
	if failRooms, ok := jsonObj["failRooms"].(string); ok && failRooms != "" {
		json.Unmarshal([]byte(failRooms), &resp.FailRooms)
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// ToggleCloseChatroomStat ToggleCloseChatroomStat
func (c *ChatroomV1ServiceImpl) ToggleCloseChatroomStat(req *ToggleCloseChatroomStatRequestV1) (*core.Result[*ToggleCloseChatroomStatResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := c.httpClient.ExecuteV1Api(ToggleCloseStat, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*ToggleCloseChatroomStatResponseV1)(nil)), nil
	}

	resp := &ToggleCloseChatroomStatResponseV1{}
	if descVal, ok := jsonObj["desc"]; ok {
		descJson, _ := json.Marshal(descVal)
		json.Unmarshal(descJson, resp)
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// UpdateDelayClosePolicy UpdateDelayClosePolicy
func (c *ChatroomV1ServiceImpl) UpdateDelayClosePolicy(req *UpdateChatroomDelayClosePolicyRequestV1) (*core.Result[*UpdateChatroomDelayClosePolicyResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := c.httpClient.ExecuteV1Api(UpdateDelayClosePolicy, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*UpdateChatroomDelayClosePolicyResponseV1)(nil)), nil
	}

	resp := &UpdateChatroomDelayClosePolicyResponseV1{}
	if dataVal, ok := jsonObj["chatroom"]; ok {
		dataJson, _ := json.Marshal(dataVal)
		json.Unmarshal(dataJson, resp)
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// UpdateChatroomInOutNotification UpdateChatroomInOutNotification
func (c *ChatroomV1ServiceImpl) UpdateChatroomInOutNotification(req *UpdateChatroomInOutNotificationRequestV1) (*core.Result[*UpdateChatroomInOutNotificationResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := c.httpClient.ExecuteV1Api(UpdateInOutNotification, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*UpdateChatroomInOutNotificationResponseV1)(nil)), nil
	}

	resp := &UpdateChatroomInOutNotificationResponseV1{}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// KickMember KickMember
func (c *ChatroomV1ServiceImpl) KickMember(req *KickMemberRequestV1) (*core.Result[*KickMemberResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := c.httpClient.ExecuteV1Api(KickMember, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*KickMemberResponseV1)(nil)), nil
	}

	resp := &KickMemberResponseV1{}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// SetMemberRole SetMemberRole
func (c *ChatroomV1ServiceImpl) SetMemberRole(req *SetMemberRoleRequestV1) (*core.Result[*SetMemberRoleResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := c.httpClient.ExecuteV1Api(SetMemberRole, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*SetMemberRoleResponseV1)(nil)), nil
	}

	resp := &SetMemberRoleResponseV1{}
	if descVal, ok := jsonObj["desc"]; ok {
		descJson, _ := json.Marshal(descVal)
		json.Unmarshal(descJson, resp)
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// UpdateMyRoomRole UpdateMyRoomRole
func (c *ChatroomV1ServiceImpl) UpdateMyRoomRole(req *UpdateMyRoomRoleRequestV1) (*core.Result[*UpdateMyRoomRoleResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := c.httpClient.ExecuteV1Api(UpdateMyRoomRole, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*UpdateMyRoomRoleResponseV1)(nil)), nil
	}

	resp := &UpdateMyRoomRoleResponseV1{}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// MembersByPage MembersByPage
func (c *ChatroomV1ServiceImpl) MembersByPage(req *QueryMembersByPageRequestV1) (*core.Result[*QueryMembersByPageResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := c.httpClient.ExecuteV1Api(MembersByPage, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*QueryMembersByPageResponseV1)(nil)), nil
	}

	resp := &QueryMembersByPageResponseV1{}
	if descStr, ok := jsonObj["desc"].(string); ok && descStr != "" {
		var descObj map[string]interface{}
		if err := json.Unmarshal([]byte(descStr), &descObj); err == nil {
			if dataVal, ok := descObj["data"]; ok {
				dataJson, _ := json.Marshal(dataVal)
				json.Unmarshal(dataJson, &resp.Data)
			}
		}
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// MembersByRoles MembersByRoles
func (c *ChatroomV1ServiceImpl) MembersByRoles(req *QueryMembersByRolesRequestV1) (*core.Result[*QueryMembersByRolesResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := c.httpClient.ExecuteV1Api(QueryMembersByRole, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*QueryMembersByRolesResponseV1)(nil)), nil
	}

	resp := &QueryMembersByRolesResponseV1{}
	if descStr, ok := jsonObj["desc"].(string); ok && descStr != "" {
		var descObj map[string]interface{}
		if err := json.Unmarshal([]byte(descStr), &descObj); err == nil {
			if dataVal, ok := descObj["data"]; ok {
				dataJson, _ := json.Marshal(dataVal)
				json.Unmarshal(dataJson, &resp.Data)
			}
		}
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// QueryMembers QueryMembers
func (c *ChatroomV1ServiceImpl) QueryMembers(req *QueryMembersRequestV1) (*core.Result[*QueryMembersResponseV1], error) {
	paramMap := utils.Convert(req)

	if len(req.Accids) > 0 {
		accidsJson, _ := json.Marshal(req.Accids)
		paramMap["accids"] = string(accidsJson)
	}

	apiResponse, err := c.httpClient.ExecuteV1Api(QueryMembers, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*QueryMembersResponseV1)(nil)), nil
	}

	resp := &QueryMembersResponseV1{}
	if descStr, ok := jsonObj["desc"].(string); ok && descStr != "" {
		var descObj map[string]interface{}
		if err := json.Unmarshal([]byte(descStr), &descObj); err == nil {
			if dataVal, ok := descObj["data"]; ok {
				dataJson, _ := json.Marshal(dataVal)
				json.Unmarshal(dataJson, &resp.Data)
			}
		}
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// AddRobot AddRobot
func (c *ChatroomV1ServiceImpl) AddRobot(req *AddRobotRequestV1) (*core.Result[*AddRobotResponseV1], error) {
	paramMap := utils.Convert(req)

	if len(req.Accids) > 0 {
		accidsJson, _ := json.Marshal(req.Accids)
		paramMap["accids"] = string(accidsJson)
	}

	apiResponse, err := c.httpClient.ExecuteV1Api(AddRobot, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*AddRobotResponseV1)(nil)), nil
	}

	resp := &AddRobotResponseV1{}
	if descStr, ok := jsonObj["desc"].(string); ok && descStr != "" {
		var descObj map[string]interface{}
		if err := json.Unmarshal([]byte(descStr), &descObj); err == nil {
			if failAccids, ok := descObj["failAccids"]; ok {
				failJson, _ := json.Marshal(failAccids)
				json.Unmarshal(failJson, &resp.FailAccids)
			}
			if successAccids, ok := descObj["successAccids"]; ok {
				successJson, _ := json.Marshal(successAccids)
				json.Unmarshal(successJson, &resp.SuccessAccids)
			}
			if oldAccids, ok := descObj["oldAccids"]; ok {
				oldJson, _ := json.Marshal(oldAccids)
				json.Unmarshal(oldJson, &resp.OldAccids)
			}
		}
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// RemoveRobot RemoveRobot
func (c *ChatroomV1ServiceImpl) RemoveRobot(req *RemoveRobotRequestV1) (*core.Result[*RemoveRobotResponseV1], error) {
	paramMap := utils.Convert(req)

	if len(req.Accids) > 0 {
		accidsJson, _ := json.Marshal(req.Accids)
		paramMap["accids"] = string(accidsJson)
	}

	apiResponse, err := c.httpClient.ExecuteV1Api(RemoveRobot, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*RemoveRobotResponseV1)(nil)), nil
	}

	resp := &RemoveRobotResponseV1{}
	if descStr, ok := jsonObj["desc"].(string); ok && descStr != "" {
		var descObj map[string]interface{}
		if err := json.Unmarshal([]byte(descStr), &descObj); err == nil {
			if failAccids, ok := descObj["failAccids"]; ok {
				failJson, _ := json.Marshal(failAccids)
				json.Unmarshal(failJson, &resp.FailAccids)
			}
			if successAccids, ok := descObj["successAccids"]; ok {
				successJson, _ := json.Marshal(successAccids)
				json.Unmarshal(successJson, &resp.SuccessAccids)
			}
		}
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// CleanRobot CleanRobot
func (c *ChatroomV1ServiceImpl) CleanRobot(req *CleanRobotRequestV1) (*core.Result[*CleanRobotResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := c.httpClient.ExecuteV1Api(CleanRobot, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*CleanRobotResponseV1)(nil)), nil
	}

	resp := &CleanRobotResponseV1{}
	if descVal, ok := jsonObj["desc"]; ok {
		descJson, _ := json.Marshal(descVal)
		json.Unmarshal(descJson, resp)
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// TemporaryMute TemporaryMute
func (c *ChatroomV1ServiceImpl) TemporaryMute(req *TemporaryMuteRequestV1) (*core.Result[*TemporaryMuteResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := c.httpClient.ExecuteV1Api(TemporaryMute, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*TemporaryMuteResponseV1)(nil)), nil
	}

	resp := &TemporaryMuteResponseV1{}
	if descVal, ok := jsonObj["desc"]; ok {
		descJson, _ := json.Marshal(descVal)
		json.Unmarshal(descJson, resp)
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// MuteRoom MuteRoom
func (c *ChatroomV1ServiceImpl) MuteRoom(req *MuteRoomRequestV1) (*core.Result[*MuteRoomResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := c.httpClient.ExecuteV1Api(MuteRoom, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*MuteRoomResponseV1)(nil)), nil
	}

	resp := &MuteRoomResponseV1{}
	if descVal, ok := jsonObj["desc"]; ok {
		descJson, _ := json.Marshal(descVal)
		json.Unmarshal(descJson, resp)
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// TagTemporaryMute TagTemporaryMute
func (c *ChatroomV1ServiceImpl) TagTemporaryMute(req *TagTemporaryMuteRequestV1) (*core.Result[*TagTemporaryMuteResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := c.httpClient.ExecuteV1Api(TagMute, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*TagTemporaryMuteResponseV1)(nil)), nil
	}

	resp := &TagTemporaryMuteResponseV1{}
	if descVal, ok := jsonObj["desc"]; ok {
		descJson, _ := json.Marshal(descVal)
		json.Unmarshal(descJson, resp)
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// TagMembersCount TagMembersCount
func (c *ChatroomV1ServiceImpl) TagMembersCount(req *TagMembersCountRequestV1) (*core.Result[*TagMembersCountResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := c.httpClient.ExecuteV1Api(TagCount, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*TagMembersCountResponseV1)(nil)), nil
	}

	resp := &TagMembersCountResponseV1{}
	if dataVal, ok := jsonObj["data"]; ok {
		dataJson, _ := json.Marshal(dataVal)
		json.Unmarshal(dataJson, resp)
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// TagMembersQuery TagMembersQuery
func (c *ChatroomV1ServiceImpl) TagMembersQuery(req *TagMembersQueryRequestV1) (*core.Result[*TagMembersQueryResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := c.httpClient.ExecuteV1Api(TagQuery, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*TagMembersQueryResponseV1)(nil)), nil
	}

	resp := &TagMembersQueryResponseV1{}
	if descStr, ok := jsonObj["desc"].(string); ok && descStr != "" {
		var descObj map[string]interface{}
		if err := json.Unmarshal([]byte(descStr), &descObj); err == nil {
			if dataVal, ok := descObj["data"]; ok {
				dataJson, _ := json.Marshal(dataVal)
				json.Unmarshal(dataJson, &resp.Data)
			}
		}
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// QueryTagHistoryMsg QueryTagHistoryMsg
func (c *ChatroomV1ServiceImpl) QueryTagHistoryMsg(req *QueryTagHistoryMsgRequestV1) (*core.Result[*QueryTagHistoryMsgResponseV1], error) {
	paramMap := utils.Convert(req)

	if len(req.Tags) > 0 {
		tagsJson, _ := json.Marshal(req.Tags)
		paramMap["tags"] = string(tagsJson)
	}

	apiResponse, err := c.httpClient.ExecuteV1Api(QueryTagMsg, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*QueryTagHistoryMsgResponseV1)(nil)), nil
	}

	resp := &QueryTagHistoryMsgResponseV1{}
	if msgs, ok := jsonObj["msgs"].(string); ok && msgs != "" {
		json.Unmarshal([]byte(msgs), &resp.Msgs)
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// UpdateChatRoomRoleTag UpdateChatRoomRoleTag
func (c *ChatroomV1ServiceImpl) UpdateChatRoomRoleTag(req *UpdateChatRoomRoleTagRequestV1) (*core.Result[*UpdateChatRoomRoleTagResponseV1], error) {
	paramMap := utils.Convert(req)

	if len(req.Tags) > 0 {
		tagsJson, _ := json.Marshal(req.Tags)
		paramMap["tags"] = string(tagsJson)
	}

	apiResponse, err := c.httpClient.ExecuteV1Api(UpdateChatRoomRoleTag, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*UpdateChatRoomRoleTagResponseV1)(nil)), nil
	}

	resp := &UpdateChatRoomRoleTagResponseV1{}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// QueryUserRoomIds QueryUserRoomIds
func (c *ChatroomV1ServiceImpl) QueryUserRoomIds(req *QueryUserRoomIdsRequestV1) (*core.Result[*QueryUserRoomIdsResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := c.httpClient.ExecuteV1Api(QueryUserRoomIds, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*QueryUserRoomIdsResponseV1)(nil)), nil
	}

	resp := &QueryUserRoomIdsResponseV1{}
	if descObj, ok := jsonObj["desc"].(map[string]interface{}); ok {
		if roomidsStr, ok := descObj["roomids"].(string); ok && roomidsStr != "" {
			json.Unmarshal([]byte(roomidsStr), &resp.RoomIds)
		}
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// QueueInit QueueInit
func (c *ChatroomV1ServiceImpl) QueueInit(req *QueueInitRequestV1) (*core.Result[*QueueInitResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := c.httpClient.ExecuteV1Api(QueueInit, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*QueueInitResponseV1)(nil)), nil
	}

	resp := &QueueInitResponseV1{}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// QueueDrop QueueDrop
func (c *ChatroomV1ServiceImpl) QueueDrop(req *QueueDropRequestV1) (*core.Result[*QueueDropResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := c.httpClient.ExecuteV1Api(QueueDrop, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*QueueDropResponseV1)(nil)), nil
	}

	resp := &QueueDropResponseV1{}
	if descVal, ok := jsonObj["desc"]; ok {
		descJson, _ := json.Marshal(descVal)
		json.Unmarshal(descJson, resp)
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// QueueOffer QueueOffer
func (c *ChatroomV1ServiceImpl) QueueOffer(req *QueueOfferRequestV1) (*core.Result[*QueueOfferResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := c.httpClient.ExecuteV1Api(QueueOffer, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*QueueOfferResponseV1)(nil)), nil
	}

	resp := &QueueOfferResponseV1{}
	if descVal, ok := jsonObj["desc"]; ok {
		descJson, _ := json.Marshal(descVal)
		json.Unmarshal(descJson, resp)
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// QueueBatchOffer QueueBatchOffer
func (c *ChatroomV1ServiceImpl) QueueBatchOffer(req *QueueBatchOfferRequestV1) (*core.Result[*QueueBatchOfferResponseV1], error) {
	paramMap := utils.Convert(req)

	if len(req.Elements) > 0 {
		elementsJson, _ := json.Marshal(req.Elements)
		paramMap["elements"] = string(elementsJson)
	}

	apiResponse, err := c.httpClient.ExecuteV1Api(QueueBatchOffer, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*QueueBatchOfferResponseV1)(nil)), nil
	}

	resp := &QueueBatchOfferResponseV1{}
	if descStr, ok := jsonObj["desc"].(string); ok && descStr != "" {
		var descObj map[string]interface{}
		if err := json.Unmarshal([]byte(descStr), &descObj); err == nil {
			if failedKeys, ok := descObj["failedKeys"].(string); ok && failedKeys != "" {
				json.Unmarshal([]byte(failedKeys), &resp.FailedKeys)
			}
			if highPriority, ok := descObj["highPriority"].(bool); ok {
				resp.HighPriority = highPriority
			}
		}
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// QueueBatchUpdate QueueBatchUpdate
func (c *ChatroomV1ServiceImpl) QueueBatchUpdate(req *QueueBatchUpdateRequestV1) (*core.Result[*QueueBatchUpdateResponseV1], error) {
	paramMap := utils.Convert(req)

	if req.Elements != nil && len(req.Elements) > 0 {
		elementsJson, _ := json.Marshal(req.Elements)
		paramMap["elements"] = string(elementsJson)
	}

	apiResponse, err := c.httpClient.ExecuteV1Api(QueueBatchUpdate, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*QueueBatchUpdateResponseV1)(nil)), nil
	}

	resp := &QueueBatchUpdateResponseV1{}
	if descStr, ok := jsonObj["desc"].(string); ok && descStr != "" {
		var descObj map[string]interface{}
		if err := json.Unmarshal([]byte(descStr), &descObj); err == nil {
			if noExistKeys, ok := descObj["noExistElementKey"]; ok {
				noExistJson, _ := json.Marshal(noExistKeys)
				json.Unmarshal(noExistJson, &resp.NoExistElementKey)
			}
			if highPriority, ok := descObj["highPriority"].(bool); ok {
				resp.HighPriority = highPriority
			}
		}
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// QueueList QueueList
func (c *ChatroomV1ServiceImpl) QueueList(req *QueueListRequestV1) (*core.Result[*QueueListResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := c.httpClient.ExecuteV1Api(QueueList, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*QueueListResponseV1)(nil)), nil
	}

	resp := &QueueListResponseV1{}
	if descStr, ok := jsonObj["desc"].(string); ok && descStr != "" {
		var descObj map[string]interface{}
		if err := json.Unmarshal([]byte(descStr), &descObj); err == nil {
			if listVal, ok := descObj["list"].([]interface{}); ok {
				for _, item := range listVal {
					if itemMap, ok := item.(map[string]interface{}); ok {
						for k, v := range itemMap {
							resp.List = append(resp.List, QueueElement{
								Key:   k,
								Value: v.(string),
							})
							break
						}
					}
				}
			}
		}
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// QueuePoll QueuePoll
func (c *ChatroomV1ServiceImpl) QueuePoll(req *QueuePollRequestV1) (*core.Result[*QueuePollResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := c.httpClient.ExecuteV1Api(QueuePoll, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*QueuePollResponseV1)(nil)), nil
	}

	resp := &QueuePollResponseV1{}
	if descStr, ok := jsonObj["desc"].(string); ok && descStr != "" {
		var descObj map[string]interface{}
		if err := json.Unmarshal([]byte(descStr), &descObj); err == nil {
			if key, ok := descObj["key"].(string); ok {
				resp.Key = key
			}
			if value, ok := descObj["value"].(string); ok {
				resp.Value = value
			}
			if highPriority, ok := descObj["highPriority"].(bool); ok {
				resp.HighPriority = highPriority
			}
		}
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// QueueGet QueueGet
func (c *ChatroomV1ServiceImpl) QueueGet(req *QueueGetRequestV1) (*core.Result[*QueueGetResponseV1], error) {
	paramMap := utils.Convert(req)

	if len(req.Keys) > 0 {
		keysJson, _ := json.Marshal(req.Keys)
		paramMap["keys"] = string(keysJson)
	}

	apiResponse, err := c.httpClient.ExecuteV1Api(QueueMultiGet, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*QueueGetResponseV1)(nil)), nil
	}

	resp := &QueueGetResponseV1{}
	if descStr, ok := jsonObj["desc"].(string); ok && descStr != "" {
		var descObj map[string]interface{}
		if err := json.Unmarshal([]byte(descStr), &descObj); err == nil {
			if listVal, ok := descObj["list"].([]interface{}); ok {
				for _, item := range listVal {
					if itemMap, ok := item.(map[string]interface{}); ok {
						for k, v := range itemMap {
							resp.List = append(resp.List, QueueElement{
								Key:   k,
								Value: v.(string),
							})
							break
						}
					}
				}
			}
		}
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}
