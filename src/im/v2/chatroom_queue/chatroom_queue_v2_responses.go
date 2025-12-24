package chatroom_queue

// InitializeChatroomQueueResponseV2 初始化聊天室队列响应
type InitializeChatroomQueueResponseV2 struct {
	// Success response will only contain status code 200
}

// QueryChatroomQueueElementsResponseV2 查询聊天室队列元素响应
type QueryChatroomQueueElementsResponseV2 struct {
	RoomId         int64                  `json:"room_id,omitempty"`          // 聊天室ID
	QueueSizeLimit int64                  `json:"queue_size_limit,omitempty"` // 队列大小限制
	QueueSize      int64                  `json:"queue_size,omitempty"`       // 队列大小
	ElementList    []QueueElementResponse `json:"element_list,omitempty"`     // 元素列表
}

// UpdateChatroomQueueResponseV2 更新聊天室队列响应
type UpdateChatroomQueueResponseV2 struct {
	RoomId            int64           `json:"room_id,omitempty"`             // 聊天室ID
	QueueSizeLimit    int64           `json:"queue_size_limit,omitempty"`    // 队列大小限制
	QueueSize         int64           `json:"queue_size,omitempty"`          // 队列大小
	InsertElementList []string        `json:"insert_element_list,omitempty"` // 插入元素列表
	UpdateElementList []string        `json:"update_element_list,omitempty"` // 更新元素列表
	FailedElementList []FailedElement `json:"failed_element_list,omitempty"` // 失败元素列表
}

// DeleteChatroomQueueResponseV2 删除聊天室队列响应
type DeleteChatroomQueueResponseV2 struct {
	RoomId       int64 `json:"room_id,omitempty"`       // 聊天室ID
	QueueSize    int   `json:"queue_size,omitempty"`    // 队列大小
	HighPriority int   `json:"high_priority,omitempty"` // 高优先级
}

// PollChatroomQueueElementResponseV2 轮询聊天室队列元素响应
type PollChatroomQueueElementResponseV2 struct {
	ElementKey       string `json:"element_key,omitempty"`        // 元素键
	ElementValue     string `json:"element_value,omitempty"`      // 元素值
	ElementAccountId string `json:"element_account_id,omitempty"` // 元素账号ID
	ElementTransient bool   `json:"element_transient,omitempty"`  // 元素是否临时
}

// QueueElementResponse 队列元素响应
type QueueElementResponse struct {
	ElementKey       string `json:"element_key,omitempty"`        // 元素键
	ElementValue     string `json:"element_value,omitempty"`      // 元素值
	ElementAccountId string `json:"element_account_id,omitempty"` // 元素账号ID
	ElementTransient bool   `json:"element_transient,omitempty"`  // 元素是否临时
	CreateTime       int64  `json:"create_time,omitempty"`        // 创建时间
	UpdateTime       int64  `json:"update_time,omitempty"`        // 更新时间
}

// FailedElement 失败元素
type FailedElement struct {
	ElementKey string `json:"element_key,omitempty"` // 元素键
	ErrorCode  int    `json:"error_code,omitempty"`  // 错误码
	ErrorMsg   string `json:"error_msg,omitempty"`   // 错误信息
}
