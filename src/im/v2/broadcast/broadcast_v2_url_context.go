package broadcast

// URL context for broadcast notification API V2
const (
	BroadcastNotification         = "/im/v2/broadcast_notification"                    // URL for sending broadcast notification
	DeleteBroadcastNotification   = "/im/v2/broadcast_notification/{broadcast_id}"     // URL for deleting broadcast notification
	QueryBroadcastNotification    = "/im/v2/broadcast_notification/{broadcast_id}"     // URL for querying broadcast notification
	ChatroomBroadcastNotification = "/im/v2.1/broadcast_notification/actions/chatroom" // URL for sending chatroom broadcast notification
)
