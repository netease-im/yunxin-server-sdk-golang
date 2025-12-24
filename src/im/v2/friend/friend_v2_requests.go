package friend

// AddFriendRequestV2 添加好友请求
type AddFriendRequestV2 struct {
	AccountId         string `json:"account_id"`                   // 账号ID
	FriendAccountId   string `json:"friend_account_id"`            // 好友账号ID
	Alias             string `json:"alias,omitempty"`              // 好友别名
	Extension         string `json:"extension,omitempty"`          // 扩展字段
	CustomerExtension string `json:"customer_extension,omitempty"` // 客户端扩展字段
	Postscript        string `json:"postscript,omitempty"`         // 附言
}

// DeleteFriendRequestV2 删除好友请求
type DeleteFriendRequestV2 struct {
	AccountId       string `json:"-"`                      // 账号ID，用于构造URL
	FriendAccountId string `json:"friend_account_id"`      // 好友账号ID
	DeleteAlias     bool   `json:"delete_alias,omitempty"` // 是否删除别名
}

// UpdateFriendRequestV2 更新好友请求
type UpdateFriendRequestV2 struct {
	AccountId         string `json:"-"`                            // 账号ID，用于构造URL
	FriendAccountId   string `json:"friend_account_id"`            // 好友账号ID
	Alias             string `json:"alias,omitempty"`              // 好友别名
	Extension         string `json:"extension,omitempty"`          // 扩展字段
	CustomerExtension string `json:"customer_extension,omitempty"` // 客户端扩展字段
}

// GetFriendRequestV2 获取好友请求
type GetFriendRequestV2 struct {
	AccountId       string `json:"-"` // 账号ID，用于构造URL
	FriendAccountId string `json:"-"` // 好友账号ID，用于构造查询参数
}

// ListFriendsRequestV2 列出好友请求
type ListFriendsRequestV2 struct {
	AccountId string `json:"-"` // 账号ID，用于构造查询参数
	PageToken string `json:"-"` // 分页token，用于构造查询参数
	Limit     int    `json:"-"` // 限制数量，用于构造查询参数
}

// HandleFriendAdditionRequestV2 处理好友添加请求
type HandleFriendAdditionRequestV2 struct {
	AccountId         string `json:"account_id"`                   // 账号ID
	FriendAccountId   string `json:"friend_account_id"`            // 好友账号ID
	Alias             string `json:"alias,omitempty"`              // 好友别名
	Extension         string `json:"extension,omitempty"`          // 扩展字段
	CustomerExtension string `json:"customer_extension,omitempty"` // 客户端扩展字段
	Postscript        string `json:"postscript,omitempty"`         // 附言
	HandleType        int    `json:"handle_type"`                  // 处理类型：1-接受，2-拒绝
}
