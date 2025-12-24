package users

// URL context for User V2 API endpoints
const (
	UpdateUser           = "/im/v2/users/{account_id}"          // Update user endpoint
	GetUser              = "/im/v2/users/{account_id}"          // Get user endpoint
	BatchGetUsers        = "/im/v2/users"                       // Batch get users endpoint
	GetUsersOnlineStatus = "/im/v2/users/actions/online_status" // Get users online status endpoint
)
