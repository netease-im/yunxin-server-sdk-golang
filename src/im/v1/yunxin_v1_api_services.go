package v1

import (
	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/im/v1/account"
	"github.com/netease-im/yunxin-server-sdk-golang/src/im/v1/chatroom"
	"github.com/netease-im/yunxin-server-sdk-golang/src/im/v1/chatroom_message"
	"github.com/netease-im/yunxin-server-sdk-golang/src/im/v1/event_subscribe"
	"github.com/netease-im/yunxin-server-sdk-golang/src/im/v1/friend"
	"github.com/netease-im/yunxin-server-sdk-golang/src/im/v1/history"
	"github.com/netease-im/yunxin-server-sdk-golang/src/im/v1/message"
	"github.com/netease-im/yunxin-server-sdk-golang/src/im/v1/signal"
	"github.com/netease-im/yunxin-server-sdk-golang/src/im/v1/special_relation"
	"github.com/netease-im/yunxin-server-sdk-golang/src/im/v1/super_team"
	"github.com/netease-im/yunxin-server-sdk-golang/src/im/v1/system_notification"
	"github.com/netease-im/yunxin-server-sdk-golang/src/im/v1/team"
	"github.com/netease-im/yunxin-server-sdk-golang/src/im/v1/translations"
)

// YunxinV1ApiServices 云信V1 API服务聚合器
type YunxinV1ApiServices struct {
	accountService            account.AccountV1Service
	friendService             friend.FriendV1Service
	historyService            history.HistoryV1Service
	teamService               team.TeamV1Service
	chatRoomService           chatroom.ChatroomV1Service
	superTeamService          super_team.SuperTeamV1Service
	chatroomMessageService    chatroom_message.ChatroomMessageV1Service
	translationService        translations.TranslationV1Service
	specialRelationService    special_relation.SpecialRelationV1Service
	messageService            message.MessageV1Service
	systemNotificationService system_notification.SystemNotificationV1Service
	signalService             signal.SignalV1Service
	eventSubscribeService     event_subscribe.EventSubscribeV1Service
}

// NewYunxinV1ApiServices 创建云信V1 API服务聚合器
func NewYunxinV1ApiServices(yunxinApiHttpClient core.YunxinApiHttpClient) *YunxinV1ApiServices {
	return &YunxinV1ApiServices{
		accountService:            account.NewAccountV1Service(yunxinApiHttpClient),
		friendService:             friend.NewFriendV1Service(yunxinApiHttpClient),
		historyService:            history.NewHistoryV1Service(yunxinApiHttpClient),
		teamService:               team.NewTeamV1Service(yunxinApiHttpClient),
		chatRoomService:           chatroom.NewChatroomV1Service(yunxinApiHttpClient),
		superTeamService:          super_team.NewSuperTeamV1Service(yunxinApiHttpClient),
		chatroomMessageService:    chatroom_message.NewChatroomMessageV1Service(yunxinApiHttpClient),
		translationService:        translations.NewTranslationsV1Service(yunxinApiHttpClient),
		specialRelationService:    special_relation.NewSpecialRelationV1Service(yunxinApiHttpClient),
		messageService:            message.NewMessageV1Service(yunxinApiHttpClient),
		systemNotificationService: system_notification.NewSystemNotificationV1Service(yunxinApiHttpClient),
		signalService:             signal.NewSignalV1Service(yunxinApiHttpClient),
		eventSubscribeService:     event_subscribe.NewEventSubscribeV1Service(yunxinApiHttpClient),
	}
}

// GetAccountService 获取账户服务
func (s *YunxinV1ApiServices) GetAccountService() account.AccountV1Service {
	return s.accountService
}

// GetFriendService 获取好友服务
func (s *YunxinV1ApiServices) GetFriendService() friend.FriendV1Service {
	return s.friendService
}

// GetHistoryService 获取历史记录服务
func (s *YunxinV1ApiServices) GetHistoryService() history.HistoryV1Service {
	return s.historyService
}

// GetTeamService 获取群组服务
func (s *YunxinV1ApiServices) GetTeamService() team.TeamV1Service {
	return s.teamService
}

// GetChatRoomService 获取聊天室服务
func (s *YunxinV1ApiServices) GetChatRoomService() chatroom.ChatroomV1Service {
	return s.chatRoomService
}

// GetSuperTeamService 获取超大群服务
func (s *YunxinV1ApiServices) GetSuperTeamService() super_team.SuperTeamV1Service {
	return s.superTeamService
}

// GetChatroomMessageService 获取聊天室消息服务
func (s *YunxinV1ApiServices) GetChatroomMessageService() chatroom_message.ChatroomMessageV1Service {
	return s.chatroomMessageService
}

// GetTranslationService 获取翻译服务
func (s *YunxinV1ApiServices) GetTranslationService() translations.TranslationV1Service {
	return s.translationService
}

// GetSpecialRelationService 获取特殊关系服务
func (s *YunxinV1ApiServices) GetSpecialRelationService() special_relation.SpecialRelationV1Service {
	return s.specialRelationService
}

// GetMessageService 获取消息服务
func (s *YunxinV1ApiServices) GetMessageService() message.MessageV1Service {
	return s.messageService
}

// GetSystemNotificationService 获取系统通知服务
func (s *YunxinV1ApiServices) GetSystemNotificationService() system_notification.SystemNotificationV1Service {
	return s.systemNotificationService
}

// GetEventSubscribeService 获取事件订阅服务
func (s *YunxinV1ApiServices) GetEventSubscribeService() event_subscribe.EventSubscribeV1Service {
	return s.eventSubscribeService
}

// GetSignalService 获取信令服务
func (s *YunxinV1ApiServices) GetSignalService() signal.SignalV1Service {
	return s.signalService
}
