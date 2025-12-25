package livewallsolution

import (
	. "github.com/netease-im/yunxin-server-sdk-golang/src/rtc/result"
)

// LivewallSolutionService 安全通服务接口
type LivewallSolutionService interface {
	// CreateTask 创建安全通审核任务
	CreateTask(req *CreateTaskRequest) (*RtcResult[*CreateTaskResponse], error)

	// StopTask 停止安全通审核任务
	StopTask(req *StopTaskRequest) (*RtcResult[*StopTaskResponse], error)

	// QueryImage 查询审核视频截图
	QueryImage(req *QueryImageRequest) (*RtcResult[*QueryImageResponse], error)

	// QueryAudioTask 查询审核音频断句
	QueryAudioTask(req *QueryAudioTaskRequest) (*RtcResult[*QueryAudioTaskResponse], error)
}
