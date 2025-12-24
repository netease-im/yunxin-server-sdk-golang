package chatroom

// URL context for Chatroom V2 API endpoints
const (
	CreateChatroom          = "/im/v2/chatrooms"                                       // Create chatroom endpoint
	GetChatroomInfo         = "/im/v2/chatrooms/{room_id}"                             // Get chatroom info endpoint
	UpdateChatroomInfo      = "/im/v2/chatrooms/{room_id}"                             // Update chatroom info endpoint
	ToggleChatroomMute      = "/im/v2/chatrooms/{room_id}/actions/chat_banned"         // Toggle chatroom mute status endpoint
	GetChatroomAddress      = "/im/v2/chatrooms/{room_id}/actions/address"             // Get chatroom address endpoint
	UpdateChatroomStatus    = "/im/v2/chatrooms/{room_id}/actions/update_status"       // Update chatroom status endpoint (open/close)
	ToggleInOutNotification = "/im/v2/chatrooms/{room_id}/actions/in_out_notification" // Toggle in/out notification endpoint
	QueryOpenChatrooms      = "/im/v2/chatrooms/actions/opend_chatrooms"               // Query open chatrooms endpoint
	ListOnlineMembers       = "/im/v2/chatrooms/{room_id}/actions/list_online_members" // List online chatroom members endpoint
	ListFixedMembers        = "/im/v2/chatrooms/{room_id}/actions/list_members"        // List fixed chatroom members endpoint
)
