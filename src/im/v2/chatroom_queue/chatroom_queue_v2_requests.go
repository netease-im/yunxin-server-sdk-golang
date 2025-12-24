package chatroom_queue

// InitializeChatroomQueueRequestV2 初始化聊天室队列请求
type InitializeChatroomQueueRequestV2 struct {
	RoomId         int64        `json:"room_id"`                    // 聊天室ID
	QueueSizeLimit int          `json:"queue_size_limit,omitempty"` // 队列大小限制
	RouteConfig    *RouteConfig `json:"route_config,omitempty"`     // 路由配置
}

// QueryChatroomQueueElementsRequestV2 查询聊天室队列元素请求
type QueryChatroomQueueElementsRequestV2 struct {
	RoomId         int64        `json:"-"`                      // 聊天室ID，用于构造URL
	ElementKeyList []string     `json:"element_key_list"`       // 元素键列表
	RouteConfig    *RouteConfig `json:"route_config,omitempty"` // 路由配置
}

// UpdateChatroomQueueRequestV2 更新聊天室队列请求
type UpdateChatroomQueueRequestV2 struct {
	RoomId             int64               `json:"-"`                             // 聊天室ID，用于构造URL
	OperatorAccountId  string              `json:"operator_account_id"`           // 操作者账号ID
	ElementList        []QueueElement      `json:"element_list,omitempty"`        // 元素列表
	NotificationConfig *NotificationConfig `json:"notification_config,omitempty"` // 通知配置
	RouteConfig        *RouteConfig        `json:"route_config,omitempty"`        // 路由配置
}

// DeleteChatroomQueueRequestV2 删除聊天室队列请求
type DeleteChatroomQueueRequestV2 struct {
	RoomId             int64               `json:"-"`                             // 聊天室ID，用于构造URL
	NotificationConfig *NotificationConfig `json:"notification_config,omitempty"` // 通知配置
	RouteConfig        *RouteConfig        `json:"route_config,omitempty"`        // 路由配置
}

// PollChatroomQueueElementRequestV2 轮询聊天室队列元素请求
type PollChatroomQueueElementRequestV2 struct {
	RoomId             int64               `json:"-"`                             // 聊天室ID，用于构造URL
	ElementKey         string              `json:"element_key"`                   // 元素键
	NotificationConfig *NotificationConfig `json:"notification_config,omitempty"` // 通知配置
	RouteConfig        *RouteConfig        `json:"route_config,omitempty"`        // 路由配置
}

// QueueElement 队列元素
type QueueElement struct {
	ElementKey       string `json:"element_key"`                  // 元素键
	ElementValue     string `json:"element_value"`                // 元素值
	ElementAccountId string `json:"element_account_id,omitempty"` // 元素账号ID
	ElementTransient bool   `json:"element_transient,omitempty"`  // 元素是否临时
	ElementAddPolicy int    `json:"element_add_policy,omitempty"` // 元素添加策略
}

// NotificationConfig 通知配置
type NotificationConfig struct {
	NotificationEnabled   bool   `json:"notification_enabled,omitempty"`   // 是否启用通知
	NotificationExtension string `json:"notification_extension,omitempty"` // 通知扩展字段
	HighPriority          int    `json:"high_priority,omitempty"`          // 高优先级
	HighPriorityPolicy    int    `json:"high_priority_policy,omitempty"`   // 高优先级策略
}

// RouteConfig 路由配置
type RouteConfig struct {
	RouteEnabled     bool   `json:"route_enabled,omitempty"`     // 是否启用路由
	RouteEnvironment string `json:"route_environment,omitempty"` // 路由环境
}
