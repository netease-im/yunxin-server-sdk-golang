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
	// 按照 Java SDK 标准：从根对象解析size和msgs
	if sizeVal, ok := jsonObj["size"]; ok {
		if sizeFloat, ok := sizeVal.(float64); ok {
			resp.Size = int(sizeFloat)
		}
	}
	// 按照 Java SDK 标准：msgs是JSON字符串，需要先解析成字符串再解析成数组
	if msgsVal, ok := jsonObj["msgs"]; ok && msgsVal != nil {
		var msgsStr string
		if msgsStrVal, ok := msgsVal.(string); ok {
			msgsStr = msgsStrVal
		} else {
			msgsJson, _ := json.Marshal(msgsVal)
			msgsStr = string(msgsJson)
		}
		if msgsStr != "" {
			var msgsArray []map[string]interface{}
			json.Unmarshal([]byte(msgsStr), &msgsArray)
			var messages []Message
			for _, msgMap := range msgsArray {
				// 处理body字段：如果是对象，转换为JSON字符串
				if bodyVal, exists := msgMap["body"]; exists && bodyVal != nil {
					if bodyMap, ok := bodyVal.(map[string]interface{}); ok {
						bodyJson, _ := json.Marshal(bodyMap)
						msgMap["body"] = string(bodyJson)
					} else if bodyStr, ok := bodyVal.(string); ok {
						msgMap["body"] = bodyStr
					}
				}
				var msg Message
				msgJson, _ := json.Marshal(msgMap)
				json.Unmarshal(msgJson, &msg)
				messages = append(messages, msg)
			}
			resp.Msgs = messages
		}
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
	// 按照 Java SDK 标准：从根对象解析size和msgs
	if sizeVal, ok := jsonObj["size"]; ok {
		if sizeFloat, ok := sizeVal.(float64); ok {
			resp.Size = int(sizeFloat)
		}
	}
	// 按照 Java SDK 标准：msgs是JSON字符串，需要先解析成字符串再解析成数组
	if msgsVal, ok := jsonObj["msgs"]; ok && msgsVal != nil {
		var msgsStr string
		if msgsStrVal, ok := msgsVal.(string); ok {
			msgsStr = msgsStrVal
		} else {
			msgsJson, _ := json.Marshal(msgsVal)
			msgsStr = string(msgsJson)
		}
		if msgsStr != "" {
			var msgsArray []map[string]interface{}
			json.Unmarshal([]byte(msgsStr), &msgsArray)
			var messages []Message
			for _, msgMap := range msgsArray {
				// 处理body字段：如果是对象，转换为JSON字符串
				if bodyVal, exists := msgMap["body"]; exists && bodyVal != nil {
					if bodyMap, ok := bodyVal.(map[string]interface{}); ok {
						bodyJson, _ := json.Marshal(bodyMap)
						msgMap["body"] = string(bodyJson)
					} else if bodyStr, ok := bodyVal.(string); ok {
						msgMap["body"] = bodyStr
					}
				}
				var msg Message
				msgJson, _ := json.Marshal(msgMap)
				json.Unmarshal(msgJson, &msg)
				messages = append(messages, msg)
			}
			resp.Msgs = messages
		}
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
	// 按照 Java SDK 标准：从根对象解析size和msgs
	if sizeVal, ok := jsonObj["size"]; ok {
		if sizeFloat, ok := sizeVal.(float64); ok {
			resp.Size = int(sizeFloat)
		}
	}
	// 按照 Java SDK 标准：msgs是JSON字符串，需要先解析成字符串再解析成数组
	if msgsVal, ok := jsonObj["msgs"]; ok && msgsVal != nil {
		var msgsStr string
		if msgsStrVal, ok := msgsVal.(string); ok {
			msgsStr = msgsStrVal
		} else {
			msgsJson, _ := json.Marshal(msgsVal)
			msgsStr = string(msgsJson)
		}
		if msgsStr != "" {
			var msgsArray []map[string]interface{}
			json.Unmarshal([]byte(msgsStr), &msgsArray)
			var messages []ChatroomMessage
			for _, msgMap := range msgsArray {
				// 处理body字段：如果是对象，转换为JSON字符串
				if bodyVal, exists := msgMap["body"]; exists && bodyVal != nil {
					if bodyMap, ok := bodyVal.(map[string]interface{}); ok {
						bodyJson, _ := json.Marshal(bodyMap)
						msgMap["body"] = string(bodyJson)
					} else if bodyStr, ok := bodyVal.(string); ok {
						msgMap["body"] = bodyStr
					}
				}
				var msg ChatroomMessage
				msgJson, _ := json.Marshal(msgMap)
				json.Unmarshal(msgJson, &msg)
				messages = append(messages, msg)
			}
			resp.Msgs = messages
		}
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
	// 按照 Java SDK 标准：从data对象解析
	if dataVal, ok := jsonObj["data"]; ok {
		if dataMap, ok := dataVal.(map[string]interface{}); ok {
			// 解析hasMore字段
			if hasMoreVal, ok := dataMap["hasMore"]; ok {
				if hasMoreBool, ok := hasMoreVal.(bool); ok {
					resp.HasMore = hasMoreBool
				}
			}
			// 解析sessions数组，需要特殊处理null字段
			if sessionsVal, ok := dataMap["sessions"]; ok {
				if sessionsArray, ok := sessionsVal.([]interface{}); ok {
					var sessions []Session
					for _, sessionItem := range sessionsArray {
						if sessionMap, ok := sessionItem.(map[string]interface{}); ok {
							var session Session
							if sessionTypeVal, ok := sessionMap["sessionType"]; ok {
								if sessionTypeFloat, ok := sessionTypeVal.(float64); ok {
									session.SessionType = int(sessionTypeFloat)
								}
							}
							// 只解析非null的accid
							if accidVal, ok := sessionMap["accid"]; ok && accidVal != nil {
								if accidStr, ok := accidVal.(string); ok {
									session.Accid = accidStr
								}
							}
							// 只解析非null的tid
							if tidVal, ok := sessionMap["tid"]; ok && tidVal != nil {
								if tidFloat, ok := tidVal.(float64); ok {
									session.Tid = int64(tidFloat)
								}
							}
							if updateTimeVal, ok := sessionMap["updateTime"]; ok {
								if updateTimeFloat, ok := updateTimeVal.(float64); ok {
									session.UpdateTime = int64(updateTimeFloat)
								}
							}
							if extVal, ok := sessionMap["ext"]; ok {
								if extStr, ok := extVal.(string); ok {
									session.Ext = extStr
								}
							}
							if lastMsgTypeVal, ok := sessionMap["lastMsgType"]; ok {
								if lastMsgTypeStr, ok := lastMsgTypeVal.(string); ok {
									session.LastMsgType = lastMsgTypeStr
								}
							}
							if lastMsgVal, ok := sessionMap["lastMsg"]; ok {
								if lastMsgStr, ok := lastMsgVal.(string); ok {
									session.LastMsg = lastMsgStr
								}
							}
							sessions = append(sessions, session)
						}
					}
					resp.Sessions = sessions
				}
			}
		}
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
	// Java SDK从根对象的msg字段解析
	if msgVal, ok := jsonObj["msg"]; ok && msgVal != nil {
		msgJson, _ := json.Marshal(msgVal)
		var broadcastMsg BroadcastMessage
		json.Unmarshal(msgJson, &broadcastMsg)
		resp.Msg = broadcastMsg
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
	// Java SDK从根对象的msgs字段解析
	if msgsVal, ok := jsonObj["msgs"]; ok && msgsVal != nil {
		msgsJson, _ := json.Marshal(msgsVal)
		json.Unmarshal(msgsJson, &resp.Msgs)
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
	// 按照 Java SDK 标准：从根对象解析events和size
	if eventsVal, ok := jsonObj["events"]; ok {
		if eventsArray, ok := eventsVal.([]interface{}); ok {
			var userEvents []UserEvent
			for _, evt := range eventsArray {
				var userEvent UserEvent
				evtJson, _ := json.Marshal(evt)
				json.Unmarshal(evtJson, &userEvent)
				userEvents = append(userEvents, userEvent)
			}
			resp.Events = userEvents
		}
	}
	if sizeVal, ok := jsonObj["size"]; ok {
		if sizeFloat, ok := sizeVal.(float64); ok {
			resp.Size = int(sizeFloat)
		}
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}
