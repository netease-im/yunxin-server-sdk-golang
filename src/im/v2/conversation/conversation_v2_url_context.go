package conversation

// URL context for Conversation V2 API endpoints
const (
	CreateConversation       = "/im/v2.1/conversations"                                                // Create conversation endpoint
	UpdateConversation       = "/im/v2/conversations/{conversation_id}"                                // Update conversation endpoint
	DeleteConversation       = "/im/v2/conversations/{conversation_id}"                                // Delete conversation endpoint
	BatchDeleteConversations = "/im/v2/conversations/actions/conversation_ids"                         // Batch delete conversations endpoint
	GetConversation          = "/im/v2.1/conversations/{conversation_id}"                              // Get conversation endpoint
	BatchGetConversations    = "/im/v2.1/conversations/actions/conversation_ids"                       // Batch get conversations endpoint
	ListConversations        = "/im/v2.1/conversations"                                                // List conversations endpoint (paginated)
	StickTopConversation     = "/im/v2/conversations/{conversation_id}/actions/stick_top_conversation" // Stick top conversation endpoint
)
