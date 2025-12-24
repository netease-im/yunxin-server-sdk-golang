package friend

import (
	"encoding/json"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/utils"
)

// Add 添加好友
func (f *FriendV1ServiceImpl) Add(req *AddFriendRequestV1) (*core.Result[*AddFriendResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := f.httpClient.ExecuteV1Api(Add, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*AddFriendResponseV1)(nil)), nil
	}

	resp := &AddFriendResponseV1{}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// Update 更新好友相关信息
func (f *FriendV1ServiceImpl) Update(req *UpdateFriendRequestV1) (*core.Result[*UpdateFriendResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := f.httpClient.ExecuteV1Api(Update, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*UpdateFriendResponseV1)(nil)), nil
	}

	resp := &UpdateFriendResponseV1{}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// Delete 删除好友关系
func (f *FriendV1ServiceImpl) Delete(req *DeleteFriendRequestV1) (*core.Result[*DeleteFriendResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := f.httpClient.ExecuteV1Api(Delete, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*DeleteFriendResponseV1)(nil)), nil
	}

	resp := &DeleteFriendResponseV1{}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// Get 获取好友列表
func (f *FriendV1ServiceImpl) Get(req *GetFriendListRequestV1) (*core.Result[*GetFriendListResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := f.httpClient.ExecuteV1Api(Get, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*GetFriendListResponseV1)(nil)), nil
	}

	resp := &GetFriendListResponseV1{}
	if friendsVal, ok := jsonObj["friends"]; ok {
		friendsJson, _ := json.Marshal(friendsVal)
		json.Unmarshal(friendsJson, &resp.Friends)
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// GetByAccid 获取好友关系
func (f *FriendV1ServiceImpl) GetByAccid(req *GetFriendRelationshipRequestV1) (*core.Result[*GetFriendRelationshipResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := f.httpClient.ExecuteV1Api(GetByAccid, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*GetFriendRelationshipResponseV1)(nil)), nil
	}

	resp := &GetFriendRelationshipResponseV1{}
	if friendVal, ok := jsonObj["friend"]; ok {
		friendJson, _ := json.Marshal(friendVal)
		json.Unmarshal(friendJson, resp)
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}
