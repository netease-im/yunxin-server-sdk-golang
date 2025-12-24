package meeting

import (
	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/meeting/user"
)

// YunxinMeetingApiServices 云信会议API服务
type YunxinMeetingApiServices struct {
	meetingUserService user.MeetingUserService
}

// NewYunxinMeetingApiServices 创建云信会议API服务
func NewYunxinMeetingApiServices(yunxinApiHttpClient core.YunxinApiHttpClient) *YunxinMeetingApiServices {
	return &YunxinMeetingApiServices{
		meetingUserService: user.NewMeetingUserService(yunxinApiHttpClient),
	}
}

// GetMeetingUserService 获取会议用户服务
func (s *YunxinMeetingApiServices) GetMeetingUserService() user.MeetingUserService {
	return s.meetingUserService
}
