package friend

// Friend V1 Request Structures

// AddFriendRequestV1 AddFriend请求
type AddFriendRequestV1 struct {
	Accid    string `json:"accid,omitempty"`
	Faccid   string `json:"faccid,omitempty"`
	Type     int    `json:"type,omitempty"`
	Msg      string `json:"msg,omitempty"`
	Serverex string `json:"serverex,omitempty"`
}

// DeleteFriendRequestV1 DeleteFriend请求
type DeleteFriendRequestV1 struct {
	Accid       string `json:"accid,omitempty"`
	Faccid      string `json:"faccid,omitempty"`
	DeleteAlias bool   `json:"isDeleteAlias,omitempty"`
}

// GetFriendListRequestV1 GetFriendList请求
type GetFriendListRequestV1 struct {
	Accid      string `json:"accid,omitempty"`
	Updatetime int64  `json:"updatetime,omitempty"`
}

// GetFriendRelationshipRequestV1 GetFriendRelationship请求
type GetFriendRelationshipRequestV1 struct {
	Accid  string `json:"accid,omitempty"`
	Faccid string `json:"faccid,omitempty"`
}

// UpdateFriendRequestV1 UpdateFriend请求
type UpdateFriendRequestV1 struct {
	Accid    string `json:"accid,omitempty"`
	Faccid   string `json:"faccid,omitempty"`
	Alias    string `json:"alias,omitempty"`
	Ex       string `json:"ex,omitempty"`
	Serverex string `json:"serverex,omitempty"`
}
