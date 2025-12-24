package event_subscribe

// URL路径常量，用于所有事件订阅相关的API请求
const (
	AddEventSubscribe         = "/event/subscribe/add.action"      // 添加事件订阅
	QueryEventSubscribe       = "/event/subscribe/query.action"    // 查询事件订阅
	DeleteEventSubscribe      = "/event/subscribe/delete.action"   // 删除特定事件订阅
	BatchDeleteEventSubscribe = "/event/subscribe/batchdel.action" // 批量删除所有事件订阅
)
