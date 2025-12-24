package mute

// AddMuteContactRequestV2 添加免打扰联系人请求
type AddMuteContactRequestV2 struct {
	AccountId        string `json:"account_id"`         // 账号ID
	ContactAccountId string `json:"contact_account_id"` // 联系人账号ID
}

// RemoveMuteContactRequestV2 移除免打扰联系人请求
type RemoveMuteContactRequestV2 struct {
	AccountId        string `json:"-"`                  // 账号ID，用于构造URL
	ContactAccountId string `json:"contact_account_id"` // 联系人账号ID
}

// ListMuteContactsRequestV2 列出免打扰联系人请求
type ListMuteContactsRequestV2 struct {
	AccountId string `json:"-"` // 账号ID，用于构造查询参数
	PageToken string `json:"-"` // 分页token，用于构造查询参数
	Limit     int    `json:"-"` // 限制数量，用于构造查询参数
}
