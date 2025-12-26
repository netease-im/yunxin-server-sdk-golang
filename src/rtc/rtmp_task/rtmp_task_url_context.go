package rtmp_task

// URL路径常量,用于所有旁路推流任务相关的API请求
const (
	CreateRtmpTaskV2    = "/v2/api/rooms/{cid}/task"          // V2版本创建旁路推流任务
	CreateRtmpTaskV3    = "/v3/api/rooms/task"                // V3版本创建旁路推流任务
	UpdateRtmpTaskV2    = "/v2/api/rooms/{cid}/update/task"   // V2版本更新旁路推流任务
	UpdateRtmpTaskV3    = "/v3/api/rooms/update/task"         // V3版本更新旁路推流任务
	QueryRtmpTaskV2     = "/v2/api/rooms/{cid}/task/{taskId}" // V2版本查询指定旁路推流任务
	QueryRtmpTaskV3     = "/v3/api/rooms/task"                // V3版本查询指定旁路推流任务
	QueryAllRtmpTasksV2 = "/v2/api/rooms/{cid}/tasks"         // V2版本查询所有旁路推流任务
	QueryAllRtmpTasksV3 = "/v3/api/rooms/tasks"               // V3版本查询所有旁路推流任务
	DeleteRtmpTaskV2    = "/v2/api/rooms/{cid}/task"          // V2版本停止旁路推流任务
	DeleteRtmpTaskV3    = "/v3/api/rooms/task"                // V3版本停止旁路推流任务
)
