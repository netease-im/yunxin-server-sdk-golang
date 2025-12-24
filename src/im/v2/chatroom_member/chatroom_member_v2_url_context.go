package chatroom_member

// URL context for Chatroom Member V2 API endpoints
const (
	SetMemberRole              = "/im/v2/room_members/{account_id}"                           // Set chatroom member role endpoint
	UpdateOnlineMemberInfo     = "/im/v2/room_members/{account_id}"                           // Update chatroom online member information endpoint
	ToggleChatBan              = "/im/v2/room_members/{account_id}/actions/chat_banned"       // Ban/unban chatroom member from chatting endpoint
	ModifyMemberTags           = "/im/v2/room_members/{account_id}/actions/tags"              // Modify chatroom online member tags endpoint
	QueryTaggedMembersCount    = "/im/v2/room_members/{room_id}/actions/tagged_members_count" // Query tagged members count endpoint
	ListTagMembers             = "/im/v2/room_members/{room_id}/actions/list_tag_members"     // List tagged members with pagination endpoint
	ToggleTempChatBan          = "/im/v2/room_members/{account_id}/actions/temp_chat_banned"  // Temporarily ban/unban chatroom member from chatting endpoint
	QueryChatroomBanList       = "/im/v2/room_members/{room_id}/actions/banned_members"       // Query chatroom ban list endpoint
	ToggleBlocked              = "/im/v2/room_members/{account_id}/actions/blocked"           // Block/unblock chatroom member endpoint
	QueryChatroomBlacklist     = "/im/v2.1/room_members/{room_id}/actions/blocked"            // Query chatroom blacklist endpoint
	ToggleTaggedMembersChatBan = "/im/v2/room_members/actions/chat_banned_tagged_members"     // Toggle chatroom tagged members chat ban endpoint
	BatchQueryChatroomMembers  = "/im/v2/room_members/{room_id}/actions/batch"                // Batch query chatroom fixed members information endpoint
	AddVirtualMembers          = "/im/v2/room_members/actions/virtual_members"                // Add virtual members to chatroom endpoint
	DeleteVirtualMembers       = "/im/v2/room_members/actions/virtual_members"                // Delete virtual members from chatroom endpoint
	ClearVirtualMembers        = "/im/v2/room_members/actions/clear_virtual_members"          // Clear all virtual members from chatroom endpoint
	QueryVirtualMembers        = "/im/v2.1/room_members/actions/virtual_members"              // Query virtual members in chatroom endpoint
	QueryChatBanned            = "/im/v2.1/room_members/{room_id}/actions/chat_banned"        // Query chat banned members endpoint
)
