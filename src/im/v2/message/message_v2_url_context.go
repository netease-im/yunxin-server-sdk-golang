package message

// URL context for Message V2 API endpoints
const (
	SendMessage                = "/im/v2/conversations/{conversation_id}/messages"                        // Send message endpoint
	BatchSendP2PMessage        = "/im/v2/conversations/messages"                                          // Batch send P2P messages endpoint
	ModifyMessage              = "/im/v2/messages/actions/modifyMsg"                                      // Modify (update) message endpoint
	WithdrawMessage            = "/im/v2/conversations/{conversation_id}/messages/{message_server_id}"    // Withdraw (recall or delete) message endpoint
	DeleteConversationMessages = "/im/v2/conversations/{conversation_id}/messages"                        // Delete all messages in a conversation endpoint
	SendP2PReadReceipt         = "/im/v2/messages/actions/p2p_read_receipt"                               // Send P2P read receipt endpoint
	SendTeamReadReceipt        = "/im/v2/messages/actions/team_read_receipt"                              // Send team read receipt endpoint
	QueryTeamReadReceipt       = "/im/v2/messages/actions/team_read_receipt"                              // Query team message read receipt details endpoint
	StreamMessage              = "/im/v2/conversations/{conversation_id}/messages/actions/stream_message" // Send streaming message endpoint
	QueryMessage               = "/im/v2.1/conversations/{conversation_id}/messages/{message_server_id}"  // Query single message details endpoint
	SearchMessages             = "/im/v2.1/messages/actions/search_messages"                              // Search messages endpoint
	QueryConversationMessages  = "/im/v2.1/conversations/{conversation_id}/messages"                      // Query conversation messages with pagination endpoint
	BatchQueryMessages         = "/im/v2.1/conversations/{conversation_id}/batch_messages"                // Batch query messages by message IDs endpoint
	QueryThreadMessages        = "/im/v2.1/messages/actions/thread_messages"                              // Query thread messages endpoint
	AddQuickComment            = "/im/v2/messages/actions/quick_comment"                                  // Add quick comment to message endpoint
	DeleteQuickComment         = "/im/v2/messages/actions/quick_comment"                                  // Delete quick comment from message endpoint
	BatchQueryQuickComments    = "/im/v2/messages/actions/quick_comment"                                  // Batch query quick comments endpoint
)
