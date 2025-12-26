package rtmp_task

import . "github.com/netease-im/yunxin-server-sdk-golang/src/rtc/result"

// RtmpTaskService 旁路推流任务服务接口
type RtmpTaskService interface {
	// CreateRtmpTaskV2 V2版本创建旁路推流任务
	CreateRtmpTaskV2(request *RtmpTaskCreateRequestV2) (*RtcResult[*RtmpTaskCreateResponse], error)

	// CreateRtmpTaskV3 V3版本创建旁路推流任务
	CreateRtmpTaskV3(request *RtmpTaskCreateRequestV3) (*RtcResult[*RtmpTaskCreateResponse], error)

	// UpdateRtmpTaskV2 V2版本更新旁路推流任务
	UpdateRtmpTaskV2(request *RtmpTaskUpdateRequestV2) (*RtcResult[*RtmpTaskUpdateResponse], error)

	// UpdateRtmpTaskV3 V3版本更新旁路推流任务
	UpdateRtmpTaskV3(request *RtmpTaskUpdateRequestV3) (*RtcResult[*RtmpTaskUpdateResponse], error)

	// QueryRtmpTaskV2 V2版本查询指定旁路推流任务
	QueryRtmpTaskV2(request *RtmpTaskQueryRequestV2) (*RtcResult[*RtmpTaskQueryResponse], error)

	// QueryRtmpTaskV3 V3版本查询指定旁路推流任务
	QueryRtmpTaskV3(request *RtmpTaskQueryRequestV3) (*RtcResult[*RtmpTaskQueryResponse], error)

	// QueryAllRtmpTasksV2 V2版本查询所有旁路推流任务
	QueryAllRtmpTasksV2(request *RtmpTaskQueryAllRequestV2) (*RtcResult[*RtmpTaskQueryAllResponse], error)

	// QueryAllRtmpTasksV3 V3版本查询所有旁路推流任务
	QueryAllRtmpTasksV3(request *RtmpTaskQueryAllRequestV3) (*RtcResult[*RtmpTaskQueryAllResponse], error)

	// DeleteRtmpTaskV2 V2版本停止旁路推流任务
	DeleteRtmpTaskV2(request *RtmpTaskDeleteRequestV2) (*RtcResult[*RtmpTaskDeleteResponse], error)

	// DeleteRtmpTaskV3 V3版本停止旁路推流任务
	DeleteRtmpTaskV3(request *RtmpTaskDeleteRequestV3) (*RtcResult[*RtmpTaskDeleteResponse], error)
}
