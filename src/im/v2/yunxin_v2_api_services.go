package v2

import (
	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/im/v2/account"
	"github.com/netease-im/yunxin-server-sdk-golang/src/im/v2/ai"
	"github.com/netease-im/yunxin-server-sdk-golang/src/im/v2/block"
	"github.com/netease-im/yunxin-server-sdk-golang/src/im/v2/broadcast"
	"github.com/netease-im/yunxin-server-sdk-golang/src/im/v2/chatroom"
	"github.com/netease-im/yunxin-server-sdk-golang/src/im/v2/chatroom_member"
	"github.com/netease-im/yunxin-server-sdk-golang/src/im/v2/chatroom_message"
	"github.com/netease-im/yunxin-server-sdk-golang/src/im/v2/chatroom_queue"
	"github.com/netease-im/yunxin-server-sdk-golang/src/im/v2/conversation"
	"github.com/netease-im/yunxin-server-sdk-golang/src/im/v2/conversation_group"
	"github.com/netease-im/yunxin-server-sdk-golang/src/im/v2/conversation_unread"
	"github.com/netease-im/yunxin-server-sdk-golang/src/im/v2/friend"
	"github.com/netease-im/yunxin-server-sdk-golang/src/im/v2/message"
	"github.com/netease-im/yunxin-server-sdk-golang/src/im/v2/mute"
	"github.com/netease-im/yunxin-server-sdk-golang/src/im/v2/signal"
	"github.com/netease-im/yunxin-server-sdk-golang/src/im/v2/subscription"
	"github.com/netease-im/yunxin-server-sdk-golang/src/im/v2/system_notification"
	"github.com/netease-im/yunxin-server-sdk-golang/src/im/v2/team"
	"github.com/netease-im/yunxin-server-sdk-golang/src/im/v2/team_member"
	"github.com/netease-im/yunxin-server-sdk-golang/src/im/v2/tool"
	"github.com/netease-im/yunxin-server-sdk-golang/src/im/v2/users"
)

// YunxinV2ApiServices 云信V2 API服务聚合器
type YunxinV2ApiServices struct {
	accountService            account.AccountV2Service
	userService               users.UserV2Service
	friendService             friend.FriendV2Service
	teamService               team.TeamV2Service
	teamMemberService         team_member.TeamMemberV2Service
	chatroomService           chatroom.ChatroomV2Service
	chatroomMemberService     chatroom_member.ChatroomMemberV2Service
	chatroomMessageService    chatroom_message.ChatroomMessageV2Service
	chatroomQueueService      chatroom_queue.ChatroomQueueV2Service
	messageService            message.MessageV2Service
	conversationService       conversation.ConversationV2Service
	conversationGroupService  conversation_group.ConversationGroupV2Service
	conversationUnreadService conversation_unread.ConversationUnreadV2Service
	muteService               mute.MuteV2Service
	blockService              block.BlockV2Service
	broadcastService          broadcast.BroadcastV2Service
	systemNotificationService system_notification.CustomNotificationV2Service
	subscriptionService       subscription.SubscriptionV2Service
	aiService                 ai.AiV2Service
	toolService               tool.ToolV2Service
	signalService             signal.SignalV2Service
}

// NewYunxinV2ApiServices 创建云信V2 API服务聚合器
func NewYunxinV2ApiServices(yunxinApiHttpClient core.YunxinApiHttpClient) *YunxinV2ApiServices {
	return &YunxinV2ApiServices{
		accountService:            account.NewAccountV2Service(yunxinApiHttpClient),
		userService:               users.NewUserV2Service(yunxinApiHttpClient),
		friendService:             friend.NewFriendV2Service(yunxinApiHttpClient),
		teamService:               team.NewTeamV2Service(yunxinApiHttpClient),
		teamMemberService:         team_member.NewTeamMemberV2Service(yunxinApiHttpClient),
		chatroomService:           chatroom.NewChatroomV2Service(yunxinApiHttpClient),
		chatroomMemberService:     chatroom_member.NewChatroomMemberV2Service(yunxinApiHttpClient),
		chatroomMessageService:    chatroom_message.NewChatroomMessageV2Service(yunxinApiHttpClient),
		chatroomQueueService:      chatroom_queue.NewChatroomQueueV2Service(yunxinApiHttpClient),
		messageService:            message.NewMessageV2Service(yunxinApiHttpClient),
		conversationService:       conversation.NewConversationV2Service(yunxinApiHttpClient),
		conversationGroupService:  conversation_group.NewConversationGroupV2Service(yunxinApiHttpClient),
		conversationUnreadService: conversation_unread.NewConversationUnreadV2Service(yunxinApiHttpClient),
		muteService:               mute.NewMuteV2Service(yunxinApiHttpClient),
		blockService:              block.NewBlockV2Service(yunxinApiHttpClient),
		broadcastService:          broadcast.NewBroadcastV2Service(yunxinApiHttpClient),
		systemNotificationService: system_notification.NewCustomNotificationV2Service(yunxinApiHttpClient),
		subscriptionService:       subscription.NewSubscriptionV2Service(yunxinApiHttpClient),
		aiService:                 ai.NewAiV2Service(yunxinApiHttpClient),
		toolService:               tool.NewToolV2Service(yunxinApiHttpClient),
		signalService:             signal.NewSignalV2Service(yunxinApiHttpClient),
	}
}

// GetAccountService 获取账户服务
func (s *YunxinV2ApiServices) GetAccountService() account.AccountV2Service {
	return s.accountService
}

// GetUserService 获取用户服务
func (s *YunxinV2ApiServices) GetUserService() users.UserV2Service {
	return s.userService
}

// GetFriendService 获取好友服务
func (s *YunxinV2ApiServices) GetFriendService() friend.FriendV2Service {
	return s.friendService
}

// GetTeamService 获取群组服务
func (s *YunxinV2ApiServices) GetTeamService() team.TeamV2Service {
	return s.teamService
}

// GetTeamMemberService 获取群成员服务
func (s *YunxinV2ApiServices) GetTeamMemberService() team_member.TeamMemberV2Service {
	return s.teamMemberService
}

// GetChatroomService 获取聊天室服务
func (s *YunxinV2ApiServices) GetChatroomService() chatroom.ChatroomV2Service {
	return s.chatroomService
}

// GetChatroomMemberService 获取聊天室成员服务
func (s *YunxinV2ApiServices) GetChatroomMemberService() chatroom_member.ChatroomMemberV2Service {
	return s.chatroomMemberService
}

// GetChatroomMessageService 获取聊天室消息服务
func (s *YunxinV2ApiServices) GetChatroomMessageService() chatroom_message.ChatroomMessageV2Service {
	return s.chatroomMessageService
}

// GetChatroomQueueService 获取聊天室队列服务
func (s *YunxinV2ApiServices) GetChatroomQueueService() chatroom_queue.ChatroomQueueV2Service {
	return s.chatroomQueueService
}

// GetMessageService 获取消息服务
func (s *YunxinV2ApiServices) GetMessageService() message.MessageV2Service {
	return s.messageService
}

// GetConversationService 获取会话服务
func (s *YunxinV2ApiServices) GetConversationService() conversation.ConversationV2Service {
	return s.conversationService
}

// GetConversationGroupService 获取会话分组服务
func (s *YunxinV2ApiServices) GetConversationGroupService() conversation_group.ConversationGroupV2Service {
	return s.conversationGroupService
}

// GetConversationUnreadService 获取会话未读服务
func (s *YunxinV2ApiServices) GetConversationUnreadService() conversation_unread.ConversationUnreadV2Service {
	return s.conversationUnreadService
}

// GetMuteService 获取静音服务
func (s *YunxinV2ApiServices) GetMuteService() mute.MuteV2Service {
	return s.muteService
}

// GetBlockService 获取黑名单服务
func (s *YunxinV2ApiServices) GetBlockService() block.BlockV2Service {
	return s.blockService
}

// GetBroadcastService 获取广播服务
func (s *YunxinV2ApiServices) GetBroadcastService() broadcast.BroadcastV2Service {
	return s.broadcastService
}

// GetSystemNotificationService 获取系统通知服务
func (s *YunxinV2ApiServices) GetSystemNotificationService() system_notification.CustomNotificationV2Service {
	return s.systemNotificationService
}

// GetSubscriptionService 获取订阅服务
func (s *YunxinV2ApiServices) GetSubscriptionService() subscription.SubscriptionV2Service {
	return s.subscriptionService
}

// GetAiService 获取AI服务
func (s *YunxinV2ApiServices) GetAiService() ai.AiV2Service {
	return s.aiService
}

// GetToolService 获取工具服务
func (s *YunxinV2ApiServices) GetToolService() tool.ToolV2Service {
	return s.toolService
}

// GetSignalService 获取信令服务
func (s *YunxinV2ApiServices) GetSignalService() signal.SignalV2Service {
	return s.signalService
}
