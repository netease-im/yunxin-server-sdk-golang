package upload

// VodUploadInitResponse VOD上传初始化响应
type VodUploadInitResponse struct {
	XNosToken string `json:"xNosToken"`
	Bucket    string `json:"bucket"`
	Object    string `json:"object"`
}
