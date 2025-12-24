package mute

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/http"
	"github.com/netease-im/yunxin-server-sdk-golang/src/im/v2/util"
)

// AddMuteContact 添加免打扰联系人
func (m *MuteV2ServiceImpl) AddMuteContact(req *AddMuteContactRequestV2) (*core.Result[*AddMuteContactResponseV2], error) {
	// Validate required parameters
	if req.AccountId == "" {
		return nil, fmt.Errorf("account ID cannot be empty")
	}

	if req.ContactAccountId == "" {
		return nil, fmt.Errorf("contact account ID cannot be empty")
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := m.httpClient.ExecuteV2Api(http.POST, AddMuteContact, nil, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*AddMuteContactResponseV2](apiResponse)
}

// RemoveMuteContact 移除免打扰联系人
func (m *MuteV2ServiceImpl) RemoveMuteContact(req *RemoveMuteContactRequestV2) (*core.Result[*RemoveMuteContactResponseV2], error) {
	// Validate required parameters
	if req.AccountId == "" {
		return nil, fmt.Errorf("account ID cannot be empty")
	}

	if req.ContactAccountId == "" {
		return nil, fmt.Errorf("contact account ID cannot be empty")
	}

	pathParams := map[string]string{
		"account_id": req.AccountId,
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := m.httpClient.ExecuteV2Api(http.DELETE, RemoveMuteContact, pathParams, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*RemoveMuteContactResponseV2](apiResponse)
}

// ListMuteContacts 列出免打扰联系人
func (m *MuteV2ServiceImpl) ListMuteContacts(req *ListMuteContactsRequestV2) (*core.Result[*ListMuteContactsResponseV2], error) {
	// Validate required parameters
	if req.AccountId == "" {
		return nil, fmt.Errorf("account ID cannot be empty")
	}

	queryParams := map[string]string{
		"account_id": req.AccountId,
	}

	if req.PageToken != "" {
		queryParams["page_token"] = req.PageToken
	}

	if req.Limit > 0 {
		queryParams["limit"] = strconv.Itoa(req.Limit)
	}

	apiResponse, err := m.httpClient.ExecuteV2Api(http.GET, ListMuteContacts, nil, queryParams, "")
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*ListMuteContactsResponseV2](apiResponse)
}
