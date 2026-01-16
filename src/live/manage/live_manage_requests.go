package manage

// LiveCreateChannelRequest 创建直播频道请求
type LiveCreateChannelRequest struct {
	Name string `json:"name"`
	Type int    `json:"type"`
}
