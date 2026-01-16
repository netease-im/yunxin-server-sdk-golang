package user

// CreateMeetingAccountRequest 创建会议账号请求
type CreateMeetingAccountRequest struct {
	UserUuid                string   `json:"userUuid,omitempty"`
	ImToken                 string   `json:"imToken,omitempty"`
	Name                    string   `json:"name,omitempty"`
	ShortMeetingNum         string   `json:"shortMeetingNum,omitempty"`
	SipCid                  string   `json:"sipCid,omitempty"`
	Avatar                  string   `json:"avatar,omitempty"`
	PhoneNumber             string   `json:"phoneNumber,omitempty"`
	Email                   string   `json:"email,omitempty"`
	ReturnGeneratedPassword bool     `json:"returnGeneratedPassword,omitempty"`
	Departments             []string `json:"departments,omitempty"`
}
