package block

import (
	"encoding/json"
	"strconv"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/http"
	"github.com/netease-im/yunxin-server-sdk-golang/src/im/v2/util"
)

// AddBlockContact 添加黑名单联系人
func (b *BlockV2ServiceImpl) AddBlockContact(req *AddBlockContactRequestV2) (*core.Result[*AddBlockContactResponseV2], error) {
	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	apiResponse, err := b.httpClient.ExecuteV2Api(http.POST, AddBlockContact, nil, nil, string(requestBody))
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*AddBlockContactResponseV2](apiResponse)
}

// RemoveBlockContact 移除黑名单联系人
func (b *BlockV2ServiceImpl) RemoveBlockContact(req *RemoveBlockContactRequestV2) (*core.Result[*RemoveBlockContactResponseV2], error) {
	pathParams := map[string]string{
		"account_id": req.AccountId,
	}

	queryParams := map[string]string{
		"contact_account_id": req.ContactAccountId,
	}

	apiResponse, err := b.httpClient.ExecuteV2Api(http.DELETE, RemoveBlockContact, pathParams, queryParams, "")
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*RemoveBlockContactResponseV2](apiResponse)
}

// ListBlockContacts 查询黑名单列表
func (b *BlockV2ServiceImpl) ListBlockContacts(req *ListBlockContactsRequestV2) (*core.Result[*ListBlockContactsResponseV2], error) {
	queryParams := map[string]string{
		"account_id": req.AccountId,
	}

	if req.PageToken != "" {
		queryParams["page_token"] = req.PageToken
	}

	if req.Limit > 0 {
		queryParams["limit"] = strconv.Itoa(req.Limit)
	}

	apiResponse, err := b.httpClient.ExecuteV2Api(http.GET, ListBlockContacts, nil, queryParams, "")
	if err != nil {
		return nil, err
	}

	return util.ConvertWithClass[*ListBlockContactsResponseV2](apiResponse)
}
