package chatroom_queue

// URL context for Chatroom Queue V2 API endpoints
const (
	InitializeChatroomQueue    = "/im/v2/room_queues/{room_id}"               // Initialize chatroom queue endpoint
	QueryChatroomQueueElements = "/im/v2/room_queues/{room_id}/actions/query" // Query chatroom queue elements endpoint
	UpdateChatroomQueue        = "/im/v2/room_queues/{room_id}"               // Update chatroom queue endpoint
	DeleteChatroomQueue        = "/im/v2/room_queues/{room_id}"               // Delete chatroom queue endpoint
	PollChatroomQueueElement   = "/im/v2/room_queues/{room_id}/actions/poll"  // Poll element from chatroom queue endpoint
)
