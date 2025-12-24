package subscription

// URL context for user status subscription API V2
const (
	SubscribeUserStatus         = "/im/v2/subscription/{account_id}" // URL template for subscribing to user status events, {account_id} should be replaced with the subscriber's account ID
	UnsubscribeUserStatus       = "/im/v2/subscription/{account_id}" // URL template for unsubscribing from user status events, {account_id} should be replaced with the subscriber's account ID, Note: This uses the same URL as subscribe but with DELETE method
	QueryUserStatusSubscription = "/im/v2/subscription/{account_id}" // URL template for querying user status subscriptions, {account_id} should be replaced with the subscriber's account ID, Note: This uses the same URL as subscribe/unsubscribe but with PATCH method
)
