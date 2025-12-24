package mute

// AddMuteContactResponseV2 添加免打扰联系人响应
type AddMuteContactResponseV2 struct {
	// Success response will only contain status code 200
}

// RemoveMuteContactResponseV2 移除免打扰联系人响应
type RemoveMuteContactResponseV2 struct {
	// Success response will only contain status code 200
}

// ListMuteContactsResponseV2 列出免打扰联系人响应
type ListMuteContactsResponseV2 struct {
	HasMore   bool     `json:"has_more,omitempty"`   // 是否有更多数据
	NextToken string   `json:"next_token,omitempty"` // 下一页token
	Items     []string `json:"items,omitempty"`      // 免打扰联系人列表
}
