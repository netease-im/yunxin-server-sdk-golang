package task

import . "github.com/netease-im/yunxin-server-sdk-golang/src/rtc/result"

// RtcTaskService 云端播放任务服务接口
type RtcTaskService interface {
	// CreateTask 创建云端播放任务
	CreateTask(request *TaskCreateRequest) (*RtcResult[*TaskCreateResponse], error)

	// UpdateTaskV2 V2版本更新云端播放任务
	UpdateTaskV2(request *TaskUpdateRequestV2) (*RtcResult[*TaskUpdateResponse], error)

	// UpdateTaskV3 V3版本更新云端播放任务
	UpdateTaskV3(request *TaskUpdateRequestV3) (*RtcResult[*TaskUpdateResponse], error)

	// QueryTaskV2 V2版本查询云端播放任务
	QueryTaskV2(request *TaskQueryRequestV2) (*RtcResult[*TaskQueryResponse], error)

	// QueryTaskV3 V3版本查询云端播放任务
	QueryTaskV3(request *TaskQueryRequestV3) (*RtcResult[*TaskQueryResponse], error)

	// PauseTaskV2 V2版本暂停云端播放任务
	PauseTaskV2(request *TaskPauseRequestV2) (*RtcResult[*TaskPauseResponse], error)

	// PauseTaskV3 V3版本暂停云端播放任务
	PauseTaskV3(request *TaskPauseRequestV3) (*RtcResult[*TaskPauseResponse], error)

	// ResumeTaskV2 V2版本恢复云端播放任务
	ResumeTaskV2(request *TaskResumeRequestV2) (*RtcResult[*TaskResumeResponse], error)

	// ResumeTaskV3 V3版本恢复云端播放任务
	ResumeTaskV3(request *TaskResumeRequestV3) (*RtcResult[*TaskResumeResponse], error)

	// DestroyTaskV2 V2版本销毁云端播放任务
	DestroyTaskV2(request *TaskDestroyRequestV2) (*RtcResult[*TaskDestroyResponse], error)

	// DestroyTaskV3 V3版本销毁云端播放任务
	DestroyTaskV3(request *TaskDestroyRequestV3) (*RtcResult[*TaskDestroyResponse], error)
}
