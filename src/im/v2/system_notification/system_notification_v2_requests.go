package system_notification

// NotificationConfig 通知配置
type NotificationConfig struct {
	OfflineEnabled bool `json:"offline_enabled"` // 通知是否存离线，默认true
	UnreadEnabled  bool `json:"unread_enabled"`  // 通知是否计入未读数，默认true
}

// PushConfig 推送配置
type PushConfig struct {
	PushEnabled          bool     `json:"push_enabled"`                     // 是否启用APNs推送或Android系统通知栏推送，默认true
	PushNickEnabled      bool     `json:"push_nick_enabled"`                // 推送内容是否包含昵称，默认true
	PushContent          string   `json:"push_content,omitempty"`           // 推送内容文本，最长500字符
	PushPayload          string   `json:"push_payload,omitempty"`           // 推送payload，JSON格式，最长2048字符
	PushForcepushEnable  bool     `json:"push_forcepush_enable"`            // 是否启用强制推送(@操作)，默认false
	PushForcepushAll     bool     `json:"push_forcepush_all"`               // 是否强制推送给群里所有成员，默认false
	PushForcepushIds     []string `json:"push_forcepush_ids,omitempty"`     // 强制推送的账号ID列表，最多100个
	PushForcepushContent string   `json:"push_forcepush_content,omitempty"` // 强制推送内容文本，最长500字符
}

// RouteConfig 路由配置
type RouteConfig struct {
	RouteEnabled     bool   `json:"route_enabled"`               // 是否启用消息转发到应用服务器，默认true
	RouteEnvironment string `json:"route_environment,omitempty"` // 消息转发的环境名称
}

// SendCustomNotificationRequestV2 发送自定义系统通知请求
type SendCustomNotificationRequestV2 struct {
	SenderId           string              `json:"sender_id"`                     // 发送者IM账号ID (必填)
	Type               int                 `json:"type"`                          // 自定义通知类型: 1=P2P, 2=群(高级群), 3=超大群 (必填)
	ReceiverId         string              `json:"receiver_id"`                   // 接收者ID，type=1时为账号ID，type=2或3时为群ID (必填)
	Content            string              `json:"content"`                       // 自定义通知内容，JSON字符串，最长4096字符 (必填)
	Sound              string              `json:"sound,omitempty"`               // 客户端本地声音文件名，最长30字符 (可选)
	NotificationConfig *NotificationConfig `json:"notification_config,omitempty"` // 通知配置 (可选)
	PushConfig         *PushConfig         `json:"push_config,omitempty"`         // 推送配置 (可选)
	RouteConfig        *RouteConfig        `json:"route_config,omitempty"`        // 路由配置 (可选)
}

// SendBatchCustomNotificationRequestV2 批量发送自定义系统通知请求
type SendBatchCustomNotificationRequestV2 struct {
	SenderId           string              `json:"sender_id"`                     // 发送者IM账号ID (必填)
	ReceiverIds        []string            `json:"receiver_ids"`                  // 接收者账号ID列表，最多500个 (必填)
	Content            string              `json:"content"`                       // 自定义通知内容，JSON字符串，最长4096字符 (必填)
	Sound              string              `json:"sound,omitempty"`               // 客户端本地声音文件名，最长30字符 (可选)
	NotificationConfig *NotificationConfig `json:"notification_config,omitempty"` // 通知配置 (可选)
	PushConfig         *PushConfig         `json:"push_config,omitempty"`         // 推送配置 (可选)
	RouteConfig        *RouteConfig        `json:"route_config,omitempty"`        // 路由配置 (可选)
}
