package chatroom_message

import (
	"encoding/json"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/utils"
)

// SendMsg 发送聊天室消息
func (c *ChatroomMessageV1ServiceImpl) SendMsg(req *SendChatroomMsgRequestV1) (*core.Result[*SendChatroomMsgResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := c.httpClient.ExecuteV1Api(SendMsg, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*SendChatroomMsgResponseV1)(nil)), nil
	}

	resp := &SendChatroomMsgResponseV1{}
	if descVal, ok := jsonObj["desc"]; ok {
		descJson, _ := json.Marshal(descVal)
		json.Unmarshal(descJson, resp)
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// BatchSendMsg 批量发送聊天室消息
func (c *ChatroomMessageV1ServiceImpl) BatchSendMsg(req *BatchSendChatroomMsgRequestV1) (*core.Result[*BatchSendChatroomMsgResponseV1], error) {
	paramMap := utils.Convert(req)

	// Convert msgList to JSON string
	if len(req.MsgList) > 0 {
		msgListJson, _ := json.Marshal(req.MsgList)
		paramMap["msgList"] = string(msgListJson)
	}

	apiResponse, err := c.httpClient.ExecuteV1Api(BatchSendMsg, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*BatchSendChatroomMsgResponseV1)(nil)), nil
	}

	resp := &BatchSendChatroomMsgResponseV1{}

	// Parse fail list
	if failVal, ok := jsonObj["fail"]; ok {
		failJson, _ := json.Marshal(failVal)
		var failArray []map[string]interface{}
		json.Unmarshal(failJson, &failArray)

		for _, failObj := range failArray {
			for key, value := range failObj {
				failMsg := FailedMessage{
					ClientMsgId: key,
					Reason:      value.(string),
				}
				resp.FailList = append(resp.FailList, failMsg)
			}
		}
	}

	// Parse success list
	if successVal, ok := jsonObj["success"]; ok {
		successJson, _ := json.Marshal(successVal)
		var successArray []map[string]interface{}
		json.Unmarshal(successJson, &successArray)

		for _, successObj := range successArray {
			for key, value := range successObj {
				valueJson, _ := json.Marshal(value)
				var msg SuccessfulMessage
				json.Unmarshal(valueJson, &msg)
				msg.ClientMsgId = key
				resp.SuccessList = append(resp.SuccessList, msg)
			}
		}
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// Recall 撤回聊天室消息
func (c *ChatroomMessageV1ServiceImpl) Recall(req *RecallChatroomMsgRequestV1) (*core.Result[*RecallChatroomMsgResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := c.httpClient.ExecuteV1Api(Recall, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*RecallChatroomMsgResponseV1)(nil)), nil
	}

	resp := &RecallChatroomMsgResponseV1{}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// SendMsgToSomeone 发送聊天室定向消息
func (c *ChatroomMessageV1ServiceImpl) SendMsgToSomeone(req *ChatroomTargetMsgRequestV1) (*core.Result[*ChatroomTargetMsgResponseV1], error) {
	paramMap := utils.Convert(req)

	// Convert toAccids to JSON string
	if len(req.ToAccids) > 0 {
		toAccidsJson, _ := json.Marshal(req.ToAccids)
		paramMap["toAccids"] = string(toAccidsJson)
	}

	apiResponse, err := c.httpClient.ExecuteV1Api(SendMsgToSomeone, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*ChatroomTargetMsgResponseV1)(nil)), nil
	}

	resp := &ChatroomTargetMsgResponseV1{}
	if descVal, ok := jsonObj["desc"]; ok {
		descJson, _ := json.Marshal(descVal)
		json.Unmarshal(descJson, resp)
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// BatchSendMsgToSomeone 批量发送聊天室定向消息
func (c *ChatroomMessageV1ServiceImpl) BatchSendMsgToSomeone(req *BatchChatroomTargetMsgRequestV1) (*core.Result[*BatchChatroomTargetMsgResponseV1], error) {
	paramMap := utils.Convert(req)

	// Convert msgList to JSON string
	if len(req.MsgList) > 0 {
		msgListJson, _ := json.Marshal(req.MsgList)
		paramMap["msgList"] = string(msgListJson)
	}

	// Convert toAccids to JSON string
	if len(req.ToAccids) > 0 {
		toAccidsJson, _ := json.Marshal(req.ToAccids)
		paramMap["toAccids"] = string(toAccidsJson)
	}

	apiResponse, err := c.httpClient.ExecuteV1Api(BatchSendMsgToSomeone, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*BatchChatroomTargetMsgResponseV1)(nil)), nil
	}

	resp := &BatchChatroomTargetMsgResponseV1{}

	// Parse fail list
	if failVal, ok := jsonObj["fail"]; ok {
		failJson, _ := json.Marshal(failVal)
		var failArray []map[string]interface{}
		json.Unmarshal(failJson, &failArray)

		for _, failObj := range failArray {
			for key, value := range failObj {
				failMsg := FailedMessage{
					ClientMsgId: key,
					Reason:      value.(string),
				}
				resp.FailList = append(resp.FailList, failMsg)
			}
		}
	}

	// Parse success list
	if successVal, ok := jsonObj["success"]; ok {
		successJson, _ := json.Marshal(successVal)
		var successArray []map[string]interface{}
		json.Unmarshal(successJson, &successArray)

		for _, successObj := range successArray {
			for key, value := range successObj {
				valueJson, _ := json.Marshal(value)
				var msg ChatroomTargetMsgResponseV1
				json.Unmarshal(valueJson, &msg)
				msg.MsgidClient = key
				resp.SuccessList = append(resp.SuccessList, msg)
			}
		}
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// Broadcast 发送聊天室全服广播消息
func (c *ChatroomMessageV1ServiceImpl) Broadcast(req *ChatroomBroadcastRequestV1) (*core.Result[*SendChatroomMsgResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := c.httpClient.ExecuteV1Api(Broadcast, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*SendChatroomMsgResponseV1)(nil)), nil
	}

	resp := &SendChatroomMsgResponseV1{}
	if descVal, ok := jsonObj["desc"]; ok {
		descJson, _ := json.Marshal(descVal)
		json.Unmarshal(descJson, resp)
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}
