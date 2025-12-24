package chatroom_message

// URL context for Chatroom Message V2 API endpoints
const (
	SendChatroomMessage           = "/im/v2/chatrooms/{room_id}/messages"                     // Send chatroom message endpoint
	BatchSendChatroomMessages     = "/im/v2/chatrooms/messages"                               // Batch send chatroom messages endpoint
	RecallOrDeleteChatroomMessage = "/im/v2/chatrooms/{room_id}/messages/{message_client_id}" // Recall or delete chatroom message endpoint
	QueryChatroomHistoryMessages  = "/im/v2.1/chatrooms/{room_id}/messages"                   // Query chatroom history messages endpoint
)
