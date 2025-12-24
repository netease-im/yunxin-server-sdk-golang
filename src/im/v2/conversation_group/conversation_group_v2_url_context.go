package conversation_group

// URL context for Conversation Group V2 API endpoints
const (
	CreateConversationGroup    = "/im/v2/conversation_groups"                   // Create conversation group endpoint
	UpdateConversationGroup    = "/im/v2/conversation_groups/{group_id}"        // Update conversation group endpoint
	DeleteConversationGroup    = "/im/v2/conversation_groups/{group_id}"        // Delete conversation group endpoint
	GetConversationGroup       = "/im/v2/conversation_groups/{group_id}"        // Get conversation group endpoint
	BatchGetConversationGroups = "/im/v2/conversation_groups/actions/group_ids" // Batch get conversation groups endpoint
	ListAllConversationGroups  = "/im/v2/conversation_groups"                   // List all conversation groups endpoint
)
