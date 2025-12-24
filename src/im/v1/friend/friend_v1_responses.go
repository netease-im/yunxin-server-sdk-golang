package friend

// Friend V1 Response Structures

// AddFriendResponseV1 AddFriend响应
type AddFriendResponseV1 struct {
}

// DeleteFriendResponseV1 DeleteFriend响应
type DeleteFriendResponseV1 struct {
}

// GetFriendListResponseV1 GetFriendList响应
type GetFriendListResponseV1 struct {
	Friends []FriendInfo `json:"friends,omitempty"`
}

// FriendInfo FriendInfo
type FriendInfo struct {
	Createtime  int64  `json:"createtime,omitempty"`
	Bidirection bool   `json:"bidirection,omitempty"`
	Faccid      string `json:"faccid,omitempty"`
	Alias       string `json:"alias,omitempty"`
}

// GetFriendRelationshipResponseV1 GetFriendRelationship响应
type GetFriendRelationshipResponseV1 struct {
	Createtime  int64  `json:"createtime,omitempty"`
	Ex          string `json:"ex,omitempty"`
	Bidirection bool   `json:"bidirection,omitempty"`
	Faccid      string `json:"faccid,omitempty"`
	Serverex    string `json:"serverex,omitempty"`
	Updatetime  int64  `json:"updatetime,omitempty"`
}

// UpdateFriendResponseV1 UpdateFriend响应
type UpdateFriendResponseV1 struct {
}
