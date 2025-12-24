package upload

// VodUploadInitRequest VOD上传初始化请求
type VodUploadInitRequest struct {
	OriginFileName string `json:"originFileName"`        // 原始文件名
	UserDefInfo    string `json:"userDefInfo,omitempty"` // 用户自定义信息
	TypeId         int    `json:"typeId,omitempty"`      // 分类ID
	PresetId       int    `json:"presetId,omitempty"`    // 转码模板ID
	CallbackUrl    string `json:"callbackUrl,omitempty"` // 回调地址
	Description    string `json:"description,omitempty"` // 视频描述
}
