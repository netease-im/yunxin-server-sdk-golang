package system_notification

// FailedNotification 失败的通知详情
type FailedNotification struct {
	AccountId string `json:"account_id"` // 发送失败的账号ID
	ErrorMsg  string `json:"error_msg"`  // 错误消息
	ErrorCode int    `json:"error_code"` // 错误码
}

// SendCustomNotificationResponseV2 发送自定义系统通知响应
type SendCustomNotificationResponseV2 struct {
	// API只返回成功状态码(200)，无额外数据返回
}

// SendBatchCustomNotificationResponseV2 批量发送自定义系统通知响应
type SendBatchCustomNotificationResponseV2 struct {
	SuccessList []string             `json:"success_list"` // 成功发送的账号ID列表
	FailedList  []FailedNotification `json:"failed_list"`  // 失败的通知列表
}
