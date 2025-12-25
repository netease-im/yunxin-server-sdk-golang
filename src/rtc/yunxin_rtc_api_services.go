package rtc

import (
	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/rtc/cloud_record"
	"github.com/netease-im/yunxin-server-sdk-golang/src/rtc/room"
	"github.com/netease-im/yunxin-server-sdk-golang/src/rtc/task"
)

// YunxinRtcApiServices 云信RTC API服务
type YunxinRtcApiServices struct {
	rtcRoomService        room.RtcRoomService
	rtcCloudRecordService cloud_record.CloudRecordService
	rtcTaskService        task.RtcTaskService
}

// NewYunxinRtcApiServices 创建云信RTC API服务实例
func NewYunxinRtcApiServices(yunxinApiHttpClient core.YunxinApiHttpClient) *YunxinRtcApiServices {
	return &YunxinRtcApiServices{
		rtcRoomService:        room.NewRtcRoomService(yunxinApiHttpClient),
		rtcCloudRecordService: cloud_record.NewCloudRecordService(yunxinApiHttpClient),
		rtcTaskService:        task.NewRtcTaskService(yunxinApiHttpClient),
	}
}

// GetRtcRoomService 获取RTC房间服务
func (s *YunxinRtcApiServices) GetRtcRoomService() room.RtcRoomService {
	return s.rtcRoomService
}

// GetRtcCloudRecordService 获取RTC云端录制服务
func (s *YunxinRtcApiServices) GetRtcCloudRecordService() cloud_record.CloudRecordService {
	return s.rtcCloudRecordService
}

// GetRtcTaskService 获取RTC云端播放服务
func (s *YunxinRtcApiServices) GetRtcTaskService() task.RtcTaskService {
	return s.rtcTaskService
}
