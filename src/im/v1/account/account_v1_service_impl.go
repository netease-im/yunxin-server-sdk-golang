package account

import (
	"encoding/json"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/utils"
)

// CreateAccount 创建账户
func (a *AccountV1ServiceImpl) CreateAccount(req *CreateAccountRequestV1) (*core.Result[*CreateAccountResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := a.httpClient.ExecuteV1Api(CreateAccount, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*CreateAccountResponseV1)(nil)), nil
	}

	resp := &CreateAccountResponseV1{}
	if dataVal, ok := jsonObj["info"]; ok {
		dataJson, _ := json.Marshal(dataVal)
		json.Unmarshal(dataJson, resp)
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// UpdateToken 更新Token
func (a *AccountV1ServiceImpl) UpdateToken(req *UpdateTokenRequestV1) (*core.Result[*UpdateTokenResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := a.httpClient.ExecuteV1Api(UpdateToken, paramMap)
	if err != nil {
		return nil, err
	}

	var jsonObj map[string]interface{}
	if err := json.Unmarshal([]byte(apiResponse.GetData()), &jsonObj); err != nil {
		return nil, err
	}

	code := int(jsonObj["code"].(float64))
	desc := ""
	if descVal, ok := jsonObj["desc"]; ok {
		desc = descVal.(string)
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, &UpdateTokenResponseV1{}), nil
}

// RefreshToken 刷新Token
func (a *AccountV1ServiceImpl) RefreshToken(req *RefreshTokenRequestV1) (*core.Result[*RefreshTokenResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := a.httpClient.ExecuteV1Api(RefreshToken, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*RefreshTokenResponseV1)(nil)), nil
	}

	resp := &RefreshTokenResponseV1{}
	if dataVal, ok := jsonObj["info"]; ok {
		dataJson, _ := json.Marshal(dataVal)
		json.Unmarshal(dataJson, resp)
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// Block 封禁账户
func (a *AccountV1ServiceImpl) Block(req *BlockAccountRequestV1) (*core.Result[*BlockAccountResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := a.httpClient.ExecuteV1Api(BlockAccount, paramMap)
	if err != nil {
		return nil, err
	}

	var jsonObj map[string]interface{}
	if err := json.Unmarshal([]byte(apiResponse.GetData()), &jsonObj); err != nil {
		return nil, err
	}

	code := int(jsonObj["code"].(float64))
	desc := ""
	if descVal, ok := jsonObj["desc"]; ok {
		desc = descVal.(string)
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, &BlockAccountResponseV1{}), nil
}

// Unblock 解除账户封禁
func (a *AccountV1ServiceImpl) Unblock(req *UnBlockAccountRequestV1) (*core.Result[*UnBlockAccountResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := a.httpClient.ExecuteV1Api(UnblockAccount, paramMap)
	if err != nil {
		return nil, err
	}

	var jsonObj map[string]interface{}
	if err := json.Unmarshal([]byte(apiResponse.GetData()), &jsonObj); err != nil {
		return nil, err
	}

	code := int(jsonObj["code"].(float64))
	desc := ""
	if descVal, ok := jsonObj["desc"]; ok {
		desc = descVal.(string)
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, &UnBlockAccountResponseV1{}), nil
}

// Mute 账户禁言
func (a *AccountV1ServiceImpl) Mute(req *MuteAccountRequestV1) (*core.Result[*MuteAccountResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := a.httpClient.ExecuteV1Api(MuteAccount, paramMap)
	if err != nil {
		return nil, err
	}

	var jsonObj map[string]interface{}
	if err := json.Unmarshal([]byte(apiResponse.GetData()), &jsonObj); err != nil {
		return nil, err
	}

	code := int(jsonObj["code"].(float64))
	desc := ""
	if descVal, ok := jsonObj["desc"]; ok {
		desc = descVal.(string)
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, &MuteAccountResponseV1{}), nil
}

// MuteModule 模块禁言
func (a *AccountV1ServiceImpl) MuteModule(req *MuteModuleRequestV1) (*core.Result[*MuteModuleResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := a.httpClient.ExecuteV1Api(MuteModule, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*MuteModuleResponseV1)(nil)), nil
	}

	resp := &MuteModuleResponseV1{}
	if dataVal, ok := jsonObj["data"]; ok {
		dataJson, _ := json.Marshal(dataVal)
		json.Unmarshal(dataJson, resp)
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// SetDonnop 设置免打扰
func (a *AccountV1ServiceImpl) SetDonnop(req *SetDonnopRequestV1) (*core.Result[*SetDonnopResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := a.httpClient.ExecuteV1Api(SetDonnop, paramMap)
	if err != nil {
		return nil, err
	}

	var jsonObj map[string]interface{}
	if err := json.Unmarshal([]byte(apiResponse.GetData()), &jsonObj); err != nil {
		return nil, err
	}

	code := int(jsonObj["code"].(float64))
	desc := ""
	if descVal, ok := jsonObj["desc"]; ok {
		desc = descVal.(string)
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, &SetDonnopResponseV1{}), nil
}

// QueryOnlineStatus 查询账户在线状态
func (a *AccountV1ServiceImpl) QueryOnlineStatus(req *QueryAccountOnlineStatusRequestV1) (*core.Result[*QueryAccountOnlineStatusResponseV1], error) {
	paramMap := make(map[string]string)
	if len(req.Accids) > 0 {
		accidsJson, _ := json.Marshal(req.Accids)
		paramMap["accids"] = string(accidsJson)
	}

	apiResponse, err := a.httpClient.ExecuteV1Api(QueryOnlineStatus, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*QueryAccountOnlineStatusResponseV1)(nil)), nil
	}

	resp := &QueryAccountOnlineStatusResponseV1{}
	if dataVal, ok := jsonObj["data"]; ok {
		dataJson, _ := json.Marshal(dataVal)
		json.Unmarshal(dataJson, resp)
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// QueryUserInfos 获取用户信息
func (a *AccountV1ServiceImpl) QueryUserInfos(req *QueryUserInfosRequestV1) (*core.Result[*QueryUserInfosResponseV1], error) {
	paramMap := make(map[string]string)
	if len(req.Accids) > 0 {
		accidsJson, _ := json.Marshal(req.Accids)
		paramMap["accids"] = string(accidsJson)
	}
	if req.MuteStatus {
		paramMap["muteStatus"] = "true"
	}

	apiResponse, err := a.httpClient.ExecuteV1Api(QueryUserInfos, paramMap)
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
		return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, (*QueryUserInfosResponseV1)(nil)), nil
	}

	resp := &QueryUserInfosResponseV1{}
	if dataVal, ok := jsonObj["uinfos"]; ok {
		dataJson, _ := json.Marshal(dataVal)
		json.Unmarshal(dataJson, &resp.UserInfoList)
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), "", resp), nil
}

// UpdateUinfo 更新用户信息
func (a *AccountV1ServiceImpl) UpdateUinfo(req *UpdateUinfoRequestV1) (*core.Result[*UpdateUinfoResponseV1], error) {
	paramMap := utils.Convert(req)

	apiResponse, err := a.httpClient.ExecuteV1Api(UpdateUinfo, paramMap)
	if err != nil {
		return nil, err
	}

	var jsonObj map[string]interface{}
	if err := json.Unmarshal([]byte(apiResponse.GetData()), &jsonObj); err != nil {
		return nil, err
	}

	code := int(jsonObj["code"].(float64))
	desc := ""
	if descVal, ok := jsonObj["desc"]; ok {
		desc = descVal.(string)
	}

	return core.NewResult(apiResponse.GetEndpoint(), code, apiResponse.GetTraceId(), desc, &UpdateUinfoResponseV1{}), nil
}
