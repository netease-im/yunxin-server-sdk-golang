package manage

// LiveCreateChannelRequest 创建直播频道请求
type LiveCreateChannelRequest struct {
	Name        string `json:"name"`                  // 频道名称
	Type        int    `json:"type"`                  // 频道类型，0:rtmp,1:hls
	Quality     int    `json:"quality,omitempty"`     // 画质等级，可选
	Record      bool   `json:"record,omitempty"`      // 是否录制，可选
	Format      string `json:"format,omitempty"`      // 录制格式，可选
	Duration    int    `json:"duration,omitempty"`    // 录制时长，可选，单位分钟
	Filename    string `json:"filename,omitempty"`    // 录制文件名，可选
	CallbackUrl string `json:"callbackUrl,omitempty"` // 回调地址，可选
}
