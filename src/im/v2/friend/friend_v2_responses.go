package friend

// AddFriendResponseV2 添加好友响应
type AddFriendResponseV2 struct {
	FriendAccountId   string `json:"friend_account_id,omitempty"`  // 好友账号ID
	Alias             string `json:"alias,omitempty"`              // 好友别名
	Extension         string `json:"extension,omitempty"`          // 扩展字段
	CustomerExtension string `json:"customer_extension,omitempty"` // 客户端扩展字段
}

// DeleteFriendResponseV2 删除好友响应
type DeleteFriendResponseV2 struct {
	// Success response will only contain status code 200
}

// UpdateFriendResponseV2 更新好友响应
type UpdateFriendResponseV2 struct {
	FriendAccountId   string `json:"friend_account_id,omitempty"`  // 好友账号ID
	Alias             string `json:"alias,omitempty"`              // 好友别名
	Extension         string `json:"extension,omitempty"`          // 扩展字段
	CustomerExtension string `json:"customer_extension,omitempty"` // 客户端扩展字段
}

// GetFriendResponseV2 获取好友响应
type GetFriendResponseV2 struct {
	FriendAccountId   string `json:"friend_account_id,omitempty"`  // 好友账号ID
	Alias             string `json:"alias,omitempty"`              // 好友别名
	Extension         string `json:"extension,omitempty"`          // 扩展字段
	CustomerExtension string `json:"customer_extension,omitempty"` // 客户端扩展字段
}

// ListFriendsResponseV2 列出好友响应
type ListFriendsResponseV2 struct {
	HasMore   bool           `json:"has_more,omitempty"`   // 是否有更多数据
	NextToken string         `json:"next_token,omitempty"` // 下一页token
	Items     []FriendItemV2 `json:"items,omitempty"`      // 好友列表
	Total     int            `json:"total,omitempty"`      // 总数
}

// HandleFriendAdditionResponseV2 处理好友添加响应
type HandleFriendAdditionResponseV2 struct {
	FriendAccountId   string `json:"friend_account_id,omitempty"`  // 好友账号ID
	Alias             string `json:"alias,omitempty"`              // 好友别名
	Extension         string `json:"extension,omitempty"`          // 扩展字段
	CustomerExtension string `json:"customer_extension,omitempty"` // 客户端扩展字段
}

// FriendItemV2 好友项
type FriendItemV2 struct {
	FriendAccountId   string `json:"friend_account_id,omitempty"`  // 好友账号ID
	Alias             string `json:"alias,omitempty"`              // 好友别名
	Extension         string `json:"extension,omitempty"`          // 扩展字段
	CustomerExtension string `json:"customer_extension,omitempty"` // 客户端扩展字段
	CreateTime        int64  `json:"create_time,omitempty"`        // 创建时间
	UpdateTime        int64  `json:"update_time,omitempty"`        // 更新时间
}
