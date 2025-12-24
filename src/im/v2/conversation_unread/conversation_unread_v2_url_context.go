package conversation_unread

// URL context for Conversation Unread V2 API endpoints
const (
	OverviewsConversation   = "/im/v2/conversation_overviews/{account_id}"                               // Get conversation overview endpoint for an account
	ClearConversationUnread = "/im/v2/conversations/{conversation_id}/actions/clear_conversation_unread" // Clear conversation unread count endpoint
)
