package cloud_record

// URL路径常量,用于所有云端录制相关的API请求
const (
	CreateCloudRecordV2  = "/v2/api/cloudrecord/submit"     // V2版本创建云端录制任务
	CreateCloudRecordV3  = "/v3/api/cloudrecord/submit"     // V3版本创建云端录制任务
	UpdateLayoutV2       = "/v2/api/cloudrecord/update"     // V2版本更新录制布局
	UpdateLayoutV3       = "/v3/api/cloudrecord/update"     // V3版本更新录制布局
	QueryTaskV2          = "/v2/api/cloudrecord/tasks"      // V2版本查询云端录制任务
	QueryTaskV3          = "/v3/api/cloudrecord/tasks"      // V3版本查询云端录制任务
	UpdateSubscriptionV2 = "/v2/api/cloudrecord/updateSubs" // V2版本更新订阅名单
	UpdateSubscriptionV3 = "/v3/api/cloudrecord/updateSubs" // V3版本更新订阅名单
	StopRecordV2         = "/v2/api/cloudrecord/stop"       // V2版本停止录制任务
	StopRecordV3         = "/v3/api/cloudrecord/stop"       // V3版本停止录制任务
)
