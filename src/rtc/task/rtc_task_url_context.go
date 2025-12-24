package task

// URL路径常量,用于所有云端播放任务相关的API请求
const (
	CreateTaskV2  = "/v2/api/task/create" // V2版本创建云端播放任务
	UpdateTaskV2  = "/v2/api/task/update" // V2版本更新云端播放任务
	UpdateTaskV3  = "/v3/api/task/update" // V3版本更新云端播放任务
	QueryTaskV2   = "/v2/api/task/list"   // V2版本查询云端播放任务
	QueryTaskV3   = "/v3/api/task/list"   // V3版本查询云端播放任务
	PauseTaskV2   = "/v2/api/task/pause"  // V2版本暂停云端播放任务
	PauseTaskV3   = "/v3/api/task/pause"  // V3版本暂停云端播放任务
	ResumeTaskV2  = "/v2/api/task/resume" // V2版本恢复云端播放任务
	ResumeTaskV3  = "/v3/api/task/resume" // V3版本恢复云端播放任务
	DestroyTaskV2 = "/v2/api/task/close"  // V2版本销毁云端播放任务
	DestroyTaskV3 = "/v3/api/task/close"  // V3版本销毁云端播放任务
)
