package manage

// LiveCreateChannelRequest 创建直播频道请求
type LiveCreateChannelRequest struct {
	Name string `json:"name,omitempty"`
	Type int    `json:"type,omitempty"`
}
