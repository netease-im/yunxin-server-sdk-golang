package account

// URL context for Account V2 API endpoints
const (
	Accounts       = "/im/v2/accounts"                                    // Accounts endpoint
	AccountWithID  = "/im/v2/accounts/{account_id}"                       // Account with ID endpoint template
	SetPushConfig  = "/im/v2/accounts/{account_id}/actions/push_config"   // Set push config endpoint template
	DisableAccount = "/im/v2/accounts/{account_id}/actions/disable"       // Disable account endpoint template
	KickAccount    = "/im/v2/accounts/{account_id}/actions/kick"          // Kick account endpoint template
	RefreshToken   = "/im/v2/accounts/{account_id}/actions/refresh_token" // Refresh token URL pattern
)
