package friend

// URL context for Friend V2 API endpoints
const (
	Friends              = "/im/v2.1/friends"                                // Base friend endpoint
	DeleteFriend         = "/im/v2/friends/{account_id}"                     // Delete friend endpoint
	UpdateFriend         = "/im/v2.1/friends/{account_id}"                   // Update friend endpoint
	GetFriend            = "/im/v2.1/friends/{account_id}"                   // Get friend endpoint
	ListFriends          = "/im/v2.1/friends"                                // List friends endpoint
	HandleFriendAddition = "/im/v2.1/friends/actions/handle_friend_addition" // Handle friend addition endpoint (accept/reject)
)
