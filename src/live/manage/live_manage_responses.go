package manage

// LiveCreateChannelResponse 创建直播频道响应
type LiveCreateChannelResponse struct {
	Cid         string `json:"cid"`         // 频道唯一标识
	Name        string `json:"name"`        // 频道名称
	PushUrl     string `json:"pushUrl"`     // 推流地址
	RtmpPullUrl string `json:"rtmpPullUrl"` // rtmp拉流地址
	RtsPullUrl  string `json:"rtsPullUrl"`  // rts拉流地址
	HttpPullUrl string `json:"httpPullUrl"` // HTTP拉流地址
	HlsPullUrl  string `json:"hlsPullUrl"`  // HLS拉流地址
	Ctime       int64  `json:"ctime"`       // 创建时间戳
}
