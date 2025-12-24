package cloud_record

import . "github.com/netease-im/yunxin-server-sdk-golang/src/rtc/result"

// CloudRecordService 云端录制服务接口
type CloudRecordService interface {
	// CreateCloudRecordV2 V2版本创建云端录制任务
	CreateCloudRecordV2(request *CloudRecordCreateRequestV2) (*RtcResult[*CloudRecordCreateResponse], error)

	// CreateCloudRecordV3 V3版本创建云端录制任务
	CreateCloudRecordV3(request *CloudRecordCreateRequestV3) (*RtcResult[*CloudRecordCreateResponse], error)

	// UpdateLayoutV2 V2版本更新录制布局
	UpdateLayoutV2(request *CloudRecordUpdateLayoutRequestV2) (*RtcResult[*CloudRecordUpdateLayoutResponse], error)

	// UpdateLayoutV3 V3版本更新录制布局
	UpdateLayoutV3(request *CloudRecordUpdateLayoutRequestV3) (*RtcResult[*CloudRecordUpdateLayoutResponse], error)

	// QueryTaskV2 V2版本查询云端录制任务
	QueryTaskV2(request *CloudRecordQueryTaskRequestV2) (*RtcResult[*CloudRecordQueryTaskResponse], error)

	// QueryTaskV3 V3版本查询云端录制任务
	QueryTaskV3(request *CloudRecordQueryTaskRequestV3) (*RtcResult[*CloudRecordQueryTaskResponse], error)

	// UpdateSubscriptionV2 V2版本更新订阅名单
	UpdateSubscriptionV2(request *CloudRecordUpdateSubscriptionRequestV2) (*RtcResult[*CloudRecordUpdateSubscriptionResponse], error)

	// UpdateSubscriptionV3 V3版本更新订阅名单
	UpdateSubscriptionV3(request *CloudRecordUpdateSubscriptionRequestV3) (*RtcResult[*CloudRecordUpdateSubscriptionResponse], error)

	// StopRecordV2 V2版本停止录制任务
	StopRecordV2(request *CloudRecordStopRequestV2) (*RtcResult[*CloudRecordStopResponse], error)

	// StopRecordV3 V3版本停止录制任务
	StopRecordV3(request *CloudRecordStopRequestV3) (*RtcResult[*CloudRecordStopResponse], error)
}
