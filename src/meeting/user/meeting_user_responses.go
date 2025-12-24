package user

// CreateMeetingAccountResponse 创建会议账号响应
type CreateMeetingAccountResponse struct {
	UserUuid          string   `json:"userUuid"`
	UserToken         string   `json:"userToken"`
	Name              string   `json:"name"`
	PrivateMeetingNum string   `json:"privateMeetingNum"`
	ShortMeetingNum   string   `json:"shortMeetingNum"`
	SipCid            string   `json:"sipCid"`
	Avatar            string   `json:"avatar"`
	PhoneNumber       string   `json:"phoneNumber"`
	Email             string   `json:"email"`
	Password          string   `json:"password"`
	State             int      `json:"state"`
	Departments       []string `json:"departments"`
}
