package block

// AddBlockContactRequestV2 添加黑名单联系人请求
type AddBlockContactRequestV2 struct {
	AccountId        string `json:"account_id"`         // 账号ID
	ContactAccountId string `json:"contact_account_id"` // 联系人账号ID
}

// RemoveBlockContactRequestV2 移除黑名单联系人请求
type RemoveBlockContactRequestV2 struct {
	AccountId        string `json:"account_id"`         // 账号ID
	ContactAccountId string `json:"contact_account_id"` // 联系人账号ID
}

// ListBlockContactsRequestV2 查询黑名单列表请求
type ListBlockContactsRequestV2 struct {
	AccountId string `json:"account_id"`           // 账号ID
	PageToken string `json:"page_token,omitempty"` // 分页标记
	Limit     int    `json:"limit,omitempty"`      // 每页数量
}
