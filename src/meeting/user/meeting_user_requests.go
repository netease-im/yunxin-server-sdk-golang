package user

// CreateMeetingAccountRequest 创建会议账号请求
type CreateMeetingAccountRequest struct {
	UserUuid                string   `json:"userUuid"`
	ImToken                 string   `json:"imToken"`
	Name                    string   `json:"name"`
	ShortMeetingNum         string   `json:"shortMeetingNum"`
	SipCid                  string   `json:"sipCid"`
	Avatar                  string   `json:"avatar"`
	PhoneNumber             string   `json:"phoneNumber"`
	Email                   string   `json:"email"`
	ReturnGeneratedPassword bool     `json:"returnGeneratedPassword"`
	Departments             []string `json:"departments"`
}
