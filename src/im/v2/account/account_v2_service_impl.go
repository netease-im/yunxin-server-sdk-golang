package account

import (
	"encoding/json"
	"strings"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/http"
	"github.com/netease-im/yunxin-server-sdk-golang/src/im/v2/util"
)

// CreateAccount 创建账号
func (a *AccountV2ServiceImpl) CreateAccount(req *CreateAccountRequestV2) (*core.Result[*CreateAccountResponseV2], error) {
	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := a.httpClient.ExecuteV2Api(http.POST, Accounts, nil, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*CreateAccountResponseV2](apiResponse)
}

// UpdateAccount 更新账号
func (a *AccountV2ServiceImpl) UpdateAccount(req *UpdateAccountRequestV2) (*core.Result[*UpdateAccountResponseV2], error) {
	pathParams := map[string]string{
		"account_id": req.AccountId,
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := a.httpClient.ExecuteV2Api(http.PATCH, AccountWithID, pathParams, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*UpdateAccountResponseV2](apiResponse)
}

// BatchQueryAccounts 批量查询账号
func (a *AccountV2ServiceImpl) BatchQueryAccounts(req *BatchQueryAccountsRequestV2) (*core.Result[*BatchQueryAccountsResponseV2], error) {
	queryParams := map[string]string{
		"account_ids": strings.Join(req.AccountList, ","),
	}

	apiResponse, err := a.httpClient.ExecuteV2Api(http.GET, Accounts, nil, queryParams, "")
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*BatchQueryAccountsResponseV2](apiResponse)
}

// DisableAccount 禁用/启用账号
func (a *AccountV2ServiceImpl) DisableAccount(req *DisableAccountRequestV2) (*core.Result[*DisableAccountResponseV2], error) {
	pathParams := map[string]string{
		"account_id": req.AccountId,
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := a.httpClient.ExecuteV2Api(http.PATCH, DisableAccount, pathParams, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*DisableAccountResponseV2](apiResponse)
}

// SetPushConfig 设置推送配置
func (a *AccountV2ServiceImpl) SetPushConfig(req *SetPushConfigRequestV2) (*core.Result[*SetPushConfigResponseV2], error) {
	pathParams := map[string]string{
		"account_id": req.AccountId,
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := a.httpClient.ExecuteV2Api(http.PATCH, SetPushConfig, pathParams, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*SetPushConfigResponseV2](apiResponse)
}

// GetAccountDetails 获取账号详情
func (a *AccountV2ServiceImpl) GetAccountDetails(req *GetAccountDetailsRequestV2) (*core.Result[*GetAccountDetailsResponseV2], error) {
	pathParams := map[string]string{
		"account_id": req.AccountId,
	}

	apiResponse, err := a.httpClient.ExecuteV2Api(http.GET, AccountWithID, pathParams, nil, "")
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*GetAccountDetailsResponseV2](apiResponse)
}

// KickAccount 踢账号下线
func (a *AccountV2ServiceImpl) KickAccount(req *KickAccountRequestV2) (*core.Result[*KickAccountResponseV2], error) {
	pathParams := map[string]string{
		"account_id": req.AccountId,
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := a.httpClient.ExecuteV2Api(http.POST, KickAccount, pathParams, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*KickAccountResponseV2](apiResponse)
}

// RefreshToken 刷新Token
func (a *AccountV2ServiceImpl) RefreshToken(req *RefreshTokenRequestV2) (*core.Result[*RefreshTokenResponseV2], error) {
	pathParams := map[string]string{
		"account_id": req.AccountId,
	}

	apiResponse, err := a.httpClient.ExecuteV2Api(http.PATCH, RefreshToken, pathParams, nil, "")
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*RefreshTokenResponseV2](apiResponse)
}
