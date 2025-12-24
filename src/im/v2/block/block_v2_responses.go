package block

// AddBlockContactResponseV2 添加黑名单联系人响应
type AddBlockContactResponseV2 struct {
	// Empty struct as the API returns an empty object on success
}

// RemoveBlockContactResponseV2 移除黑名单联系人响应
type RemoveBlockContactResponseV2 struct {
	// Empty struct as the API returns an empty object on success
}

// ListBlockContactsResponseV2 查询黑名单列表响应
type ListBlockContactsResponseV2 struct {
	HasMore   bool     `json:"has_more,omitempty"`   // 是否有更多数据
	NextToken string   `json:"next_token,omitempty"` // 下一页标记
	Items     []string `json:"items,omitempty"`      // 黑名单联系人账号ID列表
}
